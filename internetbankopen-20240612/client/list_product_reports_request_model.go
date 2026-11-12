// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLatestOnly(v bool) *ListProductReportsRequest
	GetLatestOnly() *bool
	SetProductCode(v string) *ListProductReportsRequest
	GetProductCode() *string
}

type ListProductReportsRequest struct {
	LatestOnly *bool `json:"LatestOnly,omitempty" xml:"LatestOnly,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cloudsec.demo.v1
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListProductReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductReportsRequest) GoString() string {
	return s.String()
}

func (s *ListProductReportsRequest) GetLatestOnly() *bool {
	return s.LatestOnly
}

func (s *ListProductReportsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListProductReportsRequest) SetLatestOnly(v bool) *ListProductReportsRequest {
	s.LatestOnly = &v
	return s
}

func (s *ListProductReportsRequest) SetProductCode(v string) *ListProductReportsRequest {
	s.ProductCode = &v
	return s
}

func (s *ListProductReportsRequest) Validate() error {
	return dara.Validate(s)
}
