// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceLinkedRoleDeletionStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeletionTaskId(v string) *GetServiceLinkedRoleDeletionStatusRequest
	GetDeletionTaskId() *string
}

type GetServiceLinkedRoleDeletionStatusRequest struct {
	// The ID of the deletion task.
	//
	// example:
	//
	// task/acs-service-role/hdr.aliyuncs.com/AliyunServiceRoleForHdr/c4d22c52-247f-4ee1-83a2-6c0460bd****
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" xml:"DeletionTaskId,omitempty"`
}

func (s GetServiceLinkedRoleDeletionStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLinkedRoleDeletionStatusRequest) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleDeletionStatusRequest) GetDeletionTaskId() *string {
	return s.DeletionTaskId
}

func (s *GetServiceLinkedRoleDeletionStatusRequest) SetDeletionTaskId(v string) *GetServiceLinkedRoleDeletionStatusRequest {
	s.DeletionTaskId = &v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusRequest) Validate() error {
	return dara.Validate(s)
}
