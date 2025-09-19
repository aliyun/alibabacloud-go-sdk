// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLoginBaseConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyLoginBaseConfigResponseBody
	GetCode() *string
	SetMessage(v string) *ModifyLoginBaseConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyLoginBaseConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyLoginBaseConfigResponseBody
	GetSuccess() *bool
}

type ModifyLoginBaseConfigResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// AB96FDDF-ED29-52B1-9FAE-8203F2808F24
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyLoginBaseConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLoginBaseConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLoginBaseConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyLoginBaseConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyLoginBaseConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLoginBaseConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyLoginBaseConfigResponseBody) SetCode(v string) *ModifyLoginBaseConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyLoginBaseConfigResponseBody) SetMessage(v string) *ModifyLoginBaseConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyLoginBaseConfigResponseBody) SetRequestId(v string) *ModifyLoginBaseConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLoginBaseConfigResponseBody) SetSuccess(v bool) *ModifyLoginBaseConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyLoginBaseConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
