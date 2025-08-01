// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNacosConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetNacosConfigRequest
	GetAcceptLanguage() *string
	SetBeta(v bool) *GetNacosConfigRequest
	GetBeta() *bool
	SetDataId(v string) *GetNacosConfigRequest
	GetDataId() *string
	SetGroup(v string) *GetNacosConfigRequest
	GetGroup() *string
	SetInstanceId(v string) *GetNacosConfigRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *GetNacosConfigRequest
	GetNamespaceId() *string
}

type GetNacosConfigRequest struct {
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
	// Whether it is a Beta release.
	//
	// - `true`: Yes
	//
	// - `false`: No
	//
	// example:
	//
	// true
	Beta *bool `json:"Beta,omitempty" xml:"Beta,omitempty"`
	// Data ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// halvie-mp-item****
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// Configuration group information.
	//
	// This parameter is required.
	//
	// example:
	//
	// common
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse-cn-i7m2h0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Namespace ID.
	//
	// example:
	//
	// ddaf8f12-****-b1c1-86e7c72e266b
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s GetNacosConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNacosConfigRequest) GoString() string {
	return s.String()
}

func (s *GetNacosConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetNacosConfigRequest) GetBeta() *bool {
	return s.Beta
}

func (s *GetNacosConfigRequest) GetDataId() *string {
	return s.DataId
}

func (s *GetNacosConfigRequest) GetGroup() *string {
	return s.Group
}

func (s *GetNacosConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetNacosConfigRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetNacosConfigRequest) SetAcceptLanguage(v string) *GetNacosConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetNacosConfigRequest) SetBeta(v bool) *GetNacosConfigRequest {
	s.Beta = &v
	return s
}

func (s *GetNacosConfigRequest) SetDataId(v string) *GetNacosConfigRequest {
	s.DataId = &v
	return s
}

func (s *GetNacosConfigRequest) SetGroup(v string) *GetNacosConfigRequest {
	s.Group = &v
	return s
}

func (s *GetNacosConfigRequest) SetInstanceId(v string) *GetNacosConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNacosConfigRequest) SetNamespaceId(v string) *GetNacosConfigRequest {
	s.NamespaceId = &v
	return s
}

func (s *GetNacosConfigRequest) Validate() error {
	return dara.Validate(s)
}
