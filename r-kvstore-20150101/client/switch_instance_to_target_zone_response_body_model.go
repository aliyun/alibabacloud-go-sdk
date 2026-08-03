// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchInstanceToTargetZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchInstanceToTargetZoneResponseBody
	GetRequestId() *string
}

type SwitchInstanceToTargetZoneResponseBody struct {
	// example:
	//
	// 5D622714-AEDD-4609-9167-F5DDD3D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchInstanceToTargetZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchInstanceToTargetZoneResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchInstanceToTargetZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchInstanceToTargetZoneResponseBody) SetRequestId(v string) *SwitchInstanceToTargetZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchInstanceToTargetZoneResponseBody) Validate() error {
	return dara.Validate(s)
}
