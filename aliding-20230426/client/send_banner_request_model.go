// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendBannerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *SendBannerRequest
	GetContent() map[string]interface{}
	SetEndTime(v int64) *SendBannerRequest
	GetEndTime() *int64
	SetStartTime(v int64) *SendBannerRequest
	GetStartTime() *int64
	SetTenantContext(v *SendBannerRequestTenantContext) *SendBannerRequest
	GetTenantContext() *SendBannerRequestTenantContext
}

type SendBannerRequest struct {
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1693881641000L
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1693881641000L
	StartTime     *int64                          `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantContext *SendBannerRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s SendBannerRequest) String() string {
	return dara.Prettify(s)
}

func (s SendBannerRequest) GoString() string {
	return s.String()
}

func (s *SendBannerRequest) GetContent() map[string]interface{} {
	return s.Content
}

func (s *SendBannerRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *SendBannerRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SendBannerRequest) GetTenantContext() *SendBannerRequestTenantContext {
	return s.TenantContext
}

func (s *SendBannerRequest) SetContent(v map[string]interface{}) *SendBannerRequest {
	s.Content = v
	return s
}

func (s *SendBannerRequest) SetEndTime(v int64) *SendBannerRequest {
	s.EndTime = &v
	return s
}

func (s *SendBannerRequest) SetStartTime(v int64) *SendBannerRequest {
	s.StartTime = &v
	return s
}

func (s *SendBannerRequest) SetTenantContext(v *SendBannerRequestTenantContext) *SendBannerRequest {
	s.TenantContext = v
	return s
}

func (s *SendBannerRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendBannerRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SendBannerRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s SendBannerRequestTenantContext) GoString() string {
	return s.String()
}

func (s *SendBannerRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *SendBannerRequestTenantContext) SetTenantId(v string) *SendBannerRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *SendBannerRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
