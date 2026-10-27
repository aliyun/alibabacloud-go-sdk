// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCuaChatCompletionStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CuaChatCompletionStreamResponseBody
	GetRequestId() *string
}

type CuaChatCompletionStreamResponseBody struct {
	// example:
	//
	// request-20260918-001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CuaChatCompletionStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CuaChatCompletionStreamResponseBody) GoString() string {
	return s.String()
}

func (s *CuaChatCompletionStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CuaChatCompletionStreamResponseBody) SetRequestId(v string) *CuaChatCompletionStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *CuaChatCompletionStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
