// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceResourceTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxcomputeSchema(v string) *ListInstanceResourceTablesRequest
	GetMaxcomputeSchema() *string
}

type ListInstanceResourceTablesRequest struct {
	// example:
	//
	// jackal
	MaxcomputeSchema *string `json:"MaxcomputeSchema,omitempty" xml:"MaxcomputeSchema,omitempty"`
}

func (s ListInstanceResourceTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResourceTablesRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceResourceTablesRequest) GetMaxcomputeSchema() *string {
	return s.MaxcomputeSchema
}

func (s *ListInstanceResourceTablesRequest) SetMaxcomputeSchema(v string) *ListInstanceResourceTablesRequest {
	s.MaxcomputeSchema = &v
	return s
}

func (s *ListInstanceResourceTablesRequest) Validate() error {
	return dara.Validate(s)
}
