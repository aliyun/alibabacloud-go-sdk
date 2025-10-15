// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkZoneDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateNetworkZoneDescriptionResponseBody
	GetRequestId() *string
}

type UpdateNetworkZoneDescriptionResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNetworkZoneDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkZoneDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNetworkZoneDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNetworkZoneDescriptionResponseBody) SetRequestId(v string) *UpdateNetworkZoneDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNetworkZoneDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
