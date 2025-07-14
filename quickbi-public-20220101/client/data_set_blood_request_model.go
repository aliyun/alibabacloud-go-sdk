// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataSetBloodRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSetIds(v string) *DataSetBloodRequest
	GetDataSetIds() *string
	SetUserId(v string) *DataSetBloodRequest
	GetUserId() *string
	SetWorksType(v string) *DataSetBloodRequest
	GetWorksType() *string
}

type DataSetBloodRequest struct {
	// List of dataset IDs, separated by English commas.
	//
	// This parameter is required.
	//
	// example:
	//
	// 234235234,234235235,234235235
	DataSetIds *string `json:"DataSetIds,omitempty" xml:"DataSetIds,omitempty"`
	// Specify the owner of the report, which is the userId.
	//
	// example:
	//
	// dasasgaj342351
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Specify the type of report:
	//
	// - REPORT: Workbooks
	//
	// - dashboardOfflineQuery: Downloads
	//
	// - DASHBOARD: Dashboard
	//
	// - ANALYSIS: Ad Hoc Analysis
	//
	// - SCREEN: Visualization Screen
	//
	// - PAGE: Old dashboard
	//
	// example:
	//
	// PAGE
	WorksType *string `json:"WorksType,omitempty" xml:"WorksType,omitempty"`
}

func (s DataSetBloodRequest) String() string {
	return dara.Prettify(s)
}

func (s DataSetBloodRequest) GoString() string {
	return s.String()
}

func (s *DataSetBloodRequest) GetDataSetIds() *string {
	return s.DataSetIds
}

func (s *DataSetBloodRequest) GetUserId() *string {
	return s.UserId
}

func (s *DataSetBloodRequest) GetWorksType() *string {
	return s.WorksType
}

func (s *DataSetBloodRequest) SetDataSetIds(v string) *DataSetBloodRequest {
	s.DataSetIds = &v
	return s
}

func (s *DataSetBloodRequest) SetUserId(v string) *DataSetBloodRequest {
	s.UserId = &v
	return s
}

func (s *DataSetBloodRequest) SetWorksType(v string) *DataSetBloodRequest {
	s.WorksType = &v
	return s
}

func (s *DataSetBloodRequest) Validate() error {
	return dara.Validate(s)
}
