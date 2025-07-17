// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSchemaResponseBody
	GetRequestId() *string
	SetSchema(v *Schema) *GetSchemaResponseBody
	GetSchema() *Schema
	SetSuccess(v bool) *GetSchemaResponseBody
	GetSuccess() *bool
}

type GetSchemaResponseBody struct {
	// example:
	//
	// A89B5D9D-74EA-XXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Schema    *Schema `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSchemaResponseBody) GetSchema() *Schema {
	return s.Schema
}

func (s *GetSchemaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSchemaResponseBody) SetRequestId(v string) *GetSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSchemaResponseBody) SetSchema(v *Schema) *GetSchemaResponseBody {
	s.Schema = v
	return s
}

func (s *GetSchemaResponseBody) SetSuccess(v bool) *GetSchemaResponseBody {
	s.Success = &v
	return s
}

func (s *GetSchemaResponseBody) Validate() error {
	return dara.Validate(s)
}
