// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStorageDomainRoutingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateStorageDomainRoutingRuleRequest
	GetInstanceId() *string
	SetRoutes(v []*RouteItem) *UpdateStorageDomainRoutingRuleRequest
	GetRoutes() []*RouteItem
	SetRuleId(v string) *UpdateStorageDomainRoutingRuleRequest
	GetRuleId() *string
}

type UpdateStorageDomainRoutingRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Routes []*RouteItem `json:"Routes,omitempty" xml:"Routes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// crsdr-b6thg027zmk1****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s UpdateStorageDomainRoutingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStorageDomainRoutingRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateStorageDomainRoutingRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateStorageDomainRoutingRuleRequest) GetRoutes() []*RouteItem {
	return s.Routes
}

func (s *UpdateStorageDomainRoutingRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateStorageDomainRoutingRuleRequest) SetInstanceId(v string) *UpdateStorageDomainRoutingRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateStorageDomainRoutingRuleRequest) SetRoutes(v []*RouteItem) *UpdateStorageDomainRoutingRuleRequest {
	s.Routes = v
	return s
}

func (s *UpdateStorageDomainRoutingRuleRequest) SetRuleId(v string) *UpdateStorageDomainRoutingRuleRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateStorageDomainRoutingRuleRequest) Validate() error {
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
