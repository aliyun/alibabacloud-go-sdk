// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetSchemaRequest
	GetInstanceId() *string
	SetRequestId(v string) *GetSchemaRequest
	GetRequestId() *string
	SetSchemaId(v string) *GetSchemaRequest
	GetSchemaId() *string
}

type GetSchemaRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// b0eb2742-f37e-4c67-82d4-25c651c1xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

func (s GetSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSchemaRequest) GoString() string {
	return s.String()
}

func (s *GetSchemaRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSchemaRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSchemaRequest) GetSchemaId() *string {
	return s.SchemaId
}

func (s *GetSchemaRequest) SetInstanceId(v string) *GetSchemaRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSchemaRequest) SetRequestId(v string) *GetSchemaRequest {
	s.RequestId = &v
	return s
}

func (s *GetSchemaRequest) SetSchemaId(v string) *GetSchemaRequest {
	s.SchemaId = &v
	return s
}

func (s *GetSchemaRequest) Validate() error {
	return dara.Validate(s)
}
