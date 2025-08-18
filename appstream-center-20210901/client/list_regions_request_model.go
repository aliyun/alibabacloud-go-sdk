// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizSource(v string) *ListRegionsRequest
	GetBizSource() *string
	SetProductType(v string) *ListRegionsRequest
	GetProductType() *string
}

type ListRegionsRequest struct {
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// null
	BizSource *string `json:"BizSource,omitempty" xml:"BizSource,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ListRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRegionsRequest) GoString() string {
	return s.String()
}

func (s *ListRegionsRequest) GetBizSource() *string {
	return s.BizSource
}

func (s *ListRegionsRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListRegionsRequest) SetBizSource(v string) *ListRegionsRequest {
	s.BizSource = &v
	return s
}

func (s *ListRegionsRequest) SetProductType(v string) *ListRegionsRequest {
	s.ProductType = &v
	return s
}

func (s *ListRegionsRequest) Validate() error {
	return dara.Validate(s)
}
