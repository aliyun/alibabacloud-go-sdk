// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDedicatedBlockStorageClusterDiskThroughputStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest
	GetClientToken() *string
	SetQosRequestId(v string) *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest
	GetQosRequestId() *string
	SetRegionId(v string) *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest
	GetRegionId() *string
}

type QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the request of SetDedicatedBlockStorageClusterDiskThroughput api.
	//
	// This parameter is required.
	//
	// example:
	//
	// A37597B5-BB99-19B3-85EA-4C2B91F0****
	QosRequestId *string `json:"QosRequestId,omitempty" xml:"QosRequestId,omitempty"`
	// The region ID of the dedicated block storage cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest) GetQosRequestId() *string {
	return s.QosRequestId
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest) SetClientToken(v string) *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest) SetQosRequestId(v string) *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest {
	s.QosRequestId = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest) SetRegionId(v string) *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest {
	s.RegionId = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest) Validate() error {
	return dara.Validate(s)
}
