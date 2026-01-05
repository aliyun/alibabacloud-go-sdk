// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceProductArn(v string) *CopyProductRequest
	GetSourceProductArn() *string
	SetTargetProductName(v string) *CopyProductRequest
	GetTargetProductName() *string
}

type CopyProductRequest struct {
	// The Alibaba Cloud Resource Name (ARN) of the source product.
	//
	// > The source product can belong to the current account or belong to a product portfolio that is shared by another account.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:servicecatalog:cn-hangzhou:146611588617****:product/prod-bp18r7q127****
	SourceProductArn *string `json:"SourceProductArn,omitempty" xml:"SourceProductArn,omitempty"`
	// The name of the destination product.
	//
	// example:
	//
	// DEMO-ECS
	TargetProductName *string `json:"TargetProductName,omitempty" xml:"TargetProductName,omitempty"`
}

func (s CopyProductRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyProductRequest) GoString() string {
	return s.String()
}

func (s *CopyProductRequest) GetSourceProductArn() *string {
	return s.SourceProductArn
}

func (s *CopyProductRequest) GetTargetProductName() *string {
	return s.TargetProductName
}

func (s *CopyProductRequest) SetSourceProductArn(v string) *CopyProductRequest {
	s.SourceProductArn = &v
	return s
}

func (s *CopyProductRequest) SetTargetProductName(v string) *CopyProductRequest {
	s.TargetProductName = &v
	return s
}

func (s *CopyProductRequest) Validate() error {
	return dara.Validate(s)
}
