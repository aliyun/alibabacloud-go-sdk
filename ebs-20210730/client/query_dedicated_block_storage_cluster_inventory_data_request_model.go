// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDedicatedBlockStorageClusterInventoryDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *QueryDedicatedBlockStorageClusterInventoryDataRequest
	GetClientToken() *string
	SetDbscId(v string) *QueryDedicatedBlockStorageClusterInventoryDataRequest
	GetDbscId() *string
	SetEndTime(v int64) *QueryDedicatedBlockStorageClusterInventoryDataRequest
	GetEndTime() *int64
	SetPeriod(v int32) *QueryDedicatedBlockStorageClusterInventoryDataRequest
	GetPeriod() *int32
	SetRegionId(v string) *QueryDedicatedBlockStorageClusterInventoryDataRequest
	GetRegionId() *string
	SetStartTime(v int64) *QueryDedicatedBlockStorageClusterInventoryDataRequest
	GetStartTime() *int64
}

type QueryDedicatedBlockStorageClusterInventoryDataRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests.
	//
	// The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure idempotence ](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the dedicated block storage cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbsc-xxx
	DbscId *string `json:"DbscId,omitempty" xml:"DbscId,omitempty"`
	// End timestamp of trend data.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1606403800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time interval (seconds) between data retrieval points.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID of the dedicated block storage cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Start timestamp of trend data.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1606303800
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryDedicatedBlockStorageClusterInventoryDataRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDedicatedBlockStorageClusterInventoryDataRequest) GoString() string {
	return s.String()
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataRequest) GetDbscId() *string {
	return s.DbscId
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataRequest) SetClientToken(v string) *QueryDedicatedBlockStorageClusterInventoryDataRequest {
	s.ClientToken = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataRequest) SetDbscId(v string) *QueryDedicatedBlockStorageClusterInventoryDataRequest {
	s.DbscId = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataRequest) SetEndTime(v int64) *QueryDedicatedBlockStorageClusterInventoryDataRequest {
	s.EndTime = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataRequest) SetPeriod(v int32) *QueryDedicatedBlockStorageClusterInventoryDataRequest {
	s.Period = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataRequest) SetRegionId(v string) *QueryDedicatedBlockStorageClusterInventoryDataRequest {
	s.RegionId = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataRequest) SetStartTime(v int64) *QueryDedicatedBlockStorageClusterInventoryDataRequest {
	s.StartTime = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataRequest) Validate() error {
	return dara.Validate(s)
}
