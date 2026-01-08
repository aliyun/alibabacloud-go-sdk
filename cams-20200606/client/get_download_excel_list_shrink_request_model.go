// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDownloadExcelListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *GetDownloadExcelListShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *GetDownloadExcelListShrinkRequest
	GetBizExtendShrink() *string
	SetCondition(v string) *GetDownloadExcelListShrinkRequest
	GetCondition() *string
	SetCountryNamesShrink(v string) *GetDownloadExcelListShrinkRequest
	GetCountryNamesShrink() *string
	SetEndDate(v string) *GetDownloadExcelListShrinkRequest
	GetEndDate() *string
	SetGroupId(v string) *GetDownloadExcelListShrinkRequest
	GetGroupId() *string
	SetGroupIdsShrink(v string) *GetDownloadExcelListShrinkRequest
	GetGroupIdsShrink() *string
	SetOwnerId(v int64) *GetDownloadExcelListShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetDownloadExcelListShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetDownloadExcelListShrinkRequest
	GetResourceOwnerId() *int64
	SetStartDate(v string) *GetDownloadExcelListShrinkRequest
	GetStartDate() *string
}

type GetDownloadExcelListShrinkRequest struct {
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// {}
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// example:
	//
	// aa
	Condition          *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	CountryNamesShrink *string `json:"CountryNames,omitempty" xml:"CountryNames,omitempty"`
	// example:
	//
	// 2025-12-01
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 111
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupIdsShrink       *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 2025-11-01
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetDownloadExcelListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDownloadExcelListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDownloadExcelListShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *GetDownloadExcelListShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *GetDownloadExcelListShrinkRequest) GetCondition() *string {
	return s.Condition
}

func (s *GetDownloadExcelListShrinkRequest) GetCountryNamesShrink() *string {
	return s.CountryNamesShrink
}

func (s *GetDownloadExcelListShrinkRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetDownloadExcelListShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetDownloadExcelListShrinkRequest) GetGroupIdsShrink() *string {
	return s.GroupIdsShrink
}

func (s *GetDownloadExcelListShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetDownloadExcelListShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetDownloadExcelListShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetDownloadExcelListShrinkRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *GetDownloadExcelListShrinkRequest) SetBizCode(v string) *GetDownloadExcelListShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *GetDownloadExcelListShrinkRequest) SetBizExtendShrink(v string) *GetDownloadExcelListShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *GetDownloadExcelListShrinkRequest) SetCondition(v string) *GetDownloadExcelListShrinkRequest {
	s.Condition = &v
	return s
}

func (s *GetDownloadExcelListShrinkRequest) SetCountryNamesShrink(v string) *GetDownloadExcelListShrinkRequest {
	s.CountryNamesShrink = &v
	return s
}

func (s *GetDownloadExcelListShrinkRequest) SetEndDate(v string) *GetDownloadExcelListShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetDownloadExcelListShrinkRequest) SetGroupId(v string) *GetDownloadExcelListShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *GetDownloadExcelListShrinkRequest) SetGroupIdsShrink(v string) *GetDownloadExcelListShrinkRequest {
	s.GroupIdsShrink = &v
	return s
}

func (s *GetDownloadExcelListShrinkRequest) SetOwnerId(v int64) *GetDownloadExcelListShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *GetDownloadExcelListShrinkRequest) SetResourceOwnerAccount(v string) *GetDownloadExcelListShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetDownloadExcelListShrinkRequest) SetResourceOwnerId(v int64) *GetDownloadExcelListShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetDownloadExcelListShrinkRequest) SetStartDate(v string) *GetDownloadExcelListShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetDownloadExcelListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
