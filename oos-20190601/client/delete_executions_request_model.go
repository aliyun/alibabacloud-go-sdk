// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExecutionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionIds(v string) *DeleteExecutionsRequest
	GetExecutionIds() *string
	SetForce(v bool) *DeleteExecutionsRequest
	GetForce() *bool
	SetRegionId(v string) *DeleteExecutionsRequest
	GetRegionId() *string
}

type DeleteExecutionsRequest struct {
	// The execution IDs.
	//
	// You can specify multiple execution IDs in a JSON array in the format of `["xxxxxxxxx", "yyyyyyyyy", ... "zzzzzzzzz"]`. You can specify up to 100 execution IDs at a time. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["exec-xxx"]
	ExecutionIds *string `json:"ExecutionIds,omitempty" xml:"ExecutionIds,omitempty"`
	// Whether to force delete the running task, the default value is false.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteExecutionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExecutionsRequest) GoString() string {
	return s.String()
}

func (s *DeleteExecutionsRequest) GetExecutionIds() *string {
	return s.ExecutionIds
}

func (s *DeleteExecutionsRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteExecutionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteExecutionsRequest) SetExecutionIds(v string) *DeleteExecutionsRequest {
	s.ExecutionIds = &v
	return s
}

func (s *DeleteExecutionsRequest) SetForce(v bool) *DeleteExecutionsRequest {
	s.Force = &v
	return s
}

func (s *DeleteExecutionsRequest) SetRegionId(v string) *DeleteExecutionsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteExecutionsRequest) Validate() error {
	return dara.Validate(s)
}
