// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAIAgentInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopAIAgentInstanceResponseBody
	GetRequestId() *string
}

type StopAIAgentInstanceResponseBody struct {
	// example:
	//
	// 7B117AF5-2A16-412C-B127-FA6175ED1AD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopAIAgentInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopAIAgentInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopAIAgentInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopAIAgentInstanceResponseBody) SetRequestId(v string) *StopAIAgentInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopAIAgentInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
