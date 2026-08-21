// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWorkspaceQueueResponseBody
	GetRequestId() *string
}

type DeleteWorkspaceQueueResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWorkspaceQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceQueueResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkspaceQueueResponseBody) SetRequestId(v string) *DeleteWorkspaceQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkspaceQueueResponseBody) Validate() error {
	return dara.Validate(s)
}
