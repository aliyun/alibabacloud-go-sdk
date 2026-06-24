// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOpenSourceAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteOpenSourceAccountResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteOpenSourceAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteOpenSourceAccountResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteOpenSourceAccountResponseBody
	GetSuccess() *string
}

type DeleteOpenSourceAccountResponseBody struct {
	// The return code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 035F5EEF-730F-5579-86E3-F1B394A5CBB6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteOpenSourceAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpenSourceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOpenSourceAccountResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteOpenSourceAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteOpenSourceAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOpenSourceAccountResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteOpenSourceAccountResponseBody) SetCode(v int32) *DeleteOpenSourceAccountResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteOpenSourceAccountResponseBody) SetMessage(v string) *DeleteOpenSourceAccountResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteOpenSourceAccountResponseBody) SetRequestId(v string) *DeleteOpenSourceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOpenSourceAccountResponseBody) SetSuccess(v string) *DeleteOpenSourceAccountResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteOpenSourceAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
