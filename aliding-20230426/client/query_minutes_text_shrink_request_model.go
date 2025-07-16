// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMinutesTextShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *QueryMinutesTextShrinkRequest
	GetTenantContextShrink() *string
	SetConferenceId(v string) *QueryMinutesTextShrinkRequest
	GetConferenceId() *string
	SetDirection(v string) *QueryMinutesTextShrinkRequest
	GetDirection() *string
	SetMaxResults(v int64) *QueryMinutesTextShrinkRequest
	GetMaxResults() *int64
	SetNextToken(v string) *QueryMinutesTextShrinkRequest
	GetNextToken() *string
}

type QueryMinutesTextShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 607452e01401526ee39609e1
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryMinutesTextShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesTextShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryMinutesTextShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *QueryMinutesTextShrinkRequest) GetDirection() *string {
	return s.Direction
}

func (s *QueryMinutesTextShrinkRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *QueryMinutesTextShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryMinutesTextShrinkRequest) SetTenantContextShrink(v string) *QueryMinutesTextShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryMinutesTextShrinkRequest) SetConferenceId(v string) *QueryMinutesTextShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *QueryMinutesTextShrinkRequest) SetDirection(v string) *QueryMinutesTextShrinkRequest {
	s.Direction = &v
	return s
}

func (s *QueryMinutesTextShrinkRequest) SetMaxResults(v int64) *QueryMinutesTextShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryMinutesTextShrinkRequest) SetNextToken(v string) *QueryMinutesTextShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *QueryMinutesTextShrinkRequest) Validate() error {
	return dara.Validate(s)
}
