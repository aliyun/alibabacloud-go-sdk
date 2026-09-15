// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPaChatCompletionStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PaChatCompletionStreamResponseBody
	GetRequestId() *string
}

type PaChatCompletionStreamResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PaChatCompletionStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PaChatCompletionStreamResponseBody) GoString() string {
	return s.String()
}

func (s *PaChatCompletionStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PaChatCompletionStreamResponseBody) SetRequestId(v string) *PaChatCompletionStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *PaChatCompletionStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
