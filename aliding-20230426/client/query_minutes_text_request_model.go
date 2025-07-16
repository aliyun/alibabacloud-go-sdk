// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMinutesTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *QueryMinutesTextRequestTenantContext) *QueryMinutesTextRequest
	GetTenantContext() *QueryMinutesTextRequestTenantContext
	SetConferenceId(v string) *QueryMinutesTextRequest
	GetConferenceId() *string
	SetDirection(v string) *QueryMinutesTextRequest
	GetDirection() *string
	SetMaxResults(v int64) *QueryMinutesTextRequest
	GetMaxResults() *int64
	SetNextToken(v string) *QueryMinutesTextRequest
	GetNextToken() *string
}

type QueryMinutesTextRequest struct {
	TenantContext *QueryMinutesTextRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
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

func (s QueryMinutesTextRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesTextRequest) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextRequest) GetTenantContext() *QueryMinutesTextRequestTenantContext {
	return s.TenantContext
}

func (s *QueryMinutesTextRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *QueryMinutesTextRequest) GetDirection() *string {
	return s.Direction
}

func (s *QueryMinutesTextRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *QueryMinutesTextRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryMinutesTextRequest) SetTenantContext(v *QueryMinutesTextRequestTenantContext) *QueryMinutesTextRequest {
	s.TenantContext = v
	return s
}

func (s *QueryMinutesTextRequest) SetConferenceId(v string) *QueryMinutesTextRequest {
	s.ConferenceId = &v
	return s
}

func (s *QueryMinutesTextRequest) SetDirection(v string) *QueryMinutesTextRequest {
	s.Direction = &v
	return s
}

func (s *QueryMinutesTextRequest) SetMaxResults(v int64) *QueryMinutesTextRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryMinutesTextRequest) SetNextToken(v string) *QueryMinutesTextRequest {
	s.NextToken = &v
	return s
}

func (s *QueryMinutesTextRequest) Validate() error {
	return dara.Validate(s)
}

type QueryMinutesTextRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryMinutesTextRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesTextRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryMinutesTextRequestTenantContext) SetTenantId(v string) *QueryMinutesTextRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryMinutesTextRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
