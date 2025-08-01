// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMessageQueueRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateMessageQueueRouteResponseBody
	GetCode() *int32
	SetData(v string) *UpdateMessageQueueRouteResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateMessageQueueRouteResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateMessageQueueRouteResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMessageQueueRouteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMessageQueueRouteResponseBody
	GetSuccess() *bool
}

type UpdateMessageQueueRouteResponseBody struct {
	// The status code returned. The value 200 indicates that the request was successful. Other values indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data of the node.
	//
	// example:
	//
	// True
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E3919C62-876A-5926-A0BC-18351A24FA35
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMessageQueueRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageQueueRouteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMessageQueueRouteResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateMessageQueueRouteResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateMessageQueueRouteResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateMessageQueueRouteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMessageQueueRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMessageQueueRouteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMessageQueueRouteResponseBody) SetCode(v int32) *UpdateMessageQueueRouteResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMessageQueueRouteResponseBody) SetData(v string) *UpdateMessageQueueRouteResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateMessageQueueRouteResponseBody) SetHttpStatusCode(v int32) *UpdateMessageQueueRouteResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMessageQueueRouteResponseBody) SetMessage(v string) *UpdateMessageQueueRouteResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMessageQueueRouteResponseBody) SetRequestId(v string) *UpdateMessageQueueRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMessageQueueRouteResponseBody) SetSuccess(v bool) *UpdateMessageQueueRouteResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMessageQueueRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
