// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkInterfaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkInterfaceIds(v []*string) *CreateNetworkInterfaceResponseBody
	GetNetworkInterfaceIds() []*string
	SetRequestId(v string) *CreateNetworkInterfaceResponseBody
	GetRequestId() *string
}

type CreateNetworkInterfaceResponseBody struct {
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitempty" xml:"NetworkInterfaceIds,omitempty" type:"Repeated"`
	// example:
	//
	// F3B261DD-3858-4D3C-877D-303ADF374600
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNetworkInterfaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceResponseBody) GetNetworkInterfaceIds() []*string {
	return s.NetworkInterfaceIds
}

func (s *CreateNetworkInterfaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNetworkInterfaceResponseBody) SetNetworkInterfaceIds(v []*string) *CreateNetworkInterfaceResponseBody {
	s.NetworkInterfaceIds = v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) SetRequestId(v string) *CreateNetworkInterfaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) Validate() error {
	return dara.Validate(s)
}
