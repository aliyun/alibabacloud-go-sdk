// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEngineNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteEngineNamespaceRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *DeleteEngineNamespaceRequest
	GetClusterId() *string
	SetId(v string) *DeleteEngineNamespaceRequest
	GetId() *string
	SetInstanceId(v string) *DeleteEngineNamespaceRequest
	GetInstanceId() *string
}

type DeleteEngineNamespaceRequest struct {
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
	// The ID of the cluster.
	//
	// example:
	//
	// mse-0c738****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// 678ca857-****-b1bf-d0a98c5ca84b
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// mse-cn-7pp2d1****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteEngineNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEngineNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteEngineNamespaceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteEngineNamespaceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteEngineNamespaceRequest) GetId() *string {
	return s.Id
}

func (s *DeleteEngineNamespaceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteEngineNamespaceRequest) SetAcceptLanguage(v string) *DeleteEngineNamespaceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteEngineNamespaceRequest) SetClusterId(v string) *DeleteEngineNamespaceRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteEngineNamespaceRequest) SetId(v string) *DeleteEngineNamespaceRequest {
	s.Id = &v
	return s
}

func (s *DeleteEngineNamespaceRequest) SetInstanceId(v string) *DeleteEngineNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteEngineNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
