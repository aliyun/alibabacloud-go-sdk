// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetConferenceHostsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperationType(v string) *SetConferenceHostsShrinkRequest
	GetOperationType() *string
	SetTenantContextShrink(v string) *SetConferenceHostsShrinkRequest
	GetTenantContextShrink() *string
	SetUserIdsShrink(v string) *SetConferenceHostsShrinkRequest
	GetUserIdsShrink() *string
	SetConferenceId(v string) *SetConferenceHostsShrinkRequest
	GetConferenceId() *string
}

type SetConferenceHostsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// add
	OperationType       *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [ "012345"]
	UserIdsShrink *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 607452e01401526ee39609e1
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s SetConferenceHostsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetConferenceHostsShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetConferenceHostsShrinkRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *SetConferenceHostsShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *SetConferenceHostsShrinkRequest) GetUserIdsShrink() *string {
	return s.UserIdsShrink
}

func (s *SetConferenceHostsShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *SetConferenceHostsShrinkRequest) SetOperationType(v string) *SetConferenceHostsShrinkRequest {
	s.OperationType = &v
	return s
}

func (s *SetConferenceHostsShrinkRequest) SetTenantContextShrink(v string) *SetConferenceHostsShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *SetConferenceHostsShrinkRequest) SetUserIdsShrink(v string) *SetConferenceHostsShrinkRequest {
	s.UserIdsShrink = &v
	return s
}

func (s *SetConferenceHostsShrinkRequest) SetConferenceId(v string) *SetConferenceHostsShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *SetConferenceHostsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
