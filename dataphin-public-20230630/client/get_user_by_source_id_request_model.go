// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserBySourceIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetUserBySourceIdRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetUserBySourceIdRequest
	GetOpUserId() *string
	SetSourceId(v string) *GetUserBySourceIdRequest
	GetSourceId() *string
	SetSourceType(v string) *GetUserBySourceIdRequest
	GetSourceType() *string
}

type GetUserBySourceIdRequest struct {
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
	// The user source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 323131
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The user source type.
	//
	// example:
	//
	// aliyun
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s GetUserBySourceIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserBySourceIdRequest) GoString() string {
	return s.String()
}

func (s *GetUserBySourceIdRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetUserBySourceIdRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetUserBySourceIdRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *GetUserBySourceIdRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *GetUserBySourceIdRequest) SetOpTenantId(v int64) *GetUserBySourceIdRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetUserBySourceIdRequest) SetOpUserId(v string) *GetUserBySourceIdRequest {
	s.OpUserId = &v
	return s
}

func (s *GetUserBySourceIdRequest) SetSourceId(v string) *GetUserBySourceIdRequest {
	s.SourceId = &v
	return s
}

func (s *GetUserBySourceIdRequest) SetSourceType(v string) *GetUserBySourceIdRequest {
	s.SourceType = &v
	return s
}

func (s *GetUserBySourceIdRequest) Validate() error {
	return dara.Validate(s)
}
