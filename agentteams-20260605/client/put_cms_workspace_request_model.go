// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutCmsWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *PutCmsWorkspaceRequest
	GetInstanceId() *string
}

type PutCmsWorkspaceRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s PutCmsWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s PutCmsWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *PutCmsWorkspaceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PutCmsWorkspaceRequest) SetInstanceId(v string) *PutCmsWorkspaceRequest {
	s.InstanceId = &v
	return s
}

func (s *PutCmsWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
