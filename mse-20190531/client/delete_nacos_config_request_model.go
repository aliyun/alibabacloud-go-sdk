// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNacosConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteNacosConfigRequest
	GetAcceptLanguage() *string
	SetBeta(v bool) *DeleteNacosConfigRequest
	GetBeta() *bool
	SetDataId(v string) *DeleteNacosConfigRequest
	GetDataId() *string
	SetGroup(v string) *DeleteNacosConfigRequest
	GetGroup() *string
	SetInstanceId(v string) *DeleteNacosConfigRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *DeleteNacosConfigRequest
	GetNamespaceId() *string
}

type DeleteNacosConfigRequest struct {
	// Language type of the returned information:
	//
	// - zh: Chinese
	//
	// - en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Whether it is a Beta release. Default is false.
	//
	// - `true`: Yes
	//
	// - `false`: No
	//
	// example:
	//
	// true
	Beta *bool `json:"Beta,omitempty" xml:"Beta,omitempty"`
	// Configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// user-ds.yml
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// Group type.
	//
	// This parameter is required.
	//
	// example:
	//
	// HALVIE_MICRO_GROUP
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse_prepaid_public_cn-tl32epfyu18
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Namespace ID. Default is public.
	//
	// example:
	//
	// 0e9d849b-****-8435da6c21ad
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DeleteNacosConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNacosConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteNacosConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteNacosConfigRequest) GetBeta() *bool {
	return s.Beta
}

func (s *DeleteNacosConfigRequest) GetDataId() *string {
	return s.DataId
}

func (s *DeleteNacosConfigRequest) GetGroup() *string {
	return s.Group
}

func (s *DeleteNacosConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteNacosConfigRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DeleteNacosConfigRequest) SetAcceptLanguage(v string) *DeleteNacosConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteNacosConfigRequest) SetBeta(v bool) *DeleteNacosConfigRequest {
	s.Beta = &v
	return s
}

func (s *DeleteNacosConfigRequest) SetDataId(v string) *DeleteNacosConfigRequest {
	s.DataId = &v
	return s
}

func (s *DeleteNacosConfigRequest) SetGroup(v string) *DeleteNacosConfigRequest {
	s.Group = &v
	return s
}

func (s *DeleteNacosConfigRequest) SetInstanceId(v string) *DeleteNacosConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteNacosConfigRequest) SetNamespaceId(v string) *DeleteNacosConfigRequest {
	s.NamespaceId = &v
	return s
}

func (s *DeleteNacosConfigRequest) Validate() error {
	return dara.Validate(s)
}
