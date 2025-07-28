// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserAssetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssets(v []*DescribeUserAssetResponseBodyAssets) *DescribeUserAssetResponseBody
	GetAssets() []*DescribeUserAssetResponseBodyAssets
	SetRequestId(v string) *DescribeUserAssetResponseBody
	GetRequestId() *string
}

type DescribeUserAssetResponseBody struct {
	// The API statistics.
	Assets []*DescribeUserAssetResponseBodyAssets `json:"Assets,omitempty" xml:"Assets,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C1823E96-EF4B-5BD2-9E02-1D18****3ED8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserAssetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAssetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserAssetResponseBody) GetAssets() []*DescribeUserAssetResponseBodyAssets {
	return s.Assets
}

func (s *DescribeUserAssetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserAssetResponseBody) SetAssets(v []*DescribeUserAssetResponseBodyAssets) *DescribeUserAssetResponseBody {
	s.Assets = v
	return s
}

func (s *DescribeUserAssetResponseBody) SetRequestId(v string) *DescribeUserAssetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserAssetResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUserAssetResponseBodyAssets struct {
	// The number of APIs returned.
	//
	// example:
	//
	// 134
	AssetNum *int64 `json:"AssetNum,omitempty" xml:"AssetNum,omitempty"`
	// The time at which the API was called. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// example:
	//
	// 1723435200
	TimeStamp *int64 `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeUserAssetResponseBodyAssets) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAssetResponseBodyAssets) GoString() string {
	return s.String()
}

func (s *DescribeUserAssetResponseBodyAssets) GetAssetNum() *int64 {
	return s.AssetNum
}

func (s *DescribeUserAssetResponseBodyAssets) GetTimeStamp() *int64 {
	return s.TimeStamp
}

func (s *DescribeUserAssetResponseBodyAssets) SetAssetNum(v int64) *DescribeUserAssetResponseBodyAssets {
	s.AssetNum = &v
	return s
}

func (s *DescribeUserAssetResponseBodyAssets) SetTimeStamp(v int64) *DescribeUserAssetResponseBodyAssets {
	s.TimeStamp = &v
	return s
}

func (s *DescribeUserAssetResponseBodyAssets) Validate() error {
	return dara.Validate(s)
}
