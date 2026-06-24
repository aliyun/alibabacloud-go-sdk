// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureType(v int32) *DescribeUserStatusRequest
	GetFeatureType() *int32
	SetLang(v string) *DescribeUserStatusRequest
	GetLang() *string
}

type DescribeUserStatusRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh_cn**: The default value. Simplified Chinese.
	//
	// - **en_us**: American English.
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeUserStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserStatusRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *DescribeUserStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeUserStatusRequest) SetFeatureType(v int32) *DescribeUserStatusRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeUserStatusRequest) SetLang(v string) *DescribeUserStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeUserStatusRequest) Validate() error {
	return dara.Validate(s)
}
