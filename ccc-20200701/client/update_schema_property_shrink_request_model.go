// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSchemaPropertyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateSchemaPropertyShrinkRequest
	GetInstanceId() *string
	SetPropertyShrink(v string) *UpdateSchemaPropertyShrinkRequest
	GetPropertyShrink() *string
	SetRequestId(v string) *UpdateSchemaPropertyShrinkRequest
	GetRequestId() *string
	SetSchemaId(v string) *UpdateSchemaPropertyShrinkRequest
	GetSchemaId() *string
}

type UpdateSchemaPropertyShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// b0eb2742-f37e-4c67-82d4-25c651c1xxxx
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PropertyShrink *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// example:
	//
	// BC976D32-AC4C-4E0F-8AA9-F4BC6C4E2B3E
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

func (s UpdateSchemaPropertyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSchemaPropertyShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSchemaPropertyShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateSchemaPropertyShrinkRequest) GetPropertyShrink() *string {
	return s.PropertyShrink
}

func (s *UpdateSchemaPropertyShrinkRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSchemaPropertyShrinkRequest) GetSchemaId() *string {
	return s.SchemaId
}

func (s *UpdateSchemaPropertyShrinkRequest) SetInstanceId(v string) *UpdateSchemaPropertyShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateSchemaPropertyShrinkRequest) SetPropertyShrink(v string) *UpdateSchemaPropertyShrinkRequest {
	s.PropertyShrink = &v
	return s
}

func (s *UpdateSchemaPropertyShrinkRequest) SetRequestId(v string) *UpdateSchemaPropertyShrinkRequest {
	s.RequestId = &v
	return s
}

func (s *UpdateSchemaPropertyShrinkRequest) SetSchemaId(v string) *UpdateSchemaPropertyShrinkRequest {
	s.SchemaId = &v
	return s
}

func (s *UpdateSchemaPropertyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
