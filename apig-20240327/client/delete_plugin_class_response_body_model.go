// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePluginClassResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeletePluginClassResponseBody
	GetCode() *string
	SetMessage(v string) *DeletePluginClassResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeletePluginClassResponseBody
	GetRequestId() *string
}

type DeletePluginClassResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3699C3E4-19D8-5475-A9B8-4524E6C3D855
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeletePluginClassResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePluginClassResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePluginClassResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeletePluginClassResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePluginClassResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePluginClassResponseBody) SetCode(v string) *DeletePluginClassResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePluginClassResponseBody) SetMessage(v string) *DeletePluginClassResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePluginClassResponseBody) SetRequestId(v string) *DeletePluginClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePluginClassResponseBody) Validate() error {
	return dara.Validate(s)
}
