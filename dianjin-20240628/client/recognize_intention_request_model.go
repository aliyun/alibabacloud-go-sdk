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
	// example:
	//
	// false
	Analysis *bool `json:"analysis,omitempty" xml:"analysis,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// common
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	Conversation              *string                                               `json:"conversation,omitempty" xml:"conversation,omitempty"`
	GlobalIntentionList       []*RecognizeIntentionRequestGlobalIntentionList       `json:"globalIntentionList,omitempty" xml:"globalIntentionList,omitempty" type:"Repeated"`
	HierarchicalIntentionList []*RecognizeIntentionRequestHierarchicalIntentionList `json:"hierarchicalIntentionList,omitempty" xml:"hierarchicalIntentionList,omitempty" type:"Repeated"`
	IntentionDomainCode       *string                                               `json:"intentionDomainCode,omitempty" xml:"intentionDomainCode,omitempty"`
	IntentionList             []*RecognizeIntentionRequestIntentionList             `json:"intentionList,omitempty" xml:"intentionList,omitempty" type:"Repeated"`
	// example:
	//
	// common
	OpType *string `json:"opType,omitempty" xml:"opType,omitempty"`
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
	return dara.Validate(s)
}

type RecognizeIntentionRequestGlobalIntentionList struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Intention   *string `json:"intention,omitempty" xml:"intention,omitempty"`
	// example:
	//
	// 1810566978021232640
	IntentionCode   *string `json:"intentionCode,omitempty" xml:"intentionCode,omitempty"`
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
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Intention   *string `json:"intention,omitempty" xml:"intention,omitempty"`
	// example:
	//
	// 1810929291010150400
	IntentionCode   *string `json:"intentionCode,omitempty" xml:"intentionCode,omitempty"`
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
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Intention   *string `json:"intention,omitempty" xml:"intention,omitempty"`
	// example:
	//
	// 1808766224000262144
	IntentionCode   *string `json:"intentionCode,omitempty" xml:"intentionCode,omitempty"`
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
