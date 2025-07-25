// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmAccessStrategiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeGtmAccessStrategiesRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeGtmAccessStrategiesRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeGtmAccessStrategiesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGtmAccessStrategiesRequest
	GetPageSize() *int32
}

type DescribeGtmAccessStrategiesRequest struct {
	// The ID of the GTM instance whose access policies you want to query.
	//
	// example:
	//
	// instance1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language used by the user.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeGtmAccessStrategiesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmAccessStrategiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategiesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGtmAccessStrategiesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGtmAccessStrategiesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGtmAccessStrategiesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGtmAccessStrategiesRequest) SetInstanceId(v string) *DescribeGtmAccessStrategiesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGtmAccessStrategiesRequest) SetLang(v string) *DescribeGtmAccessStrategiesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmAccessStrategiesRequest) SetPageNumber(v int32) *DescribeGtmAccessStrategiesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGtmAccessStrategiesRequest) SetPageSize(v int32) *DescribeGtmAccessStrategiesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGtmAccessStrategiesRequest) Validate() error {
	return dara.Validate(s)
}
