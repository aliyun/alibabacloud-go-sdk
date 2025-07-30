// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrecisionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeletePrecisionTaskResponseBody
	GetCode() *string
	SetMessage(v string) *DeletePrecisionTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeletePrecisionTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeletePrecisionTaskResponseBody
	GetSuccess() *bool
}

type DeletePrecisionTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeletePrecisionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePrecisionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrecisionTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeletePrecisionTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePrecisionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePrecisionTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeletePrecisionTaskResponseBody) SetCode(v string) *DeletePrecisionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePrecisionTaskResponseBody) SetMessage(v string) *DeletePrecisionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePrecisionTaskResponseBody) SetRequestId(v string) *DeletePrecisionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrecisionTaskResponseBody) SetSuccess(v bool) *DeletePrecisionTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DeletePrecisionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
