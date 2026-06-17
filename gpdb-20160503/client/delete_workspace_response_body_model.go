// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWorkspaceResponseBody
	GetRequestId() *string
	SetWorkspaceId(v string) *DeleteWorkspaceResponseBody
	GetWorkspaceId() *string
}

type DeleteWorkspaceResponseBody struct {
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkspaceResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteWorkspaceResponseBody) SetRequestId(v string) *DeleteWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkspaceResponseBody) SetWorkspaceId(v string) *DeleteWorkspaceResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
