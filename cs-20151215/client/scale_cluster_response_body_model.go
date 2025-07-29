// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ScaleClusterResponseBody
	GetClusterId() *string
	SetRequestId(v string) *ScaleClusterResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ScaleClusterResponseBody
	GetTaskId() *string
}

type ScaleClusterResponseBody struct {
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	TaskId    *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s ScaleClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScaleClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleClusterResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *ScaleClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScaleClusterResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ScaleClusterResponseBody) SetClusterId(v string) *ScaleClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *ScaleClusterResponseBody) SetRequestId(v string) *ScaleClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScaleClusterResponseBody) SetTaskId(v string) *ScaleClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *ScaleClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
