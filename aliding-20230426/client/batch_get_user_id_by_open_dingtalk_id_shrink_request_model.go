// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetUserIdByOpenDingtalkIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *BatchGetUserIdByOpenDingtalkIdShrinkRequest
	GetTenantContextShrink() *string
	SetOpenDingtalkIdsShrink(v string) *BatchGetUserIdByOpenDingtalkIdShrinkRequest
	GetOpenDingtalkIdsShrink() *string
}

type BatchGetUserIdByOpenDingtalkIdShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["DTOJdYJ2IQC4HuexhtjsS8Qxxxx","D8301AKf6lmZbXiilcB4P2MVxxxx"]
	OpenDingtalkIdsShrink *string `json:"openDingtalkIds,omitempty" xml:"openDingtalkIds,omitempty"`
}

func (s BatchGetUserIdByOpenDingtalkIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetUserIdByOpenDingtalkIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchGetUserIdByOpenDingtalkIdShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *BatchGetUserIdByOpenDingtalkIdShrinkRequest) GetOpenDingtalkIdsShrink() *string {
	return s.OpenDingtalkIdsShrink
}

func (s *BatchGetUserIdByOpenDingtalkIdShrinkRequest) SetTenantContextShrink(v string) *BatchGetUserIdByOpenDingtalkIdShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *BatchGetUserIdByOpenDingtalkIdShrinkRequest) SetOpenDingtalkIdsShrink(v string) *BatchGetUserIdByOpenDingtalkIdShrinkRequest {
	s.OpenDingtalkIdsShrink = &v
	return s
}

func (s *BatchGetUserIdByOpenDingtalkIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
