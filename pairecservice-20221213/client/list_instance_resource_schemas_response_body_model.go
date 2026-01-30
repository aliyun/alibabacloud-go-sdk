// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceResourceSchemasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListInstanceResourceSchemasResponseBody
	GetRequestId() *string
	SetSchemas(v []*ListInstanceResourceSchemasResponseBodySchemas) *ListInstanceResourceSchemasResponseBody
	GetSchemas() []*ListInstanceResourceSchemasResponseBodySchemas
	SetTotalCount(v int64) *ListInstanceResourceSchemasResponseBody
	GetTotalCount() *int64
}

type ListInstanceResourceSchemasResponseBody struct {
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Schemas   []*ListInstanceResourceSchemasResponseBodySchemas `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstanceResourceSchemasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResourceSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceResourceSchemasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceResourceSchemasResponseBody) GetSchemas() []*ListInstanceResourceSchemasResponseBodySchemas {
	return s.Schemas
}

func (s *ListInstanceResourceSchemasResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListInstanceResourceSchemasResponseBody) SetRequestId(v string) *ListInstanceResourceSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceResourceSchemasResponseBody) SetSchemas(v []*ListInstanceResourceSchemasResponseBodySchemas) *ListInstanceResourceSchemasResponseBody {
	s.Schemas = v
	return s
}

func (s *ListInstanceResourceSchemasResponseBody) SetTotalCount(v int64) *ListInstanceResourceSchemasResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstanceResourceSchemasResponseBody) Validate() error {
	if s.Schemas != nil {
		for _, item := range s.Schemas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceResourceSchemasResponseBodySchemas struct {
	// example:
	//
	// default
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s ListInstanceResourceSchemasResponseBodySchemas) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResourceSchemasResponseBodySchemas) GoString() string {
	return s.String()
}

func (s *ListInstanceResourceSchemasResponseBodySchemas) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListInstanceResourceSchemasResponseBodySchemas) SetSchemaName(v string) *ListInstanceResourceSchemasResponseBodySchemas {
	s.SchemaName = &v
	return s
}

func (s *ListInstanceResourceSchemasResponseBodySchemas) Validate() error {
	return dara.Validate(s)
}
