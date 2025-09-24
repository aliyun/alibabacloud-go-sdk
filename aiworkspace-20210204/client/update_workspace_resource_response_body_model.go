// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWorkspaceResourceResponseBody
	GetRequestId() *string
	SetResourceIds(v []*string) *UpdateWorkspaceResourceResponseBody
	GetResourceIds() []*string
}

type UpdateWorkspaceResourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The updated resource IDs.
	//
	// example:
	//
	// Resource-dks******jkf
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
}

func (s UpdateWorkspaceResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkspaceResourceResponseBody) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *UpdateWorkspaceResourceResponseBody) SetRequestId(v string) *UpdateWorkspaceResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkspaceResourceResponseBody) SetResourceIds(v []*string) *UpdateWorkspaceResourceResponseBody {
	s.ResourceIds = v
	return s
}

func (s *UpdateWorkspaceResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
