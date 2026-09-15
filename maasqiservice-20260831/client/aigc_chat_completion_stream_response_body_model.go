// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAigcChatCompletionStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AigcChatCompletionStreamResponseBody
	GetRequestId() *string
}

type AigcChatCompletionStreamResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AigcChatCompletionStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AigcChatCompletionStreamResponseBody) GoString() string {
	return s.String()
}

func (s *AigcChatCompletionStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AigcChatCompletionStreamResponseBody) SetRequestId(v string) *AigcChatCompletionStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *AigcChatCompletionStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
