// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneNacosConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CloneNacosConfigRequest
	GetAcceptLanguage() *string
	SetDataIds(v string) *CloneNacosConfigRequest
	GetDataIds() *string
	SetIds(v string) *CloneNacosConfigRequest
	GetIds() *string
	SetInstanceId(v string) *CloneNacosConfigRequest
	GetInstanceId() *string
	SetOriginNamespaceId(v string) *CloneNacosConfigRequest
	GetOriginNamespaceId() *string
	SetPolicy(v string) *CloneNacosConfigRequest
	GetPolicy() *string
	SetTargetNamespaceId(v string) *CloneNacosConfigRequest
	GetTargetNamespaceId() *string
}

type CloneNacosConfigRequest struct {
	// Language type of the returned message:
	//
	// - zh: Chinese
	//
	// - en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Configuration items to be cloned, in the format of dataId+group, with multiple items separated by commas.
	//
	// example:
	//
	// test+test,test1+test1
	DataIds *string `json:"DataIds,omitempty" xml:"DataIds,omitempty"`
	// Deprecated
	//
	// List of configuration IDs.
	//
	// example:
	//
	// 253661,253662
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse_prepaid_public_cn-i7m25igg403
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Source namespace ID.
	//
	// example:
	//
	// be821963-6d48-4ea5-9910-6057d****
	OriginNamespaceId *string `json:"OriginNamespaceId,omitempty" xml:"OriginNamespaceId,omitempty"`
	// The strategy used when a write conflict occurs.
	//
	// - ABORT
	//
	// - SKIP
	//
	// - OVERWRITE
	//
	// This parameter is required.
	//
	// example:
	//
	// OVERWRITE
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// Target namespace ID.
	//
	// example:
	//
	// 08be4b5d-2d1c-4e6e-aa85-83b9****
	TargetNamespaceId *string `json:"TargetNamespaceId,omitempty" xml:"TargetNamespaceId,omitempty"`
}

func (s CloneNacosConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CloneNacosConfigRequest) GoString() string {
	return s.String()
}

func (s *CloneNacosConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CloneNacosConfigRequest) GetDataIds() *string {
	return s.DataIds
}

func (s *CloneNacosConfigRequest) GetIds() *string {
	return s.Ids
}

func (s *CloneNacosConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CloneNacosConfigRequest) GetOriginNamespaceId() *string {
	return s.OriginNamespaceId
}

func (s *CloneNacosConfigRequest) GetPolicy() *string {
	return s.Policy
}

func (s *CloneNacosConfigRequest) GetTargetNamespaceId() *string {
	return s.TargetNamespaceId
}

func (s *CloneNacosConfigRequest) SetAcceptLanguage(v string) *CloneNacosConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CloneNacosConfigRequest) SetDataIds(v string) *CloneNacosConfigRequest {
	s.DataIds = &v
	return s
}

func (s *CloneNacosConfigRequest) SetIds(v string) *CloneNacosConfigRequest {
	s.Ids = &v
	return s
}

func (s *CloneNacosConfigRequest) SetInstanceId(v string) *CloneNacosConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *CloneNacosConfigRequest) SetOriginNamespaceId(v string) *CloneNacosConfigRequest {
	s.OriginNamespaceId = &v
	return s
}

func (s *CloneNacosConfigRequest) SetPolicy(v string) *CloneNacosConfigRequest {
	s.Policy = &v
	return s
}

func (s *CloneNacosConfigRequest) SetTargetNamespaceId(v string) *CloneNacosConfigRequest {
	s.TargetNamespaceId = &v
	return s
}

func (s *CloneNacosConfigRequest) Validate() error {
	return dara.Validate(s)
}
