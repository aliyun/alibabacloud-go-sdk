// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterQueryShrink(v string) *GetStandardTemplateShrinkRequest
	GetFilterQueryShrink() *string
	SetId(v int64) *GetStandardTemplateShrinkRequest
	GetId() *int64
	SetNullable(v bool) *GetStandardTemplateShrinkRequest
	GetNullable() *bool
	SetOpTenantId(v int64) *GetStandardTemplateShrinkRequest
	GetOpTenantId() *int64
}

type GetStandardTemplateShrinkRequest struct {
	// The filter condition.
	FilterQueryShrink *string `json:"FilterQuery,omitempty" xml:"FilterQuery,omitempty"`
	// The standard template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Specifies whether to allow a null value to be returned when the template does not exist. If set to false, an exception is thrown. Default value: true.
	Nullable *bool `json:"Nullable,omitempty" xml:"Nullable,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetStandardTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateShrinkRequest) GetFilterQueryShrink() *string {
	return s.FilterQueryShrink
}

func (s *GetStandardTemplateShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *GetStandardTemplateShrinkRequest) GetNullable() *bool {
	return s.Nullable
}

func (s *GetStandardTemplateShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetStandardTemplateShrinkRequest) SetFilterQueryShrink(v string) *GetStandardTemplateShrinkRequest {
	s.FilterQueryShrink = &v
	return s
}

func (s *GetStandardTemplateShrinkRequest) SetId(v int64) *GetStandardTemplateShrinkRequest {
	s.Id = &v
	return s
}

func (s *GetStandardTemplateShrinkRequest) SetNullable(v bool) *GetStandardTemplateShrinkRequest {
	s.Nullable = &v
	return s
}

func (s *GetStandardTemplateShrinkRequest) SetOpTenantId(v int64) *GetStandardTemplateShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetStandardTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
