// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNacosHistoryConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetNacosHistoryConfigRequest
	GetAcceptLanguage() *string
	SetDataId(v string) *GetNacosHistoryConfigRequest
	GetDataId() *string
	SetGroup(v string) *GetNacosHistoryConfigRequest
	GetGroup() *string
	SetInstanceId(v string) *GetNacosHistoryConfigRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *GetNacosHistoryConfigRequest
	GetNamespaceId() *string
	SetNid(v string) *GetNacosHistoryConfigRequest
	GetNid() *string
}

type GetNacosHistoryConfigRequest struct {
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
	// msg-center.main.yaml
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The name of the group.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEFAULT_GROUP
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse_prepaid_public_cn-st220g9ka02
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// 6cf708a5-****-89f2-3ba62c5ee9ba
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The version ID of the configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 40****
	Nid *string `json:"Nid,omitempty" xml:"Nid,omitempty"`
}

func (s GetNacosHistoryConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNacosHistoryConfigRequest) GoString() string {
	return s.String()
}

func (s *GetNacosHistoryConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetNacosHistoryConfigRequest) GetDataId() *string {
	return s.DataId
}

func (s *GetNacosHistoryConfigRequest) GetGroup() *string {
	return s.Group
}

func (s *GetNacosHistoryConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetNacosHistoryConfigRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetNacosHistoryConfigRequest) GetNid() *string {
	return s.Nid
}

func (s *GetNacosHistoryConfigRequest) SetAcceptLanguage(v string) *GetNacosHistoryConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetNacosHistoryConfigRequest) SetDataId(v string) *GetNacosHistoryConfigRequest {
	s.DataId = &v
	return s
}

func (s *GetNacosHistoryConfigRequest) SetGroup(v string) *GetNacosHistoryConfigRequest {
	s.Group = &v
	return s
}

func (s *GetNacosHistoryConfigRequest) SetInstanceId(v string) *GetNacosHistoryConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNacosHistoryConfigRequest) SetNamespaceId(v string) *GetNacosHistoryConfigRequest {
	s.NamespaceId = &v
	return s
}

func (s *GetNacosHistoryConfigRequest) SetNid(v string) *GetNacosHistoryConfigRequest {
	s.Nid = &v
	return s
}

func (s *GetNacosHistoryConfigRequest) Validate() error {
	return dara.Validate(s)
}
