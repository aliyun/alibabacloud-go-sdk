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
	SetSourceId(v string) *GetUserBySourceIdRequest
	GetSourceId() *string
}

type GetUserBySourceIdRequest struct {
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
	// 323131
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
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

func (s *GetUserBySourceIdRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *GetUserBySourceIdRequest) SetOpTenantId(v int64) *GetUserBySourceIdRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetUserBySourceIdRequest) SetSourceId(v string) *GetUserBySourceIdRequest {
	s.SourceId = &v
	return s
}

func (s *GetUserBySourceIdRequest) Validate() error {
	return dara.Validate(s)
}
