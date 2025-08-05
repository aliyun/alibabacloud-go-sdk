// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserAssetIPTrafficInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetIP(v string) *DescribeUserAssetIPTrafficInfoRequest
	GetAssetIP() *string
	SetLang(v string) *DescribeUserAssetIPTrafficInfoRequest
	GetLang() *string
	SetTrafficTime(v string) *DescribeUserAssetIPTrafficInfoRequest
	GetTrafficTime() *string
}

type DescribeUserAssetIPTrafficInfoRequest struct {
	// The IP address of the asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.0.XX.XX
	AssetIP *string `json:"AssetIP,omitempty" xml:"AssetIP,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1656923760
	TrafficTime *string `json:"TrafficTime,omitempty" xml:"TrafficTime,omitempty"`
}

func (s DescribeUserAssetIPTrafficInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAssetIPTrafficInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserAssetIPTrafficInfoRequest) GetAssetIP() *string {
	return s.AssetIP
}

func (s *DescribeUserAssetIPTrafficInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeUserAssetIPTrafficInfoRequest) GetTrafficTime() *string {
	return s.TrafficTime
}

func (s *DescribeUserAssetIPTrafficInfoRequest) SetAssetIP(v string) *DescribeUserAssetIPTrafficInfoRequest {
	s.AssetIP = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoRequest) SetLang(v string) *DescribeUserAssetIPTrafficInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoRequest) SetTrafficTime(v string) *DescribeUserAssetIPTrafficInfoRequest {
	s.TrafficTime = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoRequest) Validate() error {
	return dara.Validate(s)
}
