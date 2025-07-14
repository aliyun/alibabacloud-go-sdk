// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorksByOrganizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int32) *QueryWorksByOrganizationRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryWorksByOrganizationRequest
	GetPageSize() *int32
	SetStatus(v int32) *QueryWorksByOrganizationRequest
	GetStatus() *int32
	SetThirdPartAuthFlag(v int32) *QueryWorksByOrganizationRequest
	GetThirdPartAuthFlag() *int32
	SetWorksType(v string) *QueryWorksByOrganizationRequest
	GetWorksType() *string
}

type QueryWorksByOrganizationRequest struct {
	// Page number.
	//
	// - Default value is 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Number of items per page.
	//
	// - Default value is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the work to query. Possible values:
	//
	// - 0: Unpublished
	//
	// - 1: Published
	//
	// - 2: Modified but not published
	//
	// - 3: Offline
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Third-party embedding status. Possible values:
	//
	// - 0: Embedding not enabled
	//
	// - 1: Embedding enabled
	//
	// example:
	//
	// 1
	ThirdPartAuthFlag *int32 `json:"ThirdPartAuthFlag,omitempty" xml:"ThirdPartAuthFlag,omitempty"`
	// The type of work to query. Leave blank to query all types. Possible values:
	//
	// - DATAPRODUCT: Data Portal
	//
	// - PAGE: Dashboard
	//
	// - REPORT: Spreadsheet
	//
	// - dashboardOfflineQuery: Self-service Data Retrieval
	//
	// example:
	//
	// PAGE
	WorksType *string `json:"WorksType,omitempty" xml:"WorksType,omitempty"`
}

func (s QueryWorksByOrganizationRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryWorksByOrganizationRequest) GoString() string {
	return s.String()
}

func (s *QueryWorksByOrganizationRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryWorksByOrganizationRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryWorksByOrganizationRequest) GetStatus() *int32 {
	return s.Status
}

func (s *QueryWorksByOrganizationRequest) GetThirdPartAuthFlag() *int32 {
	return s.ThirdPartAuthFlag
}

func (s *QueryWorksByOrganizationRequest) GetWorksType() *string {
	return s.WorksType
}

func (s *QueryWorksByOrganizationRequest) SetPageNum(v int32) *QueryWorksByOrganizationRequest {
	s.PageNum = &v
	return s
}

func (s *QueryWorksByOrganizationRequest) SetPageSize(v int32) *QueryWorksByOrganizationRequest {
	s.PageSize = &v
	return s
}

func (s *QueryWorksByOrganizationRequest) SetStatus(v int32) *QueryWorksByOrganizationRequest {
	s.Status = &v
	return s
}

func (s *QueryWorksByOrganizationRequest) SetThirdPartAuthFlag(v int32) *QueryWorksByOrganizationRequest {
	s.ThirdPartAuthFlag = &v
	return s
}

func (s *QueryWorksByOrganizationRequest) SetWorksType(v string) *QueryWorksByOrganizationRequest {
	s.WorksType = &v
	return s
}

func (s *QueryWorksByOrganizationRequest) Validate() error {
	return dara.Validate(s)
}
