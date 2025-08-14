// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnroutePrivateZoneInCenToVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnroutePrivateZoneInCenToVpcResponseBody
	GetRequestId() *string
}

type UnroutePrivateZoneInCenToVpcResponseBody struct {
	// example:
	//
	// C0245BEF-52AC-44A8-A776-EF96FD26A5CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnroutePrivateZoneInCenToVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnroutePrivateZoneInCenToVpcResponseBody) GoString() string {
	return s.String()
}

func (s *UnroutePrivateZoneInCenToVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnroutePrivateZoneInCenToVpcResponseBody) SetRequestId(v string) *UnroutePrivateZoneInCenToVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnroutePrivateZoneInCenToVpcResponseBody) Validate() error {
	return dara.Validate(s)
}
