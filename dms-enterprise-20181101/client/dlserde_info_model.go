// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDLSerdeInfo interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DLSerdeInfo
	GetDescription() *string
	SetDeserializerClass(v string) *DLSerdeInfo
	GetDeserializerClass() *string
	SetName(v string) *DLSerdeInfo
	GetName() *string
	SetParameters(v map[string]interface{}) *DLSerdeInfo
	GetParameters() map[string]interface{}
	SetSerdeType(v int32) *DLSerdeInfo
	GetSerdeType() *int32
	SetSerializationLib(v string) *DLSerdeInfo
	GetSerializationLib() *string
	SetSerializerClass(v string) *DLSerdeInfo
	GetSerializerClass() *string
}

type DLSerdeInfo struct {
	Description       *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	DeserializerClass *string                `json:"DeserializerClass,omitempty" xml:"DeserializerClass,omitempty"`
	Name              *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Parameters        map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	SerdeType         *int32                 `json:"SerdeType,omitempty" xml:"SerdeType,omitempty"`
	SerializationLib  *string                `json:"SerializationLib,omitempty" xml:"SerializationLib,omitempty"`
	SerializerClass   *string                `json:"SerializerClass,omitempty" xml:"SerializerClass,omitempty"`
}

func (s DLSerdeInfo) String() string {
	return dara.Prettify(s)
}

func (s DLSerdeInfo) GoString() string {
	return s.String()
}

func (s *DLSerdeInfo) GetDescription() *string {
	return s.Description
}

func (s *DLSerdeInfo) GetDeserializerClass() *string {
	return s.DeserializerClass
}

func (s *DLSerdeInfo) GetName() *string {
	return s.Name
}

func (s *DLSerdeInfo) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *DLSerdeInfo) GetSerdeType() *int32 {
	return s.SerdeType
}

func (s *DLSerdeInfo) GetSerializationLib() *string {
	return s.SerializationLib
}

func (s *DLSerdeInfo) GetSerializerClass() *string {
	return s.SerializerClass
}

func (s *DLSerdeInfo) SetDescription(v string) *DLSerdeInfo {
	s.Description = &v
	return s
}

func (s *DLSerdeInfo) SetDeserializerClass(v string) *DLSerdeInfo {
	s.DeserializerClass = &v
	return s
}

func (s *DLSerdeInfo) SetName(v string) *DLSerdeInfo {
	s.Name = &v
	return s
}

func (s *DLSerdeInfo) SetParameters(v map[string]interface{}) *DLSerdeInfo {
	s.Parameters = v
	return s
}

func (s *DLSerdeInfo) SetSerdeType(v int32) *DLSerdeInfo {
	s.SerdeType = &v
	return s
}

func (s *DLSerdeInfo) SetSerializationLib(v string) *DLSerdeInfo {
	s.SerializationLib = &v
	return s
}

func (s *DLSerdeInfo) SetSerializerClass(v string) *DLSerdeInfo {
	s.SerializerClass = &v
	return s
}

func (s *DLSerdeInfo) Validate() error {
	return dara.Validate(s)
}
