// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNacosConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateNacosConfigRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *CreateNacosConfigRequest
	GetAppName() *string
	SetBetaIps(v string) *CreateNacosConfigRequest
	GetBetaIps() *string
	SetContent(v string) *CreateNacosConfigRequest
	GetContent() *string
	SetDataId(v string) *CreateNacosConfigRequest
	GetDataId() *string
	SetDesc(v string) *CreateNacosConfigRequest
	GetDesc() *string
	SetGroup(v string) *CreateNacosConfigRequest
	GetGroup() *string
	SetInstanceId(v string) *CreateNacosConfigRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *CreateNacosConfigRequest
	GetNamespaceId() *string
	SetTags(v string) *CreateNacosConfigRequest
	GetTags() *string
	SetType(v string) *CreateNacosConfigRequest
	GetType() *string
}

type CreateNacosConfigRequest struct {
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
	// The name of the application.
	//
	// example:
	//
	// saledatacenter-task
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The list of IP addresses where the beta release of the configuration is performed.
	//
	// example:
	//
	// 100.117.XX.XX,100.117.XX.XX
	BetaIps *string `json:"BetaIps,omitempty" xml:"BetaIps,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the data.
	//
	// This parameter is required.
	//
	// example:
	//
	// common.yaml
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The description of the configuration.
	//
	// example:
	//
	// Basic module configuration.
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The ID of the group.
	//
	// This parameter is required.
	//
	// example:
	//
	// alime-bridge-aliyun
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse_prepaid_public_cn-tl32****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// 547fd2a0-d0d6-****-80db2a1afb82
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The tags of the configuration.
	//
	// example:
	//
	// Basic configurations
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The format of the configuration. Supported formats include TEXT, JSON, and XML.
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateNacosConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNacosConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateNacosConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateNacosConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateNacosConfigRequest) GetBetaIps() *string {
	return s.BetaIps
}

func (s *CreateNacosConfigRequest) GetContent() *string {
	return s.Content
}

func (s *CreateNacosConfigRequest) GetDataId() *string {
	return s.DataId
}

func (s *CreateNacosConfigRequest) GetDesc() *string {
	return s.Desc
}

func (s *CreateNacosConfigRequest) GetGroup() *string {
	return s.Group
}

func (s *CreateNacosConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateNacosConfigRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateNacosConfigRequest) GetTags() *string {
	return s.Tags
}

func (s *CreateNacosConfigRequest) GetType() *string {
	return s.Type
}

func (s *CreateNacosConfigRequest) SetAcceptLanguage(v string) *CreateNacosConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateNacosConfigRequest) SetAppName(v string) *CreateNacosConfigRequest {
	s.AppName = &v
	return s
}

func (s *CreateNacosConfigRequest) SetBetaIps(v string) *CreateNacosConfigRequest {
	s.BetaIps = &v
	return s
}

func (s *CreateNacosConfigRequest) SetContent(v string) *CreateNacosConfigRequest {
	s.Content = &v
	return s
}

func (s *CreateNacosConfigRequest) SetDataId(v string) *CreateNacosConfigRequest {
	s.DataId = &v
	return s
}

func (s *CreateNacosConfigRequest) SetDesc(v string) *CreateNacosConfigRequest {
	s.Desc = &v
	return s
}

func (s *CreateNacosConfigRequest) SetGroup(v string) *CreateNacosConfigRequest {
	s.Group = &v
	return s
}

func (s *CreateNacosConfigRequest) SetInstanceId(v string) *CreateNacosConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateNacosConfigRequest) SetNamespaceId(v string) *CreateNacosConfigRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateNacosConfigRequest) SetTags(v string) *CreateNacosConfigRequest {
	s.Tags = &v
	return s
}

func (s *CreateNacosConfigRequest) SetType(v string) *CreateNacosConfigRequest {
	s.Type = &v
	return s
}

func (s *CreateNacosConfigRequest) Validate() error {
	return dara.Validate(s)
}
