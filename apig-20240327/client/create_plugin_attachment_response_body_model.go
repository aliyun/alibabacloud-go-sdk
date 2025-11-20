// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePluginAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePluginAttachmentResponseBody
	GetCode() *string
	SetData(v *CreatePluginAttachmentResponseBodyData) *CreatePluginAttachmentResponseBody
	GetData() *CreatePluginAttachmentResponseBodyData
	SetMessage(v string) *CreatePluginAttachmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePluginAttachmentResponseBody
	GetRequestId() *string
}

type CreatePluginAttachmentResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *CreatePluginAttachmentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message returned.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EBCB8485-24F9-54CD-B258-CB15FDB27677
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreatePluginAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePluginAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePluginAttachmentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePluginAttachmentResponseBody) GetData() *CreatePluginAttachmentResponseBodyData {
	return s.Data
}

func (s *CreatePluginAttachmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePluginAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePluginAttachmentResponseBody) SetCode(v string) *CreatePluginAttachmentResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePluginAttachmentResponseBody) SetData(v *CreatePluginAttachmentResponseBodyData) *CreatePluginAttachmentResponseBody {
	s.Data = v
	return s
}

func (s *CreatePluginAttachmentResponseBody) SetMessage(v string) *CreatePluginAttachmentResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePluginAttachmentResponseBody) SetRequestId(v string) *CreatePluginAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePluginAttachmentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePluginAttachmentResponseBodyData struct {
	// The plug-in ID.
	//
	// example:
	//
	// pa-cvs7jpmm1hkgihaqv4a0
	PluginAttachmentId *string `json:"pluginAttachmentId,omitempty" xml:"pluginAttachmentId,omitempty"`
}

func (s CreatePluginAttachmentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePluginAttachmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePluginAttachmentResponseBodyData) GetPluginAttachmentId() *string {
	return s.PluginAttachmentId
}

func (s *CreatePluginAttachmentResponseBodyData) SetPluginAttachmentId(v string) *CreatePluginAttachmentResponseBodyData {
	s.PluginAttachmentId = &v
	return s
}

func (s *CreatePluginAttachmentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
