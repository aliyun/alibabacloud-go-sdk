// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkPeerConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteNetworkPeerConnectionRequest
	GetInstanceId() *string
}

type DeleteNetworkPeerConnectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pcc-5rr19av7tkpgi9os52ag1****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteNetworkPeerConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkPeerConnectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkPeerConnectionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteNetworkPeerConnectionRequest) SetInstanceId(v string) *DeleteNetworkPeerConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteNetworkPeerConnectionRequest) Validate() error {
	return dara.Validate(s)
}
