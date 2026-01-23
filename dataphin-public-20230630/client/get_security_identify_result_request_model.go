// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityIdentifyResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetSecurityIdentifyResultRequest
	GetId() *int64
	SetOpTenantId(v int64) *GetSecurityIdentifyResultRequest
	GetOpTenantId() *int64
}

type GetSecurityIdentifyResultRequest struct {
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

func (s GetSecurityIdentifyResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityIdentifyResultRequest) GoString() string {
	return s.String()
}

func (s *GetSecurityIdentifyResultRequest) GetId() *int64 {
	return s.Id
}

func (s *GetSecurityIdentifyResultRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetSecurityIdentifyResultRequest) SetId(v int64) *GetSecurityIdentifyResultRequest {
	s.Id = &v
	return s
}

func (s *GetSecurityIdentifyResultRequest) SetOpTenantId(v int64) *GetSecurityIdentifyResultRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetSecurityIdentifyResultRequest) Validate() error {
	return dara.Validate(s)
}
