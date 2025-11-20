// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMinutesSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *QueryMinutesSummaryRequestTenantContext) *QueryMinutesSummaryRequest
	GetTenantContext() *QueryMinutesSummaryRequestTenantContext
	SetConferenceId(v string) *QueryMinutesSummaryRequest
	GetConferenceId() *string
	SetSummaryTypeList(v []*string) *QueryMinutesSummaryRequest
	GetSummaryTypeList() []*string
}

type QueryMinutesSummaryRequest struct {
	TenantContext *QueryMinutesSummaryRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 607452e01401526ee39609e1
	ConferenceId    *string   `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	SummaryTypeList []*string `json:"summaryTypeList,omitempty" xml:"summaryTypeList,omitempty" type:"Repeated"`
}

func (s QueryMinutesSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesSummaryRequest) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryRequest) GetTenantContext() *QueryMinutesSummaryRequestTenantContext {
	return s.TenantContext
}

func (s *QueryMinutesSummaryRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *QueryMinutesSummaryRequest) GetSummaryTypeList() []*string {
	return s.SummaryTypeList
}

func (s *QueryMinutesSummaryRequest) SetTenantContext(v *QueryMinutesSummaryRequestTenantContext) *QueryMinutesSummaryRequest {
	s.TenantContext = v
	return s
}

func (s *QueryMinutesSummaryRequest) SetConferenceId(v string) *QueryMinutesSummaryRequest {
	s.ConferenceId = &v
	return s
}

func (s *QueryMinutesSummaryRequest) SetSummaryTypeList(v []*string) *QueryMinutesSummaryRequest {
	s.SummaryTypeList = v
	return s
}

func (s *QueryMinutesSummaryRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMinutesSummaryRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryMinutesSummaryRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesSummaryRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryMinutesSummaryRequestTenantContext) SetTenantId(v string) *QueryMinutesSummaryRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryMinutesSummaryRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
