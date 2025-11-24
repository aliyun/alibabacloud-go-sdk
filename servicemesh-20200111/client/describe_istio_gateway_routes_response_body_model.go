// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIstioGatewayRoutesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetManagementRoutes(v []*DescribeIstioGatewayRoutesResponseBodyManagementRoutes) *DescribeIstioGatewayRoutesResponseBody
	GetManagementRoutes() []*DescribeIstioGatewayRoutesResponseBodyManagementRoutes
	SetRequestId(v string) *DescribeIstioGatewayRoutesResponseBody
	GetRequestId() *string
}

type DescribeIstioGatewayRoutesResponseBody struct {
	// The routing rules.
	ManagementRoutes []*DescribeIstioGatewayRoutesResponseBodyManagementRoutes `json:"ManagementRoutes,omitempty" xml:"ManagementRoutes,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 31d3a0f0-07ed-4f6e-9004-1804498c****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIstioGatewayRoutesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRoutesResponseBody) GetManagementRoutes() []*DescribeIstioGatewayRoutesResponseBodyManagementRoutes {
	return s.ManagementRoutes
}

func (s *DescribeIstioGatewayRoutesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIstioGatewayRoutesResponseBody) SetManagementRoutes(v []*DescribeIstioGatewayRoutesResponseBodyManagementRoutes) *DescribeIstioGatewayRoutesResponseBody {
	s.ManagementRoutes = v
	return s
}

func (s *DescribeIstioGatewayRoutesResponseBody) SetRequestId(v string) *DescribeIstioGatewayRoutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIstioGatewayRoutesResponseBody) Validate() error {
	if s.ManagementRoutes != nil {
		for _, item := range s.ManagementRoutes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIstioGatewayRoutesResponseBodyManagementRoutes struct {
	// The name of the ASM gateway.
	//
	// example:
	//
	// ingressgateway
	ASMGatewayName *string `json:"ASMGatewayName,omitempty" xml:"ASMGatewayName,omitempty"`
	// The description of the routing rule.
	//
	// example:
	//
	// demo route
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Destination hosts list.
	DestinationHost []*string `json:"DestinationHost,omitempty" xml:"DestinationHost,omitempty" type:"Repeated"`
	// Destination subset list.
	DestinationSubSet []*string `json:"DestinationSubSet,omitempty" xml:"DestinationSubSet,omitempty" type:"Repeated"`
	// The namespace.
	//
	// example:
	//
	// istio-demo
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The priority of the routing rule. The value of this parameter is an integer. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The name of the routing rule.
	//
	// example:
	//
	// http-route
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// The path that is used to match request URLs.
	//
	// example:
	//
	// /reviews/v1
	RoutePath *string `json:"RoutePath,omitempty" xml:"RoutePath,omitempty"`
	// The status of the routing rule. Valid values:
	//
	// 	- `0`: The routing rule is valid.
	//
	// 	- `1`: The routing rule is invalid.
	//
	// 	- `2`: An error occurs during the creation or update of the routing rule.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeIstioGatewayRoutesResponseBodyManagementRoutes) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRoutesResponseBodyManagementRoutes) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) GetASMGatewayName() *string {
	return s.ASMGatewayName
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) GetDescription() *string {
	return s.Description
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) GetDestinationHost() []*string {
	return s.DestinationHost
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) GetDestinationSubSet() []*string {
	return s.DestinationSubSet
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) GetRouteName() *string {
	return s.RouteName
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) GetRoutePath() *string {
	return s.RoutePath
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) SetASMGatewayName(v string) *DescribeIstioGatewayRoutesResponseBodyManagementRoutes {
	s.ASMGatewayName = &v
	return s
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) SetDescription(v string) *DescribeIstioGatewayRoutesResponseBodyManagementRoutes {
	s.Description = &v
	return s
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) SetDestinationHost(v []*string) *DescribeIstioGatewayRoutesResponseBodyManagementRoutes {
	s.DestinationHost = v
	return s
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) SetDestinationSubSet(v []*string) *DescribeIstioGatewayRoutesResponseBodyManagementRoutes {
	s.DestinationSubSet = v
	return s
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) SetNamespace(v string) *DescribeIstioGatewayRoutesResponseBodyManagementRoutes {
	s.Namespace = &v
	return s
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) SetPriority(v int32) *DescribeIstioGatewayRoutesResponseBodyManagementRoutes {
	s.Priority = &v
	return s
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) SetRouteName(v string) *DescribeIstioGatewayRoutesResponseBodyManagementRoutes {
	s.RouteName = &v
	return s
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) SetRoutePath(v string) *DescribeIstioGatewayRoutesResponseBodyManagementRoutes {
	s.RoutePath = &v
	return s
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) SetStatus(v int32) *DescribeIstioGatewayRoutesResponseBodyManagementRoutes {
	s.Status = &v
	return s
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) Validate() error {
	return dara.Validate(s)
}
