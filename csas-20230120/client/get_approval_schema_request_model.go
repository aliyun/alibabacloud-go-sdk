// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApprovalSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSchemaId(v string) *GetApprovalSchemaRequest
	GetSchemaId() *string
}

type GetApprovalSchemaRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s GetApprovalSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalSchemaRequest) GoString() string {
	return s.String()
}

func (s *GetApprovalSchemaRequest) GetSchemaId() *string {
	return s.SchemaId
}

func (s *GetApprovalSchemaRequest) SetSchemaId(v string) *GetApprovalSchemaRequest {
	s.SchemaId = &v
	return s
}

func (s *GetApprovalSchemaRequest) Validate() error {
	return dara.Validate(s)
}
