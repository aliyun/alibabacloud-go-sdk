// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceLinkedRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeletionTaskId(v string) *DeleteServiceLinkedRoleResponseBody
	GetDeletionTaskId() *string
	SetRequestId(v string) *DeleteServiceLinkedRoleResponseBody
	GetRequestId() *string
}

type DeleteServiceLinkedRoleResponseBody struct {
	// The ID of the deletion task.
	//
	// example:
	//
	// task/acs-service-role/polardb.aliyuncs.com/AliyunServiceRoleForPolarDB/64c4f9cc-fac2-4692-ae1b-804ae4b9****
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" xml:"DeletionTaskId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B595E5BF-FF5F-4E7F-B95A-B90FE242FEB6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceLinkedRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceLinkedRoleResponseBody) GetDeletionTaskId() *string {
	return s.DeletionTaskId
}

func (s *DeleteServiceLinkedRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceLinkedRoleResponseBody) SetDeletionTaskId(v string) *DeleteServiceLinkedRoleResponseBody {
	s.DeletionTaskId = &v
	return s
}

func (s *DeleteServiceLinkedRoleResponseBody) SetRequestId(v string) *DeleteServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceLinkedRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
