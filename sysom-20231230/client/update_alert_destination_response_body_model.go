// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertDestinationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAlertDestinationResponseBody
	GetCode() *string
	SetData(v interface{}) *UpdateAlertDestinationResponseBody
	GetData() interface{}
	SetMessage(v string) *UpdateAlertDestinationResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAlertDestinationResponseBody
	GetRequestId() *string
}

type UpdateAlertDestinationResponseBody struct {
	// The status code.
	//
	// - If `code == Success`, the authorization is successful.
	//
	// - Other status codes indicate authorization failed. Check the `message` field for the detailed fault information.
	//
	// example:
	//
	// SysomOpenAPI.InvalidParameter
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned result.
	//
	// example:
	//
	// {
	//
	// "webhook":"",
	//
	// "sec":""
	//
	// }
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// The error message.
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the request error information.
	//
	// example:
	//
	// Success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateAlertDestinationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertDestinationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlertDestinationResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAlertDestinationResponseBody) GetData() interface{} {
	return s.Data
}

func (s *UpdateAlertDestinationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAlertDestinationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAlertDestinationResponseBody) SetCode(v string) *UpdateAlertDestinationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAlertDestinationResponseBody) SetData(v interface{}) *UpdateAlertDestinationResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAlertDestinationResponseBody) SetMessage(v string) *UpdateAlertDestinationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAlertDestinationResponseBody) SetRequestId(v string) *UpdateAlertDestinationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAlertDestinationResponseBody) Validate() error {
	return dara.Validate(s)
}
