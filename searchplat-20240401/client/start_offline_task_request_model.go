// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartOfflineTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParallelism(v int32) *StartOfflineTaskRequest
	GetParallelism() *int32
	SetTimestamp(v int64) *StartOfflineTaskRequest
	GetTimestamp() *int64
	SetRegionId(v string) *StartOfflineTaskRequest
	GetRegionId() *string
}

type StartOfflineTaskRequest struct {
	// The degree of task parallelism.
	//
	// example:
	//
	// 4
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

func (s StartOfflineTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StartOfflineTaskRequest) GoString() string {
	return s.String()
}

func (s *StartOfflineTaskRequest) GetParallelism() *int32 {
	return s.Parallelism
}

func (s *StartOfflineTaskRequest) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *StartOfflineTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartOfflineTaskRequest) SetParallelism(v int32) *StartOfflineTaskRequest {
	s.Parallelism = &v
	return s
}

func (s *StartOfflineTaskRequest) SetTimestamp(v int64) *StartOfflineTaskRequest {
	s.Timestamp = &v
	return s
}

func (s *StartOfflineTaskRequest) SetRegionId(v string) *StartOfflineTaskRequest {
	s.RegionId = &v
	return s
}

func (s *StartOfflineTaskRequest) Validate() error {
	return dara.Validate(s)
}
