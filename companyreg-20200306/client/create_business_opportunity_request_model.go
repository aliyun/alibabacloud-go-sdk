// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBusinessOpportunityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *CreateBusinessOpportunityRequest
	GetBizType() *string
	SetContactName(v string) *CreateBusinessOpportunityRequest
	GetContactName() *string
	SetMobile(v string) *CreateBusinessOpportunityRequest
	GetMobile() *string
	SetSource(v int32) *CreateBusinessOpportunityRequest
	GetSource() *int32
	SetVCode(v string) *CreateBusinessOpportunityRequest
	GetVCode() *string
}

type CreateBusinessOpportunityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// esp.hightech
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 18704330000
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// This parameter is required.
	Source *int32 `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 345674
	VCode *string `json:"VCode,omitempty" xml:"VCode,omitempty"`
}

func (s CreateBusinessOpportunityRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBusinessOpportunityRequest) GoString() string {
	return s.String()
}

func (s *CreateBusinessOpportunityRequest) GetBizType() *string {
	return s.BizType
}

func (s *CreateBusinessOpportunityRequest) GetContactName() *string {
	return s.ContactName
}

func (s *CreateBusinessOpportunityRequest) GetMobile() *string {
	return s.Mobile
}

func (s *CreateBusinessOpportunityRequest) GetSource() *int32 {
	return s.Source
}

func (s *CreateBusinessOpportunityRequest) GetVCode() *string {
	return s.VCode
}

func (s *CreateBusinessOpportunityRequest) SetBizType(v string) *CreateBusinessOpportunityRequest {
	s.BizType = &v
	return s
}

func (s *CreateBusinessOpportunityRequest) SetContactName(v string) *CreateBusinessOpportunityRequest {
	s.ContactName = &v
	return s
}

func (s *CreateBusinessOpportunityRequest) SetMobile(v string) *CreateBusinessOpportunityRequest {
	s.Mobile = &v
	return s
}

func (s *CreateBusinessOpportunityRequest) SetSource(v int32) *CreateBusinessOpportunityRequest {
	s.Source = &v
	return s
}

func (s *CreateBusinessOpportunityRequest) SetVCode(v string) *CreateBusinessOpportunityRequest {
	s.VCode = &v
	return s
}

func (s *CreateBusinessOpportunityRequest) Validate() error {
	return dara.Validate(s)
}
