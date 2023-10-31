package service

import (
	"strconv"

	"jobportal-graphql/graph/model"
	newModel "jobportal-graphql/models"
)

func (s *Service) CreateCompany(companyData model.NewCompany) (*model.Company, error) {
	companyDetails := newModel.Company{
		Name:     companyData.Companyname,
		Location: companyData.Address,
		Salary:   companyData.Sal,
	}

	companyDetails, err := s.userRepo.CreateCompany(companyDetails)
	if err != nil {
		return nil, err
	}

	return &model.Company{
		ID:          strconv.FormatUint(uint64(companyDetails.ID), 10),
		Companyname: companyData.Companyname,
		Address:     companyData.Address,
		Sal:         companyData.Sal,
	}, nil

}
func (s *Service) ViewAllCompanies() ([]*model.Company, error) {
	companies, err := s.userRepo.ViewAllCompanies()
	if err != nil {
		return nil, err
	}
	cpDatas := []*model.Company{}

	for _, v := range companies {
		cpData := model.Company{
			ID:       strconv.FormatUint(uint64(v.ID), 10),
			Companyname:  v.Name,
			Address: v.Location,
			Sal:   v.Salary,
		}
		cpDatas = append(cpDatas, &cpData)
	}
	return cpDatas, nil
}

