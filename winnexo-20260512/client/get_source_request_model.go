// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeDetails(v bool) *GetSourceRequest
	GetIncludeDetails() *bool
	SetSourceId(v string) *GetSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *GetSourceRequest
	GetTenantId() *string
}

type GetSourceRequest struct {
	// Specifies whether to return large detail fields (settings / notes / structuredTables / unstructuredDocs). Default value: False, which returns only metadata.
	//
	// example:
	//
	// false
	IncludeDetails *bool `json:"includeDetails,omitempty" xml:"includeDetails,omitempty"`
	// The primary ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID to which the task belongs.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSourceRequest) GoString() string {
	return s.String()
}

func (s *GetSourceRequest) GetIncludeDetails() *bool {
	return s.IncludeDetails
}

func (s *GetSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *GetSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetSourceRequest) SetIncludeDetails(v bool) *GetSourceRequest {
	s.IncludeDetails = &v
	return s
}

func (s *GetSourceRequest) SetSourceId(v string) *GetSourceRequest {
	s.SourceId = &v
	return s
}

func (s *GetSourceRequest) SetTenantId(v string) *GetSourceRequest {
	s.TenantId = &v
	return s
}

func (s *GetSourceRequest) Validate() error {
	return dara.Validate(s)
}
