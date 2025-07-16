// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeliveryPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *CreateDeliveryPlanRequest
	GetContent() map[string]interface{}
	SetEndTime(v int64) *CreateDeliveryPlanRequest
	GetEndTime() *int64
	SetResId(v string) *CreateDeliveryPlanRequest
	GetResId() *string
	SetStartTime(v int64) *CreateDeliveryPlanRequest
	GetStartTime() *int64
	SetTenantContext(v *CreateDeliveryPlanRequestTenantContext) *CreateDeliveryPlanRequest
	GetTenantContext() *CreateDeliveryPlanRequestTenantContext
	SetUserIdList(v []*string) *CreateDeliveryPlanRequest
	GetUserIdList() []*string
}

type CreateDeliveryPlanRequest struct {
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
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
	StartTime     *int64                                  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantContext *CreateDeliveryPlanRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	UserIdList    []*string                               `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
}

func (s CreateDeliveryPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateDeliveryPlanRequest) GetContent() map[string]interface{} {
	return s.Content
}

func (s *CreateDeliveryPlanRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateDeliveryPlanRequest) GetResId() *string {
	return s.ResId
}

func (s *CreateDeliveryPlanRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateDeliveryPlanRequest) GetTenantContext() *CreateDeliveryPlanRequestTenantContext {
	return s.TenantContext
}

func (s *CreateDeliveryPlanRequest) GetUserIdList() []*string {
	return s.UserIdList
}

func (s *CreateDeliveryPlanRequest) SetContent(v map[string]interface{}) *CreateDeliveryPlanRequest {
	s.Content = v
	return s
}

func (s *CreateDeliveryPlanRequest) SetEndTime(v int64) *CreateDeliveryPlanRequest {
	s.EndTime = &v
	return s
}

func (s *CreateDeliveryPlanRequest) SetResId(v string) *CreateDeliveryPlanRequest {
	s.ResId = &v
	return s
}

func (s *CreateDeliveryPlanRequest) SetStartTime(v int64) *CreateDeliveryPlanRequest {
	s.StartTime = &v
	return s
}

func (s *CreateDeliveryPlanRequest) SetTenantContext(v *CreateDeliveryPlanRequestTenantContext) *CreateDeliveryPlanRequest {
	s.TenantContext = v
	return s
}

func (s *CreateDeliveryPlanRequest) SetUserIdList(v []*string) *CreateDeliveryPlanRequest {
	s.UserIdList = v
	return s
}

func (s *CreateDeliveryPlanRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDeliveryPlanRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateDeliveryPlanRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryPlanRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreateDeliveryPlanRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateDeliveryPlanRequestTenantContext) SetTenantId(v string) *CreateDeliveryPlanRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CreateDeliveryPlanRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
