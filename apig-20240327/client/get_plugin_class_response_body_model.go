// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPluginClassResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPluginClassResponseBody
	GetCode() *string
	SetData(v *GetPluginClassResponseBodyData) *GetPluginClassResponseBody
	GetData() *GetPluginClassResponseBodyData
	SetMessage(v string) *GetPluginClassResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPluginClassResponseBody
	GetRequestId() *string
}

type GetPluginClassResponseBody struct {
	Code    *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GetPluginClassResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                         `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetPluginClassResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPluginClassResponseBody) GoString() string {
	return s.String()
}

func (s *GetPluginClassResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPluginClassResponseBody) GetData() *GetPluginClassResponseBodyData {
	return s.Data
}

func (s *GetPluginClassResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPluginClassResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPluginClassResponseBody) SetCode(v string) *GetPluginClassResponseBody {
	s.Code = &v
	return s
}

func (s *GetPluginClassResponseBody) SetData(v *GetPluginClassResponseBodyData) *GetPluginClassResponseBody {
	s.Data = v
	return s
}

func (s *GetPluginClassResponseBody) SetMessage(v string) *GetPluginClassResponseBody {
	s.Message = &v
	return s
}

func (s *GetPluginClassResponseBody) SetRequestId(v string) *GetPluginClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPluginClassResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPluginClassResponseBodyData struct {
	Alias         *string `json:"alias,omitempty" xml:"alias,omitempty"`
	Description   *string `json:"description,omitempty" xml:"description,omitempty"`
	Document      *string `json:"document,omitempty" xml:"document,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	PublishStatus *string `json:"publishStatus,omitempty" xml:"publishStatus,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
	WasmLanguage  *string `json:"wasmLanguage,omitempty" xml:"wasmLanguage,omitempty"`
}

func (s GetPluginClassResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPluginClassResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPluginClassResponseBodyData) GetAlias() *string {
	return s.Alias
}

func (s *GetPluginClassResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetPluginClassResponseBodyData) GetDocument() *string {
	return s.Document
}

func (s *GetPluginClassResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetPluginClassResponseBodyData) GetPublishStatus() *string {
	return s.PublishStatus
}

func (s *GetPluginClassResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetPluginClassResponseBodyData) GetWasmLanguage() *string {
	return s.WasmLanguage
}

func (s *GetPluginClassResponseBodyData) SetAlias(v string) *GetPluginClassResponseBodyData {
	s.Alias = &v
	return s
}

func (s *GetPluginClassResponseBodyData) SetDescription(v string) *GetPluginClassResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetPluginClassResponseBodyData) SetDocument(v string) *GetPluginClassResponseBodyData {
	s.Document = &v
	return s
}

func (s *GetPluginClassResponseBodyData) SetName(v string) *GetPluginClassResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetPluginClassResponseBodyData) SetPublishStatus(v string) *GetPluginClassResponseBodyData {
	s.PublishStatus = &v
	return s
}

func (s *GetPluginClassResponseBodyData) SetType(v string) *GetPluginClassResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetPluginClassResponseBodyData) SetWasmLanguage(v string) *GetPluginClassResponseBodyData {
	s.WasmLanguage = &v
	return s
}

func (s *GetPluginClassResponseBodyData) Validate() error {
	return dara.Validate(s)
}
