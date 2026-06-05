// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelComfyTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CancelComfyTaskResponseBody
	GetCode() *int64
	SetMessage(v string) *CancelComfyTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelComfyTaskResponseBody
	GetRequestId() *string
}

type CancelComfyTaskResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelComfyTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelComfyTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelComfyTaskResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CancelComfyTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelComfyTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelComfyTaskResponseBody) SetCode(v int64) *CancelComfyTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelComfyTaskResponseBody) SetMessage(v string) *CancelComfyTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CancelComfyTaskResponseBody) SetRequestId(v string) *CancelComfyTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelComfyTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
