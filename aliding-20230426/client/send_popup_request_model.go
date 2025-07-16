// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendPopupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *SendPopupRequest
	GetContent() map[string]interface{}
	SetEndTime(v int64) *SendPopupRequest
	GetEndTime() *int64
	SetStartTime(v int64) *SendPopupRequest
	GetStartTime() *int64
	SetTenantContext(v *SendPopupRequestTenantContext) *SendPopupRequest
	GetTenantContext() *SendPopupRequestTenantContext
}

type SendPopupRequest struct {
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1693881641000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1693881641000
	StartTime     *int64                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantContext *SendPopupRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s SendPopupRequest) String() string {
	return dara.Prettify(s)
}

func (s SendPopupRequest) GoString() string {
	return s.String()
}

func (s *SendPopupRequest) GetContent() map[string]interface{} {
	return s.Content
}

func (s *SendPopupRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *SendPopupRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SendPopupRequest) GetTenantContext() *SendPopupRequestTenantContext {
	return s.TenantContext
}

func (s *SendPopupRequest) SetContent(v map[string]interface{}) *SendPopupRequest {
	s.Content = v
	return s
}

func (s *SendPopupRequest) SetEndTime(v int64) *SendPopupRequest {
	s.EndTime = &v
	return s
}

func (s *SendPopupRequest) SetStartTime(v int64) *SendPopupRequest {
	s.StartTime = &v
	return s
}

func (s *SendPopupRequest) SetTenantContext(v *SendPopupRequestTenantContext) *SendPopupRequest {
	s.TenantContext = v
	return s
}

func (s *SendPopupRequest) Validate() error {
	return dara.Validate(s)
}

type SendPopupRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SendPopupRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s SendPopupRequestTenantContext) GoString() string {
	return s.String()
}

func (s *SendPopupRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *SendPopupRequestTenantContext) SetTenantId(v string) *SendPopupRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *SendPopupRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
