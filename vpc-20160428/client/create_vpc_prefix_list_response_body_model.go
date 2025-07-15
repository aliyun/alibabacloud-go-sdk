// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcPrefixListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrefixListId(v string) *CreateVpcPrefixListResponseBody
	GetPrefixListId() *string
	SetRequestId(v string) *CreateVpcPrefixListResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateVpcPrefixListResponseBody
	GetResourceGroupId() *string
}

type CreateVpcPrefixListResponseBody struct {
	// The ID of the prefix list.
	//
	// example:
	//
	// pl-0b700s2ee3***
	PrefixListId *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the prefix list belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateVpcPrefixListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcPrefixListResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcPrefixListResponseBody) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *CreateVpcPrefixListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVpcPrefixListResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateVpcPrefixListResponseBody) SetPrefixListId(v string) *CreateVpcPrefixListResponseBody {
	s.PrefixListId = &v
	return s
}

func (s *CreateVpcPrefixListResponseBody) SetRequestId(v string) *CreateVpcPrefixListResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcPrefixListResponseBody) SetResourceGroupId(v string) *CreateVpcPrefixListResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVpcPrefixListResponseBody) Validate() error {
	return dara.Validate(s)
}
