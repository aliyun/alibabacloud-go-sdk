// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIAgentDialogueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAIAgentDialogueResponseBody
	GetRequestId() *string
}

type DeleteAIAgentDialogueResponseBody struct {
	// example:
	//
	// 7B117AF5-2A1******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAIAgentDialogueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIAgentDialogueResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAIAgentDialogueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAIAgentDialogueResponseBody) SetRequestId(v string) *DeleteAIAgentDialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAIAgentDialogueResponseBody) Validate() error {
	return dara.Validate(s)
}
