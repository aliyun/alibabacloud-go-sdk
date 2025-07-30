// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchInstanceHAResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchInstanceHAResponseBody
	GetRequestId() *string
}

type SwitchInstanceHAResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5D622714-AEDD-4609-9167-F5DDD3D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchInstanceHAResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchInstanceHAResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchInstanceHAResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchInstanceHAResponseBody) SetRequestId(v string) *SwitchInstanceHAResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchInstanceHAResponseBody) Validate() error {
	return dara.Validate(s)
}
