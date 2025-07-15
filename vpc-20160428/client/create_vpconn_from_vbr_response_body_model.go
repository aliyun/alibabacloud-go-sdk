// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpconnFromVbrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVpconnFromVbrResponseBody
	GetRequestId() *string
	SetVirtualPhysicalConnection(v string) *CreateVpconnFromVbrResponseBody
	GetVirtualPhysicalConnection() *string
}

type CreateVpconnFromVbrResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5356F028-0F5C-56FC-8574-897D24379041
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the hosted connection.
	//
	// example:
	//
	// pc-bp1mrgfbtmc9brre7****
	VirtualPhysicalConnection *string `json:"VirtualPhysicalConnection,omitempty" xml:"VirtualPhysicalConnection,omitempty"`
}

func (s CreateVpconnFromVbrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVpconnFromVbrResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpconnFromVbrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVpconnFromVbrResponseBody) GetVirtualPhysicalConnection() *string {
	return s.VirtualPhysicalConnection
}

func (s *CreateVpconnFromVbrResponseBody) SetRequestId(v string) *CreateVpconnFromVbrResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpconnFromVbrResponseBody) SetVirtualPhysicalConnection(v string) *CreateVpconnFromVbrResponseBody {
	s.VirtualPhysicalConnection = &v
	return s
}

func (s *CreateVpconnFromVbrResponseBody) Validate() error {
	return dara.Validate(s)
}
