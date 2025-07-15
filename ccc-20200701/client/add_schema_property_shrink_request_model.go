// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSchemaPropertyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AddSchemaPropertyShrinkRequest
	GetInstanceId() *string
	SetPropertyShrink(v string) *AddSchemaPropertyShrinkRequest
	GetPropertyShrink() *string
	SetRequestId(v string) *AddSchemaPropertyShrinkRequest
	GetRequestId() *string
	SetSchemaId(v string) *AddSchemaPropertyShrinkRequest
	GetSchemaId() *string
}

type AddSchemaPropertyShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// b0eb2742-f37e-4c67-82d4-25c651c1xxxx
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PropertyShrink *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// example:
	//
	// 03C67DAD-EB26-41D8-949D-9B0C470FB716
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// schema id
	//
	// This parameter is required.
	//
	// example:
	//
	// profile
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s AddSchemaPropertyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSchemaPropertyShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddSchemaPropertyShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddSchemaPropertyShrinkRequest) GetPropertyShrink() *string {
	return s.PropertyShrink
}

func (s *AddSchemaPropertyShrinkRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSchemaPropertyShrinkRequest) GetSchemaId() *string {
	return s.SchemaId
}

func (s *AddSchemaPropertyShrinkRequest) SetInstanceId(v string) *AddSchemaPropertyShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *AddSchemaPropertyShrinkRequest) SetPropertyShrink(v string) *AddSchemaPropertyShrinkRequest {
	s.PropertyShrink = &v
	return s
}

func (s *AddSchemaPropertyShrinkRequest) SetRequestId(v string) *AddSchemaPropertyShrinkRequest {
	s.RequestId = &v
	return s
}

func (s *AddSchemaPropertyShrinkRequest) SetSchemaId(v string) *AddSchemaPropertyShrinkRequest {
	s.SchemaId = &v
	return s
}

func (s *AddSchemaPropertyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
