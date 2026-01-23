// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetQualityRuleRequest
	GetId() *int64
	SetOpTenantId(v int64) *GetQualityRuleRequest
	GetOpTenantId() *int64
}

type GetQualityRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetQualityRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleRequest) GoString() string {
	return s.String()
}

func (s *GetQualityRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *GetQualityRuleRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetQualityRuleRequest) SetId(v int64) *GetQualityRuleRequest {
	s.Id = &v
	return s
}

func (s *GetQualityRuleRequest) SetOpTenantId(v int64) *GetQualityRuleRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetQualityRuleRequest) Validate() error {
	return dara.Validate(s)
}
