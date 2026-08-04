// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomerLabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndtime(v string) *AddCustomerLabelRequest
	GetEndtime() *string
	SetLabelSeries(v string) *AddCustomerLabelRequest
	GetLabelSeries() *string
	SetLabelTypes(v []*string) *AddCustomerLabelRequest
	GetLabelTypes() []*string
	SetOrganization(v string) *AddCustomerLabelRequest
	GetOrganization() *string
	SetPK(v int64) *AddCustomerLabelRequest
	GetPK() *int64
	SetStartTime(v string) *AddCustomerLabelRequest
	GetStartTime() *string
	SetToken(v string) *AddCustomerLabelRequest
	GetToken() *string
	SetUserName(v string) *AddCustomerLabelRequest
	GetUserName() *string
}

type AddCustomerLabelRequest struct {
	Endtime *string `json:"Endtime,omitempty" xml:"Endtime,omitempty"`
	// This parameter is required.
	LabelSeries *string `json:"LabelSeries,omitempty" xml:"LabelSeries,omitempty"`
	// This parameter is required.
	LabelTypes []*string `json:"LabelTypes,omitempty" xml:"LabelTypes,omitempty" type:"Repeated"`
	// This parameter is required.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// This parameter is required.
	PK        *int64  `json:"PK,omitempty" xml:"PK,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// This parameter is required.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s AddCustomerLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCustomerLabelRequest) GoString() string {
	return s.String()
}

func (s *AddCustomerLabelRequest) GetEndtime() *string {
	return s.Endtime
}

func (s *AddCustomerLabelRequest) GetLabelSeries() *string {
	return s.LabelSeries
}

func (s *AddCustomerLabelRequest) GetLabelTypes() []*string {
	return s.LabelTypes
}

func (s *AddCustomerLabelRequest) GetOrganization() *string {
	return s.Organization
}

func (s *AddCustomerLabelRequest) GetPK() *int64 {
	return s.PK
}

func (s *AddCustomerLabelRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *AddCustomerLabelRequest) GetToken() *string {
	return s.Token
}

func (s *AddCustomerLabelRequest) GetUserName() *string {
	return s.UserName
}

func (s *AddCustomerLabelRequest) SetEndtime(v string) *AddCustomerLabelRequest {
	s.Endtime = &v
	return s
}

func (s *AddCustomerLabelRequest) SetLabelSeries(v string) *AddCustomerLabelRequest {
	s.LabelSeries = &v
	return s
}

func (s *AddCustomerLabelRequest) SetLabelTypes(v []*string) *AddCustomerLabelRequest {
	s.LabelTypes = v
	return s
}

func (s *AddCustomerLabelRequest) SetOrganization(v string) *AddCustomerLabelRequest {
	s.Organization = &v
	return s
}

func (s *AddCustomerLabelRequest) SetPK(v int64) *AddCustomerLabelRequest {
	s.PK = &v
	return s
}

func (s *AddCustomerLabelRequest) SetStartTime(v string) *AddCustomerLabelRequest {
	s.StartTime = &v
	return s
}

func (s *AddCustomerLabelRequest) SetToken(v string) *AddCustomerLabelRequest {
	s.Token = &v
	return s
}

func (s *AddCustomerLabelRequest) SetUserName(v string) *AddCustomerLabelRequest {
	s.UserName = &v
	return s
}

func (s *AddCustomerLabelRequest) Validate() error {
	return dara.Validate(s)
}
