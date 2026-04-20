// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGrantInstanceToTransitRouterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyGrantInstanceToTransitRouterResponseBody
	GetRequestId() *string
}

type ModifyGrantInstanceToTransitRouterResponseBody struct {
	// example:
	//
	// F7DDDC17-FA06-4AC2-8F35-59D2470FCFC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyGrantInstanceToTransitRouterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyGrantInstanceToTransitRouterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGrantInstanceToTransitRouterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyGrantInstanceToTransitRouterResponseBody) SetRequestId(v string) *ModifyGrantInstanceToTransitRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGrantInstanceToTransitRouterResponseBody) Validate() error {
	return dara.Validate(s)
}
