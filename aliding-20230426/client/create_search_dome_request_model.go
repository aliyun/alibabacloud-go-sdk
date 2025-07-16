// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSearchDomeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *CreateSearchDomeRequest
	GetContent() *string
	SetEndTime(v int64) *CreateSearchDomeRequest
	GetEndTime() *int64
	SetResId(v string) *CreateSearchDomeRequest
	GetResId() *string
	SetStartTime(v int64) *CreateSearchDomeRequest
	GetStartTime() *int64
	SetTenantContext(v *CreateSearchDomeRequestTenantContext) *CreateSearchDomeRequest
	GetTenantContext() *CreateSearchDomeRequestTenantContext
	SetUserIdList(v []*string) *CreateSearchDomeRequest
	GetUserIdList() []*string
}

type CreateSearchDomeRequest struct {
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
	// 1030
	ResId *string `json:"ResId,omitempty" xml:"ResId,omitempty"`
	// example:
	//
	// 1699265024987
	StartTime     *int64                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantContext *CreateSearchDomeRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	UserIdList    []*string                             `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
}

func (s CreateSearchDomeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchDomeRequest) GoString() string {
	return s.String()
}

func (s *CreateSearchDomeRequest) GetContent() *string {
	return s.Content
}

func (s *CreateSearchDomeRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateSearchDomeRequest) GetResId() *string {
	return s.ResId
}

func (s *CreateSearchDomeRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateSearchDomeRequest) GetTenantContext() *CreateSearchDomeRequestTenantContext {
	return s.TenantContext
}

func (s *CreateSearchDomeRequest) GetUserIdList() []*string {
	return s.UserIdList
}

func (s *CreateSearchDomeRequest) SetContent(v string) *CreateSearchDomeRequest {
	s.Content = &v
	return s
}

func (s *CreateSearchDomeRequest) SetEndTime(v int64) *CreateSearchDomeRequest {
	s.EndTime = &v
	return s
}

func (s *CreateSearchDomeRequest) SetResId(v string) *CreateSearchDomeRequest {
	s.ResId = &v
	return s
}

func (s *CreateSearchDomeRequest) SetStartTime(v int64) *CreateSearchDomeRequest {
	s.StartTime = &v
	return s
}

func (s *CreateSearchDomeRequest) SetTenantContext(v *CreateSearchDomeRequestTenantContext) *CreateSearchDomeRequest {
	s.TenantContext = v
	return s
}

func (s *CreateSearchDomeRequest) SetUserIdList(v []*string) *CreateSearchDomeRequest {
	s.UserIdList = v
	return s
}

func (s *CreateSearchDomeRequest) Validate() error {
	return dara.Validate(s)
}

type CreateSearchDomeRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateSearchDomeRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchDomeRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreateSearchDomeRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateSearchDomeRequestTenantContext) SetTenantId(v string) *CreateSearchDomeRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CreateSearchDomeRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
