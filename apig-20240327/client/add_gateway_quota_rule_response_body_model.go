// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayQuotaRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddGatewayQuotaRuleResponseBody
	GetCode() *string
	SetData(v *AddGatewayQuotaRuleResponseBodyData) *AddGatewayQuotaRuleResponseBody
	GetData() *AddGatewayQuotaRuleResponseBodyData
	SetMessage(v string) *AddGatewayQuotaRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddGatewayQuotaRuleResponseBody
	GetRequestId() *string
}

type AddGatewayQuotaRuleResponseBody struct {
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
	Data *AddGatewayQuotaRuleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The message content.
	//
	// example:
	//
	// 你好，世界！
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AddGatewayQuotaRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayQuotaRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddGatewayQuotaRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddGatewayQuotaRuleResponseBody) GetData() *AddGatewayQuotaRuleResponseBodyData {
	return s.Data
}

func (s *AddGatewayQuotaRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddGatewayQuotaRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddGatewayQuotaRuleResponseBody) SetCode(v string) *AddGatewayQuotaRuleResponseBody {
	s.Code = &v
	return s
}

func (s *AddGatewayQuotaRuleResponseBody) SetData(v *AddGatewayQuotaRuleResponseBodyData) *AddGatewayQuotaRuleResponseBody {
	s.Data = v
	return s
}

func (s *AddGatewayQuotaRuleResponseBody) SetMessage(v string) *AddGatewayQuotaRuleResponseBody {
	s.Message = &v
	return s
}

func (s *AddGatewayQuotaRuleResponseBody) SetRequestId(v string) *AddGatewayQuotaRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGatewayQuotaRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddGatewayQuotaRuleResponseBodyData struct {
	// Indicates whether the write request is accepted by the system. A value of false typically indicates a retryable scenario such as an unconfirmed conflict overwrite.
	//
	// example:
	//
	// true
	Accepted *bool `json:"accepted,omitempty" xml:"accepted,omitempty"`
	// The conflict preview.
	ConflictPreview *AddGatewayQuotaRuleResponseBodyDataConflictPreview `json:"conflictPreview,omitempty" xml:"conflictPreview,omitempty" type:"Struct"`
	// Indicates whether the request is a dry run.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// qr-xxxxx
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
}

func (s AddGatewayQuotaRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayQuotaRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddGatewayQuotaRuleResponseBodyData) GetAccepted() *bool {
	return s.Accepted
}

func (s *AddGatewayQuotaRuleResponseBodyData) GetConflictPreview() *AddGatewayQuotaRuleResponseBodyDataConflictPreview {
	return s.ConflictPreview
}

func (s *AddGatewayQuotaRuleResponseBodyData) GetDryRun() *bool {
	return s.DryRun
}

func (s *AddGatewayQuotaRuleResponseBodyData) GetRuleId() *string {
	return s.RuleId
}

func (s *AddGatewayQuotaRuleResponseBodyData) SetAccepted(v bool) *AddGatewayQuotaRuleResponseBodyData {
	s.Accepted = &v
	return s
}

func (s *AddGatewayQuotaRuleResponseBodyData) SetConflictPreview(v *AddGatewayQuotaRuleResponseBodyDataConflictPreview) *AddGatewayQuotaRuleResponseBodyData {
	s.ConflictPreview = v
	return s
}

func (s *AddGatewayQuotaRuleResponseBodyData) SetDryRun(v bool) *AddGatewayQuotaRuleResponseBodyData {
	s.DryRun = &v
	return s
}

func (s *AddGatewayQuotaRuleResponseBodyData) SetRuleId(v string) *AddGatewayQuotaRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *AddGatewayQuotaRuleResponseBodyData) Validate() error {
	if s.ConflictPreview != nil {
		if err := s.ConflictPreview.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddGatewayQuotaRuleResponseBodyDataConflictPreview struct {
	// The conflict hash.
	//
	// example:
	//
	// f8f44dc6cf369a017d56b7197eb4fb5ac4bbb6b09a92b9b41999541fxxxxxxxx
	ConflictHash *string `json:"conflictHash,omitempty" xml:"conflictHash,omitempty"`
	// The list of conflicting principals (consumers).
	Items []*AddGatewayQuotaRuleResponseBodyDataConflictPreviewItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The total number of conflicts.
	//
	// example:
	//
	// 2
	TotalConflictCount *int32 `json:"totalConflictCount,omitempty" xml:"totalConflictCount,omitempty"`
}

func (s AddGatewayQuotaRuleResponseBodyDataConflictPreview) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayQuotaRuleResponseBodyDataConflictPreview) GoString() string {
	return s.String()
}

func (s *AddGatewayQuotaRuleResponseBodyDataConflictPreview) GetConflictHash() *string {
	return s.ConflictHash
}

func (s *AddGatewayQuotaRuleResponseBodyDataConflictPreview) GetItems() []*AddGatewayQuotaRuleResponseBodyDataConflictPreviewItems {
	return s.Items
}

func (s *AddGatewayQuotaRuleResponseBodyDataConflictPreview) GetTotalConflictCount() *int32 {
	return s.TotalConflictCount
}

func (s *AddGatewayQuotaRuleResponseBodyDataConflictPreview) SetConflictHash(v string) *AddGatewayQuotaRuleResponseBodyDataConflictPreview {
	s.ConflictHash = &v
	return s
}

func (s *AddGatewayQuotaRuleResponseBodyDataConflictPreview) SetItems(v []*AddGatewayQuotaRuleResponseBodyDataConflictPreviewItems) *AddGatewayQuotaRuleResponseBodyDataConflictPreview {
	s.Items = v
	return s
}

func (s *AddGatewayQuotaRuleResponseBodyDataConflictPreview) SetTotalConflictCount(v int32) *AddGatewayQuotaRuleResponseBodyDataConflictPreview {
	s.TotalConflictCount = &v
	return s
}

func (s *AddGatewayQuotaRuleResponseBodyDataConflictPreview) Validate() error {
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

type AddGatewayQuotaRuleResponseBodyDataConflictPreviewItems struct {
	// The period type of the existing conflicting rule on the consumer. Valid values:
	//
	// - day: The period of the existing conflicting rule is day.
	//
	// - week: The period of the existing conflicting rule is week.
	//
	// - month: The period of the existing conflicting rule is month.
	//
	// example:
	//
	// week
	ConflictPeriodType *string `json:"conflictPeriodType,omitempty" xml:"conflictPeriodType,omitempty"`
	// The type of the existing conflicting rule on the consumer. Valid values:
	//
	// - calendar: The existing conflicting rule uses a calendar period.
	//
	// - epoch: The existing conflicting rule uses a custom period.
	//
	// example:
	//
	// calendar
	ConflictType *string `json:"conflictType,omitempty" xml:"conflictType,omitempty"`
	// The consumer ID.
	//
	// example:
	//
	// cs-xxxxxx
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// The consumer name.
	//
	// example:
	//
	// consumer-a
	ConsumerName *string `json:"consumerName,omitempty" xml:"consumerName,omitempty"`
}

func (s AddGatewayQuotaRuleResponseBodyDataConflictPreviewItems) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayQuotaRuleResponseBodyDataConflictPreviewItems) GoString() string {
	return s.String()
}

func (s *AddGatewayQuotaRuleResponseBodyDataConflictPreviewItems) GetConflictPeriodType() *string {
	return s.ConflictPeriodType
}

func (s *AddGatewayQuotaRuleResponseBodyDataConflictPreviewItems) GetConflictType() *string {
	return s.ConflictType
}

func (s *AddGatewayQuotaRuleResponseBodyDataConflictPreviewItems) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *AddGatewayQuotaRuleResponseBodyDataConflictPreviewItems) GetConsumerName() *string {
	return s.ConsumerName
}

func (s *AddGatewayQuotaRuleResponseBodyDataConflictPreviewItems) SetConflictPeriodType(v string) *AddGatewayQuotaRuleResponseBodyDataConflictPreviewItems {
	s.ConflictPeriodType = &v
	return s
}

func (s *AddGatewayQuotaRuleResponseBodyDataConflictPreviewItems) SetConflictType(v string) *AddGatewayQuotaRuleResponseBodyDataConflictPreviewItems {
	s.ConflictType = &v
	return s
}

func (s *AddGatewayQuotaRuleResponseBodyDataConflictPreviewItems) SetConsumerId(v string) *AddGatewayQuotaRuleResponseBodyDataConflictPreviewItems {
	s.ConsumerId = &v
	return s
}

func (s *AddGatewayQuotaRuleResponseBodyDataConflictPreviewItems) SetConsumerName(v string) *AddGatewayQuotaRuleResponseBodyDataConflictPreviewItems {
	s.ConsumerName = &v
	return s
}

func (s *AddGatewayQuotaRuleResponseBodyDataConflictPreviewItems) Validate() error {
	return dara.Validate(s)
}
