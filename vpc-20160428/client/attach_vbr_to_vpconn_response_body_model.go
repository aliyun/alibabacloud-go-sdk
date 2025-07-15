// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachVbrToVpconnResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachVbrToVpconnResponseBody
	GetRequestId() *string
	SetVirtualPhysicalConnection(v string) *AttachVbrToVpconnResponseBody
	GetVirtualPhysicalConnection() *string
}

type AttachVbrToVpconnResponseBody struct {
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

func (s AttachVbrToVpconnResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachVbrToVpconnResponseBody) GoString() string {
	return s.String()
}

func (s *AttachVbrToVpconnResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachVbrToVpconnResponseBody) GetVirtualPhysicalConnection() *string {
	return s.VirtualPhysicalConnection
}

func (s *AttachVbrToVpconnResponseBody) SetRequestId(v string) *AttachVbrToVpconnResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachVbrToVpconnResponseBody) SetVirtualPhysicalConnection(v string) *AttachVbrToVpconnResponseBody {
	s.VirtualPhysicalConnection = &v
	return s
}

func (s *AttachVbrToVpconnResponseBody) Validate() error {
	return dara.Validate(s)
}
