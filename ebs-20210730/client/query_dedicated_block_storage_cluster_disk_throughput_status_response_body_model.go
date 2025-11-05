// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody
	GetStatus() *string
}

type QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A37597A6-BB99-19B3-85EA-4C2B91F0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the throughput after setting the throughput by SetDedicatedBlockStorageClusterDiskThroughput api.
	//
	// - SUCCESS: The throughput has been successfully set.
	//
	// - RUNNING: The throughput is currently being set.
	//
	// - WAIT(): The throughput is waiting to be set.
	//
	// - FAIL(): The throughput setting has failed.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody) SetRequestId(v string) *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody) SetStatus(v string) *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody {
	s.Status = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
