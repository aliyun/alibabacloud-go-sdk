// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpnConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *CreateVpnConnectionResponseBody
	GetCreateTime() *int64
	SetName(v string) *CreateVpnConnectionResponseBody
	GetName() *string
	SetRequestId(v string) *CreateVpnConnectionResponseBody
	GetRequestId() *string
	SetVpnConnectionId(v string) *CreateVpnConnectionResponseBody
	GetVpnConnectionId() *string
}

type CreateVpnConnectionResponseBody struct {
	// The timestamp generated when the IPsec-VPN connection was established. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1544666102000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the IPsec-VPN connection.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 082AD562-B8DB-4BB2-861F-DA1FCA01FD76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the IPsec-VPN connection.
	//
	// example:
	//
	// vco-bp15oes1py4i6****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
}

func (s CreateVpnConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpnConnectionResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateVpnConnectionResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateVpnConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVpnConnectionResponseBody) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *CreateVpnConnectionResponseBody) SetCreateTime(v int64) *CreateVpnConnectionResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateVpnConnectionResponseBody) SetName(v string) *CreateVpnConnectionResponseBody {
	s.Name = &v
	return s
}

func (s *CreateVpnConnectionResponseBody) SetRequestId(v string) *CreateVpnConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpnConnectionResponseBody) SetVpnConnectionId(v string) *CreateVpnConnectionResponseBody {
	s.VpnConnectionId = &v
	return s
}

func (s *CreateVpnConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
