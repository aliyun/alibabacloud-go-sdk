// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConstraintRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *CreateConstraintRequest
	GetConfig() *string
	SetConstraintType(v string) *CreateConstraintRequest
	GetConstraintType() *string
	SetDescription(v string) *CreateConstraintRequest
	GetDescription() *string
	SetPortfolioId(v string) *CreateConstraintRequest
	GetPortfolioId() *string
	SetProductId(v string) *CreateConstraintRequest
	GetProductId() *string
}

type CreateConstraintRequest struct {
	// The configuration of the constraint.
	//
	// Format: { "LocalRoleName": "\\<role_name>" }.
	//
	// This parameter is required.
	//
	// example:
	//
	// { "LocalRoleName": "TestRole" }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The type of the constraint.
	//
	// The value is fixed as Launch, which specifies the launch constraint.
	//
	// This parameter is required.
	//
	// example:
	//
	// Launch
	ConstraintType *string `json:"ConstraintType,omitempty" xml:"ConstraintType,omitempty"`
	// The description of the constraint.
	//
	// The value must be 1 to 128 characters in length.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// Launch as local role TestRole
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the product portfolio to which the constraint belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// port-bp1yt7582g****
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the product for which the constraint is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod-bp18r7q127****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s CreateConstraintRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConstraintRequest) GoString() string {
	return s.String()
}

func (s *CreateConstraintRequest) GetConfig() *string {
	return s.Config
}

func (s *CreateConstraintRequest) GetConstraintType() *string {
	return s.ConstraintType
}

func (s *CreateConstraintRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateConstraintRequest) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *CreateConstraintRequest) GetProductId() *string {
	return s.ProductId
}

func (s *CreateConstraintRequest) SetConfig(v string) *CreateConstraintRequest {
	s.Config = &v
	return s
}

func (s *CreateConstraintRequest) SetConstraintType(v string) *CreateConstraintRequest {
	s.ConstraintType = &v
	return s
}

func (s *CreateConstraintRequest) SetDescription(v string) *CreateConstraintRequest {
	s.Description = &v
	return s
}

func (s *CreateConstraintRequest) SetPortfolioId(v string) *CreateConstraintRequest {
	s.PortfolioId = &v
	return s
}

func (s *CreateConstraintRequest) SetProductId(v string) *CreateConstraintRequest {
	s.ProductId = &v
	return s
}

func (s *CreateConstraintRequest) Validate() error {
	return dara.Validate(s)
}
