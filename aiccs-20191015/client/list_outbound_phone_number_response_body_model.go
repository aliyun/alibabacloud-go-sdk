// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOutboundPhoneNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListOutboundPhoneNumberResponseBody
	GetCode() *string
	SetData(v []*string) *ListOutboundPhoneNumberResponseBody
	GetData() []*string
	SetHttpStatusCode(v int64) *ListOutboundPhoneNumberResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *ListOutboundPhoneNumberResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListOutboundPhoneNumberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListOutboundPhoneNumberResponseBody
	GetSuccess() *bool
}

type ListOutboundPhoneNumberResponseBody struct {
	// Status code. A return value of "Success" indicates that the request succeeded.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Outbound caller phone numbers.
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListOutboundPhoneNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOutboundPhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ListOutboundPhoneNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListOutboundPhoneNumberResponseBody) GetData() []*string {
	return s.Data
}

func (s *ListOutboundPhoneNumberResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ListOutboundPhoneNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListOutboundPhoneNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOutboundPhoneNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListOutboundPhoneNumberResponseBody) SetCode(v string) *ListOutboundPhoneNumberResponseBody {
	s.Code = &v
	return s
}

func (s *ListOutboundPhoneNumberResponseBody) SetData(v []*string) *ListOutboundPhoneNumberResponseBody {
	s.Data = v
	return s
}

func (s *ListOutboundPhoneNumberResponseBody) SetHttpStatusCode(v int64) *ListOutboundPhoneNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListOutboundPhoneNumberResponseBody) SetMessage(v string) *ListOutboundPhoneNumberResponseBody {
	s.Message = &v
	return s
}

func (s *ListOutboundPhoneNumberResponseBody) SetRequestId(v string) *ListOutboundPhoneNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOutboundPhoneNumberResponseBody) SetSuccess(v bool) *ListOutboundPhoneNumberResponseBody {
	s.Success = &v
	return s
}

func (s *ListOutboundPhoneNumberResponseBody) Validate() error {
	return dara.Validate(s)
}
