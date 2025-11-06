// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListListenersByConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListListenersByConfigRequest
	GetAcceptLanguage() *string
	SetDataId(v string) *ListListenersByConfigRequest
	GetDataId() *string
	SetExtGrayRules(v []*ListListenersByConfigRequestExtGrayRules) *ListListenersByConfigRequest
	GetExtGrayRules() []*ListListenersByConfigRequestExtGrayRules
	SetGroup(v string) *ListListenersByConfigRequest
	GetGroup() *string
	SetInstanceId(v string) *ListListenersByConfigRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *ListListenersByConfigRequest
	GetNamespaceId() *string
	SetRequestPars(v string) *ListListenersByConfigRequest
	GetRequestPars() *string
}

type ListListenersByConfigRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the data.
	//
	// This parameter is required.
	//
	// example:
	//
	// zeekr-clueboss.yml
	DataId       *string                                     `json:"DataId,omitempty" xml:"DataId,omitempty"`
	ExtGrayRules []*ListListenersByConfigRequestExtGrayRules `json:"ExtGrayRules,omitempty" xml:"ExtGrayRules,omitempty" type:"Repeated"`
	// The name of the group.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse-cn-m7r1yurp00e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// aaeb4d28-c9eb-4fa2-85f5-d03ce7ee8df1
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The extended request parameters in the JSON format.
	//
	// example:
	//
	// {\\\\"appGroup\\\\":\\\\"emas-zfive_prehost\\\\",\\\\"appName\\\\":\\\\"emas-zfive\\\\",\\\\"appStage\\\\":\\\\"PRE_PUBLISH\\\\",\\\\"appUnit\\\\":\\\\"\\\\",\\\\"bucId\\\\":\\\\"225902\\\\",\\\\"bucName\\\\":\\\\"Wireless\\\\",\\\\"provider\\\\":\\\\"aliyun\\\\"}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s ListListenersByConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ListListenersByConfigRequest) GoString() string {
	return s.String()
}

func (s *ListListenersByConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListListenersByConfigRequest) GetDataId() *string {
	return s.DataId
}

func (s *ListListenersByConfigRequest) GetExtGrayRules() []*ListListenersByConfigRequestExtGrayRules {
	return s.ExtGrayRules
}

func (s *ListListenersByConfigRequest) GetGroup() *string {
	return s.Group
}

func (s *ListListenersByConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListListenersByConfigRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListListenersByConfigRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *ListListenersByConfigRequest) SetAcceptLanguage(v string) *ListListenersByConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListListenersByConfigRequest) SetDataId(v string) *ListListenersByConfigRequest {
	s.DataId = &v
	return s
}

func (s *ListListenersByConfigRequest) SetExtGrayRules(v []*ListListenersByConfigRequestExtGrayRules) *ListListenersByConfigRequest {
	s.ExtGrayRules = v
	return s
}

func (s *ListListenersByConfigRequest) SetGroup(v string) *ListListenersByConfigRequest {
	s.Group = &v
	return s
}

func (s *ListListenersByConfigRequest) SetInstanceId(v string) *ListListenersByConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ListListenersByConfigRequest) SetNamespaceId(v string) *ListListenersByConfigRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListListenersByConfigRequest) SetRequestPars(v string) *ListListenersByConfigRequest {
	s.RequestPars = &v
	return s
}

func (s *ListListenersByConfigRequest) Validate() error {
	if s.ExtGrayRules != nil {
		for _, item := range s.ExtGrayRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListListenersByConfigRequestExtGrayRules struct {
	GrayRule         *string `json:"GrayRule,omitempty" xml:"GrayRule,omitempty"`
	GrayRuleName     *string `json:"GrayRuleName,omitempty" xml:"GrayRuleName,omitempty"`
	GrayRulePriority *int32  `json:"GrayRulePriority,omitempty" xml:"GrayRulePriority,omitempty"`
	GrayRuleType     *string `json:"GrayRuleType,omitempty" xml:"GrayRuleType,omitempty"`
}

func (s ListListenersByConfigRequestExtGrayRules) String() string {
	return dara.Prettify(s)
}

func (s ListListenersByConfigRequestExtGrayRules) GoString() string {
	return s.String()
}

func (s *ListListenersByConfigRequestExtGrayRules) GetGrayRule() *string {
	return s.GrayRule
}

func (s *ListListenersByConfigRequestExtGrayRules) GetGrayRuleName() *string {
	return s.GrayRuleName
}

func (s *ListListenersByConfigRequestExtGrayRules) GetGrayRulePriority() *int32 {
	return s.GrayRulePriority
}

func (s *ListListenersByConfigRequestExtGrayRules) GetGrayRuleType() *string {
	return s.GrayRuleType
}

func (s *ListListenersByConfigRequestExtGrayRules) SetGrayRule(v string) *ListListenersByConfigRequestExtGrayRules {
	s.GrayRule = &v
	return s
}

func (s *ListListenersByConfigRequestExtGrayRules) SetGrayRuleName(v string) *ListListenersByConfigRequestExtGrayRules {
	s.GrayRuleName = &v
	return s
}

func (s *ListListenersByConfigRequestExtGrayRules) SetGrayRulePriority(v int32) *ListListenersByConfigRequestExtGrayRules {
	s.GrayRulePriority = &v
	return s
}

func (s *ListListenersByConfigRequestExtGrayRules) SetGrayRuleType(v string) *ListListenersByConfigRequestExtGrayRules {
	s.GrayRuleType = &v
	return s
}

func (s *ListListenersByConfigRequestExtGrayRules) Validate() error {
	return dara.Validate(s)
}
