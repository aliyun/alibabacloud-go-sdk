// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePluginClassResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePluginClassResponseBody
	GetCode() *string
	SetData(v *CreatePluginClassResponseBodyData) *CreatePluginClassResponseBody
	GetData() *CreatePluginClassResponseBodyData
	SetMessage(v string) *CreatePluginClassResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePluginClassResponseBody
	GetRequestId() *string
}

type CreatePluginClassResponseBody struct {
	Code    *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data    *CreatePluginClassResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                            `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreatePluginClassResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePluginClassResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePluginClassResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePluginClassResponseBody) GetData() *CreatePluginClassResponseBodyData {
	return s.Data
}

func (s *CreatePluginClassResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePluginClassResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePluginClassResponseBody) SetCode(v string) *CreatePluginClassResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePluginClassResponseBody) SetData(v *CreatePluginClassResponseBodyData) *CreatePluginClassResponseBody {
	s.Data = v
	return s
}

func (s *CreatePluginClassResponseBody) SetMessage(v string) *CreatePluginClassResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePluginClassResponseBody) SetRequestId(v string) *CreatePluginClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePluginClassResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePluginClassResponseBodyData struct {
	PluginClassId *string `json:"pluginClassId,omitempty" xml:"pluginClassId,omitempty"`
}

func (s CreatePluginClassResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePluginClassResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePluginClassResponseBodyData) GetPluginClassId() *string {
	return s.PluginClassId
}

func (s *CreatePluginClassResponseBodyData) SetPluginClassId(v string) *CreatePluginClassResponseBodyData {
	s.PluginClassId = &v
	return s
}

func (s *CreatePluginClassResponseBodyData) Validate() error {
	return dara.Validate(s)
}
