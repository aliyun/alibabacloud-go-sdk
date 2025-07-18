// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApprovalSchemasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListApprovalSchemasRequest
	GetCurrentPage() *int64
	SetPageSize(v int64) *ListApprovalSchemasRequest
	GetPageSize() *int64
	SetPolicyType(v string) *ListApprovalSchemasRequest
	GetPolicyType() *string
	SetSchemaIds(v []*string) *ListApprovalSchemasRequest
	GetSchemaIds() []*string
	SetSchemaName(v string) *ListApprovalSchemasRequest
	GetSchemaName() *string
}

type ListApprovalSchemasRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// DlpSend
	PolicyType *string   `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	SchemaIds  []*string `json:"SchemaIds,omitempty" xml:"SchemaIds,omitempty" type:"Repeated"`
	// example:
	//
	// test
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s ListApprovalSchemasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalSchemasRequest) GoString() string {
	return s.String()
}

func (s *ListApprovalSchemasRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListApprovalSchemasRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListApprovalSchemasRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListApprovalSchemasRequest) GetSchemaIds() []*string {
	return s.SchemaIds
}

func (s *ListApprovalSchemasRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListApprovalSchemasRequest) SetCurrentPage(v int64) *ListApprovalSchemasRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListApprovalSchemasRequest) SetPageSize(v int64) *ListApprovalSchemasRequest {
	s.PageSize = &v
	return s
}

func (s *ListApprovalSchemasRequest) SetPolicyType(v string) *ListApprovalSchemasRequest {
	s.PolicyType = &v
	return s
}

func (s *ListApprovalSchemasRequest) SetSchemaIds(v []*string) *ListApprovalSchemasRequest {
	s.SchemaIds = v
	return s
}

func (s *ListApprovalSchemasRequest) SetSchemaName(v string) *ListApprovalSchemasRequest {
	s.SchemaName = &v
	return s
}

func (s *ListApprovalSchemasRequest) Validate() error {
	return dara.Validate(s)
}
