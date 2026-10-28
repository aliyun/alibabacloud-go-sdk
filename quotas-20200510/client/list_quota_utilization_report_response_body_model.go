// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaUtilizationReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListQuotaUtilizationReportResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ListQuotaUtilizationReportResponseBody
	GetErrorMsg() *string
	SetGeneratedAt(v string) *ListQuotaUtilizationReportResponseBody
	GetGeneratedAt() *string
	SetMessage(v string) *ListQuotaUtilizationReportResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListQuotaUtilizationReportResponseBody
	GetNextToken() *string
	SetQuotas(v []*ListQuotaUtilizationReportResponseBodyQuotas) *ListQuotaUtilizationReportResponseBody
	GetQuotas() []*ListQuotaUtilizationReportResponseBodyQuotas
	SetRequestId(v string) *ListQuotaUtilizationReportResponseBody
	GetRequestId() *string
	SetStatus(v string) *ListQuotaUtilizationReportResponseBody
	GetStatus() *string
	SetSuccess(v bool) *ListQuotaUtilizationReportResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListQuotaUtilizationReportResponseBody
	GetTotalCount() *int32
}

type ListQuotaUtilizationReportResponseBody struct {
	// example:
	//
	// RAM.PERMISSION.DENIED
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// You are not authorized to do this action or the API input parameter is invalid.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 2026-10-26T06:24:57Z
	GeneratedAt *string `json:"GeneratedAt,omitempty" xml:"GeneratedAt,omitempty"`
	// example:
	//
	// Quota utilization data is available.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// eyJvZmZzZXQiOjEwLCJvdHNOZXh0VG9rZW4iOiJDQUVTR3dvWkNoY0tFM1J2ZEdGc1gzVnpZV2RsWDNCbGNtTmxiblFRQVJnQkl1RUJDZ2tCQUFBQUFBQUFXVUFLMHdFRHpnQUFBREZUTXpFek9ETXdNemN6TnpNNE16VXpPVE0xTXpjek5qTXpNekV6TlRNNE16Z3VVek01TXpjek5UTTJOalkyTWpZeE16Y3laRFl5TmpJek16WXhNbVF6TkRNMk16RXpOREprTmpFek9UTTRNemN5WkRZMk16WTJORE16TXpVMk16TXhNemt6TURZMk16RXpNUzVUTmpFMk16YzBOamsyWmpabE56UTNNall4TmprMll5NVRNelkyTWpZMU16a3pPRE0yTmpRek1qSmtNekV6TlRNeU5qRXlaRE0wTXpFek1UWTJNbVEyTWpZME16TXpNVEprTXpNMk16WXlOalV6TkRNMU16VXpNalkwTXpVek1qTTMiLCJ0YXNrSWQiOiI5NzU2ZmJhNy1iYjNhLTQ2MTQtYTk4Ny1mNmQzNWMxOTBmMTEiLCJ1aWQiOiIxODA3Nzg1OTU3NjMxNTg4IiwidXRpbGl6YXRpb25TdGF0dXNlcyI6IiJ9
	NextToken *string                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Quotas    []*ListQuotaUtilizationReportResponseBodyQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// example:
	//
	// 7ED584FB-ECBF-4A2A-969D-F54D25EFABF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQuotaUtilizationReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaUtilizationReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotaUtilizationReportResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListQuotaUtilizationReportResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListQuotaUtilizationReportResponseBody) GetGeneratedAt() *string {
	return s.GeneratedAt
}

func (s *ListQuotaUtilizationReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListQuotaUtilizationReportResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQuotaUtilizationReportResponseBody) GetQuotas() []*ListQuotaUtilizationReportResponseBodyQuotas {
	return s.Quotas
}

func (s *ListQuotaUtilizationReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQuotaUtilizationReportResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListQuotaUtilizationReportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListQuotaUtilizationReportResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListQuotaUtilizationReportResponseBody) SetErrorCode(v string) *ListQuotaUtilizationReportResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBody) SetErrorMsg(v string) *ListQuotaUtilizationReportResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBody) SetGeneratedAt(v string) *ListQuotaUtilizationReportResponseBody {
	s.GeneratedAt = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBody) SetMessage(v string) *ListQuotaUtilizationReportResponseBody {
	s.Message = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBody) SetNextToken(v string) *ListQuotaUtilizationReportResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBody) SetQuotas(v []*ListQuotaUtilizationReportResponseBodyQuotas) *ListQuotaUtilizationReportResponseBody {
	s.Quotas = v
	return s
}

func (s *ListQuotaUtilizationReportResponseBody) SetRequestId(v string) *ListQuotaUtilizationReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBody) SetStatus(v string) *ListQuotaUtilizationReportResponseBody {
	s.Status = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBody) SetSuccess(v bool) *ListQuotaUtilizationReportResponseBody {
	s.Success = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBody) SetTotalCount(v int32) *ListQuotaUtilizationReportResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBody) Validate() error {
	if s.Quotas != nil {
		for _, item := range s.Quotas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQuotaUtilizationReportResponseBodyQuotas struct {
	// example:
	//
	// true
	Adjustable *bool `json:"Adjustable,omitempty" xml:"Adjustable,omitempty"`
	// example:
	//
	// true
	Consumable      *bool                                                        `json:"Consumable,omitempty" xml:"Consumable,omitempty"`
	Dimensions      map[string]*string                                           `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	FlowControlAlgo *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgo `json:"FlowControlAlgo,omitempty" xml:"FlowControlAlgo,omitempty" type:"Struct"`
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// example:
	//
	// acs:quotas:cn-hangzhou:120886317861****:quota/ecs/q_security-groups
	QuotaArn *string `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// example:
	//
	// The maximum number of security groups that can be owned by the current account.
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	// example:
	//
	// ecs.g5.2xlarge
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// example:
	//
	// 10000
	TotalQuota *float64 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// example:
	//
	// 1009
	TotalUsage *float64 `json:"TotalUsage,omitempty" xml:"TotalUsage,omitempty"`
	// example:
	//
	// applicationProcess
	UnadjustableDetail *string                                                  `json:"UnadjustableDetail,omitempty" xml:"UnadjustableDetail,omitempty"`
	UsageMetric        *ListQuotaUtilizationReportResponseBodyQuotasUsageMetric `json:"UsageMetric,omitempty" xml:"UsageMetric,omitempty" type:"Struct"`
	// example:
	//
	// 10.1
	Utilization *float64 `json:"Utilization,omitempty" xml:"Utilization,omitempty"`
	// example:
	//
	// WARN
	UtilizationStatus *string `json:"UtilizationStatus,omitempty" xml:"UtilizationStatus,omitempty"`
}

func (s ListQuotaUtilizationReportResponseBodyQuotas) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaUtilizationReportResponseBodyQuotas) GoString() string {
	return s.String()
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) GetAdjustable() *bool {
	return s.Adjustable
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) GetConsumable() *bool {
	return s.Consumable
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) GetDimensions() map[string]*string {
	return s.Dimensions
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) GetFlowControlAlgo() *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgo {
	return s.FlowControlAlgo
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) GetQuotaArn() *string {
	return s.QuotaArn
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) GetQuotaCategory() *string {
	return s.QuotaCategory
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) GetQuotaDescription() *string {
	return s.QuotaDescription
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) GetQuotaName() *string {
	return s.QuotaName
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) GetTotalQuota() *float64 {
	return s.TotalQuota
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) GetTotalUsage() *float64 {
	return s.TotalUsage
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) GetUnadjustableDetail() *string {
	return s.UnadjustableDetail
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) GetUsageMetric() *ListQuotaUtilizationReportResponseBodyQuotasUsageMetric {
	return s.UsageMetric
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) GetUtilization() *float64 {
	return s.Utilization
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) GetUtilizationStatus() *string {
	return s.UtilizationStatus
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) SetAdjustable(v bool) *ListQuotaUtilizationReportResponseBodyQuotas {
	s.Adjustable = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) SetConsumable(v bool) *ListQuotaUtilizationReportResponseBodyQuotas {
	s.Consumable = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) SetDimensions(v map[string]*string) *ListQuotaUtilizationReportResponseBodyQuotas {
	s.Dimensions = v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) SetFlowControlAlgo(v *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgo) *ListQuotaUtilizationReportResponseBodyQuotas {
	s.FlowControlAlgo = v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) SetProductCode(v string) *ListQuotaUtilizationReportResponseBodyQuotas {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) SetQuotaActionCode(v string) *ListQuotaUtilizationReportResponseBodyQuotas {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) SetQuotaArn(v string) *ListQuotaUtilizationReportResponseBodyQuotas {
	s.QuotaArn = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) SetQuotaCategory(v string) *ListQuotaUtilizationReportResponseBodyQuotas {
	s.QuotaCategory = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) SetQuotaDescription(v string) *ListQuotaUtilizationReportResponseBodyQuotas {
	s.QuotaDescription = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) SetQuotaName(v string) *ListQuotaUtilizationReportResponseBodyQuotas {
	s.QuotaName = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) SetTotalQuota(v float64) *ListQuotaUtilizationReportResponseBodyQuotas {
	s.TotalQuota = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) SetTotalUsage(v float64) *ListQuotaUtilizationReportResponseBodyQuotas {
	s.TotalUsage = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) SetUnadjustableDetail(v string) *ListQuotaUtilizationReportResponseBodyQuotas {
	s.UnadjustableDetail = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) SetUsageMetric(v *ListQuotaUtilizationReportResponseBodyQuotasUsageMetric) *ListQuotaUtilizationReportResponseBodyQuotas {
	s.UsageMetric = v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) SetUtilization(v float64) *ListQuotaUtilizationReportResponseBodyQuotas {
	s.Utilization = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) SetUtilizationStatus(v string) *ListQuotaUtilizationReportResponseBodyQuotas {
	s.UtilizationStatus = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotas) Validate() error {
	if s.FlowControlAlgo != nil {
		if err := s.FlowControlAlgo.Validate(); err != nil {
			return err
		}
	}
	if s.UsageMetric != nil {
		if err := s.UsageMetric.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgo struct {
	// example:
	//
	// FixedWindow
	AlgoType *string                                                            `json:"AlgoType,omitempty" xml:"AlgoType,omitempty"`
	Period   *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgoPeriod `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	// example:
	//
	// 1
	RefillInterval *int32 `json:"RefillInterval,omitempty" xml:"RefillInterval,omitempty"`
	// example:
	//
	// 2
	TokensPerRefill *int32 `json:"TokensPerRefill,omitempty" xml:"TokensPerRefill,omitempty"`
	// example:
	//
	// 10
	TotalQuota *float64 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
}

func (s ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgo) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgo) GoString() string {
	return s.String()
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgo) GetAlgoType() *string {
	return s.AlgoType
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgo) GetPeriod() *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgoPeriod {
	return s.Period
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgo) GetRefillInterval() *int32 {
	return s.RefillInterval
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgo) GetTokensPerRefill() *int32 {
	return s.TokensPerRefill
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgo) GetTotalQuota() *float64 {
	return s.TotalQuota
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgo) SetAlgoType(v string) *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgo {
	s.AlgoType = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgo) SetPeriod(v *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgoPeriod) *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgo {
	s.Period = v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgo) SetRefillInterval(v int32) *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgo {
	s.RefillInterval = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgo) SetTokensPerRefill(v int32) *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgo {
	s.TokensPerRefill = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgo) SetTotalQuota(v float64) *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgo {
	s.TotalQuota = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgo) Validate() error {
	if s.Period != nil {
		if err := s.Period.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgoPeriod struct {
	// example:
	//
	// second
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// example:
	//
	// 1
	PeriodValue *int32 `json:"PeriodValue,omitempty" xml:"PeriodValue,omitempty"`
}

func (s ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgoPeriod) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgoPeriod) GoString() string {
	return s.String()
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgoPeriod) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgoPeriod) GetPeriodValue() *int32 {
	return s.PeriodValue
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgoPeriod) SetPeriodUnit(v string) *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgoPeriod {
	s.PeriodUnit = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgoPeriod) SetPeriodValue(v int32) *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgoPeriod {
	s.PeriodValue = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasFlowControlAlgoPeriod) Validate() error {
	return dara.Validate(s)
}

type ListQuotaUtilizationReportResponseBodyQuotasUsageMetric struct {
	MetricDimensions map[string]*string `json:"MetricDimensions,omitempty" xml:"MetricDimensions,omitempty"`
	// example:
	//
	// Usage
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// example:
	//
	// acs_quotas_flowcontrol
	MetricNamespace *string `json:"MetricNamespace,omitempty" xml:"MetricNamespace,omitempty"`
}

func (s ListQuotaUtilizationReportResponseBodyQuotasUsageMetric) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaUtilizationReportResponseBodyQuotasUsageMetric) GoString() string {
	return s.String()
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasUsageMetric) GetMetricDimensions() map[string]*string {
	return s.MetricDimensions
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasUsageMetric) GetMetricName() *string {
	return s.MetricName
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasUsageMetric) GetMetricNamespace() *string {
	return s.MetricNamespace
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasUsageMetric) SetMetricDimensions(v map[string]*string) *ListQuotaUtilizationReportResponseBodyQuotasUsageMetric {
	s.MetricDimensions = v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasUsageMetric) SetMetricName(v string) *ListQuotaUtilizationReportResponseBodyQuotasUsageMetric {
	s.MetricName = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasUsageMetric) SetMetricNamespace(v string) *ListQuotaUtilizationReportResponseBodyQuotasUsageMetric {
	s.MetricNamespace = &v
	return s
}

func (s *ListQuotaUtilizationReportResponseBodyQuotasUsageMetric) Validate() error {
	return dara.Validate(s)
}
