// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDownloadExcelListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *GetDownloadExcelListRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *GetDownloadExcelListRequest
	GetBizExtend() map[string]interface{}
	SetCondition(v string) *GetDownloadExcelListRequest
	GetCondition() *string
	SetCountryNames(v []*string) *GetDownloadExcelListRequest
	GetCountryNames() []*string
	SetEndDate(v string) *GetDownloadExcelListRequest
	GetEndDate() *string
	SetGroupId(v string) *GetDownloadExcelListRequest
	GetGroupId() *string
	SetGroupIds(v []*string) *GetDownloadExcelListRequest
	GetGroupIds() []*string
	SetOwnerId(v int64) *GetDownloadExcelListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetDownloadExcelListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetDownloadExcelListRequest
	GetResourceOwnerId() *int64
	SetStartDate(v string) *GetDownloadExcelListRequest
	GetStartDate() *string
}

type GetDownloadExcelListRequest struct {
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// {}
	BizExtend map[string]interface{} `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// example:
	//
	// aa
	Condition    *string   `json:"Condition,omitempty" xml:"Condition,omitempty"`
	CountryNames []*string `json:"CountryNames,omitempty" xml:"CountryNames,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-12-01
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 111
	GroupId              *string   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupIds             []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 2025-11-01
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetDownloadExcelListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDownloadExcelListRequest) GoString() string {
	return s.String()
}

func (s *GetDownloadExcelListRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *GetDownloadExcelListRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *GetDownloadExcelListRequest) GetCondition() *string {
	return s.Condition
}

func (s *GetDownloadExcelListRequest) GetCountryNames() []*string {
	return s.CountryNames
}

func (s *GetDownloadExcelListRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetDownloadExcelListRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetDownloadExcelListRequest) GetGroupIds() []*string {
	return s.GroupIds
}

func (s *GetDownloadExcelListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetDownloadExcelListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetDownloadExcelListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetDownloadExcelListRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *GetDownloadExcelListRequest) SetBizCode(v string) *GetDownloadExcelListRequest {
	s.BizCode = &v
	return s
}

func (s *GetDownloadExcelListRequest) SetBizExtend(v map[string]interface{}) *GetDownloadExcelListRequest {
	s.BizExtend = v
	return s
}

func (s *GetDownloadExcelListRequest) SetCondition(v string) *GetDownloadExcelListRequest {
	s.Condition = &v
	return s
}

func (s *GetDownloadExcelListRequest) SetCountryNames(v []*string) *GetDownloadExcelListRequest {
	s.CountryNames = v
	return s
}

func (s *GetDownloadExcelListRequest) SetEndDate(v string) *GetDownloadExcelListRequest {
	s.EndDate = &v
	return s
}

func (s *GetDownloadExcelListRequest) SetGroupId(v string) *GetDownloadExcelListRequest {
	s.GroupId = &v
	return s
}

func (s *GetDownloadExcelListRequest) SetGroupIds(v []*string) *GetDownloadExcelListRequest {
	s.GroupIds = v
	return s
}

func (s *GetDownloadExcelListRequest) SetOwnerId(v int64) *GetDownloadExcelListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetDownloadExcelListRequest) SetResourceOwnerAccount(v string) *GetDownloadExcelListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetDownloadExcelListRequest) SetResourceOwnerId(v int64) *GetDownloadExcelListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetDownloadExcelListRequest) SetStartDate(v string) *GetDownloadExcelListRequest {
	s.StartDate = &v
	return s
}

func (s *GetDownloadExcelListRequest) Validate() error {
	return dara.Validate(s)
}
