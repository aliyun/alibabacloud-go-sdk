// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchInstanceZoneFailOverResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchInstanceZoneFailOverResponseBody
	GetRequestId() *string
}

type SwitchInstanceZoneFailOverResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2D9F3768-EDA9-4811-943E-42C8006E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchInstanceZoneFailOverResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchInstanceZoneFailOverResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchInstanceZoneFailOverResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchInstanceZoneFailOverResponseBody) SetRequestId(v string) *SwitchInstanceZoneFailOverResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchInstanceZoneFailOverResponseBody) Validate() error {
	return dara.Validate(s)
}
