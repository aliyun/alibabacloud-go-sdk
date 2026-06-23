// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyResourceAccessPermissionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyContentsShrink(v string) *ApplyResourceAccessPermissionShrinkRequest
	GetApplyContentsShrink() *string
	SetClientToken(v string) *ApplyResourceAccessPermissionShrinkRequest
	GetClientToken() *string
	SetReason(v string) *ApplyResourceAccessPermissionShrinkRequest
	GetReason() *string
}

type ApplyResourceAccessPermissionShrinkRequest struct {
	// A list of permission requests.
	//
	// This parameter is required.
	ApplyContentsShrink *string `json:"ApplyContents,omitempty" xml:"ApplyContents,omitempty"`
	// The idempotency parameter, which prevents duplicate operations from repeated calls.
	//
	// example:
	//
	// ABFUOEUOTRTRJKE
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The reason for the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 业务发展需要
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ApplyResourceAccessPermissionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyResourceAccessPermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *ApplyResourceAccessPermissionShrinkRequest) GetApplyContentsShrink() *string {
	return s.ApplyContentsShrink
}

func (s *ApplyResourceAccessPermissionShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ApplyResourceAccessPermissionShrinkRequest) GetReason() *string {
	return s.Reason
}

func (s *ApplyResourceAccessPermissionShrinkRequest) SetApplyContentsShrink(v string) *ApplyResourceAccessPermissionShrinkRequest {
	s.ApplyContentsShrink = &v
	return s
}

func (s *ApplyResourceAccessPermissionShrinkRequest) SetClientToken(v string) *ApplyResourceAccessPermissionShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ApplyResourceAccessPermissionShrinkRequest) SetReason(v string) *ApplyResourceAccessPermissionShrinkRequest {
	s.Reason = &v
	return s
}

func (s *ApplyResourceAccessPermissionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
