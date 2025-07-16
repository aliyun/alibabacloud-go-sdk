// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupLiveListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetGroupLiveListRequest
	GetEndTime() *int64
	SetOpenConversationId(v string) *GetGroupLiveListRequest
	GetOpenConversationId() *string
	SetStartTime(v int64) *GetGroupLiveListRequest
	GetStartTime() *int64
	SetTenantContext(v *GetGroupLiveListRequestTenantContext) *GetGroupLiveListRequest
	GetTenantContext() *GetGroupLiveListRequestTenantContext
}

type GetGroupLiveListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1398324600000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cidt*****Xa4K10w==
	OpenConversationId *string `json:"OpenConversationId,omitempty" xml:"OpenConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1398324600000
	StartTime     *int64                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantContext *GetGroupLiveListRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetGroupLiveListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGroupLiveListRequest) GoString() string {
	return s.String()
}

func (s *GetGroupLiveListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetGroupLiveListRequest) GetOpenConversationId() *string {
	return s.OpenConversationId
}

func (s *GetGroupLiveListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetGroupLiveListRequest) GetTenantContext() *GetGroupLiveListRequestTenantContext {
	return s.TenantContext
}

func (s *GetGroupLiveListRequest) SetEndTime(v int64) *GetGroupLiveListRequest {
	s.EndTime = &v
	return s
}

func (s *GetGroupLiveListRequest) SetOpenConversationId(v string) *GetGroupLiveListRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GetGroupLiveListRequest) SetStartTime(v int64) *GetGroupLiveListRequest {
	s.StartTime = &v
	return s
}

func (s *GetGroupLiveListRequest) SetTenantContext(v *GetGroupLiveListRequestTenantContext) *GetGroupLiveListRequest {
	s.TenantContext = v
	return s
}

func (s *GetGroupLiveListRequest) Validate() error {
	return dara.Validate(s)
}

type GetGroupLiveListRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetGroupLiveListRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetGroupLiveListRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetGroupLiveListRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetGroupLiveListRequestTenantContext) SetTenantId(v string) *GetGroupLiveListRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetGroupLiveListRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
