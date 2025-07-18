// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApprovalProcessesForApprovalSchemasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSchemaIds(v []*string) *ListApprovalProcessesForApprovalSchemasRequest
	GetSchemaIds() []*string
}

type ListApprovalProcessesForApprovalSchemasRequest struct {
	// This parameter is required.
	SchemaIds []*string `json:"SchemaIds,omitempty" xml:"SchemaIds,omitempty" type:"Repeated"`
}

func (s ListApprovalProcessesForApprovalSchemasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalProcessesForApprovalSchemasRequest) GoString() string {
	return s.String()
}

func (s *ListApprovalProcessesForApprovalSchemasRequest) GetSchemaIds() []*string {
	return s.SchemaIds
}

func (s *ListApprovalProcessesForApprovalSchemasRequest) SetSchemaIds(v []*string) *ListApprovalProcessesForApprovalSchemasRequest {
	s.SchemaIds = v
	return s
}

func (s *ListApprovalProcessesForApprovalSchemasRequest) Validate() error {
	return dara.Validate(s)
}
