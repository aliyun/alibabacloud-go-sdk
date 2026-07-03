// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetGatewayQuotaRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResetGatewayQuotaRuleResponseBody
	GetCode() *string
	SetData(v *ResetGatewayQuotaRuleResponseBodyData) *ResetGatewayQuotaRuleResponseBody
	GetData() *ResetGatewayQuotaRuleResponseBodyData
	SetMessage(v string) *ResetGatewayQuotaRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResetGatewayQuotaRuleResponseBody
	GetRequestId() *string
}

type ResetGatewayQuotaRuleResponseBody struct {
	// The status code or error code.
	//
	// example:
	//
	// 200, 404, 500
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	//
	// example:
	//
	// {\\"key\\": \\"value\\"}
	Data *ResetGatewayQuotaRuleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The message content.
	//
	// example:
	//
	// 你好，世界！
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The unique identifier of the request.
	//
	// example:
	//
	// 1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ResetGatewayQuotaRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetGatewayQuotaRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ResetGatewayQuotaRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResetGatewayQuotaRuleResponseBody) GetData() *ResetGatewayQuotaRuleResponseBodyData {
	return s.Data
}

func (s *ResetGatewayQuotaRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResetGatewayQuotaRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetGatewayQuotaRuleResponseBody) SetCode(v string) *ResetGatewayQuotaRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ResetGatewayQuotaRuleResponseBody) SetData(v *ResetGatewayQuotaRuleResponseBodyData) *ResetGatewayQuotaRuleResponseBody {
	s.Data = v
	return s
}

func (s *ResetGatewayQuotaRuleResponseBody) SetMessage(v string) *ResetGatewayQuotaRuleResponseBody {
	s.Message = &v
	return s
}

func (s *ResetGatewayQuotaRuleResponseBody) SetRequestId(v string) *ResetGatewayQuotaRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetGatewayQuotaRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResetGatewayQuotaRuleResponseBodyData struct {
	// Indicates whether the write request semantics are accepted by the system. A value of false typically indicates a retryable scenario such as an unconfirmed conflict overwrite.
	//
	// example:
	//
	// true
	Accepted *bool `json:"accepted,omitempty" xml:"accepted,omitempty"`
	// The conflict preview.
	ConflictPreview *ResetGatewayQuotaRuleResponseBodyDataConflictPreview `json:"conflictPreview,omitempty" xml:"conflictPreview,omitempty" type:"Struct"`
	// Indicates whether this is a dry run.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// qr-d8j7fpmm1hks65xxxx
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
}

func (s ResetGatewayQuotaRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ResetGatewayQuotaRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ResetGatewayQuotaRuleResponseBodyData) GetAccepted() *bool {
	return s.Accepted
}

func (s *ResetGatewayQuotaRuleResponseBodyData) GetConflictPreview() *ResetGatewayQuotaRuleResponseBodyDataConflictPreview {
	return s.ConflictPreview
}

func (s *ResetGatewayQuotaRuleResponseBodyData) GetDryRun() *bool {
	return s.DryRun
}

func (s *ResetGatewayQuotaRuleResponseBodyData) GetRuleId() *string {
	return s.RuleId
}

func (s *ResetGatewayQuotaRuleResponseBodyData) SetAccepted(v bool) *ResetGatewayQuotaRuleResponseBodyData {
	s.Accepted = &v
	return s
}

func (s *ResetGatewayQuotaRuleResponseBodyData) SetConflictPreview(v *ResetGatewayQuotaRuleResponseBodyDataConflictPreview) *ResetGatewayQuotaRuleResponseBodyData {
	s.ConflictPreview = v
	return s
}

func (s *ResetGatewayQuotaRuleResponseBodyData) SetDryRun(v bool) *ResetGatewayQuotaRuleResponseBodyData {
	s.DryRun = &v
	return s
}

func (s *ResetGatewayQuotaRuleResponseBodyData) SetRuleId(v string) *ResetGatewayQuotaRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *ResetGatewayQuotaRuleResponseBodyData) Validate() error {
	if s.ConflictPreview != nil {
		if err := s.ConflictPreview.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResetGatewayQuotaRuleResponseBodyDataConflictPreview struct {
	// The conflict hash.
	//
	// example:
	//
	// f8f44dc6cf369a017d56b7197eb4fb5ac4bbb6b09a92b9b41999541fxxxxxxxx
	ConflictHash *string `json:"conflictHash,omitempty" xml:"conflictHash,omitempty"`
	// The list of conflicting entities (consumers).
	Items []*ResetGatewayQuotaRuleResponseBodyDataConflictPreviewItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The total number of conflicts.
	//
	// example:
	//
	// 20
	TotalConflictCount *int32 `json:"totalConflictCount,omitempty" xml:"totalConflictCount,omitempty"`
}

func (s ResetGatewayQuotaRuleResponseBodyDataConflictPreview) String() string {
	return dara.Prettify(s)
}

func (s ResetGatewayQuotaRuleResponseBodyDataConflictPreview) GoString() string {
	return s.String()
}

func (s *ResetGatewayQuotaRuleResponseBodyDataConflictPreview) GetConflictHash() *string {
	return s.ConflictHash
}

func (s *ResetGatewayQuotaRuleResponseBodyDataConflictPreview) GetItems() []*ResetGatewayQuotaRuleResponseBodyDataConflictPreviewItems {
	return s.Items
}

func (s *ResetGatewayQuotaRuleResponseBodyDataConflictPreview) GetTotalConflictCount() *int32 {
	return s.TotalConflictCount
}

func (s *ResetGatewayQuotaRuleResponseBodyDataConflictPreview) SetConflictHash(v string) *ResetGatewayQuotaRuleResponseBodyDataConflictPreview {
	s.ConflictHash = &v
	return s
}

func (s *ResetGatewayQuotaRuleResponseBodyDataConflictPreview) SetItems(v []*ResetGatewayQuotaRuleResponseBodyDataConflictPreviewItems) *ResetGatewayQuotaRuleResponseBodyDataConflictPreview {
	s.Items = v
	return s
}

func (s *ResetGatewayQuotaRuleResponseBodyDataConflictPreview) SetTotalConflictCount(v int32) *ResetGatewayQuotaRuleResponseBodyDataConflictPreview {
	s.TotalConflictCount = &v
	return s
}

func (s *ResetGatewayQuotaRuleResponseBodyDataConflictPreview) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ResetGatewayQuotaRuleResponseBodyDataConflictPreviewItems struct {
	// The period type of the existing conflicting rule on the consumer. Valid values: day (daily period), week (weekly period), or month (monthly period).
	//
	// example:
	//
	// week
	ConflictPeriodType *string `json:"conflictPeriodType,omitempty" xml:"conflictPeriodType,omitempty"`
	// The type of the existing conflicting rule on the consumer. Valid values: calendar (the conflicting rule uses a calendar period) and epoch (the conflicting rule uses a custom period).
	//
	// example:
	//
	// calendar
	ConflictType *string `json:"conflictType,omitempty" xml:"conflictType,omitempty"`
	// The consumer ID.
	//
	// example:
	//
	// cs-d82n1g6m1hkm3xxxxxxx
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// The consumer name.
	//
	// example:
	//
	// consumer-a
	ConsumerName *string `json:"consumerName,omitempty" xml:"consumerName,omitempty"`
}

func (s ResetGatewayQuotaRuleResponseBodyDataConflictPreviewItems) String() string {
	return dara.Prettify(s)
}

func (s ResetGatewayQuotaRuleResponseBodyDataConflictPreviewItems) GoString() string {
	return s.String()
}

func (s *ResetGatewayQuotaRuleResponseBodyDataConflictPreviewItems) GetConflictPeriodType() *string {
	return s.ConflictPeriodType
}

func (s *ResetGatewayQuotaRuleResponseBodyDataConflictPreviewItems) GetConflictType() *string {
	return s.ConflictType
}

func (s *ResetGatewayQuotaRuleResponseBodyDataConflictPreviewItems) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *ResetGatewayQuotaRuleResponseBodyDataConflictPreviewItems) GetConsumerName() *string {
	return s.ConsumerName
}

func (s *ResetGatewayQuotaRuleResponseBodyDataConflictPreviewItems) SetConflictPeriodType(v string) *ResetGatewayQuotaRuleResponseBodyDataConflictPreviewItems {
	s.ConflictPeriodType = &v
	return s
}

func (s *ResetGatewayQuotaRuleResponseBodyDataConflictPreviewItems) SetConflictType(v string) *ResetGatewayQuotaRuleResponseBodyDataConflictPreviewItems {
	s.ConflictType = &v
	return s
}

func (s *ResetGatewayQuotaRuleResponseBodyDataConflictPreviewItems) SetConsumerId(v string) *ResetGatewayQuotaRuleResponseBodyDataConflictPreviewItems {
	s.ConsumerId = &v
	return s
}

func (s *ResetGatewayQuotaRuleResponseBodyDataConflictPreviewItems) SetConsumerName(v string) *ResetGatewayQuotaRuleResponseBodyDataConflictPreviewItems {
	s.ConsumerName = &v
	return s
}

func (s *ResetGatewayQuotaRuleResponseBodyDataConflictPreviewItems) Validate() error {
	return dara.Validate(s)
}
