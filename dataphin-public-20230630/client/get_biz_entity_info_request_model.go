// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBizEntityInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetBizEntityInfoRequest
	GetId() *int64
	SetOpTenantId(v int64) *GetBizEntityInfoRequest
	GetOpTenantId() *int64
	SetType(v string) *GetBizEntityInfoRequest
	GetType() *string
}

type GetBizEntityInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 101001201
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BIZ_OBJECT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetBizEntityInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBizEntityInfoRequest) GoString() string {
	return s.String()
}

func (s *GetBizEntityInfoRequest) GetId() *int64 {
	return s.Id
}

func (s *GetBizEntityInfoRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetBizEntityInfoRequest) GetType() *string {
	return s.Type
}

func (s *GetBizEntityInfoRequest) SetId(v int64) *GetBizEntityInfoRequest {
	s.Id = &v
	return s
}

func (s *GetBizEntityInfoRequest) SetOpTenantId(v int64) *GetBizEntityInfoRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetBizEntityInfoRequest) SetType(v string) *GetBizEntityInfoRequest {
	s.Type = &v
	return s
}

func (s *GetBizEntityInfoRequest) Validate() error {
	return dara.Validate(s)
}
