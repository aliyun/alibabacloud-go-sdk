// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertDestinationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAlertDestinationResponseBody
	GetCode() *string
	SetData(v interface{}) *DeleteAlertDestinationResponseBody
	GetData() interface{}
	SetMessage(v string) *DeleteAlertDestinationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAlertDestinationResponseBody
	GetRequestId() *string
}

type DeleteAlertDestinationResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// {}
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// SysomOpenAPIException: SysomOpenAPI.InvalidParameter Invalid params, should be json string or dict
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteAlertDestinationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertDestinationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertDestinationResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAlertDestinationResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DeleteAlertDestinationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAlertDestinationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAlertDestinationResponseBody) SetCode(v string) *DeleteAlertDestinationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAlertDestinationResponseBody) SetData(v interface{}) *DeleteAlertDestinationResponseBody {
	s.Data = v
	return s
}

func (s *DeleteAlertDestinationResponseBody) SetMessage(v string) *DeleteAlertDestinationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAlertDestinationResponseBody) SetRequestId(v string) *DeleteAlertDestinationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlertDestinationResponseBody) Validate() error {
	return dara.Validate(s)
}
