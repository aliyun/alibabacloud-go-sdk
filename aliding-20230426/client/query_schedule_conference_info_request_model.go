// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryScheduleConferenceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *QueryScheduleConferenceInfoRequest
	GetMaxResults() *int32
	SetNextToken(v string) *QueryScheduleConferenceInfoRequest
	GetNextToken() *string
	SetScheduleConferenceId(v string) *QueryScheduleConferenceInfoRequest
	GetScheduleConferenceId() *string
	SetTenantContext(v *QueryScheduleConferenceInfoRequestTenantContext) *QueryScheduleConferenceInfoRequest
	GetTenantContext() *QueryScheduleConferenceInfoRequestTenantContext
}

type QueryScheduleConferenceInfoRequest struct {
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
	ScheduleConferenceId *string                                          `json:"ScheduleConferenceId,omitempty" xml:"ScheduleConferenceId,omitempty"`
	TenantContext        *QueryScheduleConferenceInfoRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s QueryScheduleConferenceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryScheduleConferenceInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceInfoRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryScheduleConferenceInfoRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryScheduleConferenceInfoRequest) GetScheduleConferenceId() *string {
	return s.ScheduleConferenceId
}

func (s *QueryScheduleConferenceInfoRequest) GetTenantContext() *QueryScheduleConferenceInfoRequestTenantContext {
	return s.TenantContext
}

func (s *QueryScheduleConferenceInfoRequest) SetMaxResults(v int32) *QueryScheduleConferenceInfoRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryScheduleConferenceInfoRequest) SetNextToken(v string) *QueryScheduleConferenceInfoRequest {
	s.NextToken = &v
	return s
}

func (s *QueryScheduleConferenceInfoRequest) SetScheduleConferenceId(v string) *QueryScheduleConferenceInfoRequest {
	s.ScheduleConferenceId = &v
	return s
}

func (s *QueryScheduleConferenceInfoRequest) SetTenantContext(v *QueryScheduleConferenceInfoRequestTenantContext) *QueryScheduleConferenceInfoRequest {
	s.TenantContext = v
	return s
}

func (s *QueryScheduleConferenceInfoRequest) Validate() error {
	return dara.Validate(s)
}

type QueryScheduleConferenceInfoRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryScheduleConferenceInfoRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryScheduleConferenceInfoRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceInfoRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryScheduleConferenceInfoRequestTenantContext) SetTenantId(v string) *QueryScheduleConferenceInfoRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryScheduleConferenceInfoRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
