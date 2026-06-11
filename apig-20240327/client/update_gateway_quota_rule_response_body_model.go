// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayQuotaRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateGatewayQuotaRuleResponseBody
	GetCode() *string
	SetData(v *UpdateGatewayQuotaRuleResponseBodyData) *UpdateGatewayQuotaRuleResponseBody
	GetData() *UpdateGatewayQuotaRuleResponseBodyData
	SetMessage(v string) *UpdateGatewayQuotaRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayQuotaRuleResponseBody
	GetRequestId() *string
}

type UpdateGatewayQuotaRuleResponseBody struct {
	// example:
	//
	// 200, 404, 500
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// {\\"key\\": \\"value\\"}
	Data *UpdateGatewayQuotaRuleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 你好，世界！
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateGatewayQuotaRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayQuotaRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayQuotaRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateGatewayQuotaRuleResponseBody) GetData() *UpdateGatewayQuotaRuleResponseBodyData {
	return s.Data
}

func (s *UpdateGatewayQuotaRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayQuotaRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayQuotaRuleResponseBody) SetCode(v string) *UpdateGatewayQuotaRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayQuotaRuleResponseBody) SetData(v *UpdateGatewayQuotaRuleResponseBodyData) *UpdateGatewayQuotaRuleResponseBody {
	s.Data = v
	return s
}

func (s *UpdateGatewayQuotaRuleResponseBody) SetMessage(v string) *UpdateGatewayQuotaRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayQuotaRuleResponseBody) SetRequestId(v string) *UpdateGatewayQuotaRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayQuotaRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateGatewayQuotaRuleResponseBodyData struct {
	Accepted        *bool                                                  `json:"accepted,omitempty" xml:"accepted,omitempty"`
	ConflictPreview *UpdateGatewayQuotaRuleResponseBodyDataConflictPreview `json:"conflictPreview,omitempty" xml:"conflictPreview,omitempty" type:"Struct"`
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// example:
	//
	// qr-xxxxxx
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
}

func (s UpdateGatewayQuotaRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayQuotaRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateGatewayQuotaRuleResponseBodyData) GetAccepted() *bool {
	return s.Accepted
}

func (s *UpdateGatewayQuotaRuleResponseBodyData) GetConflictPreview() *UpdateGatewayQuotaRuleResponseBodyDataConflictPreview {
	return s.ConflictPreview
}

func (s *UpdateGatewayQuotaRuleResponseBodyData) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateGatewayQuotaRuleResponseBodyData) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateGatewayQuotaRuleResponseBodyData) SetAccepted(v bool) *UpdateGatewayQuotaRuleResponseBodyData {
	s.Accepted = &v
	return s
}

func (s *UpdateGatewayQuotaRuleResponseBodyData) SetConflictPreview(v *UpdateGatewayQuotaRuleResponseBodyDataConflictPreview) *UpdateGatewayQuotaRuleResponseBodyData {
	s.ConflictPreview = v
	return s
}

func (s *UpdateGatewayQuotaRuleResponseBodyData) SetDryRun(v bool) *UpdateGatewayQuotaRuleResponseBodyData {
	s.DryRun = &v
	return s
}

func (s *UpdateGatewayQuotaRuleResponseBodyData) SetRuleId(v string) *UpdateGatewayQuotaRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *UpdateGatewayQuotaRuleResponseBodyData) Validate() error {
	if s.ConflictPreview != nil {
		if err := s.ConflictPreview.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateGatewayQuotaRuleResponseBodyDataConflictPreview struct {
	// example:
	//
	// f8f44dc6cf369a017d56b7197eb4fb5ac4bbb6b09a92b9b41999541f50xxxxxx
	ConflictHash *string                                                       `json:"conflictHash,omitempty" xml:"conflictHash,omitempty"`
	Items        []*UpdateGatewayQuotaRuleResponseBodyDataConflictPreviewItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	TotalConflictCount *int32 `json:"totalConflictCount,omitempty" xml:"totalConflictCount,omitempty"`
}

func (s UpdateGatewayQuotaRuleResponseBodyDataConflictPreview) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayQuotaRuleResponseBodyDataConflictPreview) GoString() string {
	return s.String()
}

func (s *UpdateGatewayQuotaRuleResponseBodyDataConflictPreview) GetConflictHash() *string {
	return s.ConflictHash
}

func (s *UpdateGatewayQuotaRuleResponseBodyDataConflictPreview) GetItems() []*UpdateGatewayQuotaRuleResponseBodyDataConflictPreviewItems {
	return s.Items
}

func (s *UpdateGatewayQuotaRuleResponseBodyDataConflictPreview) GetTotalConflictCount() *int32 {
	return s.TotalConflictCount
}

func (s *UpdateGatewayQuotaRuleResponseBodyDataConflictPreview) SetConflictHash(v string) *UpdateGatewayQuotaRuleResponseBodyDataConflictPreview {
	s.ConflictHash = &v
	return s
}

func (s *UpdateGatewayQuotaRuleResponseBodyDataConflictPreview) SetItems(v []*UpdateGatewayQuotaRuleResponseBodyDataConflictPreviewItems) *UpdateGatewayQuotaRuleResponseBodyDataConflictPreview {
	s.Items = v
	return s
}

func (s *UpdateGatewayQuotaRuleResponseBodyDataConflictPreview) SetTotalConflictCount(v int32) *UpdateGatewayQuotaRuleResponseBodyDataConflictPreview {
	s.TotalConflictCount = &v
	return s
}

func (s *UpdateGatewayQuotaRuleResponseBodyDataConflictPreview) Validate() error {
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

type UpdateGatewayQuotaRuleResponseBodyDataConflictPreviewItems struct {
	// example:
	//
	// cs-d82n1g6m1hkm375xxxxx
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// example:
	//
	// consumer-a
	ConsumerName *string `json:"consumerName,omitempty" xml:"consumerName,omitempty"`
}

func (s UpdateGatewayQuotaRuleResponseBodyDataConflictPreviewItems) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayQuotaRuleResponseBodyDataConflictPreviewItems) GoString() string {
	return s.String()
}

func (s *UpdateGatewayQuotaRuleResponseBodyDataConflictPreviewItems) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *UpdateGatewayQuotaRuleResponseBodyDataConflictPreviewItems) GetConsumerName() *string {
	return s.ConsumerName
}

func (s *UpdateGatewayQuotaRuleResponseBodyDataConflictPreviewItems) SetConsumerId(v string) *UpdateGatewayQuotaRuleResponseBodyDataConflictPreviewItems {
	s.ConsumerId = &v
	return s
}

func (s *UpdateGatewayQuotaRuleResponseBodyDataConflictPreviewItems) SetConsumerName(v string) *UpdateGatewayQuotaRuleResponseBodyDataConflictPreviewItems {
	s.ConsumerName = &v
	return s
}

func (s *UpdateGatewayQuotaRuleResponseBodyDataConflictPreviewItems) Validate() error {
	return dara.Validate(s)
}
