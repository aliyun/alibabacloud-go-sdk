// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmartCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *SmartCallResponseBody
	GetCallId() *string
	SetCode(v string) *SmartCallResponseBody
	GetCode() *string
	SetMessage(v string) *SmartCallResponseBody
	GetMessage() *string
	SetRequestId(v string) *SmartCallResponseBody
	GetRequestId() *string
}

type SmartCallResponseBody struct {
	// The unique receipt ID for this call.
	//
	// You can call the [QueryCallDetailByCallId](~~QueryCallDetailByCallId~~) operation to query the details of the call based on the receipt ID.
	//
	// example:
	//
	// 116012854210^10281427****
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
	// A90E4451-FED7-49D2-87C8-00700A8C4D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SmartCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SmartCallResponseBody) GoString() string {
	return s.String()
}

func (s *SmartCallResponseBody) GetCallId() *string {
	return s.CallId
}

func (s *SmartCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *SmartCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SmartCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SmartCallResponseBody) SetCallId(v string) *SmartCallResponseBody {
	s.CallId = &v
	return s
}

func (s *SmartCallResponseBody) SetCode(v string) *SmartCallResponseBody {
	s.Code = &v
	return s
}

func (s *SmartCallResponseBody) SetMessage(v string) *SmartCallResponseBody {
	s.Message = &v
	return s
}

func (s *SmartCallResponseBody) SetRequestId(v string) *SmartCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *SmartCallResponseBody) Validate() error {
	return dara.Validate(s)
}
