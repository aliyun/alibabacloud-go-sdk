// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMinutesSummaryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *QueryMinutesSummaryShrinkRequest
	GetTenantContextShrink() *string
	SetConferenceId(v string) *QueryMinutesSummaryShrinkRequest
	GetConferenceId() *string
	SetSummaryTypeListShrink(v string) *QueryMinutesSummaryShrinkRequest
	GetSummaryTypeListShrink() *string
}

type QueryMinutesSummaryShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 607452e01401526ee39609e1
	ConferenceId          *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	SummaryTypeListShrink *string `json:"summaryTypeList,omitempty" xml:"summaryTypeList,omitempty"`
}

func (s QueryMinutesSummaryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesSummaryShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryMinutesSummaryShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *QueryMinutesSummaryShrinkRequest) GetSummaryTypeListShrink() *string {
	return s.SummaryTypeListShrink
}

func (s *QueryMinutesSummaryShrinkRequest) SetTenantContextShrink(v string) *QueryMinutesSummaryShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryMinutesSummaryShrinkRequest) SetConferenceId(v string) *QueryMinutesSummaryShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *QueryMinutesSummaryShrinkRequest) SetSummaryTypeListShrink(v string) *QueryMinutesSummaryShrinkRequest {
	s.SummaryTypeListShrink = &v
	return s
}

func (s *QueryMinutesSummaryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
