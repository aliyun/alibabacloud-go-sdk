// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudConnectNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCcnId(v string) *CreateCloudConnectNetworkResponseBody
	GetCcnId() *string
	SetCidrBlock(v string) *CreateCloudConnectNetworkResponseBody
	GetCidrBlock() *string
	SetDescription(v string) *CreateCloudConnectNetworkResponseBody
	GetDescription() *string
	SetName(v string) *CreateCloudConnectNetworkResponseBody
	GetName() *string
	SetRequestId(v string) *CreateCloudConnectNetworkResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateCloudConnectNetworkResponseBody
	GetResourceGroupId() *string
	SetSnatCidrBlock(v string) *CreateCloudConnectNetworkResponseBody
	GetSnatCidrBlock() *string
	SetStatus(v string) *CreateCloudConnectNetworkResponseBody
	GetStatus() *string
}

type CreateCloudConnectNetworkResponseBody struct {
	// The ID of the CCN instance.
	//
	// example:
	//
	// ccn-l9340rlu5ens*****
	CcnId *string `json:"CcnId,omitempty" xml:"CcnId,omitempty"`
	// The private CIDR block.
	//
	// example:
	//
	// 172.XX.XX.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The description about the CCN instance.
	//
	// example:
	//
	// ccndesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the CCN instance.
	//
	// example:
	//
	// ccnname
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C48E8EB2-37A4-495B-A95C-29CA1FD26C82
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the CCN instance belongs.
	//
	// example:
	//
	// rg-acfm2iu4fnc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The private CIDR block used for SNAT.
	//
	// example:
	//
	// 172.XX.XX.0/25
	SnatCidrBlock *string `json:"SnatCidrBlock,omitempty" xml:"SnatCidrBlock,omitempty"`
	// The status of the CCN instance. Valid values:
	//
	// 	- **Active**: indicates the CNN instance running as expected.
	//
	// 	- **Pending**: indicates the CCN instance is to be created.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateCloudConnectNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudConnectNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudConnectNetworkResponseBody) GetCcnId() *string {
	return s.CcnId
}

func (s *CreateCloudConnectNetworkResponseBody) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *CreateCloudConnectNetworkResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateCloudConnectNetworkResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateCloudConnectNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloudConnectNetworkResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateCloudConnectNetworkResponseBody) GetSnatCidrBlock() *string {
	return s.SnatCidrBlock
}

func (s *CreateCloudConnectNetworkResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateCloudConnectNetworkResponseBody) SetCcnId(v string) *CreateCloudConnectNetworkResponseBody {
	s.CcnId = &v
	return s
}

func (s *CreateCloudConnectNetworkResponseBody) SetCidrBlock(v string) *CreateCloudConnectNetworkResponseBody {
	s.CidrBlock = &v
	return s
}

func (s *CreateCloudConnectNetworkResponseBody) SetDescription(v string) *CreateCloudConnectNetworkResponseBody {
	s.Description = &v
	return s
}

func (s *CreateCloudConnectNetworkResponseBody) SetName(v string) *CreateCloudConnectNetworkResponseBody {
	s.Name = &v
	return s
}

func (s *CreateCloudConnectNetworkResponseBody) SetRequestId(v string) *CreateCloudConnectNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudConnectNetworkResponseBody) SetResourceGroupId(v string) *CreateCloudConnectNetworkResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateCloudConnectNetworkResponseBody) SetSnatCidrBlock(v string) *CreateCloudConnectNetworkResponseBody {
	s.SnatCidrBlock = &v
	return s
}

func (s *CreateCloudConnectNetworkResponseBody) SetStatus(v string) *CreateCloudConnectNetworkResponseBody {
	s.Status = &v
	return s
}

func (s *CreateCloudConnectNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
