// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConferenceMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *QueryConferenceMembersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *QueryConferenceMembersRequest
	GetNextToken() *string
	SetTenantContext(v *QueryConferenceMembersRequestTenantContext) *QueryConferenceMembersRequest
	GetTenantContext() *QueryConferenceMembersRequestTenantContext
	SetConferenceId(v string) *QueryConferenceMembersRequest
	GetConferenceId() *string
}

type QueryConferenceMembersRequest struct {
	// example:
	//
	// 300
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 123000000
	NextToken     *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TenantContext *QueryConferenceMembersRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s QueryConferenceMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceMembersRequest) GoString() string {
	return s.String()
}

func (s *QueryConferenceMembersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryConferenceMembersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryConferenceMembersRequest) GetTenantContext() *QueryConferenceMembersRequestTenantContext {
	return s.TenantContext
}

func (s *QueryConferenceMembersRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *QueryConferenceMembersRequest) SetMaxResults(v int32) *QueryConferenceMembersRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryConferenceMembersRequest) SetNextToken(v string) *QueryConferenceMembersRequest {
	s.NextToken = &v
	return s
}

func (s *QueryConferenceMembersRequest) SetTenantContext(v *QueryConferenceMembersRequestTenantContext) *QueryConferenceMembersRequest {
	s.TenantContext = v
	return s
}

func (s *QueryConferenceMembersRequest) SetConferenceId(v string) *QueryConferenceMembersRequest {
	s.ConferenceId = &v
	return s
}

func (s *QueryConferenceMembersRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryConferenceMembersRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryConferenceMembersRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceMembersRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryConferenceMembersRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryConferenceMembersRequestTenantContext) SetTenantId(v string) *QueryConferenceMembersRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryConferenceMembersRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
