// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStreamingJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteStreamingJobRequest
	GetDBInstanceId() *string
	SetJobId(v int32) *DeleteStreamingJobRequest
	GetJobId() *int32
	SetRegionId(v string) *DeleteStreamingJobRequest
	GetRegionId() *string
}

type DeleteStreamingJobRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp10g78o9807yv9h3
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	JobId *int32 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteStreamingJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStreamingJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteStreamingJobRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteStreamingJobRequest) GetJobId() *int32 {
	return s.JobId
}

func (s *DeleteStreamingJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteStreamingJobRequest) SetDBInstanceId(v string) *DeleteStreamingJobRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteStreamingJobRequest) SetJobId(v int32) *DeleteStreamingJobRequest {
	s.JobId = &v
	return s
}

func (s *DeleteStreamingJobRequest) SetRegionId(v string) *DeleteStreamingJobRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteStreamingJobRequest) Validate() error {
	return dara.Validate(s)
}
