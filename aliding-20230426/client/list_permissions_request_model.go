// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryUuid(v string) *ListPermissionsRequest
	GetDentryUuid() *string
	SetOption(v *ListPermissionsRequestOption) *ListPermissionsRequest
	GetOption() *ListPermissionsRequestOption
	SetTenantContext(v *ListPermissionsRequestTenantContext) *ListPermissionsRequest
	GetTenantContext() *ListPermissionsRequestTenantContext
}

type ListPermissionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// KGZLxjv9VGkoG9YwHE5wx7k2V6EDybno
	DentryUuid    *string                              `json:"DentryUuid,omitempty" xml:"DentryUuid,omitempty"`
	Option        *ListPermissionsRequestOption        `json:"Option,omitempty" xml:"Option,omitempty" type:"Struct"`
	TenantContext *ListPermissionsRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s ListPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListPermissionsRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *ListPermissionsRequest) GetOption() *ListPermissionsRequestOption {
	return s.Option
}

func (s *ListPermissionsRequest) GetTenantContext() *ListPermissionsRequestTenantContext {
	return s.TenantContext
}

func (s *ListPermissionsRequest) SetDentryUuid(v string) *ListPermissionsRequest {
	s.DentryUuid = &v
	return s
}

func (s *ListPermissionsRequest) SetOption(v *ListPermissionsRequestOption) *ListPermissionsRequest {
	s.Option = v
	return s
}

func (s *ListPermissionsRequest) SetTenantContext(v *ListPermissionsRequestTenantContext) *ListPermissionsRequest {
	s.TenantContext = v
	return s
}

func (s *ListPermissionsRequest) Validate() error {
	if s.Option != nil {
		if err := s.Option.Validate(); err != nil {
			return err
		}
	}
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPermissionsRequestOption struct {
	FilterRoleIds []*string `json:"FilterRoleIds,omitempty" xml:"FilterRoleIds,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListPermissionsRequestOption) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionsRequestOption) GoString() string {
	return s.String()
}

func (s *ListPermissionsRequestOption) GetFilterRoleIds() []*string {
	return s.FilterRoleIds
}

func (s *ListPermissionsRequestOption) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPermissionsRequestOption) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPermissionsRequestOption) SetFilterRoleIds(v []*string) *ListPermissionsRequestOption {
	s.FilterRoleIds = v
	return s
}

func (s *ListPermissionsRequestOption) SetMaxResults(v int32) *ListPermissionsRequestOption {
	s.MaxResults = &v
	return s
}

func (s *ListPermissionsRequestOption) SetNextToken(v string) *ListPermissionsRequestOption {
	s.NextToken = &v
	return s
}

func (s *ListPermissionsRequestOption) Validate() error {
	return dara.Validate(s)
}

type ListPermissionsRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListPermissionsRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionsRequestTenantContext) GoString() string {
	return s.String()
}

func (s *ListPermissionsRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *ListPermissionsRequestTenantContext) SetTenantId(v string) *ListPermissionsRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *ListPermissionsRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
