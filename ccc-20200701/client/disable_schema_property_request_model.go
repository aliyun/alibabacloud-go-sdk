// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSchemaPropertyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DisableSchemaPropertyRequest
	GetInstanceId() *string
	SetPropertyName(v string) *DisableSchemaPropertyRequest
	GetPropertyName() *string
	SetRequestId(v string) *DisableSchemaPropertyRequest
	GetRequestId() *string
	SetSchemaId(v string) *DisableSchemaPropertyRequest
	GetSchemaId() *string
}

type DisableSchemaPropertyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ef1e71e9-ae9d-487c-96ad-9181d85cf802
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// name
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	// example:
	//
	// 2263B273-AC1B-44EB-BA98-87F2322C6780
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

func (s DisableSchemaPropertyRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableSchemaPropertyRequest) GoString() string {
	return s.String()
}

func (s *DisableSchemaPropertyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableSchemaPropertyRequest) GetPropertyName() *string {
	return s.PropertyName
}

func (s *DisableSchemaPropertyRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableSchemaPropertyRequest) GetSchemaId() *string {
	return s.SchemaId
}

func (s *DisableSchemaPropertyRequest) SetInstanceId(v string) *DisableSchemaPropertyRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableSchemaPropertyRequest) SetPropertyName(v string) *DisableSchemaPropertyRequest {
	s.PropertyName = &v
	return s
}

func (s *DisableSchemaPropertyRequest) SetRequestId(v string) *DisableSchemaPropertyRequest {
	s.RequestId = &v
	return s
}

func (s *DisableSchemaPropertyRequest) SetSchemaId(v string) *DisableSchemaPropertyRequest {
	s.SchemaId = &v
	return s
}

func (s *DisableSchemaPropertyRequest) Validate() error {
	return dara.Validate(s)
}
