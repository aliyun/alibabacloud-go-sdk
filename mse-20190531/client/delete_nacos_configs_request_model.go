// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNacosConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteNacosConfigsRequest
	GetAcceptLanguage() *string
	SetIds(v string) *DeleteNacosConfigsRequest
	GetIds() *string
	SetInstanceId(v string) *DeleteNacosConfigsRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *DeleteNacosConfigsRequest
	GetNamespaceId() *string
}

type DeleteNacosConfigsRequest struct {
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
	// The IDs of configurations.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20024,20025,20026,20027,20034,20104,20394
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse_prepaid_public_cn-i7m2e32pd0n
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ef93a21-3487-4367-a859-857d8****
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DeleteNacosConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNacosConfigsRequest) GoString() string {
	return s.String()
}

func (s *DeleteNacosConfigsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteNacosConfigsRequest) GetIds() *string {
	return s.Ids
}

func (s *DeleteNacosConfigsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteNacosConfigsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DeleteNacosConfigsRequest) SetAcceptLanguage(v string) *DeleteNacosConfigsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteNacosConfigsRequest) SetIds(v string) *DeleteNacosConfigsRequest {
	s.Ids = &v
	return s
}

func (s *DeleteNacosConfigsRequest) SetInstanceId(v string) *DeleteNacosConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteNacosConfigsRequest) SetNamespaceId(v string) *DeleteNacosConfigsRequest {
	s.NamespaceId = &v
	return s
}

func (s *DeleteNacosConfigsRequest) Validate() error {
	return dara.Validate(s)
}
