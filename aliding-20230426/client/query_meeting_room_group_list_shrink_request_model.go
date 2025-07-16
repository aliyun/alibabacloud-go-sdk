// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomGroupListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequestShrink(v string) *QueryMeetingRoomGroupListShrinkRequest
	GetRequestShrink() *string
	SetTenantContextShrink(v string) *QueryMeetingRoomGroupListShrinkRequest
	GetTenantContextShrink() *string
}

type QueryMeetingRoomGroupListShrinkRequest struct {
	RequestShrink       *string `json:"Request,omitempty" xml:"Request,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s QueryMeetingRoomGroupListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomGroupListShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupListShrinkRequest) GetRequestShrink() *string {
	return s.RequestShrink
}

func (s *QueryMeetingRoomGroupListShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryMeetingRoomGroupListShrinkRequest) SetRequestShrink(v string) *QueryMeetingRoomGroupListShrinkRequest {
	s.RequestShrink = &v
	return s
}

func (s *QueryMeetingRoomGroupListShrinkRequest) SetTenantContextShrink(v string) *QueryMeetingRoomGroupListShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryMeetingRoomGroupListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
