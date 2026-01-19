// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCampaignId(v string) *ListCasesRequest
	GetCampaignId() *string
	SetInstanceId(v string) *ListCasesRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *ListCasesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListCasesRequest
	GetPageSize() *int64
	SetPhoneNumber(v string) *ListCasesRequest
	GetPhoneNumber() *string
}

type ListCasesRequest struct {
	CampaignId  *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber  *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s ListCasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCasesRequest) GoString() string {
	return s.String()
}

func (s *ListCasesRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *ListCasesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCasesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListCasesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCasesRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ListCasesRequest) SetCampaignId(v string) *ListCasesRequest {
	s.CampaignId = &v
	return s
}

func (s *ListCasesRequest) SetInstanceId(v string) *ListCasesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCasesRequest) SetPageNumber(v int64) *ListCasesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCasesRequest) SetPageSize(v int64) *ListCasesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCasesRequest) SetPhoneNumber(v string) *ListCasesRequest {
	s.PhoneNumber = &v
	return s
}

func (s *ListCasesRequest) Validate() error {
	return dara.Validate(s)
}
