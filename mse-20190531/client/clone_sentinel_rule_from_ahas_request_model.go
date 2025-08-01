// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneSentinelRuleFromAhasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CloneSentinelRuleFromAhasRequest
	GetAcceptLanguage() *string
	SetAhasNamespace(v string) *CloneSentinelRuleFromAhasRequest
	GetAhasNamespace() *string
	SetAppName(v string) *CloneSentinelRuleFromAhasRequest
	GetAppName() *string
	SetIsAHASPublicRegion(v bool) *CloneSentinelRuleFromAhasRequest
	GetIsAHASPublicRegion() *bool
	SetMseAppName(v string) *CloneSentinelRuleFromAhasRequest
	GetMseAppName() *string
	SetNamespace(v string) *CloneSentinelRuleFromAhasRequest
	GetNamespace() *string
}

type CloneSentinelRuleFromAhasRequest struct {
	// The language in which you want to display the results. Valid values: zh and en. zh indicates Chinese, which is the default value. en indicates English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The namespace (environment) of Application High Availability Service (AHAS).
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	AhasNamespace *string `json:"AhasNamespace,omitempty" xml:"AhasNamespace,omitempty"`
	// The application name.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Specifies whether AHAS is deployed in the Internet region.
	//
	// example:
	//
	// false
	IsAHASPublicRegion *bool `json:"IsAHASPublicRegion,omitempty" xml:"IsAHASPublicRegion,omitempty"`
	// The name of the MSE application after migration. If this parameter is not specified, the name of the Application High Availability Service (AHAS) application is used by default.
	//
	// example:
	//
	// spring-cloud-a
	MseAppName *string `json:"MseAppName,omitempty" xml:"MseAppName,omitempty"`
	// The namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s CloneSentinelRuleFromAhasRequest) String() string {
	return dara.Prettify(s)
}

func (s CloneSentinelRuleFromAhasRequest) GoString() string {
	return s.String()
}

func (s *CloneSentinelRuleFromAhasRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CloneSentinelRuleFromAhasRequest) GetAhasNamespace() *string {
	return s.AhasNamespace
}

func (s *CloneSentinelRuleFromAhasRequest) GetAppName() *string {
	return s.AppName
}

func (s *CloneSentinelRuleFromAhasRequest) GetIsAHASPublicRegion() *bool {
	return s.IsAHASPublicRegion
}

func (s *CloneSentinelRuleFromAhasRequest) GetMseAppName() *string {
	return s.MseAppName
}

func (s *CloneSentinelRuleFromAhasRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CloneSentinelRuleFromAhasRequest) SetAcceptLanguage(v string) *CloneSentinelRuleFromAhasRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CloneSentinelRuleFromAhasRequest) SetAhasNamespace(v string) *CloneSentinelRuleFromAhasRequest {
	s.AhasNamespace = &v
	return s
}

func (s *CloneSentinelRuleFromAhasRequest) SetAppName(v string) *CloneSentinelRuleFromAhasRequest {
	s.AppName = &v
	return s
}

func (s *CloneSentinelRuleFromAhasRequest) SetIsAHASPublicRegion(v bool) *CloneSentinelRuleFromAhasRequest {
	s.IsAHASPublicRegion = &v
	return s
}

func (s *CloneSentinelRuleFromAhasRequest) SetMseAppName(v string) *CloneSentinelRuleFromAhasRequest {
	s.MseAppName = &v
	return s
}

func (s *CloneSentinelRuleFromAhasRequest) SetNamespace(v string) *CloneSentinelRuleFromAhasRequest {
	s.Namespace = &v
	return s
}

func (s *CloneSentinelRuleFromAhasRequest) Validate() error {
	return dara.Validate(s)
}
