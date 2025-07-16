// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordMinutesUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *QueryRecordMinutesUrlRequest
	GetBizType() *string
	SetConferenceId(v string) *QueryRecordMinutesUrlRequest
	GetConferenceId() *string
	SetTenantContext(v *QueryRecordMinutesUrlRequestTenantContext) *QueryRecordMinutesUrlRequest
	GetTenantContext() *QueryRecordMinutesUrlRequestTenantContext
}

type QueryRecordMinutesUrlRequest struct {
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
	ConferenceId  *string                                    `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	TenantContext *QueryRecordMinutesUrlRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s QueryRecordMinutesUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordMinutesUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordMinutesUrlRequest) GetBizType() *string {
	return s.BizType
}

func (s *QueryRecordMinutesUrlRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *QueryRecordMinutesUrlRequest) GetTenantContext() *QueryRecordMinutesUrlRequestTenantContext {
	return s.TenantContext
}

func (s *QueryRecordMinutesUrlRequest) SetBizType(v string) *QueryRecordMinutesUrlRequest {
	s.BizType = &v
	return s
}

func (s *QueryRecordMinutesUrlRequest) SetConferenceId(v string) *QueryRecordMinutesUrlRequest {
	s.ConferenceId = &v
	return s
}

func (s *QueryRecordMinutesUrlRequest) SetTenantContext(v *QueryRecordMinutesUrlRequestTenantContext) *QueryRecordMinutesUrlRequest {
	s.TenantContext = v
	return s
}

func (s *QueryRecordMinutesUrlRequest) Validate() error {
	return dara.Validate(s)
}

type QueryRecordMinutesUrlRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryRecordMinutesUrlRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordMinutesUrlRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryRecordMinutesUrlRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryRecordMinutesUrlRequestTenantContext) SetTenantId(v string) *QueryRecordMinutesUrlRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryRecordMinutesUrlRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
