// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateServiceSourceRequest
	GetAcceptLanguage() *string
	SetAddress(v string) *UpdateServiceSourceRequest
	GetAddress() *string
	SetGatewayId(v int64) *UpdateServiceSourceRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *UpdateServiceSourceRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *UpdateServiceSourceRequest
	GetId() *int64
	SetIngressOptionsRequest(v *UpdateServiceSourceRequestIngressOptionsRequest) *UpdateServiceSourceRequest
	GetIngressOptionsRequest() *UpdateServiceSourceRequestIngressOptionsRequest
	SetName(v string) *UpdateServiceSourceRequest
	GetName() *string
	SetPathList(v []*string) *UpdateServiceSourceRequest
	GetPathList() []*string
	SetSource(v string) *UpdateServiceSourceRequest
	GetSource() *string
	SetType(v string) *UpdateServiceSourceRequest
	GetType() *string
}

type UpdateServiceSourceRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese.
	//
	// 	- en: English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The address.
	//
	// example:
	//
	// c9ad2a0717032427e920754e25b49e3b5
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The ID of the gateway.
	//
	// example:
	//
	// 429
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-cdd1bb9bfb8341e9805f931a3ba1f4c6
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the service source.
	//
	// example:
	//
	// 63
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The configurations of Ingress resources.
	//
	// example:
	//
	// zh
	IngressOptionsRequest *UpdateServiceSourceRequestIngressOptionsRequest `json:"IngressOptionsRequest,omitempty" xml:"IngressOptionsRequest,omitempty" type:"Struct"`
	// The name.
	//
	// example:
	//
	// istio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// An array of service root paths.
	PathList []*string `json:"PathList,omitempty" xml:"PathList,omitempty" type:"Repeated"`
	// The service source. Valid values:
	//
	// 	- K8s: ACK cluster.
	//
	// 	- MSE: Nacos instance.
	//
	// example:
	//
	// K8s
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the service source. Valid values:
	//
	// 	- K8s: ACK cluster.
	//
	// 	- NACOS: Nacos instance.
	//
	// example:
	//
	// K8s
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateServiceSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceSourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceSourceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateServiceSourceRequest) GetAddress() *string {
	return s.Address
}

func (s *UpdateServiceSourceRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateServiceSourceRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateServiceSourceRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateServiceSourceRequest) GetIngressOptionsRequest() *UpdateServiceSourceRequestIngressOptionsRequest {
	return s.IngressOptionsRequest
}

func (s *UpdateServiceSourceRequest) GetName() *string {
	return s.Name
}

func (s *UpdateServiceSourceRequest) GetPathList() []*string {
	return s.PathList
}

func (s *UpdateServiceSourceRequest) GetSource() *string {
	return s.Source
}

func (s *UpdateServiceSourceRequest) GetType() *string {
	return s.Type
}

func (s *UpdateServiceSourceRequest) SetAcceptLanguage(v string) *UpdateServiceSourceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateServiceSourceRequest) SetAddress(v string) *UpdateServiceSourceRequest {
	s.Address = &v
	return s
}

func (s *UpdateServiceSourceRequest) SetGatewayId(v int64) *UpdateServiceSourceRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateServiceSourceRequest) SetGatewayUniqueId(v string) *UpdateServiceSourceRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateServiceSourceRequest) SetId(v int64) *UpdateServiceSourceRequest {
	s.Id = &v
	return s
}

func (s *UpdateServiceSourceRequest) SetIngressOptionsRequest(v *UpdateServiceSourceRequestIngressOptionsRequest) *UpdateServiceSourceRequest {
	s.IngressOptionsRequest = v
	return s
}

func (s *UpdateServiceSourceRequest) SetName(v string) *UpdateServiceSourceRequest {
	s.Name = &v
	return s
}

func (s *UpdateServiceSourceRequest) SetPathList(v []*string) *UpdateServiceSourceRequest {
	s.PathList = v
	return s
}

func (s *UpdateServiceSourceRequest) SetSource(v string) *UpdateServiceSourceRequest {
	s.Source = &v
	return s
}

func (s *UpdateServiceSourceRequest) SetType(v string) *UpdateServiceSourceRequest {
	s.Type = &v
	return s
}

func (s *UpdateServiceSourceRequest) Validate() error {
	if s.IngressOptionsRequest != nil {
		if err := s.IngressOptionsRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateServiceSourceRequestIngressOptionsRequest struct {
	// Specifies whether to enable Ingress.
	//
	// example:
	//
	// true
	EnableIngress *bool `json:"EnableIngress,omitempty" xml:"EnableIngress,omitempty"`
	// Specifies whether to update the Ingress status.
	//
	// example:
	//
	// true
	EnableStatus *bool `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// Specifies whether to monitor Ingress classes.
	//
	// example:
	//
	// com.test.xxx
	IngressClass *string `json:"IngressClass,omitempty" xml:"IngressClass,omitempty"`
	// The namespace whose resources you want to monitor.
	//
	// example:
	//
	// default
	WatchNamespace *string `json:"WatchNamespace,omitempty" xml:"WatchNamespace,omitempty"`
}

func (s UpdateServiceSourceRequestIngressOptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceSourceRequestIngressOptionsRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceSourceRequestIngressOptionsRequest) GetEnableIngress() *bool {
	return s.EnableIngress
}

func (s *UpdateServiceSourceRequestIngressOptionsRequest) GetEnableStatus() *bool {
	return s.EnableStatus
}

func (s *UpdateServiceSourceRequestIngressOptionsRequest) GetIngressClass() *string {
	return s.IngressClass
}

func (s *UpdateServiceSourceRequestIngressOptionsRequest) GetWatchNamespace() *string {
	return s.WatchNamespace
}

func (s *UpdateServiceSourceRequestIngressOptionsRequest) SetEnableIngress(v bool) *UpdateServiceSourceRequestIngressOptionsRequest {
	s.EnableIngress = &v
	return s
}

func (s *UpdateServiceSourceRequestIngressOptionsRequest) SetEnableStatus(v bool) *UpdateServiceSourceRequestIngressOptionsRequest {
	s.EnableStatus = &v
	return s
}

func (s *UpdateServiceSourceRequestIngressOptionsRequest) SetIngressClass(v string) *UpdateServiceSourceRequestIngressOptionsRequest {
	s.IngressClass = &v
	return s
}

func (s *UpdateServiceSourceRequestIngressOptionsRequest) SetWatchNamespace(v string) *UpdateServiceSourceRequestIngressOptionsRequest {
	s.WatchNamespace = &v
	return s
}

func (s *UpdateServiceSourceRequestIngressOptionsRequest) Validate() error {
	return dara.Validate(s)
}
