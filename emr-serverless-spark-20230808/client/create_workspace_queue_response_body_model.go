// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateWorkspaceQueueResponseBody
	GetRequestId() *string
}

type CreateWorkspaceQueueResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateWorkspaceQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceQueueResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkspaceQueueResponseBody) SetRequestId(v string) *CreateWorkspaceQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkspaceQueueResponseBody) Validate() error {
	return dara.Validate(s)
}
