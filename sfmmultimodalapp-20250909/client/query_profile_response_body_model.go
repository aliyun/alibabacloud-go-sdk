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
	// example:
	//
	// 200
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryProfileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EDD6123F-0122-5FBF-9A7E-097F319CF478
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Attributes []*QueryProfileResponseBodyDataAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// example:
	//
	// 29C606055
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// ttt
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 384dc4786b9d4f5a8cab0d83112cd5a8
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
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
	// example:
	//
	// 147235
	AttributeId *string `json:"AttributeId,omitempty" xml:"AttributeId,omitempty"`
	// example:
	//
	// []
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// 21D419945-01
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// False
	Immutable *string `json:"Immutable,omitempty" xml:"Immutable,omitempty"`
	// example:
	//
	// 20250724
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
