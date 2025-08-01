// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPullServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *PullServicesRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *PullServicesRequest
	GetGatewayUniqueId() *string
	SetNamespace(v string) *PullServicesRequest
	GetNamespace() *string
	SetSourceId(v int64) *PullServicesRequest
	GetSourceId() *int64
	SetSourceType(v string) *PullServicesRequest
	GetSourceType() *string
}

type PullServicesRequest struct {
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
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-eccf313e2224438ba53d95d039e5****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The namespace.
	//
	// example:
	//
	// public
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// 100
	SourceId *int64 `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The type of the service source.
	//
	// example:
	//
	// K8s
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s PullServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s PullServicesRequest) GoString() string {
	return s.String()
}

func (s *PullServicesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *PullServicesRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *PullServicesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *PullServicesRequest) GetSourceId() *int64 {
	return s.SourceId
}

func (s *PullServicesRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *PullServicesRequest) SetAcceptLanguage(v string) *PullServicesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *PullServicesRequest) SetGatewayUniqueId(v string) *PullServicesRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *PullServicesRequest) SetNamespace(v string) *PullServicesRequest {
	s.Namespace = &v
	return s
}

func (s *PullServicesRequest) SetSourceId(v int64) *PullServicesRequest {
	s.SourceId = &v
	return s
}

func (s *PullServicesRequest) SetSourceType(v string) *PullServicesRequest {
	s.SourceType = &v
	return s
}

func (s *PullServicesRequest) Validate() error {
	return dara.Validate(s)
}
