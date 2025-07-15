// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChatRoutingProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateChatRoutingProfileResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateChatRoutingProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateChatRoutingProfileResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateChatRoutingProfileResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateChatRoutingProfileResponseBody
	GetRequestId() *string
}

type UpdateChatRoutingProfileResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// BC976D32-AC4C-4E0F-8AA9-F4BC6C4E2B3E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateChatRoutingProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateChatRoutingProfileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateChatRoutingProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateChatRoutingProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateChatRoutingProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateChatRoutingProfileResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateChatRoutingProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateChatRoutingProfileResponseBody) SetCode(v string) *UpdateChatRoutingProfileResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateChatRoutingProfileResponseBody) SetHttpStatusCode(v int32) *UpdateChatRoutingProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateChatRoutingProfileResponseBody) SetMessage(v string) *UpdateChatRoutingProfileResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateChatRoutingProfileResponseBody) SetParams(v []*string) *UpdateChatRoutingProfileResponseBody {
	s.Params = v
	return s
}

func (s *UpdateChatRoutingProfileResponseBody) SetRequestId(v string) *UpdateChatRoutingProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateChatRoutingProfileResponseBody) Validate() error {
	return dara.Validate(s)
}
