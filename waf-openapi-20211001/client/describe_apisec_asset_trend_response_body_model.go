// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecAssetTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeApisecAssetTrendResponseBodyData) *DescribeApisecAssetTrendResponseBody
	GetData() []*DescribeApisecAssetTrendResponseBodyData
	SetRequestId(v string) *DescribeApisecAssetTrendResponseBody
	GetRequestId() *string
}

type DescribeApisecAssetTrendResponseBody struct {
	// The data returned.
	Data []*DescribeApisecAssetTrendResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Id of the request.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19****5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApisecAssetTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecAssetTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisecAssetTrendResponseBody) GetData() []*DescribeApisecAssetTrendResponseBodyData {
	return s.Data
}

func (s *DescribeApisecAssetTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisecAssetTrendResponseBody) SetData(v []*DescribeApisecAssetTrendResponseBodyData) *DescribeApisecAssetTrendResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApisecAssetTrendResponseBody) SetRequestId(v string) *DescribeApisecAssetTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisecAssetTrendResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApisecAssetTrendResponseBodyData struct {
	// The number of active assets.
	//
	// example:
	//
	// 60
	AssetActive *int64 `json:"AssetActive,omitempty" xml:"AssetActive,omitempty"`
	// The total number of assets.
	//
	// example:
	//
	// 80
	AssetCount *int64 `json:"AssetCount,omitempty" xml:"AssetCount,omitempty"`
	// The number of deactivated assets.
	//
	// example:
	//
	// 20
	AssetOffline *int64 `json:"AssetOffline,omitempty" xml:"AssetOffline,omitempty"`
	// The time for statistics. Specify a UNIX timestamp in UTC. Unit: seconds.
	//
	// example:
	//
	// 1683600042
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeApisecAssetTrendResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecAssetTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApisecAssetTrendResponseBodyData) GetAssetActive() *int64 {
	return s.AssetActive
}

func (s *DescribeApisecAssetTrendResponseBodyData) GetAssetCount() *int64 {
	return s.AssetCount
}

func (s *DescribeApisecAssetTrendResponseBodyData) GetAssetOffline() *int64 {
	return s.AssetOffline
}

func (s *DescribeApisecAssetTrendResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeApisecAssetTrendResponseBodyData) SetAssetActive(v int64) *DescribeApisecAssetTrendResponseBodyData {
	s.AssetActive = &v
	return s
}

func (s *DescribeApisecAssetTrendResponseBodyData) SetAssetCount(v int64) *DescribeApisecAssetTrendResponseBodyData {
	s.AssetCount = &v
	return s
}

func (s *DescribeApisecAssetTrendResponseBodyData) SetAssetOffline(v int64) *DescribeApisecAssetTrendResponseBodyData {
	s.AssetOffline = &v
	return s
}

func (s *DescribeApisecAssetTrendResponseBodyData) SetTimestamp(v int64) *DescribeApisecAssetTrendResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *DescribeApisecAssetTrendResponseBodyData) Validate() error {
	return dara.Validate(s)
}
