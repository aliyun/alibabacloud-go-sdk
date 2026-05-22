// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByOpenDingtalkIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *GetUserIdByOpenDingtalkIdShrinkRequest
	GetTenantContextShrink() *string
	SetOpenDingtalkId(v string) *GetUserIdByOpenDingtalkIdShrinkRequest
	GetOpenDingtalkId() *string
}

type GetUserIdByOpenDingtalkIdShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DTOJdYJ2IQC4HuexhtjsSXXXX
	OpenDingtalkId *string `json:"openDingtalkId,omitempty" xml:"openDingtalkId,omitempty"`
}

func (s GetUserIdByOpenDingtalkIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByOpenDingtalkIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetUserIdByOpenDingtalkIdShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetUserIdByOpenDingtalkIdShrinkRequest) GetOpenDingtalkId() *string {
	return s.OpenDingtalkId
}

func (s *GetUserIdByOpenDingtalkIdShrinkRequest) SetTenantContextShrink(v string) *GetUserIdByOpenDingtalkIdShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetUserIdByOpenDingtalkIdShrinkRequest) SetOpenDingtalkId(v string) *GetUserIdByOpenDingtalkIdShrinkRequest {
	s.OpenDingtalkId = &v
	return s
}

func (s *GetUserIdByOpenDingtalkIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
