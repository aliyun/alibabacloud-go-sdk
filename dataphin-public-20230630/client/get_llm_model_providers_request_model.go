// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLlmModelProvidersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetLlmModelProvidersRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetLlmModelProvidersRequest
	GetOpUserId() *string
}

type GetLlmModelProvidersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s GetLlmModelProvidersRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLlmModelProvidersRequest) GoString() string {
	return s.String()
}

func (s *GetLlmModelProvidersRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetLlmModelProvidersRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetLlmModelProvidersRequest) SetOpTenantId(v int64) *GetLlmModelProvidersRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetLlmModelProvidersRequest) SetOpUserId(v string) *GetLlmModelProvidersRequest {
	s.OpUserId = &v
	return s
}

func (s *GetLlmModelProvidersRequest) Validate() error {
	return dara.Validate(s)
}
