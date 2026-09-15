// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGuiChatCompletionStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GuiChatCompletionStreamResponseBody
	GetRequestId() *string
}

type GuiChatCompletionStreamResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GuiChatCompletionStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GuiChatCompletionStreamResponseBody) GoString() string {
	return s.String()
}

func (s *GuiChatCompletionStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GuiChatCompletionStreamResponseBody) SetRequestId(v string) *GuiChatCompletionStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *GuiChatCompletionStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
