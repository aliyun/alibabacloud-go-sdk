// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBizEntityInfoByVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetBizEntityInfoByVersionRequest
	GetId() *int64
	SetOpTenantId(v int64) *GetBizEntityInfoByVersionRequest
	GetOpTenantId() *int64
	SetType(v string) *GetBizEntityInfoByVersionRequest
	GetType() *string
	SetVersionId(v int64) *GetBizEntityInfoByVersionRequest
	GetVersionId() *int64
}

type GetBizEntityInfoByVersionRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetBizEntityInfoByVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBizEntityInfoByVersionRequest) GoString() string {
	return s.String()
}

func (s *GetBizEntityInfoByVersionRequest) GetId() *int64 {
	return s.Id
}

func (s *GetBizEntityInfoByVersionRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetBizEntityInfoByVersionRequest) GetType() *string {
	return s.Type
}

func (s *GetBizEntityInfoByVersionRequest) GetVersionId() *int64 {
	return s.VersionId
}

func (s *GetBizEntityInfoByVersionRequest) SetId(v int64) *GetBizEntityInfoByVersionRequest {
	s.Id = &v
	return s
}

func (s *GetBizEntityInfoByVersionRequest) SetOpTenantId(v int64) *GetBizEntityInfoByVersionRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetBizEntityInfoByVersionRequest) SetType(v string) *GetBizEntityInfoByVersionRequest {
	s.Type = &v
	return s
}

func (s *GetBizEntityInfoByVersionRequest) SetVersionId(v int64) *GetBizEntityInfoByVersionRequest {
	s.VersionId = &v
	return s
}

func (s *GetBizEntityInfoByVersionRequest) Validate() error {
	return dara.Validate(s)
}
