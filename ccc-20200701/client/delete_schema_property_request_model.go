// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSchemaPropertyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteSchemaPropertyRequest
	GetInstanceId() *string
	SetPropertyName(v string) *DeleteSchemaPropertyRequest
	GetPropertyName() *string
	SetRequestId(v string) *DeleteSchemaPropertyRequest
	GetRequestId() *string
	SetSchemaId(v string) *DeleteSchemaPropertyRequest
	GetSchemaId() *string
}

type DeleteSchemaPropertyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 9cfad875-6260-4a53-ab6e-b13e3fb31f7d
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// name
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	// example:
	//
	// 7BEEA660-A45A-45E3-98CC-AFC65E715C23
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

func (s DeleteSchemaPropertyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchemaPropertyRequest) GoString() string {
	return s.String()
}

func (s *DeleteSchemaPropertyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteSchemaPropertyRequest) GetPropertyName() *string {
	return s.PropertyName
}

func (s *DeleteSchemaPropertyRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSchemaPropertyRequest) GetSchemaId() *string {
	return s.SchemaId
}

func (s *DeleteSchemaPropertyRequest) SetInstanceId(v string) *DeleteSchemaPropertyRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSchemaPropertyRequest) SetPropertyName(v string) *DeleteSchemaPropertyRequest {
	s.PropertyName = &v
	return s
}

func (s *DeleteSchemaPropertyRequest) SetRequestId(v string) *DeleteSchemaPropertyRequest {
	s.RequestId = &v
	return s
}

func (s *DeleteSchemaPropertyRequest) SetSchemaId(v string) *DeleteSchemaPropertyRequest {
	s.SchemaId = &v
	return s
}

func (s *DeleteSchemaPropertyRequest) Validate() error {
	return dara.Validate(s)
}
