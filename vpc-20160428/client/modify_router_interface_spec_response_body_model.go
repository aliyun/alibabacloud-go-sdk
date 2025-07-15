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
	// 	- **Mini.2**: 2 Mbit/s
	//
	// 	- **Mini.5**: 5 Mbit/s
	//
	// 	- **Small.1**: 10 Mbit/s
	//
	// 	- **Small.2**: 20 Mbit/s
	//
	// 	- **Small.5**: 50 Mbit/s
	//
	// 	- **Middle.1**: 100 Mbit/s
	//
	// 	- **Middle.2**: 200 Mbit/s
	//
	// 	- **Middle.5**: 500 Mbit/s
	//
	// 	- **Large.1**: 1,000 Mbit/s
	//
	// 	- **Large.2**: 2,000 Mbit/s
	//
	// 	- **Large.5**: 5,000 Mbit/s
	//
	// 	- **Xlarge.1**: 10,000 Mbit/s
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
