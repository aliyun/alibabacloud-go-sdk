// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiMeteringRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int32) *DescribeApiMeteringRequest
	GetPageNum() *int32
	SetProductCode(v string) *DescribeApiMeteringRequest
	GetProductCode() *string
	SetType(v int32) *DescribeApiMeteringRequest
	GetType() *int32
}

type DescribeApiMeteringRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// cmapi0004****
	ProductCode *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeApiMeteringRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiMeteringRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiMeteringRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeApiMeteringRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeApiMeteringRequest) GetType() *int32 {
	return s.Type
}

func (s *DescribeApiMeteringRequest) SetPageNum(v int32) *DescribeApiMeteringRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeApiMeteringRequest) SetProductCode(v string) *DescribeApiMeteringRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeApiMeteringRequest) SetType(v int32) *DescribeApiMeteringRequest {
	s.Type = &v
	return s
}

func (s *DescribeApiMeteringRequest) Validate() error {
	return dara.Validate(s)
}
