// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryScheduleConferenceInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *QueryScheduleConferenceInfoShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *QueryScheduleConferenceInfoShrinkRequest
	GetNextToken() *string
	SetScheduleConferenceId(v string) *QueryScheduleConferenceInfoShrinkRequest
	GetScheduleConferenceId() *string
	SetTenantContextShrink(v string) *QueryScheduleConferenceInfoShrinkRequest
	GetTenantContextShrink() *string
}

type QueryScheduleConferenceInfoShrinkRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 19
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 636cf59f2b032f014ae32902
	ScheduleConferenceId *string `json:"ScheduleConferenceId,omitempty" xml:"ScheduleConferenceId,omitempty"`
	TenantContextShrink  *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s QueryScheduleConferenceInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryScheduleConferenceInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceInfoShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryScheduleConferenceInfoShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryScheduleConferenceInfoShrinkRequest) GetScheduleConferenceId() *string {
	return s.ScheduleConferenceId
}

func (s *QueryScheduleConferenceInfoShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryScheduleConferenceInfoShrinkRequest) SetMaxResults(v int32) *QueryScheduleConferenceInfoShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryScheduleConferenceInfoShrinkRequest) SetNextToken(v string) *QueryScheduleConferenceInfoShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *QueryScheduleConferenceInfoShrinkRequest) SetScheduleConferenceId(v string) *QueryScheduleConferenceInfoShrinkRequest {
	s.ScheduleConferenceId = &v
	return s
}

func (s *QueryScheduleConferenceInfoShrinkRequest) SetTenantContextShrink(v string) *QueryScheduleConferenceInfoShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryScheduleConferenceInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
