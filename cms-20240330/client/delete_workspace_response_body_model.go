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
}

type DeleteWorkspaceResponseBody struct {
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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

func (s *DeleteWorkspaceResponseBody) SetRequestId(v string) *DeleteWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
