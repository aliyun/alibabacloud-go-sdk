// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSingleCallByTtsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *SingleCallByTtsResponseBody
	GetCallId() *string
	SetCode(v string) *SingleCallByTtsResponseBody
	GetCode() *string
	SetMessage(v string) *SingleCallByTtsResponseBody
	GetMessage() *string
	SetRequestId(v string) *SingleCallByTtsResponseBody
	GetRequestId() *string
}

type SingleCallByTtsResponseBody struct {
	// The unique receipt ID of the call.
	//
	// You can call the [QueryCallDetailByCallId](https://help.aliyun.com/document_detail/393529.html) operation to query the details of the call based on the receipt ID.
	//
	// example:
	//
	// 116012354148^10281378****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/112502.html).
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
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SingleCallByTtsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SingleCallByTtsResponseBody) GoString() string {
	return s.String()
}

func (s *SingleCallByTtsResponseBody) GetCallId() *string {
	return s.CallId
}

func (s *SingleCallByTtsResponseBody) GetCode() *string {
	return s.Code
}

func (s *SingleCallByTtsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SingleCallByTtsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SingleCallByTtsResponseBody) SetCallId(v string) *SingleCallByTtsResponseBody {
	s.CallId = &v
	return s
}

func (s *SingleCallByTtsResponseBody) SetCode(v string) *SingleCallByTtsResponseBody {
	s.Code = &v
	return s
}

func (s *SingleCallByTtsResponseBody) SetMessage(v string) *SingleCallByTtsResponseBody {
	s.Message = &v
	return s
}

func (s *SingleCallByTtsResponseBody) SetRequestId(v string) *SingleCallByTtsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SingleCallByTtsResponseBody) Validate() error {
	return dara.Validate(s)
}
