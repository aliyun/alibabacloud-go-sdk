// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkId(v string) *CreateNetworkResponseBody
	GetNetworkId() *string
	SetRequestId(v string) *CreateNetworkResponseBody
	GetRequestId() *string
}

type CreateNetworkResponseBody struct {
	// The ID of the network.
	//
	// example:
	//
	// n-5***
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkResponseBody) GetNetworkId() *string {
	return s.NetworkId
}

func (s *CreateNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNetworkResponseBody) SetNetworkId(v string) *CreateNetworkResponseBody {
	s.NetworkId = &v
	return s
}

func (s *CreateNetworkResponseBody) SetRequestId(v string) *CreateNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
