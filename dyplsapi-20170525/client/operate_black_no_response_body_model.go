// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateBlackNoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OperateBlackNoResponseBody
	GetCode() *string
	SetMessage(v string) *OperateBlackNoResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateBlackNoResponseBody
	GetRequestId() *string
}

type OperateBlackNoResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other status codes indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2D1AEB96-96D0-454E-B0DC-AE2A8DF08020
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateBlackNoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateBlackNoResponseBody) GoString() string {
	return s.String()
}

func (s *OperateBlackNoResponseBody) GetCode() *string {
	return s.Code
}

func (s *OperateBlackNoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateBlackNoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateBlackNoResponseBody) SetCode(v string) *OperateBlackNoResponseBody {
	s.Code = &v
	return s
}

func (s *OperateBlackNoResponseBody) SetMessage(v string) *OperateBlackNoResponseBody {
	s.Message = &v
	return s
}

func (s *OperateBlackNoResponseBody) SetRequestId(v string) *OperateBlackNoResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateBlackNoResponseBody) Validate() error {
	return dara.Validate(s)
}
