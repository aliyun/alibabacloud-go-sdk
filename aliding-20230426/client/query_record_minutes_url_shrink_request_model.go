// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordMinutesUrlShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *QueryRecordMinutesUrlShrinkRequest
	GetBizType() *string
	SetConferenceId(v string) *QueryRecordMinutesUrlShrinkRequest
	GetConferenceId() *string
	SetTenantContextShrink(v string) *QueryRecordMinutesUrlShrinkRequest
	GetTenantContextShrink() *string
}

type QueryRecordMinutesUrlShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// minutes
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1dddwrqrq
	ConferenceId        *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s QueryRecordMinutesUrlShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordMinutesUrlShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordMinutesUrlShrinkRequest) GetBizType() *string {
	return s.BizType
}

func (s *QueryRecordMinutesUrlShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *QueryRecordMinutesUrlShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryRecordMinutesUrlShrinkRequest) SetBizType(v string) *QueryRecordMinutesUrlShrinkRequest {
	s.BizType = &v
	return s
}

func (s *QueryRecordMinutesUrlShrinkRequest) SetConferenceId(v string) *QueryRecordMinutesUrlShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *QueryRecordMinutesUrlShrinkRequest) SetTenantContextShrink(v string) *QueryRecordMinutesUrlShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryRecordMinutesUrlShrinkRequest) Validate() error {
	return dara.Validate(s)
}
