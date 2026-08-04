// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomerLabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelSeries(v string) *DeleteCustomerLabelRequest
	GetLabelSeries() *string
	SetLabelTypes(v []*string) *DeleteCustomerLabelRequest
	GetLabelTypes() []*string
	SetOrganization(v string) *DeleteCustomerLabelRequest
	GetOrganization() *string
	SetPK(v int64) *DeleteCustomerLabelRequest
	GetPK() *int64
	SetToken(v string) *DeleteCustomerLabelRequest
	GetToken() *string
	SetUserName(v string) *DeleteCustomerLabelRequest
	GetUserName() *string
}

type DeleteCustomerLabelRequest struct {
	// This parameter is required.
	LabelSeries *string `json:"LabelSeries,omitempty" xml:"LabelSeries,omitempty"`
	// This parameter is required.
	LabelTypes []*string `json:"LabelTypes,omitempty" xml:"LabelTypes,omitempty" type:"Repeated"`
	// This parameter is required.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// This parameter is required.
	PK    *int64  `json:"PK,omitempty" xml:"PK,omitempty"`
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// This parameter is required.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DeleteCustomerLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomerLabelRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomerLabelRequest) GetLabelSeries() *string {
	return s.LabelSeries
}

func (s *DeleteCustomerLabelRequest) GetLabelTypes() []*string {
	return s.LabelTypes
}

func (s *DeleteCustomerLabelRequest) GetOrganization() *string {
	return s.Organization
}

func (s *DeleteCustomerLabelRequest) GetPK() *int64 {
	return s.PK
}

func (s *DeleteCustomerLabelRequest) GetToken() *string {
	return s.Token
}

func (s *DeleteCustomerLabelRequest) GetUserName() *string {
	return s.UserName
}

func (s *DeleteCustomerLabelRequest) SetLabelSeries(v string) *DeleteCustomerLabelRequest {
	s.LabelSeries = &v
	return s
}

func (s *DeleteCustomerLabelRequest) SetLabelTypes(v []*string) *DeleteCustomerLabelRequest {
	s.LabelTypes = v
	return s
}

func (s *DeleteCustomerLabelRequest) SetOrganization(v string) *DeleteCustomerLabelRequest {
	s.Organization = &v
	return s
}

func (s *DeleteCustomerLabelRequest) SetPK(v int64) *DeleteCustomerLabelRequest {
	s.PK = &v
	return s
}

func (s *DeleteCustomerLabelRequest) SetToken(v string) *DeleteCustomerLabelRequest {
	s.Token = &v
	return s
}

func (s *DeleteCustomerLabelRequest) SetUserName(v string) *DeleteCustomerLabelRequest {
	s.UserName = &v
	return s
}

func (s *DeleteCustomerLabelRequest) Validate() error {
	return dara.Validate(s)
}
