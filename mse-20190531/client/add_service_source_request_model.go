// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddServiceSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddServiceSourceRequest
	GetAcceptLanguage() *string
	SetAddress(v string) *AddServiceSourceRequest
	GetAddress() *string
	SetGatewayUniqueId(v string) *AddServiceSourceRequest
	GetGatewayUniqueId() *string
	SetGroupList(v []*string) *AddServiceSourceRequest
	GetGroupList() []*string
	SetIngressOptionsRequest(v *AddServiceSourceRequestIngressOptionsRequest) *AddServiceSourceRequest
	GetIngressOptionsRequest() *AddServiceSourceRequestIngressOptionsRequest
	SetName(v string) *AddServiceSourceRequest
	GetName() *string
	SetPathList(v []*string) *AddServiceSourceRequest
	GetPathList() []*string
	SetSource(v string) *AddServiceSourceRequest
	GetSource() *string
	SetToAuthorizeSecurityGroups(v []*AddServiceSourceRequestToAuthorizeSecurityGroups) *AddServiceSourceRequest
	GetToAuthorizeSecurityGroups() []*AddServiceSourceRequestToAuthorizeSecurityGroups
	SetType(v string) *AddServiceSourceRequest
	GetType() *string
}

type AddServiceSourceRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh-CN (default): Chinese
	//
	// 	- en-US: English
	//
	// 	- ja: Japanese
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Specifies whether to monitor Ingress classes.
	//
	// example:
	//
	// c9ad2a0717032427e920754e25b49e3b5
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// Specifies whether to update the Ingress status.
	//
	// example:
	//
	// gw-c70622ff52fe49beb29bea9a6f52****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The data structure.
	GroupList []*string `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
	// The list of service groups.
	IngressOptionsRequest *AddServiceSourceRequestIngressOptionsRequest `json:"IngressOptionsRequest,omitempty" xml:"IngressOptionsRequest,omitempty" type:"Struct"`
	// The namespace whose resources you want to monitor.
	//
	// example:
	//
	// istio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The HTTP status code returned.
	PathList []*string `json:"PathList,omitempty" xml:"PathList,omitempty" type:"Repeated"`
	// The service source.
	//
	// 	- K8s: ACK cluster
	//
	// 	- NACOS: MSE Nacos instance
	//
	// example:
	//
	// K8s,MSE
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The list of security groups to be authorized. You can specify security groups to allow backend services to access data sources that you create.
	ToAuthorizeSecurityGroups []*AddServiceSourceRequestToAuthorizeSecurityGroups `json:"ToAuthorizeSecurityGroups,omitempty" xml:"ToAuthorizeSecurityGroups,omitempty" type:"Repeated"`
	// The type of the service source.
	//
	// 	- K8s: Container Service for Kubernetes (ACK) cluster
	//
	// 	- NACOS: Nacos instance
	//
	// example:
	//
	// The Ingress configuration.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddServiceSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddServiceSourceRequest) GoString() string {
	return s.String()
}

func (s *AddServiceSourceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddServiceSourceRequest) GetAddress() *string {
	return s.Address
}

func (s *AddServiceSourceRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *AddServiceSourceRequest) GetGroupList() []*string {
	return s.GroupList
}

func (s *AddServiceSourceRequest) GetIngressOptionsRequest() *AddServiceSourceRequestIngressOptionsRequest {
	return s.IngressOptionsRequest
}

func (s *AddServiceSourceRequest) GetName() *string {
	return s.Name
}

func (s *AddServiceSourceRequest) GetPathList() []*string {
	return s.PathList
}

func (s *AddServiceSourceRequest) GetSource() *string {
	return s.Source
}

func (s *AddServiceSourceRequest) GetToAuthorizeSecurityGroups() []*AddServiceSourceRequestToAuthorizeSecurityGroups {
	return s.ToAuthorizeSecurityGroups
}

func (s *AddServiceSourceRequest) GetType() *string {
	return s.Type
}

func (s *AddServiceSourceRequest) SetAcceptLanguage(v string) *AddServiceSourceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddServiceSourceRequest) SetAddress(v string) *AddServiceSourceRequest {
	s.Address = &v
	return s
}

func (s *AddServiceSourceRequest) SetGatewayUniqueId(v string) *AddServiceSourceRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *AddServiceSourceRequest) SetGroupList(v []*string) *AddServiceSourceRequest {
	s.GroupList = v
	return s
}

func (s *AddServiceSourceRequest) SetIngressOptionsRequest(v *AddServiceSourceRequestIngressOptionsRequest) *AddServiceSourceRequest {
	s.IngressOptionsRequest = v
	return s
}

func (s *AddServiceSourceRequest) SetName(v string) *AddServiceSourceRequest {
	s.Name = &v
	return s
}

func (s *AddServiceSourceRequest) SetPathList(v []*string) *AddServiceSourceRequest {
	s.PathList = v
	return s
}

func (s *AddServiceSourceRequest) SetSource(v string) *AddServiceSourceRequest {
	s.Source = &v
	return s
}

func (s *AddServiceSourceRequest) SetToAuthorizeSecurityGroups(v []*AddServiceSourceRequestToAuthorizeSecurityGroups) *AddServiceSourceRequest {
	s.ToAuthorizeSecurityGroups = v
	return s
}

func (s *AddServiceSourceRequest) SetType(v string) *AddServiceSourceRequest {
	s.Type = &v
	return s
}

func (s *AddServiceSourceRequest) Validate() error {
	if s.IngressOptionsRequest != nil {
		if err := s.IngressOptionsRequest.Validate(); err != nil {
			return err
		}
	}
	if s.ToAuthorizeSecurityGroups != nil {
		for _, item := range s.ToAuthorizeSecurityGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddServiceSourceRequestIngressOptionsRequest struct {
	// The group to which the service belongs.
	//
	// example:
	//
	// true
	EnableIngress *bool `json:"EnableIngress,omitempty" xml:"EnableIngress,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- zh-CN: Chinese. This is the default value.
	//
	// 	- en-US: English.
	//
	// 	- ja: Japanese.
	//
	// example:
	//
	// true
	EnableStatus *bool `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// An array of service root paths.
	//
	// example:
	//
	// com.test.xxx
	IngressClass *string `json:"IngressClass,omitempty" xml:"IngressClass,omitempty"`
	// The root path of the service.
	//
	// example:
	//
	// default
	WatchNamespace *string `json:"WatchNamespace,omitempty" xml:"WatchNamespace,omitempty"`
}

func (s AddServiceSourceRequestIngressOptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddServiceSourceRequestIngressOptionsRequest) GoString() string {
	return s.String()
}

func (s *AddServiceSourceRequestIngressOptionsRequest) GetEnableIngress() *bool {
	return s.EnableIngress
}

func (s *AddServiceSourceRequestIngressOptionsRequest) GetEnableStatus() *bool {
	return s.EnableStatus
}

func (s *AddServiceSourceRequestIngressOptionsRequest) GetIngressClass() *string {
	return s.IngressClass
}

func (s *AddServiceSourceRequestIngressOptionsRequest) GetWatchNamespace() *string {
	return s.WatchNamespace
}

func (s *AddServiceSourceRequestIngressOptionsRequest) SetEnableIngress(v bool) *AddServiceSourceRequestIngressOptionsRequest {
	s.EnableIngress = &v
	return s
}

func (s *AddServiceSourceRequestIngressOptionsRequest) SetEnableStatus(v bool) *AddServiceSourceRequestIngressOptionsRequest {
	s.EnableStatus = &v
	return s
}

func (s *AddServiceSourceRequestIngressOptionsRequest) SetIngressClass(v string) *AddServiceSourceRequestIngressOptionsRequest {
	s.IngressClass = &v
	return s
}

func (s *AddServiceSourceRequestIngressOptionsRequest) SetWatchNamespace(v string) *AddServiceSourceRequestIngressOptionsRequest {
	s.WatchNamespace = &v
	return s
}

func (s *AddServiceSourceRequestIngressOptionsRequest) Validate() error {
	return dara.Validate(s)
}

type AddServiceSourceRequestToAuthorizeSecurityGroups struct {
	// The description of the authorization record.
	//
	// example:
	//
	// rule for xxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The authorized port range of the security group. You can select multiple port ranges. Separate each port range with a comma (,).
	//
	// example:
	//
	// 8080/8080,9000/10000
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-bp1acepclex0vmi1****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s AddServiceSourceRequestToAuthorizeSecurityGroups) String() string {
	return dara.Prettify(s)
}

func (s AddServiceSourceRequestToAuthorizeSecurityGroups) GoString() string {
	return s.String()
}

func (s *AddServiceSourceRequestToAuthorizeSecurityGroups) GetDescription() *string {
	return s.Description
}

func (s *AddServiceSourceRequestToAuthorizeSecurityGroups) GetPortRange() *string {
	return s.PortRange
}

func (s *AddServiceSourceRequestToAuthorizeSecurityGroups) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *AddServiceSourceRequestToAuthorizeSecurityGroups) SetDescription(v string) *AddServiceSourceRequestToAuthorizeSecurityGroups {
	s.Description = &v
	return s
}

func (s *AddServiceSourceRequestToAuthorizeSecurityGroups) SetPortRange(v string) *AddServiceSourceRequestToAuthorizeSecurityGroups {
	s.PortRange = &v
	return s
}

func (s *AddServiceSourceRequestToAuthorizeSecurityGroups) SetSecurityGroupId(v string) *AddServiceSourceRequestToAuthorizeSecurityGroups {
	s.SecurityGroupId = &v
	return s
}

func (s *AddServiceSourceRequestToAuthorizeSecurityGroups) Validate() error {
	return dara.Validate(s)
}
