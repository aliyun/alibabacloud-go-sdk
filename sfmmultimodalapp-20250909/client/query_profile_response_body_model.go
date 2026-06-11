// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryProfileResponseBody
	GetCode() *string
	SetData(v *QueryProfileResponseBodyData) *QueryProfileResponseBody
	GetData() *QueryProfileResponseBodyData
	SetHttpStatusCode(v int32) *QueryProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryProfileResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryProfileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryProfileResponseBody
	GetSuccess() *bool
}

type QueryProfileResponseBody struct {
	Code           *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *QueryProfileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryProfileResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryProfileResponseBody) GetData() *QueryProfileResponseBodyData {
	return s.Data
}

func (s *QueryProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryProfileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryProfileResponseBody) SetCode(v string) *QueryProfileResponseBody {
	s.Code = &v
	return s
}

func (s *QueryProfileResponseBody) SetData(v *QueryProfileResponseBodyData) *QueryProfileResponseBody {
	s.Data = v
	return s
}

func (s *QueryProfileResponseBody) SetHttpStatusCode(v int32) *QueryProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryProfileResponseBody) SetMessage(v string) *QueryProfileResponseBody {
	s.Message = &v
	return s
}

func (s *QueryProfileResponseBody) SetRequestId(v string) *QueryProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryProfileResponseBody) SetSuccess(v bool) *QueryProfileResponseBody {
	s.Success = &v
	return s
}

func (s *QueryProfileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryProfileResponseBodyData struct {
	Attributes  []*QueryProfileResponseBodyDataAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Description *string                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	SchemaId    *string                                   `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s QueryProfileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryProfileResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryProfileResponseBodyData) GetAttributes() []*QueryProfileResponseBodyDataAttributes {
	return s.Attributes
}

func (s *QueryProfileResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *QueryProfileResponseBodyData) GetName() *string {
	return s.Name
}

func (s *QueryProfileResponseBodyData) GetSchemaId() *string {
	return s.SchemaId
}

func (s *QueryProfileResponseBodyData) SetAttributes(v []*QueryProfileResponseBodyDataAttributes) *QueryProfileResponseBodyData {
	s.Attributes = v
	return s
}

func (s *QueryProfileResponseBodyData) SetDescription(v string) *QueryProfileResponseBodyData {
	s.Description = &v
	return s
}

func (s *QueryProfileResponseBodyData) SetName(v string) *QueryProfileResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryProfileResponseBodyData) SetSchemaId(v string) *QueryProfileResponseBodyData {
	s.SchemaId = &v
	return s
}

func (s *QueryProfileResponseBodyData) Validate() error {
	if s.Attributes != nil {
		for _, item := range s.Attributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryProfileResponseBodyDataAttributes struct {
	AttributeId  *string `json:"AttributeId,omitempty" xml:"AttributeId,omitempty"`
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Immutable    *string `json:"Immutable,omitempty" xml:"Immutable,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryProfileResponseBodyDataAttributes) String() string {
	return dara.Prettify(s)
}

func (s QueryProfileResponseBodyDataAttributes) GoString() string {
	return s.String()
}

func (s *QueryProfileResponseBodyDataAttributes) GetAttributeId() *string {
	return s.AttributeId
}

func (s *QueryProfileResponseBodyDataAttributes) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *QueryProfileResponseBodyDataAttributes) GetDescription() *string {
	return s.Description
}

func (s *QueryProfileResponseBodyDataAttributes) GetImmutable() *string {
	return s.Immutable
}

func (s *QueryProfileResponseBodyDataAttributes) GetName() *string {
	return s.Name
}

func (s *QueryProfileResponseBodyDataAttributes) SetAttributeId(v string) *QueryProfileResponseBodyDataAttributes {
	s.AttributeId = &v
	return s
}

func (s *QueryProfileResponseBodyDataAttributes) SetDefaultValue(v string) *QueryProfileResponseBodyDataAttributes {
	s.DefaultValue = &v
	return s
}

func (s *QueryProfileResponseBodyDataAttributes) SetDescription(v string) *QueryProfileResponseBodyDataAttributes {
	s.Description = &v
	return s
}

func (s *QueryProfileResponseBodyDataAttributes) SetImmutable(v string) *QueryProfileResponseBodyDataAttributes {
	s.Immutable = &v
	return s
}

func (s *QueryProfileResponseBodyDataAttributes) SetName(v string) *QueryProfileResponseBodyDataAttributes {
	s.Name = &v
	return s
}

func (s *QueryProfileResponseBodyDataAttributes) Validate() error {
	return dara.Validate(s)
}
