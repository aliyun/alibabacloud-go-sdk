// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetQualityTemplateRequest
	GetId() *int64
	SetOpTenantId(v int64) *GetQualityTemplateRequest
	GetOpTenantId() *int64
}

type GetQualityTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetQualityTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetQualityTemplateRequest) GetId() *int64 {
	return s.Id
}

func (s *GetQualityTemplateRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetQualityTemplateRequest) SetId(v int64) *GetQualityTemplateRequest {
	s.Id = &v
	return s
}

func (s *GetQualityTemplateRequest) SetOpTenantId(v int64) *GetQualityTemplateRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetQualityTemplateRequest) Validate() error {
	return dara.Validate(s)
}
