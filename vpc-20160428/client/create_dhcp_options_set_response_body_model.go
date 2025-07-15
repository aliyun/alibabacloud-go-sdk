// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDhcpOptionsSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDhcpOptionsSetId(v string) *CreateDhcpOptionsSetResponseBody
	GetDhcpOptionsSetId() *string
	SetRequestId(v string) *CreateDhcpOptionsSetResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateDhcpOptionsSetResponseBody
	GetResourceGroupId() *string
}

type CreateDhcpOptionsSetResponseBody struct {
	// The ID of the DHCP options set that is created.
	//
	// example:
	//
	// dopt-o6w0df4epg9zo8isy****
	DhcpOptionsSetId *string `json:"DhcpOptionsSetId,omitempty" xml:"DhcpOptionsSetId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the DHCP options set belongs.
	//
	// example:
	//
	// rg-acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateDhcpOptionsSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDhcpOptionsSetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDhcpOptionsSetResponseBody) GetDhcpOptionsSetId() *string {
	return s.DhcpOptionsSetId
}

func (s *CreateDhcpOptionsSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDhcpOptionsSetResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDhcpOptionsSetResponseBody) SetDhcpOptionsSetId(v string) *CreateDhcpOptionsSetResponseBody {
	s.DhcpOptionsSetId = &v
	return s
}

func (s *CreateDhcpOptionsSetResponseBody) SetRequestId(v string) *CreateDhcpOptionsSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDhcpOptionsSetResponseBody) SetResourceGroupId(v string) *CreateDhcpOptionsSetResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDhcpOptionsSetResponseBody) Validate() error {
	return dara.Validate(s)
}
