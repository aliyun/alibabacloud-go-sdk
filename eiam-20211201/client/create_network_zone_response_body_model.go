// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkZoneId(v string) *CreateNetworkZoneResponseBody
	GetNetworkZoneId() *string
	SetRequestId(v string) *CreateNetworkZoneResponseBody
	GetRequestId() *string
}

type CreateNetworkZoneResponseBody struct {
	// example:
	//
	// network_m6a57vre4g3h7m725yrq6pxxxx
	NetworkZoneId *string `json:"NetworkZoneId,omitempty" xml:"NetworkZoneId,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNetworkZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkZoneResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkZoneResponseBody) GetNetworkZoneId() *string {
	return s.NetworkZoneId
}

func (s *CreateNetworkZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNetworkZoneResponseBody) SetNetworkZoneId(v string) *CreateNetworkZoneResponseBody {
	s.NetworkZoneId = &v
	return s
}

func (s *CreateNetworkZoneResponseBody) SetRequestId(v string) *CreateNetworkZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkZoneResponseBody) Validate() error {
	return dara.Validate(s)
}
