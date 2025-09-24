// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWorkspaceResourceResponseBody
	GetRequestId() *string
	SetResourceIds(v []*string) *DeleteWorkspaceResourceResponseBody
	GetResourceIds() []*string
}

type DeleteWorkspaceResourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource IDs.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
}

func (s DeleteWorkspaceResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkspaceResourceResponseBody) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *DeleteWorkspaceResourceResponseBody) SetRequestId(v string) *DeleteWorkspaceResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkspaceResourceResponseBody) SetResourceIds(v []*string) *DeleteWorkspaceResourceResponseBody {
	s.ResourceIds = v
	return s
}

func (s *DeleteWorkspaceResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
