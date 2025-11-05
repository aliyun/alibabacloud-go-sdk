// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerName(v string) *CreateCustomerRequest
	GetCustomerName() *string
	SetCustomerSource(v string) *CreateCustomerRequest
	GetCustomerSource() *string
	SetCustomerSubTrade(v string) *CreateCustomerRequest
	GetCustomerSubTrade() *string
	SetCustomerTrade(v string) *CreateCustomerRequest
	GetCustomerTrade() *string
	SetNation(v string) *CreateCustomerRequest
	GetNation() *string
}

type CreateCustomerRequest struct {
	// Customer\\"s name.
	//
	// This parameter is required.
	//
	// example:
	//
	// DoorBell Marketing
	CustomerName *string `json:"CustomerName,omitempty" xml:"CustomerName,omitempty"`
	// The source/channel that allow client to connected with us. Please enumerate with Customer Source.
	//
	// This parameter is required.
	//
	// example:
	//
	// website
	CustomerSource *string `json:"CustomerSource,omitempty" xml:"CustomerSource,omitempty"`
	// The sub-industry that Customer\\"s business belongs to. Please enumerate with Customer Trade.
	//
	// example:
	//
	// 0101
	CustomerSubTrade *string `json:"CustomerSubTrade,omitempty" xml:"CustomerSubTrade,omitempty"`
	// The industry that Customer\\"s business belongs to. Please enumerate with Customer Trade.
	//
	// This parameter is required.
	//
	// example:
	//
	// 01
	CustomerTrade *string `json:"CustomerTrade,omitempty" xml:"CustomerTrade,omitempty"`
	// The region that Customer choose to launch the Cloud Service. Please use ListCountries to confirm the valid region list for current UID.
	//
	// This parameter is required.
	//
	// example:
	//
	// AR
	Nation *string `json:"Nation,omitempty" xml:"Nation,omitempty"`
}

func (s CreateCustomerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomerRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomerRequest) GetCustomerName() *string {
	return s.CustomerName
}

func (s *CreateCustomerRequest) GetCustomerSource() *string {
	return s.CustomerSource
}

func (s *CreateCustomerRequest) GetCustomerSubTrade() *string {
	return s.CustomerSubTrade
}

func (s *CreateCustomerRequest) GetCustomerTrade() *string {
	return s.CustomerTrade
}

func (s *CreateCustomerRequest) GetNation() *string {
	return s.Nation
}

func (s *CreateCustomerRequest) SetCustomerName(v string) *CreateCustomerRequest {
	s.CustomerName = &v
	return s
}

func (s *CreateCustomerRequest) SetCustomerSource(v string) *CreateCustomerRequest {
	s.CustomerSource = &v
	return s
}

func (s *CreateCustomerRequest) SetCustomerSubTrade(v string) *CreateCustomerRequest {
	s.CustomerSubTrade = &v
	return s
}

func (s *CreateCustomerRequest) SetCustomerTrade(v string) *CreateCustomerRequest {
	s.CustomerTrade = &v
	return s
}

func (s *CreateCustomerRequest) SetNation(v string) *CreateCustomerRequest {
	s.Nation = &v
	return s
}

func (s *CreateCustomerRequest) Validate() error {
	return dara.Validate(s)
}
