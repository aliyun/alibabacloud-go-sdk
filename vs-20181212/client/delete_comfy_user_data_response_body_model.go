// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComfyUserDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DeleteComfyUserDataResponseBody
	GetCode() *int64
	SetMessage(v string) *DeleteComfyUserDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteComfyUserDataResponseBody
	GetRequestId() *string
}

type DeleteComfyUserDataResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteComfyUserDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteComfyUserDataResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteComfyUserDataResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DeleteComfyUserDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteComfyUserDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteComfyUserDataResponseBody) SetCode(v int64) *DeleteComfyUserDataResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteComfyUserDataResponseBody) SetMessage(v string) *DeleteComfyUserDataResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteComfyUserDataResponseBody) SetRequestId(v string) *DeleteComfyUserDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteComfyUserDataResponseBody) Validate() error {
	return dara.Validate(s)
}
