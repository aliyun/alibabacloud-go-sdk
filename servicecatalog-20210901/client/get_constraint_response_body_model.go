// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConstraintResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConstraintDetail(v *GetConstraintResponseBodyConstraintDetail) *GetConstraintResponseBody
	GetConstraintDetail() *GetConstraintResponseBodyConstraintDetail
	SetRequestId(v string) *GetConstraintResponseBody
	GetRequestId() *string
}

type GetConstraintResponseBody struct {
	// The details of the constraint.
	ConstraintDetail *GetConstraintResponseBodyConstraintDetail `json:"ConstraintDetail,omitempty" xml:"ConstraintDetail,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetConstraintResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConstraintResponseBody) GoString() string {
	return s.String()
}

func (s *GetConstraintResponseBody) GetConstraintDetail() *GetConstraintResponseBodyConstraintDetail {
	return s.ConstraintDetail
}

func (s *GetConstraintResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConstraintResponseBody) SetConstraintDetail(v *GetConstraintResponseBodyConstraintDetail) *GetConstraintResponseBody {
	s.ConstraintDetail = v
	return s
}

func (s *GetConstraintResponseBody) SetRequestId(v string) *GetConstraintResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConstraintResponseBody) Validate() error {
	if s.ConstraintDetail != nil {
		if err := s.ConstraintDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConstraintResponseBodyConstraintDetail struct {
	// The configuration of the constraint.
	//
	// Format: { "LocalRoleName": "\\<role_name>" }.
	//
	// example:
	//
	// { "LocalRoleName": "TestRole" }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The ID of the constraint.
	//
	// example:
	//
	// cons-bp1yx7x42v****
	ConstraintId *string `json:"ConstraintId,omitempty" xml:"ConstraintId,omitempty"`
	// The type of the constraint.
	//
	// The value is fixed as Launch, which indicates the launch constraint.
	//
	// example:
	//
	// Launch
	ConstraintType *string `json:"ConstraintType,omitempty" xml:"ConstraintType,omitempty"`
	// The time when the constraint was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-04-12T06:11:26Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the constraint.
	//
	// example:
	//
	// Launch as local role TestRole
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the product portfolio to which the constraint belongs.
	//
	// example:
	//
	// port-bp1yt7582g****
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the product for which the constraint is created.
	//
	// example:
	//
	// prod-bp18r7q127****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The name of the product.
	//
	// example:
	//
	// DEMO-Create an ECS instance.
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
}

func (s GetConstraintResponseBodyConstraintDetail) String() string {
	return dara.Prettify(s)
}

func (s GetConstraintResponseBodyConstraintDetail) GoString() string {
	return s.String()
}

func (s *GetConstraintResponseBodyConstraintDetail) GetConfig() *string {
	return s.Config
}

func (s *GetConstraintResponseBodyConstraintDetail) GetConstraintId() *string {
	return s.ConstraintId
}

func (s *GetConstraintResponseBodyConstraintDetail) GetConstraintType() *string {
	return s.ConstraintType
}

func (s *GetConstraintResponseBodyConstraintDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetConstraintResponseBodyConstraintDetail) GetDescription() *string {
	return s.Description
}

func (s *GetConstraintResponseBodyConstraintDetail) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *GetConstraintResponseBodyConstraintDetail) GetProductId() *string {
	return s.ProductId
}

func (s *GetConstraintResponseBodyConstraintDetail) GetProductName() *string {
	return s.ProductName
}

func (s *GetConstraintResponseBodyConstraintDetail) SetConfig(v string) *GetConstraintResponseBodyConstraintDetail {
	s.Config = &v
	return s
}

func (s *GetConstraintResponseBodyConstraintDetail) SetConstraintId(v string) *GetConstraintResponseBodyConstraintDetail {
	s.ConstraintId = &v
	return s
}

func (s *GetConstraintResponseBodyConstraintDetail) SetConstraintType(v string) *GetConstraintResponseBodyConstraintDetail {
	s.ConstraintType = &v
	return s
}

func (s *GetConstraintResponseBodyConstraintDetail) SetCreateTime(v string) *GetConstraintResponseBodyConstraintDetail {
	s.CreateTime = &v
	return s
}

func (s *GetConstraintResponseBodyConstraintDetail) SetDescription(v string) *GetConstraintResponseBodyConstraintDetail {
	s.Description = &v
	return s
}

func (s *GetConstraintResponseBodyConstraintDetail) SetPortfolioId(v string) *GetConstraintResponseBodyConstraintDetail {
	s.PortfolioId = &v
	return s
}

func (s *GetConstraintResponseBodyConstraintDetail) SetProductId(v string) *GetConstraintResponseBodyConstraintDetail {
	s.ProductId = &v
	return s
}

func (s *GetConstraintResponseBodyConstraintDetail) SetProductName(v string) *GetConstraintResponseBodyConstraintDetail {
	s.ProductName = &v
	return s
}

func (s *GetConstraintResponseBodyConstraintDetail) Validate() error {
	return dara.Validate(s)
}
