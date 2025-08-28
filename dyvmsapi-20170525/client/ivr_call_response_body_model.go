// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIvrCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *IvrCallResponseBody
	GetCallId() *string
	SetCode(v string) *IvrCallResponseBody
	GetCode() *string
	SetMessage(v string) *IvrCallResponseBody
	GetMessage() *string
	SetRequestId(v string) *IvrCallResponseBody
	GetRequestId() *string
}

type IvrCallResponseBody struct {
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

func (s IvrCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IvrCallResponseBody) GoString() string {
	return s.String()
}

func (s *IvrCallResponseBody) GetCallId() *string {
	return s.CallId
}

func (s *IvrCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *IvrCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IvrCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IvrCallResponseBody) SetCallId(v string) *IvrCallResponseBody {
	s.CallId = &v
	return s
}

func (s *IvrCallResponseBody) SetCode(v string) *IvrCallResponseBody {
	s.Code = &v
	return s
}

func (s *IvrCallResponseBody) SetMessage(v string) *IvrCallResponseBody {
	s.Message = &v
	return s
}

func (s *IvrCallResponseBody) SetRequestId(v string) *IvrCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *IvrCallResponseBody) Validate() error {
	return dara.Validate(s)
}
