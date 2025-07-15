// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirtualBorderRouterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVirtualBorderRouterResponseBody
	GetRequestId() *string
	SetVbrId(v string) *CreateVirtualBorderRouterResponseBody
	GetVbrId() *string
}

type CreateVirtualBorderRouterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the VBR.
	//
	// example:
	//
	// vbr-bp1jcg5cmxjbl9xgc****
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
}

func (s CreateVirtualBorderRouterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualBorderRouterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVirtualBorderRouterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVirtualBorderRouterResponseBody) GetVbrId() *string {
	return s.VbrId
}

func (s *CreateVirtualBorderRouterResponseBody) SetRequestId(v string) *CreateVirtualBorderRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVirtualBorderRouterResponseBody) SetVbrId(v string) *CreateVirtualBorderRouterResponseBody {
	s.VbrId = &v
	return s
}

func (s *CreateVirtualBorderRouterResponseBody) Validate() error {
	return dara.Validate(s)
}
