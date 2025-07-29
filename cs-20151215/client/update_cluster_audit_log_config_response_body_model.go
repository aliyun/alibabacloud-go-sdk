// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterAuditLogConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateClusterAuditLogConfigResponseBody
	GetClusterId() *string
	SetRequestId(v string) *UpdateClusterAuditLogConfigResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpdateClusterAuditLogConfigResponseBody
	GetTaskId() *string
}

type UpdateClusterAuditLogConfigResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// c93095129fc41463aa455d89444fd****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 48BD70F6-A7E6-543D-9F23-08DEB764C92E
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// T-5faa48fb31b6b8078d00****
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s UpdateClusterAuditLogConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterAuditLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClusterAuditLogConfigResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateClusterAuditLogConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateClusterAuditLogConfigResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateClusterAuditLogConfigResponseBody) SetClusterId(v string) *UpdateClusterAuditLogConfigResponseBody {
	s.ClusterId = &v
	return s
}

func (s *UpdateClusterAuditLogConfigResponseBody) SetRequestId(v string) *UpdateClusterAuditLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateClusterAuditLogConfigResponseBody) SetTaskId(v string) *UpdateClusterAuditLogConfigResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpdateClusterAuditLogConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
