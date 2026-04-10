// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUsageResponseBody
	GetCode() *string
	SetData(v *GetUsageResponseBodyData) *GetUsageResponseBody
	GetData() *GetUsageResponseBodyData
	SetMessage(v string) *GetUsageResponseBody
	GetMessage() *string
	SetRetryAble(v bool) *GetUsageResponseBody
	GetRetryAble() *bool
	SetSuccess(v bool) *GetUsageResponseBody
	GetSuccess() *bool
}

type GetUsageResponseBody struct {
	// example:
	//
	// 0
	Code *string                   `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetUsageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// false
	RetryAble *bool `json:"retryAble,omitempty" xml:"retryAble,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetUsageResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUsageResponseBody) GetData() *GetUsageResponseBodyData {
	return s.Data
}

func (s *GetUsageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUsageResponseBody) GetRetryAble() *bool {
	return s.RetryAble
}

func (s *GetUsageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUsageResponseBody) SetCode(v string) *GetUsageResponseBody {
	s.Code = &v
	return s
}

func (s *GetUsageResponseBody) SetData(v *GetUsageResponseBodyData) *GetUsageResponseBody {
	s.Data = v
	return s
}

func (s *GetUsageResponseBody) SetMessage(v string) *GetUsageResponseBody {
	s.Message = &v
	return s
}

func (s *GetUsageResponseBody) SetRetryAble(v bool) *GetUsageResponseBody {
	s.RetryAble = &v
	return s
}

func (s *GetUsageResponseBody) SetSuccess(v bool) *GetUsageResponseBody {
	s.Success = &v
	return s
}

func (s *GetUsageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUsageResponseBodyData struct {
	Entitlements []*GetUsageResponseBodyDataEntitlements `json:"entitlements,omitempty" xml:"entitlements,omitempty" type:"Repeated"`
	ModelStats   []*GetUsageResponseBodyDataModelStats   `json:"modelStats,omitempty" xml:"modelStats,omitempty" type:"Repeated"`
	Summary      *GetUsageResponseBodyDataSummary        `json:"summary,omitempty" xml:"summary,omitempty" type:"Struct"`
}

func (s GetUsageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUsageResponseBodyData) GetEntitlements() []*GetUsageResponseBodyDataEntitlements {
	return s.Entitlements
}

func (s *GetUsageResponseBodyData) GetModelStats() []*GetUsageResponseBodyDataModelStats {
	return s.ModelStats
}

func (s *GetUsageResponseBodyData) GetSummary() *GetUsageResponseBodyDataSummary {
	return s.Summary
}

func (s *GetUsageResponseBodyData) SetEntitlements(v []*GetUsageResponseBodyDataEntitlements) *GetUsageResponseBodyData {
	s.Entitlements = v
	return s
}

func (s *GetUsageResponseBodyData) SetModelStats(v []*GetUsageResponseBodyDataModelStats) *GetUsageResponseBodyData {
	s.ModelStats = v
	return s
}

func (s *GetUsageResponseBodyData) SetSummary(v *GetUsageResponseBodyDataSummary) *GetUsageResponseBodyData {
	s.Summary = v
	return s
}

func (s *GetUsageResponseBodyData) Validate() error {
	if s.Entitlements != nil {
		for _, item := range s.Entitlements {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ModelStats != nil {
		for _, item := range s.ModelStats {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Summary != nil {
		if err := s.Summary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUsageResponseBodyDataEntitlements struct {
	AllowedModels []*string `json:"allowedModels,omitempty" xml:"allowedModels,omitempty" type:"Repeated"`
	// example:
	//
	// 238746
	BindingId *int64 `json:"bindingId,omitempty" xml:"bindingId,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	EffectiveAt *string `json:"effectiveAt,omitempty" xml:"effectiveAt,omitempty"`
	// example:
	//
	// 2024-01-31T00:00:00Z
	ExpireAt *string `json:"expireAt,omitempty" xml:"expireAt,omitempty"`
	// example:
	//
	// 1000
	QuotaInitial *int64 `json:"quotaInitial,omitempty" xml:"quotaInitial,omitempty"`
	// example:
	//
	// 500
	QuotaRemaining *int64 `json:"quotaRemaining,omitempty" xml:"quotaRemaining,omitempty"`
	// example:
	//
	// 500
	QuotaUsed *int64 `json:"quotaUsed,omitempty" xml:"quotaUsed,omitempty"`
	// example:
	//
	// ACTIVE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 893724
	TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// example:
	//
	// 1
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s GetUsageResponseBodyDataEntitlements) String() string {
	return dara.Prettify(s)
}

func (s GetUsageResponseBodyDataEntitlements) GoString() string {
	return s.String()
}

func (s *GetUsageResponseBodyDataEntitlements) GetAllowedModels() []*string {
	return s.AllowedModels
}

func (s *GetUsageResponseBodyDataEntitlements) GetBindingId() *int64 {
	return s.BindingId
}

func (s *GetUsageResponseBodyDataEntitlements) GetEffectiveAt() *string {
	return s.EffectiveAt
}

func (s *GetUsageResponseBodyDataEntitlements) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *GetUsageResponseBodyDataEntitlements) GetQuotaInitial() *int64 {
	return s.QuotaInitial
}

func (s *GetUsageResponseBodyDataEntitlements) GetQuotaRemaining() *int64 {
	return s.QuotaRemaining
}

func (s *GetUsageResponseBodyDataEntitlements) GetQuotaUsed() *int64 {
	return s.QuotaUsed
}

func (s *GetUsageResponseBodyDataEntitlements) GetStatus() *string {
	return s.Status
}

func (s *GetUsageResponseBodyDataEntitlements) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *GetUsageResponseBodyDataEntitlements) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetUsageResponseBodyDataEntitlements) SetAllowedModels(v []*string) *GetUsageResponseBodyDataEntitlements {
	s.AllowedModels = v
	return s
}

func (s *GetUsageResponseBodyDataEntitlements) SetBindingId(v int64) *GetUsageResponseBodyDataEntitlements {
	s.BindingId = &v
	return s
}

func (s *GetUsageResponseBodyDataEntitlements) SetEffectiveAt(v string) *GetUsageResponseBodyDataEntitlements {
	s.EffectiveAt = &v
	return s
}

func (s *GetUsageResponseBodyDataEntitlements) SetExpireAt(v string) *GetUsageResponseBodyDataEntitlements {
	s.ExpireAt = &v
	return s
}

func (s *GetUsageResponseBodyDataEntitlements) SetQuotaInitial(v int64) *GetUsageResponseBodyDataEntitlements {
	s.QuotaInitial = &v
	return s
}

func (s *GetUsageResponseBodyDataEntitlements) SetQuotaRemaining(v int64) *GetUsageResponseBodyDataEntitlements {
	s.QuotaRemaining = &v
	return s
}

func (s *GetUsageResponseBodyDataEntitlements) SetQuotaUsed(v int64) *GetUsageResponseBodyDataEntitlements {
	s.QuotaUsed = &v
	return s
}

func (s *GetUsageResponseBodyDataEntitlements) SetStatus(v string) *GetUsageResponseBodyDataEntitlements {
	s.Status = &v
	return s
}

func (s *GetUsageResponseBodyDataEntitlements) SetTemplateId(v int64) *GetUsageResponseBodyDataEntitlements {
	s.TemplateId = &v
	return s
}

func (s *GetUsageResponseBodyDataEntitlements) SetTemplateName(v string) *GetUsageResponseBodyDataEntitlements {
	s.TemplateName = &v
	return s
}

func (s *GetUsageResponseBodyDataEntitlements) Validate() error {
	return dara.Validate(s)
}

type GetUsageResponseBodyDataModelStats struct {
	// example:
	//
	// 1000
	InputUsage *int64 `json:"inputUsage,omitempty" xml:"inputUsage,omitempty"`
	// example:
	//
	// qwen-turbo
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// example:
	//
	// 500
	OutputUsage *int64 `json:"outputUsage,omitempty" xml:"outputUsage,omitempty"`
	// example:
	//
	// 1
	Requests *int64 `json:"requests,omitempty" xml:"requests,omitempty"`
	// example:
	//
	// 1500
	TotalUsage *int64 `json:"totalUsage,omitempty" xml:"totalUsage,omitempty"`
}

func (s GetUsageResponseBodyDataModelStats) String() string {
	return dara.Prettify(s)
}

func (s GetUsageResponseBodyDataModelStats) GoString() string {
	return s.String()
}

func (s *GetUsageResponseBodyDataModelStats) GetInputUsage() *int64 {
	return s.InputUsage
}

func (s *GetUsageResponseBodyDataModelStats) GetModel() *string {
	return s.Model
}

func (s *GetUsageResponseBodyDataModelStats) GetOutputUsage() *int64 {
	return s.OutputUsage
}

func (s *GetUsageResponseBodyDataModelStats) GetRequests() *int64 {
	return s.Requests
}

func (s *GetUsageResponseBodyDataModelStats) GetTotalUsage() *int64 {
	return s.TotalUsage
}

func (s *GetUsageResponseBodyDataModelStats) SetInputUsage(v int64) *GetUsageResponseBodyDataModelStats {
	s.InputUsage = &v
	return s
}

func (s *GetUsageResponseBodyDataModelStats) SetModel(v string) *GetUsageResponseBodyDataModelStats {
	s.Model = &v
	return s
}

func (s *GetUsageResponseBodyDataModelStats) SetOutputUsage(v int64) *GetUsageResponseBodyDataModelStats {
	s.OutputUsage = &v
	return s
}

func (s *GetUsageResponseBodyDataModelStats) SetRequests(v int64) *GetUsageResponseBodyDataModelStats {
	s.Requests = &v
	return s
}

func (s *GetUsageResponseBodyDataModelStats) SetTotalUsage(v int64) *GetUsageResponseBodyDataModelStats {
	s.TotalUsage = &v
	return s
}

func (s *GetUsageResponseBodyDataModelStats) Validate() error {
	return dara.Validate(s)
}

type GetUsageResponseBodyDataSummary struct {
	// example:
	//
	// 500
	TotalInputUsage *int64 `json:"totalInputUsage,omitempty" xml:"totalInputUsage,omitempty"`
	// example:
	//
	// 500
	TotalOutputUsage *int64 `json:"totalOutputUsage,omitempty" xml:"totalOutputUsage,omitempty"`
	// example:
	//
	// 1
	TotalRequests *int64 `json:"totalRequests,omitempty" xml:"totalRequests,omitempty"`
	// example:
	//
	// 1000
	TotalUsage *int64 `json:"totalUsage,omitempty" xml:"totalUsage,omitempty"`
}

func (s GetUsageResponseBodyDataSummary) String() string {
	return dara.Prettify(s)
}

func (s GetUsageResponseBodyDataSummary) GoString() string {
	return s.String()
}

func (s *GetUsageResponseBodyDataSummary) GetTotalInputUsage() *int64 {
	return s.TotalInputUsage
}

func (s *GetUsageResponseBodyDataSummary) GetTotalOutputUsage() *int64 {
	return s.TotalOutputUsage
}

func (s *GetUsageResponseBodyDataSummary) GetTotalRequests() *int64 {
	return s.TotalRequests
}

func (s *GetUsageResponseBodyDataSummary) GetTotalUsage() *int64 {
	return s.TotalUsage
}

func (s *GetUsageResponseBodyDataSummary) SetTotalInputUsage(v int64) *GetUsageResponseBodyDataSummary {
	s.TotalInputUsage = &v
	return s
}

func (s *GetUsageResponseBodyDataSummary) SetTotalOutputUsage(v int64) *GetUsageResponseBodyDataSummary {
	s.TotalOutputUsage = &v
	return s
}

func (s *GetUsageResponseBodyDataSummary) SetTotalRequests(v int64) *GetUsageResponseBodyDataSummary {
	s.TotalRequests = &v
	return s
}

func (s *GetUsageResponseBodyDataSummary) SetTotalUsage(v int64) *GetUsageResponseBodyDataSummary {
	s.TotalUsage = &v
	return s
}

func (s *GetUsageResponseBodyDataSummary) Validate() error {
	return dara.Validate(s)
}
