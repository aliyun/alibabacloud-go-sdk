// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePortfolioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePortfolioRequest
	GetDescription() *string
	SetPortfolioName(v string) *CreatePortfolioRequest
	GetPortfolioName() *string
	SetProviderName(v string) *CreatePortfolioRequest
	GetProviderName() *string
}

type CreatePortfolioRequest struct {
	// The description of the product portfolio.
	//
	// The value must be 1 to 128 characters in length.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// The description of the product portfolio.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the product portfolio.
	//
	// The value must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEMO-IT services
	PortfolioName *string `json:"PortfolioName,omitempty" xml:"PortfolioName,omitempty"`
	// The provider of the product portfolio.
	//
	// The value must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// IT team
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
}

func (s CreatePortfolioRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePortfolioRequest) GoString() string {
	return s.String()
}

func (s *CreatePortfolioRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePortfolioRequest) GetPortfolioName() *string {
	return s.PortfolioName
}

func (s *CreatePortfolioRequest) GetProviderName() *string {
	return s.ProviderName
}

func (s *CreatePortfolioRequest) SetDescription(v string) *CreatePortfolioRequest {
	s.Description = &v
	return s
}

func (s *CreatePortfolioRequest) SetPortfolioName(v string) *CreatePortfolioRequest {
	s.PortfolioName = &v
	return s
}

func (s *CreatePortfolioRequest) SetProviderName(v string) *CreatePortfolioRequest {
	s.ProviderName = &v
	return s
}

func (s *CreatePortfolioRequest) Validate() error {
	return dara.Validate(s)
}
