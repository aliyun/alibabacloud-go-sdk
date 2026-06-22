// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleObjectScanEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchType(v string) *HandleObjectScanEventRequest
	GetBatchType() *string
	SetEventId(v string) *HandleObjectScanEventRequest
	GetEventId() *string
	SetEventIdList(v []*int64) *HandleObjectScanEventRequest
	GetEventIdList() []*int64
	SetLang(v string) *HandleObjectScanEventRequest
	GetLang() *string
	SetRemark(v string) *HandleObjectScanEventRequest
	GetRemark() *string
	SetRuleConditionList(v []*HandleObjectScanEventRequestRuleConditionList) *HandleObjectScanEventRequest
	GetRuleConditionList() []*HandleObjectScanEventRequestRuleConditionList
	SetStatus(v int32) *HandleObjectScanEventRequest
	GetStatus() *int32
}

type HandleObjectScanEventRequest struct {
	// Specifies the type for batch processing of similar alerts. Valid values:
	//
	// - **sha256**: by file content
	//
	// - **eventName**: by alert name.
	//
	// example:
	//
	// sha256
	BatchType *string `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 81****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The list of event IDs.
	EventIdList []*int64 `json:"EventIdList,omitempty" xml:"EventIdList,omitempty" type:"Repeated"`
	// The language of the content in the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The remarks.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The list of whitelist rules. This parameter takes effect only when the alert is whitelisted.
	RuleConditionList []*HandleObjectScanEventRequestRuleConditionList `json:"RuleConditionList,omitempty" xml:"RuleConditionList,omitempty" type:"Repeated"`
	// The target status. Valid values:
	//
	// - **0**: Unhandled.
	//
	// - **1**: Manually handled.
	//
	// - **2**: Whitelisted.
	//
	// - **3**: Ignored.
	//
	// - **4**: Access denied.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s HandleObjectScanEventRequest) String() string {
	return dara.Prettify(s)
}

func (s HandleObjectScanEventRequest) GoString() string {
	return s.String()
}

func (s *HandleObjectScanEventRequest) GetBatchType() *string {
	return s.BatchType
}

func (s *HandleObjectScanEventRequest) GetEventId() *string {
	return s.EventId
}

func (s *HandleObjectScanEventRequest) GetEventIdList() []*int64 {
	return s.EventIdList
}

func (s *HandleObjectScanEventRequest) GetLang() *string {
	return s.Lang
}

func (s *HandleObjectScanEventRequest) GetRemark() *string {
	return s.Remark
}

func (s *HandleObjectScanEventRequest) GetRuleConditionList() []*HandleObjectScanEventRequestRuleConditionList {
	return s.RuleConditionList
}

func (s *HandleObjectScanEventRequest) GetStatus() *int32 {
	return s.Status
}

func (s *HandleObjectScanEventRequest) SetBatchType(v string) *HandleObjectScanEventRequest {
	s.BatchType = &v
	return s
}

func (s *HandleObjectScanEventRequest) SetEventId(v string) *HandleObjectScanEventRequest {
	s.EventId = &v
	return s
}

func (s *HandleObjectScanEventRequest) SetEventIdList(v []*int64) *HandleObjectScanEventRequest {
	s.EventIdList = v
	return s
}

func (s *HandleObjectScanEventRequest) SetLang(v string) *HandleObjectScanEventRequest {
	s.Lang = &v
	return s
}

func (s *HandleObjectScanEventRequest) SetRemark(v string) *HandleObjectScanEventRequest {
	s.Remark = &v
	return s
}

func (s *HandleObjectScanEventRequest) SetRuleConditionList(v []*HandleObjectScanEventRequestRuleConditionList) *HandleObjectScanEventRequest {
	s.RuleConditionList = v
	return s
}

func (s *HandleObjectScanEventRequest) SetStatus(v int32) *HandleObjectScanEventRequest {
	s.Status = &v
	return s
}

func (s *HandleObjectScanEventRequest) Validate() error {
	if s.RuleConditionList != nil {
		for _, item := range s.RuleConditionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type HandleObjectScanEventRequestRuleConditionList struct {
	// The whitelist field. Valid values:
	//
	// - **ossKey**: file path
	//
	// - **bucketName**: bucket name
	//
	// - **md5**: file MD5
	//
	// - **sha256**: file SHA-256.
	//
	// example:
	//
	// ossKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The operator. Valid values:
	//
	// - **contains**: Contains.
	//
	// - **not_contains**: Does not contain.
	//
	// - **str_equal**: Equals.
	//
	// - **str_not_equal**: Does not equal.
	//
	// - **regex**: Regular expression.
	//
	// example:
	//
	// contains
	Operate *string `json:"Operate,omitempty" xml:"Operate,omitempty"`
	// The value to match.
	//
	// example:
	//
	// sshe
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s HandleObjectScanEventRequestRuleConditionList) String() string {
	return dara.Prettify(s)
}

func (s HandleObjectScanEventRequestRuleConditionList) GoString() string {
	return s.String()
}

func (s *HandleObjectScanEventRequestRuleConditionList) GetKey() *string {
	return s.Key
}

func (s *HandleObjectScanEventRequestRuleConditionList) GetOperate() *string {
	return s.Operate
}

func (s *HandleObjectScanEventRequestRuleConditionList) GetValue() *string {
	return s.Value
}

func (s *HandleObjectScanEventRequestRuleConditionList) SetKey(v string) *HandleObjectScanEventRequestRuleConditionList {
	s.Key = &v
	return s
}

func (s *HandleObjectScanEventRequestRuleConditionList) SetOperate(v string) *HandleObjectScanEventRequestRuleConditionList {
	s.Operate = &v
	return s
}

func (s *HandleObjectScanEventRequestRuleConditionList) SetValue(v string) *HandleObjectScanEventRequestRuleConditionList {
	s.Value = &v
	return s
}

func (s *HandleObjectScanEventRequestRuleConditionList) Validate() error {
	return dara.Validate(s)
}
