// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateNetworkZoneResponseBody
	GetRequestId() *string
}

type UpdateNetworkZoneResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNetworkZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkZoneResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNetworkZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNetworkZoneResponseBody) SetRequestId(v string) *UpdateNetworkZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNetworkZoneResponseBody) Validate() error {
	return dara.Validate(s)
}
