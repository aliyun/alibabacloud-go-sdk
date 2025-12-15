// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateClusterResponseBody
	GetClusterId() *string
	SetRequestId(v string) *CreateClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateClusterResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *CreateClusterResponseBody
	GetTaskId() *string
}

type CreateClusterResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F6757FA4-8FED-4602-B7F5-3550C0******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The task ID.
	//
	// example:
	//
	// F6757FA4-8FED-4602-B7F5-3550C0******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateClusterResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateClusterResponseBody) SetClusterId(v string) *CreateClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterResponseBody) SetSuccess(v bool) *CreateClusterResponseBody {
	s.Success = &v
	return s
}

func (s *CreateClusterResponseBody) SetTaskId(v string) *CreateClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
