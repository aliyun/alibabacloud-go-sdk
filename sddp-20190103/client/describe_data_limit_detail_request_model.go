// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataLimitDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureType(v int32) *DescribeDataLimitDetailRequest
	GetFeatureType() *int32
	SetId(v int64) *DescribeDataLimitDetailRequest
	GetId() *int64
	SetLang(v string) *DescribeDataLimitDetailRequest
	GetLang() *string
	SetNetworkType(v int32) *DescribeDataLimitDetailRequest
	GetNetworkType() *int32
}

type DescribeDataLimitDetailRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The unique ID of the data asset that you want to query.
	//
	// > You can call the [DescribeDataLimits](~~DescribeDataLimits~~) operation to query the ID of the data asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12300
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Simplified Chinese.
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The network type of the data asset that you want to query. Valid values:
	//
	// 	- **1**: virtual private cloud (VPC)
	//
	// 	- **2**: classic network
	//
	// example:
	//
	// 1
	NetworkType *int32 `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
}

func (s DescribeDataLimitDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataLimitDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitDetailRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *DescribeDataLimitDetailRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeDataLimitDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDataLimitDetailRequest) GetNetworkType() *int32 {
	return s.NetworkType
}

func (s *DescribeDataLimitDetailRequest) SetFeatureType(v int32) *DescribeDataLimitDetailRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeDataLimitDetailRequest) SetId(v int64) *DescribeDataLimitDetailRequest {
	s.Id = &v
	return s
}

func (s *DescribeDataLimitDetailRequest) SetLang(v string) *DescribeDataLimitDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataLimitDetailRequest) SetNetworkType(v int32) *DescribeDataLimitDetailRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeDataLimitDetailRequest) Validate() error {
	return dara.Validate(s)
}
