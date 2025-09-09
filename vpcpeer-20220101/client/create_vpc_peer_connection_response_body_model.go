// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcPeerConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateVpcPeerConnectionResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *CreateVpcPeerConnectionResponseBody
	GetRequestId() *string
}

type CreateVpcPeerConnectionResponseBody struct {
	// The ID of the instance on which the VPC peering connection is created.
	//
	// example:
	//
	// pcc-lnk0m24khwvtkm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVpcPeerConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcPeerConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcPeerConnectionResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateVpcPeerConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVpcPeerConnectionResponseBody) SetInstanceId(v string) *CreateVpcPeerConnectionResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateVpcPeerConnectionResponseBody) SetRequestId(v string) *CreateVpcPeerConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcPeerConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
