// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleOutClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ScaleOutClusterResponseBody
	GetClusterId() *string
	SetRequestId(v string) *ScaleOutClusterResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ScaleOutClusterResponseBody
	GetTaskId() *string
}

type ScaleOutClusterResponseBody struct {
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	TaskId    *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s ScaleOutClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScaleOutClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleOutClusterResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *ScaleOutClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScaleOutClusterResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ScaleOutClusterResponseBody) SetClusterId(v string) *ScaleOutClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *ScaleOutClusterResponseBody) SetRequestId(v string) *ScaleOutClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScaleOutClusterResponseBody) SetTaskId(v string) *ScaleOutClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *ScaleOutClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
