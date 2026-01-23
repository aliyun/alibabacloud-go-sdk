// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStorageDomainRoutingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateStorageDomainRoutingRuleRequest
	GetInstanceId() *string
	SetRoutes(v []*RouteItem) *CreateStorageDomainRoutingRuleRequest
	GetRoutes() []*RouteItem
}

type CreateStorageDomainRoutingRuleRequest struct {
	// The instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The route list
	//
	// This parameter is required.
	Routes []*RouteItem `json:"Routes,omitempty" xml:"Routes,omitempty" type:"Repeated"`
}

func (s CreateStorageDomainRoutingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStorageDomainRoutingRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateStorageDomainRoutingRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateStorageDomainRoutingRuleRequest) GetRoutes() []*RouteItem {
	return s.Routes
}

func (s *CreateStorageDomainRoutingRuleRequest) SetInstanceId(v string) *CreateStorageDomainRoutingRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateStorageDomainRoutingRuleRequest) SetRoutes(v []*RouteItem) *CreateStorageDomainRoutingRuleRequest {
	s.Routes = v
	return s
}

func (s *CreateStorageDomainRoutingRuleRequest) Validate() error {
	if s.Routes != nil {
		for _, item := range s.Routes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
