// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSingleCallByVoiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *SingleCallByVoiceResponseBody
	GetCallId() *string
	SetCode(v string) *SingleCallByVoiceResponseBody
	GetCode() *string
	SetMessage(v string) *SingleCallByVoiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *SingleCallByVoiceResponseBody
	GetRequestId() *string
}

type SingleCallByVoiceResponseBody struct {
	// The unique receipt ID for the call.
	//
	// You can call the [QueryCallDetailByCallId](https://help.aliyun.com/document_detail/393529.html) operation to query the details of the call.
	//
	// example:
	//
	// 116004767703^102806****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.****
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
	// E50FFA85-0B79-4421-A7BD-00B0A271966F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SingleCallByVoiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SingleCallByVoiceResponseBody) GoString() string {
	return s.String()
}

func (s *SingleCallByVoiceResponseBody) GetCallId() *string {
	return s.CallId
}

func (s *SingleCallByVoiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *SingleCallByVoiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SingleCallByVoiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SingleCallByVoiceResponseBody) SetCallId(v string) *SingleCallByVoiceResponseBody {
	s.CallId = &v
	return s
}

func (s *SingleCallByVoiceResponseBody) SetCode(v string) *SingleCallByVoiceResponseBody {
	s.Code = &v
	return s
}

func (s *SingleCallByVoiceResponseBody) SetMessage(v string) *SingleCallByVoiceResponseBody {
	s.Message = &v
	return s
}

func (s *SingleCallByVoiceResponseBody) SetRequestId(v string) *SingleCallByVoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SingleCallByVoiceResponseBody) Validate() error {
	return dara.Validate(s)
}
