// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAIAgentInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAIAgentInstanceResponseBody
	GetRequestId() *string
}

type UpdateAIAgentInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7B117AF5-2A16-412C-B127-FA6175ED1AD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAIAgentInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAIAgentInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAIAgentInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAIAgentInstanceResponseBody) SetRequestId(v string) *UpdateAIAgentInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAIAgentInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
