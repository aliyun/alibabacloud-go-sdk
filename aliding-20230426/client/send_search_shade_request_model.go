// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendSearchShadeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *SendSearchShadeRequest
	GetContent() map[string]interface{}
	SetEndTime(v int64) *SendSearchShadeRequest
	GetEndTime() *int64
	SetStartTime(v int64) *SendSearchShadeRequest
	GetStartTime() *int64
	SetTenantContext(v *SendSearchShadeRequestTenantContext) *SendSearchShadeRequest
	GetTenantContext() *SendSearchShadeRequestTenantContext
}

type SendSearchShadeRequest struct {
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1693881641000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1693881641000
	StartTime     *int64                               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantContext *SendSearchShadeRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s SendSearchShadeRequest) String() string {
	return dara.Prettify(s)
}

func (s SendSearchShadeRequest) GoString() string {
	return s.String()
}

func (s *SendSearchShadeRequest) GetContent() map[string]interface{} {
	return s.Content
}

func (s *SendSearchShadeRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *SendSearchShadeRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SendSearchShadeRequest) GetTenantContext() *SendSearchShadeRequestTenantContext {
	return s.TenantContext
}

func (s *SendSearchShadeRequest) SetContent(v map[string]interface{}) *SendSearchShadeRequest {
	s.Content = v
	return s
}

func (s *SendSearchShadeRequest) SetEndTime(v int64) *SendSearchShadeRequest {
	s.EndTime = &v
	return s
}

func (s *SendSearchShadeRequest) SetStartTime(v int64) *SendSearchShadeRequest {
	s.StartTime = &v
	return s
}

func (s *SendSearchShadeRequest) SetTenantContext(v *SendSearchShadeRequestTenantContext) *SendSearchShadeRequest {
	s.TenantContext = v
	return s
}

func (s *SendSearchShadeRequest) Validate() error {
	return dara.Validate(s)
}

type SendSearchShadeRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SendSearchShadeRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s SendSearchShadeRequestTenantContext) GoString() string {
	return s.String()
}

func (s *SendSearchShadeRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *SendSearchShadeRequestTenantContext) SetTenantId(v string) *SendSearchShadeRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *SendSearchShadeRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
