// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteSchemaRequest
	GetInstanceId() *string
	SetRequestId(v string) *DeleteSchemaRequest
	GetRequestId() *string
	SetSchemaId(v string) *DeleteSchemaRequest
	GetSchemaId() *string
}

type DeleteSchemaRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 9cfad875-6260-4a53-ab6e-b13e3fb31f7d
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 0630E5DF-CEB0-445B-8626-D5C7481181C3
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

func (s DeleteSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchemaRequest) GoString() string {
	return s.String()
}

func (s *DeleteSchemaRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteSchemaRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSchemaRequest) GetSchemaId() *string {
	return s.SchemaId
}

func (s *DeleteSchemaRequest) SetInstanceId(v string) *DeleteSchemaRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSchemaRequest) SetRequestId(v string) *DeleteSchemaRequest {
	s.RequestId = &v
	return s
}

func (s *DeleteSchemaRequest) SetSchemaId(v string) *DeleteSchemaRequest {
	s.SchemaId = &v
	return s
}

func (s *DeleteSchemaRequest) Validate() error {
	return dara.Validate(s)
}
