// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteJobGroupRequest
	GetInstanceId() *string
	SetJobGroupId(v string) *DeleteJobGroupRequest
	GetJobGroupId() *string
}

type DeleteJobGroupRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8fa1953f-4a84-46d8-b80c-8ce9cf684fb3
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fc1fb484-4fe8-4031-b662-5b87ea88590b
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
}

func (s DeleteJobGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteJobGroupRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *DeleteJobGroupRequest) SetInstanceId(v string) *DeleteJobGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteJobGroupRequest) SetJobGroupId(v string) *DeleteJobGroupRequest {
	s.JobGroupId = &v
	return s
}

func (s *DeleteJobGroupRequest) Validate() error {
	return dara.Validate(s)
}
