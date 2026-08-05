// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopOfflineTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParallelism(v int32) *StopOfflineTaskRequest
	GetParallelism() *int32
	SetTimestamp(v int64) *StopOfflineTaskRequest
	GetTimestamp() *int64
	SetRegionId(v string) *StopOfflineTaskRequest
	GetRegionId() *string
}

type StopOfflineTaskRequest struct {
	// The parallelism of the node.
	//
	// example:
	//
	// 2
	Parallelism *int32 `json:"parallelism,omitempty" xml:"parallelism,omitempty"`
	// The start offset.
	//
	// example:
	//
	// 1747900639
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s StopOfflineTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopOfflineTaskRequest) GoString() string {
	return s.String()
}

func (s *StopOfflineTaskRequest) GetParallelism() *int32 {
	return s.Parallelism
}

func (s *StopOfflineTaskRequest) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *StopOfflineTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopOfflineTaskRequest) SetParallelism(v int32) *StopOfflineTaskRequest {
	s.Parallelism = &v
	return s
}

func (s *StopOfflineTaskRequest) SetTimestamp(v int64) *StopOfflineTaskRequest {
	s.Timestamp = &v
	return s
}

func (s *StopOfflineTaskRequest) SetRegionId(v string) *StopOfflineTaskRequest {
	s.RegionId = &v
	return s
}

func (s *StopOfflineTaskRequest) Validate() error {
	return dara.Validate(s)
}
