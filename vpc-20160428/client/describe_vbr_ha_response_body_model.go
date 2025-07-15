// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVbrHaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreationTime(v string) *DescribeVbrHaResponseBody
	GetCreationTime() *string
	SetDescription(v string) *DescribeVbrHaResponseBody
	GetDescription() *string
	SetName(v string) *DescribeVbrHaResponseBody
	GetName() *string
	SetPeerVbrId(v string) *DescribeVbrHaResponseBody
	GetPeerVbrId() *string
	SetRegionId(v string) *DescribeVbrHaResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeVbrHaResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeVbrHaResponseBody
	GetStatus() *string
	SetVbrHaId(v string) *DescribeVbrHaResponseBody
	GetVbrHaId() *string
	SetVbrId(v string) *DescribeVbrHaResponseBody
	GetVbrId() *string
}

type DescribeVbrHaResponseBody struct {
	// The time when the VBR was created.
	//
	// example:
	//
	// 2021-06-08T12:20:55
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the VBR failover group.
	//
	// The description must be 2 to 256 characters in length. It must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the VBR failover group.
	//
	// example:
	//
	// VBRHa
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the other VBR in the VBR failover group.
	//
	// example:
	//
	// vbr-bp12mw1f8k3jgygk9****
	PeerVbrId *string `json:"PeerVbrId,omitempty" xml:"PeerVbrId,omitempty"`
	// The ID of the region in which the VBR is deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the VBR failover group.
	//
	// 	- **Creating**
	//
	// 	- **Active**
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the VBR failover group.
	//
	// example:
	//
	// vbrha-sa1sxheuxtd98****
	VbrHaId *string `json:"VbrHaId,omitempty" xml:"VbrHaId,omitempty"`
	// The VBR ID.
	//
	// example:
	//
	// vbr-bp1jcg5cmxjbl9xgc****
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
}

func (s DescribeVbrHaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVbrHaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVbrHaResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeVbrHaResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeVbrHaResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeVbrHaResponseBody) GetPeerVbrId() *string {
	return s.PeerVbrId
}

func (s *DescribeVbrHaResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVbrHaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVbrHaResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeVbrHaResponseBody) GetVbrHaId() *string {
	return s.VbrHaId
}

func (s *DescribeVbrHaResponseBody) GetVbrId() *string {
	return s.VbrId
}

func (s *DescribeVbrHaResponseBody) SetCreationTime(v string) *DescribeVbrHaResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeVbrHaResponseBody) SetDescription(v string) *DescribeVbrHaResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeVbrHaResponseBody) SetName(v string) *DescribeVbrHaResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeVbrHaResponseBody) SetPeerVbrId(v string) *DescribeVbrHaResponseBody {
	s.PeerVbrId = &v
	return s
}

func (s *DescribeVbrHaResponseBody) SetRegionId(v string) *DescribeVbrHaResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeVbrHaResponseBody) SetRequestId(v string) *DescribeVbrHaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVbrHaResponseBody) SetStatus(v string) *DescribeVbrHaResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeVbrHaResponseBody) SetVbrHaId(v string) *DescribeVbrHaResponseBody {
	s.VbrHaId = &v
	return s
}

func (s *DescribeVbrHaResponseBody) SetVbrId(v string) *DescribeVbrHaResponseBody {
	s.VbrId = &v
	return s
}

func (s *DescribeVbrHaResponseBody) Validate() error {
	return dara.Validate(s)
}
