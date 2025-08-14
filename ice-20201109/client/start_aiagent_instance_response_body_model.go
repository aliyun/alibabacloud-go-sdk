// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAIAgentInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StartAIAgentInstanceResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *StartAIAgentInstanceResponseBody
	GetRequestId() *string
}

type StartAIAgentInstanceResponseBody struct {
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 7B117AF5-2A16-412C-B127-FA6175ED1AD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartAIAgentInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartAIAgentInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartAIAgentInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartAIAgentInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartAIAgentInstanceResponseBody) SetInstanceId(v string) *StartAIAgentInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *StartAIAgentInstanceResponseBody) SetRequestId(v string) *StartAIAgentInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartAIAgentInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
