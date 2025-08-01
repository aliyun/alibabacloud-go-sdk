// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListListenersByConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListListenersByConfigShrinkRequest
	GetAcceptLanguage() *string
	SetDataId(v string) *ListListenersByConfigShrinkRequest
	GetDataId() *string
	SetExtGrayRulesShrink(v string) *ListListenersByConfigShrinkRequest
	GetExtGrayRulesShrink() *string
	SetGroup(v string) *ListListenersByConfigShrinkRequest
	GetGroup() *string
	SetInstanceId(v string) *ListListenersByConfigShrinkRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *ListListenersByConfigShrinkRequest
	GetNamespaceId() *string
	SetRequestPars(v string) *ListListenersByConfigShrinkRequest
	GetRequestPars() *string
}

type ListListenersByConfigShrinkRequest struct {
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
	DataId             *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	ExtGrayRulesShrink *string `json:"ExtGrayRules,omitempty" xml:"ExtGrayRules,omitempty"`
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

func (s ListListenersByConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListListenersByConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListListenersByConfigShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListListenersByConfigShrinkRequest) GetDataId() *string {
	return s.DataId
}

func (s *ListListenersByConfigShrinkRequest) GetExtGrayRulesShrink() *string {
	return s.ExtGrayRulesShrink
}

func (s *ListListenersByConfigShrinkRequest) GetGroup() *string {
	return s.Group
}

func (s *ListListenersByConfigShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListListenersByConfigShrinkRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListListenersByConfigShrinkRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *ListListenersByConfigShrinkRequest) SetAcceptLanguage(v string) *ListListenersByConfigShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListListenersByConfigShrinkRequest) SetDataId(v string) *ListListenersByConfigShrinkRequest {
	s.DataId = &v
	return s
}

func (s *ListListenersByConfigShrinkRequest) SetExtGrayRulesShrink(v string) *ListListenersByConfigShrinkRequest {
	s.ExtGrayRulesShrink = &v
	return s
}

func (s *ListListenersByConfigShrinkRequest) SetGroup(v string) *ListListenersByConfigShrinkRequest {
	s.Group = &v
	return s
}

func (s *ListListenersByConfigShrinkRequest) SetInstanceId(v string) *ListListenersByConfigShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListListenersByConfigShrinkRequest) SetNamespaceId(v string) *ListListenersByConfigShrinkRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListListenersByConfigShrinkRequest) SetRequestPars(v string) *ListListenersByConfigShrinkRequest {
	s.RequestPars = &v
	return s
}

func (s *ListListenersByConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
