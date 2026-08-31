// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLlmModelProviderByNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetLlmModelProviderByNameRequest
	GetName() *string
	SetOpTenantId(v int64) *GetLlmModelProviderByNameRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetLlmModelProviderByNameRequest
	GetOpUserId() *string
}

type GetLlmModelProviderByNameRequest struct {
	// The name of the large language model service provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// bailian
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s GetLlmModelProviderByNameRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLlmModelProviderByNameRequest) GoString() string {
	return s.String()
}

func (s *GetLlmModelProviderByNameRequest) GetName() *string {
	return s.Name
}

func (s *GetLlmModelProviderByNameRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetLlmModelProviderByNameRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetLlmModelProviderByNameRequest) SetName(v string) *GetLlmModelProviderByNameRequest {
	s.Name = &v
	return s
}

func (s *GetLlmModelProviderByNameRequest) SetOpTenantId(v int64) *GetLlmModelProviderByNameRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetLlmModelProviderByNameRequest) SetOpUserId(v string) *GetLlmModelProviderByNameRequest {
	s.OpUserId = &v
	return s
}

func (s *GetLlmModelProviderByNameRequest) Validate() error {
	return dara.Validate(s)
}
