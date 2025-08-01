// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNacosConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateNacosConfigRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *UpdateNacosConfigRequest
	GetAppName() *string
	SetBetaIps(v string) *UpdateNacosConfigRequest
	GetBetaIps() *string
	SetContent(v string) *UpdateNacosConfigRequest
	GetContent() *string
	SetDataId(v string) *UpdateNacosConfigRequest
	GetDataId() *string
	SetDesc(v string) *UpdateNacosConfigRequest
	GetDesc() *string
	SetEncryptedDataKey(v string) *UpdateNacosConfigRequest
	GetEncryptedDataKey() *string
	SetGroup(v string) *UpdateNacosConfigRequest
	GetGroup() *string
	SetInstanceId(v string) *UpdateNacosConfigRequest
	GetInstanceId() *string
	SetMd5(v string) *UpdateNacosConfigRequest
	GetMd5() *string
	SetNamespaceId(v string) *UpdateNacosConfigRequest
	GetNamespaceId() *string
	SetTags(v string) *UpdateNacosConfigRequest
	GetTags() *string
	SetType(v string) *UpdateNacosConfigRequest
	GetType() *string
}

type UpdateNacosConfigRequest struct {
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
	// postoffice
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The list of IP addresses where the beta release of the configuration is performed.
	//
	// example:
	//
	// 196.168.XX.XX
	BetaIps *string `json:"BetaIps,omitempty" xml:"BetaIps,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// ky-check-****.yml
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The description of the configuration.
	//
	// example:
	//
	// Basic configurations
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The encryption key.
	//
	// example:
	//
	// 122wdwe****
	EncryptedDataKey *string `json:"EncryptedDataKey,omitempty" xml:"EncryptedDataKey,omitempty"`
	// The name of the group.
	//
	// This parameter is required.
	//
	// example:
	//
	// resource
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse-cn-7pp2a****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The MD5 value of the configuration.
	//
	// example:
	//
	// 045439703a273a94306422b****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// 78b7af66-d15f-4541-b886-11ed81ecb6c0
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The list of tags.
	//
	// example:
	//
	// 2021-10-20
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The format of the configuration. Supported formats include TEXT, JSON, XML, and HTML.
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateNacosConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNacosConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateNacosConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateNacosConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateNacosConfigRequest) GetBetaIps() *string {
	return s.BetaIps
}

func (s *UpdateNacosConfigRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateNacosConfigRequest) GetDataId() *string {
	return s.DataId
}

func (s *UpdateNacosConfigRequest) GetDesc() *string {
	return s.Desc
}

func (s *UpdateNacosConfigRequest) GetEncryptedDataKey() *string {
	return s.EncryptedDataKey
}

func (s *UpdateNacosConfigRequest) GetGroup() *string {
	return s.Group
}

func (s *UpdateNacosConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateNacosConfigRequest) GetMd5() *string {
	return s.Md5
}

func (s *UpdateNacosConfigRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateNacosConfigRequest) GetTags() *string {
	return s.Tags
}

func (s *UpdateNacosConfigRequest) GetType() *string {
	return s.Type
}

func (s *UpdateNacosConfigRequest) SetAcceptLanguage(v string) *UpdateNacosConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetAppName(v string) *UpdateNacosConfigRequest {
	s.AppName = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetBetaIps(v string) *UpdateNacosConfigRequest {
	s.BetaIps = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetContent(v string) *UpdateNacosConfigRequest {
	s.Content = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetDataId(v string) *UpdateNacosConfigRequest {
	s.DataId = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetDesc(v string) *UpdateNacosConfigRequest {
	s.Desc = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetEncryptedDataKey(v string) *UpdateNacosConfigRequest {
	s.EncryptedDataKey = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetGroup(v string) *UpdateNacosConfigRequest {
	s.Group = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetInstanceId(v string) *UpdateNacosConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetMd5(v string) *UpdateNacosConfigRequest {
	s.Md5 = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetNamespaceId(v string) *UpdateNacosConfigRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetTags(v string) *UpdateNacosConfigRequest {
	s.Tags = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetType(v string) *UpdateNacosConfigRequest {
	s.Type = &v
	return s
}

func (s *UpdateNacosConfigRequest) Validate() error {
	return dara.Validate(s)
}
