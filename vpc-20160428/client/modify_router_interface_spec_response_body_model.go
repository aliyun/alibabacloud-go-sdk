// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRouterInterfaceSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRouterInterfaceSpecResponseBody
	GetRequestId() *string
	SetSpec(v string) *ModifyRouterInterfaceSpecResponseBody
	GetSpec() *string
}

type ModifyRouterInterfaceSpecResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The specification of the router interface. Valid values:
	//
	// 	- **Mini.2**: 2 Mbps
	//
	// 	- **Mini.5**: 5 Mbps
	//
	// 	- **Small.1**: 10 Mbps
	//
	// 	- **Small.2**: 20 Mbps
	//
	// 	- **Small.5**: 50 Mbps
	//
	// 	- **Middle.1**: 100 Mbps
	//
	// 	- **Middle.2**: 200 Mbps
	//
	// 	- **Middle.5**: 500 Mbps
	//
	// 	- **Large.1**: 1000 Mbps
	//
	// 	- **Large.2**: 2000 Mbps
	//
	// 	- **Large.5**: 5000 Mbps
	//
	// 	- **Xlarge.1**: 10000 Mbps.
	//
	// example:
	//
	// Small.1
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s ModifyRouterInterfaceSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRouterInterfaceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRouterInterfaceSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRouterInterfaceSpecResponseBody) GetSpec() *string {
	return s.Spec
}

func (s *ModifyRouterInterfaceSpecResponseBody) SetRequestId(v string) *ModifyRouterInterfaceSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRouterInterfaceSpecResponseBody) SetSpec(v string) *ModifyRouterInterfaceSpecResponseBody {
	s.Spec = &v
	return s
}

func (s *ModifyRouterInterfaceSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
