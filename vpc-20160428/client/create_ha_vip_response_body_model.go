// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHaVipResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHaVipId(v string) *CreateHaVipResponseBody
	GetHaVipId() *string
	SetIpAddress(v string) *CreateHaVipResponseBody
	GetIpAddress() *string
	SetRequestId(v string) *CreateHaVipResponseBody
	GetRequestId() *string
}

type CreateHaVipResponseBody struct {
	// The ID of the HaVip.
	//
	// example:
	//
	// havip-2zeo05qre24nhrqpy****
	HaVipId *string `json:"HaVipId,omitempty" xml:"HaVipId,omitempty"`
	// The IP address of the HaVip.
	//
	// example:
	//
	// 192.XX.XX.10
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C44F62BE-9CE7-4277-B117-69243F3988BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHaVipResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHaVipResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHaVipResponseBody) GetHaVipId() *string {
	return s.HaVipId
}

func (s *CreateHaVipResponseBody) GetIpAddress() *string {
	return s.IpAddress
}

func (s *CreateHaVipResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHaVipResponseBody) SetHaVipId(v string) *CreateHaVipResponseBody {
	s.HaVipId = &v
	return s
}

func (s *CreateHaVipResponseBody) SetIpAddress(v string) *CreateHaVipResponseBody {
	s.IpAddress = &v
	return s
}

func (s *CreateHaVipResponseBody) SetRequestId(v string) *CreateHaVipResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHaVipResponseBody) Validate() error {
	return dara.Validate(s)
}
