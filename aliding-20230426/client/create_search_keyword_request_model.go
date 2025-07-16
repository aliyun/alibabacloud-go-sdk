// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSearchKeywordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *CreateSearchKeywordRequest
	GetContent() *string
	SetEndTime(v int64) *CreateSearchKeywordRequest
	GetEndTime() *int64
	SetResId(v string) *CreateSearchKeywordRequest
	GetResId() *string
	SetStartTime(v int64) *CreateSearchKeywordRequest
	GetStartTime() *int64
	SetTenantContext(v *CreateSearchKeywordRequestTenantContext) *CreateSearchKeywordRequest
	GetTenantContext() *CreateSearchKeywordRequestTenantContext
	SetUserIdList(v []*string) *CreateSearchKeywordRequest
	GetUserIdList() []*string
}

type CreateSearchKeywordRequest struct {
	// example:
	//
	// {}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1699265024987
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1028
	ResId *string `json:"ResId,omitempty" xml:"ResId,omitempty"`
	// example:
	//
	// 1699265024987
	StartTime     *int64                                   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantContext *CreateSearchKeywordRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	UserIdList    []*string                                `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
}

func (s CreateSearchKeywordRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchKeywordRequest) GoString() string {
	return s.String()
}

func (s *CreateSearchKeywordRequest) GetContent() *string {
	return s.Content
}

func (s *CreateSearchKeywordRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateSearchKeywordRequest) GetResId() *string {
	return s.ResId
}

func (s *CreateSearchKeywordRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateSearchKeywordRequest) GetTenantContext() *CreateSearchKeywordRequestTenantContext {
	return s.TenantContext
}

func (s *CreateSearchKeywordRequest) GetUserIdList() []*string {
	return s.UserIdList
}

func (s *CreateSearchKeywordRequest) SetContent(v string) *CreateSearchKeywordRequest {
	s.Content = &v
	return s
}

func (s *CreateSearchKeywordRequest) SetEndTime(v int64) *CreateSearchKeywordRequest {
	s.EndTime = &v
	return s
}

func (s *CreateSearchKeywordRequest) SetResId(v string) *CreateSearchKeywordRequest {
	s.ResId = &v
	return s
}

func (s *CreateSearchKeywordRequest) SetStartTime(v int64) *CreateSearchKeywordRequest {
	s.StartTime = &v
	return s
}

func (s *CreateSearchKeywordRequest) SetTenantContext(v *CreateSearchKeywordRequestTenantContext) *CreateSearchKeywordRequest {
	s.TenantContext = v
	return s
}

func (s *CreateSearchKeywordRequest) SetUserIdList(v []*string) *CreateSearchKeywordRequest {
	s.UserIdList = v
	return s
}

func (s *CreateSearchKeywordRequest) Validate() error {
	return dara.Validate(s)
}

type CreateSearchKeywordRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateSearchKeywordRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchKeywordRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreateSearchKeywordRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateSearchKeywordRequestTenantContext) SetTenantId(v string) *CreateSearchKeywordRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CreateSearchKeywordRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
