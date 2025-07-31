// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUsersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetUsersShrinkRequest
	GetOpTenantId() *int64
	SetUserIdListShrink(v string) *GetUsersShrinkRequest
	GetUserIdListShrink() *string
}

type GetUsersShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UserIdListShrink *string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty"`
}

func (s GetUsersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetUsersShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetUsersShrinkRequest) GetUserIdListShrink() *string {
	return s.UserIdListShrink
}

func (s *GetUsersShrinkRequest) SetOpTenantId(v int64) *GetUsersShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetUsersShrinkRequest) SetUserIdListShrink(v string) *GetUsersShrinkRequest {
	s.UserIdListShrink = &v
	return s
}

func (s *GetUsersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
