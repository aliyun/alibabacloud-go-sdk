// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHibernateClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *HibernateClusterResponseBody
	GetRequestId() *string
	SetTaskId(v string) *HibernateClusterResponseBody
	GetTaskId() *string
}

type HibernateClusterResponseBody struct {
	// example:
	//
	// 49511F2D-D56A-5C24-B9AE-C8491E09B***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// T-67d7ec016ce37c0106000***
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s HibernateClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HibernateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *HibernateClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HibernateClusterResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *HibernateClusterResponseBody) SetRequestId(v string) *HibernateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *HibernateClusterResponseBody) SetTaskId(v string) *HibernateClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *HibernateClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
