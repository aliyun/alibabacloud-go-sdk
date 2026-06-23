// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeIntentionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysis(v bool) *RecognizeIntentionRequest
	GetAnalysis() *bool
	SetBizType(v string) *RecognizeIntentionRequest
	GetBizType() *string
	SetConversation(v string) *RecognizeIntentionRequest
	GetConversation() *string
	SetGlobalIntentionList(v []*RecognizeIntentionRequestGlobalIntentionList) *RecognizeIntentionRequest
	GetGlobalIntentionList() []*RecognizeIntentionRequestGlobalIntentionList
	SetHierarchicalIntentionList(v []*RecognizeIntentionRequestHierarchicalIntentionList) *RecognizeIntentionRequest
	GetHierarchicalIntentionList() []*RecognizeIntentionRequestHierarchicalIntentionList
	SetIntentionDomainCode(v string) *RecognizeIntentionRequest
	GetIntentionDomainCode() *string
	SetIntentionList(v []*RecognizeIntentionRequestIntentionList) *RecognizeIntentionRequest
	GetIntentionList() []*RecognizeIntentionRequestIntentionList
	SetOpType(v string) *RecognizeIntentionRequest
	GetOpType() *string
	SetRecommend(v bool) *RecognizeIntentionRequest
	GetRecommend() *bool
}

type RecognizeIntentionRequest struct {
	// Enable analysis.
	//
	// example:
	//
	// false
	Analysis *bool `json:"analysis,omitempty" xml:"analysis,omitempty"`
	// Business type.
	//
	// This parameter is required.
	//
	// example:
	//
	// common
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// Conversation text.
	//
	// This parameter is required.
	//
	// example:
	//
	// ##客服##：您好，请问是朱杰先生吗？这里是诚信财务的周莉。我们发现您有一项款项昨天是账单日，但您还没还款，这很可能是一个小小的疏忽。来电是提醒您尽快完成还款，避免影响您的信用记录。\\n ##客户##：今天天气怎么样呢？
	Conversation *string `json:"conversation,omitempty" xml:"conversation,omitempty"`
	// Global intent list. Required when opType is hierarchical.
	GlobalIntentionList []*RecognizeIntentionRequestGlobalIntentionList `json:"globalIntentionList,omitempty" xml:"globalIntentionList,omitempty" type:"Repeated"`
	// Hierarchical intent list. Required when opType is hierarchical.
	HierarchicalIntentionList []*RecognizeIntentionRequestHierarchicalIntentionList `json:"hierarchicalIntentionList,omitempty" xml:"hierarchicalIntentionList,omitempty" type:"Repeated"`
	// Intent library: Local intent library code.
	//
	// example:
	//
	// collection
	IntentionDomainCode *string `json:"intentionDomainCode,omitempty" xml:"intentionDomainCode,omitempty"`
	// Intent list. Required when bizType is not attitude.
	IntentionList []*RecognizeIntentionRequestIntentionList `json:"intentionList,omitempty" xml:"intentionList,omitempty" type:"Repeated"`
	// Operation type.
	//
	// example:
	//
	// common
	OpType *string `json:"opType,omitempty" xml:"opType,omitempty"`
	// Recommend intent.
	//
	// example:
	//
	// false
	Recommend *bool `json:"recommend,omitempty" xml:"recommend,omitempty"`
}

func (s RecognizeIntentionRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeIntentionRequest) GoString() string {
	return s.String()
}

func (s *RecognizeIntentionRequest) GetAnalysis() *bool {
	return s.Analysis
}

func (s *RecognizeIntentionRequest) GetBizType() *string {
	return s.BizType
}

func (s *RecognizeIntentionRequest) GetConversation() *string {
	return s.Conversation
}

func (s *RecognizeIntentionRequest) GetGlobalIntentionList() []*RecognizeIntentionRequestGlobalIntentionList {
	return s.GlobalIntentionList
}

func (s *RecognizeIntentionRequest) GetHierarchicalIntentionList() []*RecognizeIntentionRequestHierarchicalIntentionList {
	return s.HierarchicalIntentionList
}

func (s *RecognizeIntentionRequest) GetIntentionDomainCode() *string {
	return s.IntentionDomainCode
}

func (s *RecognizeIntentionRequest) GetIntentionList() []*RecognizeIntentionRequestIntentionList {
	return s.IntentionList
}

func (s *RecognizeIntentionRequest) GetOpType() *string {
	return s.OpType
}

func (s *RecognizeIntentionRequest) GetRecommend() *bool {
	return s.Recommend
}

func (s *RecognizeIntentionRequest) SetAnalysis(v bool) *RecognizeIntentionRequest {
	s.Analysis = &v
	return s
}

func (s *RecognizeIntentionRequest) SetBizType(v string) *RecognizeIntentionRequest {
	s.BizType = &v
	return s
}

func (s *RecognizeIntentionRequest) SetConversation(v string) *RecognizeIntentionRequest {
	s.Conversation = &v
	return s
}

func (s *RecognizeIntentionRequest) SetGlobalIntentionList(v []*RecognizeIntentionRequestGlobalIntentionList) *RecognizeIntentionRequest {
	s.GlobalIntentionList = v
	return s
}

func (s *RecognizeIntentionRequest) SetHierarchicalIntentionList(v []*RecognizeIntentionRequestHierarchicalIntentionList) *RecognizeIntentionRequest {
	s.HierarchicalIntentionList = v
	return s
}

func (s *RecognizeIntentionRequest) SetIntentionDomainCode(v string) *RecognizeIntentionRequest {
	s.IntentionDomainCode = &v
	return s
}

func (s *RecognizeIntentionRequest) SetIntentionList(v []*RecognizeIntentionRequestIntentionList) *RecognizeIntentionRequest {
	s.IntentionList = v
	return s
}

func (s *RecognizeIntentionRequest) SetOpType(v string) *RecognizeIntentionRequest {
	s.OpType = &v
	return s
}

func (s *RecognizeIntentionRequest) SetRecommend(v bool) *RecognizeIntentionRequest {
	s.Recommend = &v
	return s
}

func (s *RecognizeIntentionRequest) Validate() error {
	if s.GlobalIntentionList != nil {
		for _, item := range s.GlobalIntentionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HierarchicalIntentionList != nil {
		for _, item := range s.HierarchicalIntentionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IntentionList != nil {
		for _, item := range s.IntentionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RecognizeIntentionRequestGlobalIntentionList struct {
	// Intent description.
	//
	// example:
	//
	// 正常付款3
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Intent name.
	//
	// example:
	//
	// 正常付款3
	Intention *string `json:"intention,omitempty" xml:"intention,omitempty"`
	// Intent code.
	//
	// example:
	//
	// 1810566978021232640
	IntentionCode *string `json:"intentionCode,omitempty" xml:"intentionCode,omitempty"`
	// Intent script.
	//
	// example:
	//
	// 好的，那先不打扰您了，祝您生活愉快！再见！
	IntentionScript *string `json:"intentionScript,omitempty" xml:"intentionScript,omitempty"`
}

func (s RecognizeIntentionRequestGlobalIntentionList) String() string {
	return dara.Prettify(s)
}

func (s RecognizeIntentionRequestGlobalIntentionList) GoString() string {
	return s.String()
}

func (s *RecognizeIntentionRequestGlobalIntentionList) GetDescription() *string {
	return s.Description
}

func (s *RecognizeIntentionRequestGlobalIntentionList) GetIntention() *string {
	return s.Intention
}

func (s *RecognizeIntentionRequestGlobalIntentionList) GetIntentionCode() *string {
	return s.IntentionCode
}

func (s *RecognizeIntentionRequestGlobalIntentionList) GetIntentionScript() *string {
	return s.IntentionScript
}

func (s *RecognizeIntentionRequestGlobalIntentionList) SetDescription(v string) *RecognizeIntentionRequestGlobalIntentionList {
	s.Description = &v
	return s
}

func (s *RecognizeIntentionRequestGlobalIntentionList) SetIntention(v string) *RecognizeIntentionRequestGlobalIntentionList {
	s.Intention = &v
	return s
}

func (s *RecognizeIntentionRequestGlobalIntentionList) SetIntentionCode(v string) *RecognizeIntentionRequestGlobalIntentionList {
	s.IntentionCode = &v
	return s
}

func (s *RecognizeIntentionRequestGlobalIntentionList) SetIntentionScript(v string) *RecognizeIntentionRequestGlobalIntentionList {
	s.IntentionScript = &v
	return s
}

func (s *RecognizeIntentionRequestGlobalIntentionList) Validate() error {
	return dara.Validate(s)
}

type RecognizeIntentionRequestHierarchicalIntentionList struct {
	// Intent description.
	//
	// example:
	//
	// 询问股票价格
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Intent name.
	//
	// example:
	//
	// 询问股票价格
	Intention *string `json:"intention,omitempty" xml:"intention,omitempty"`
	// Intent code.
	//
	// example:
	//
	// 1810929291010150400
	IntentionCode *string `json:"intentionCode,omitempty" xml:"intentionCode,omitempty"`
	// Intent script.
	//
	// example:
	//
	// 好的，那先不打扰您了，祝您生活愉快！再见！
	IntentionScript *string `json:"intentionScript,omitempty" xml:"intentionScript,omitempty"`
}

func (s RecognizeIntentionRequestHierarchicalIntentionList) String() string {
	return dara.Prettify(s)
}

func (s RecognizeIntentionRequestHierarchicalIntentionList) GoString() string {
	return s.String()
}

func (s *RecognizeIntentionRequestHierarchicalIntentionList) GetDescription() *string {
	return s.Description
}

func (s *RecognizeIntentionRequestHierarchicalIntentionList) GetIntention() *string {
	return s.Intention
}

func (s *RecognizeIntentionRequestHierarchicalIntentionList) GetIntentionCode() *string {
	return s.IntentionCode
}

func (s *RecognizeIntentionRequestHierarchicalIntentionList) GetIntentionScript() *string {
	return s.IntentionScript
}

func (s *RecognizeIntentionRequestHierarchicalIntentionList) SetDescription(v string) *RecognizeIntentionRequestHierarchicalIntentionList {
	s.Description = &v
	return s
}

func (s *RecognizeIntentionRequestHierarchicalIntentionList) SetIntention(v string) *RecognizeIntentionRequestHierarchicalIntentionList {
	s.Intention = &v
	return s
}

func (s *RecognizeIntentionRequestHierarchicalIntentionList) SetIntentionCode(v string) *RecognizeIntentionRequestHierarchicalIntentionList {
	s.IntentionCode = &v
	return s
}

func (s *RecognizeIntentionRequestHierarchicalIntentionList) SetIntentionScript(v string) *RecognizeIntentionRequestHierarchicalIntentionList {
	s.IntentionScript = &v
	return s
}

func (s *RecognizeIntentionRequestHierarchicalIntentionList) Validate() error {
	return dara.Validate(s)
}

type RecognizeIntentionRequestIntentionList struct {
	// Intent description.
	//
	// example:
	//
	// 客户表示忘记还款
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Intent name.
	//
	// example:
	//
	// 客户表示忘记还款
	Intention *string `json:"intention,omitempty" xml:"intention,omitempty"`
	// Intent code.
	//
	// example:
	//
	// 1808766224000262144
	IntentionCode *string `json:"intentionCode,omitempty" xml:"intentionCode,omitempty"`
	// Intent script.
	//
	// example:
	//
	// 好的，那先不打扰您了，祝您生活愉快！再见！
	IntentionScript *string `json:"intentionScript,omitempty" xml:"intentionScript,omitempty"`
}

func (s RecognizeIntentionRequestIntentionList) String() string {
	return dara.Prettify(s)
}

func (s RecognizeIntentionRequestIntentionList) GoString() string {
	return s.String()
}

func (s *RecognizeIntentionRequestIntentionList) GetDescription() *string {
	return s.Description
}

func (s *RecognizeIntentionRequestIntentionList) GetIntention() *string {
	return s.Intention
}

func (s *RecognizeIntentionRequestIntentionList) GetIntentionCode() *string {
	return s.IntentionCode
}

func (s *RecognizeIntentionRequestIntentionList) GetIntentionScript() *string {
	return s.IntentionScript
}

func (s *RecognizeIntentionRequestIntentionList) SetDescription(v string) *RecognizeIntentionRequestIntentionList {
	s.Description = &v
	return s
}

func (s *RecognizeIntentionRequestIntentionList) SetIntention(v string) *RecognizeIntentionRequestIntentionList {
	s.Intention = &v
	return s
}

func (s *RecognizeIntentionRequestIntentionList) SetIntentionCode(v string) *RecognizeIntentionRequestIntentionList {
	s.IntentionCode = &v
	return s
}

func (s *RecognizeIntentionRequestIntentionList) SetIntentionScript(v string) *RecognizeIntentionRequestIntentionList {
	s.IntentionScript = &v
	return s
}

func (s *RecognizeIntentionRequestIntentionList) Validate() error {
	return dara.Validate(s)
}
