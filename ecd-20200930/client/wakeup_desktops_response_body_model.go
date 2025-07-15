// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWakeupDesktopsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *WakeupDesktopsResponseBody
	GetRequestId() *string
}

type WakeupDesktopsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6557DBA9-CF3E-5C1B-B1F1-68FDA599****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s WakeupDesktopsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WakeupDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *WakeupDesktopsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WakeupDesktopsResponseBody) SetRequestId(v string) *WakeupDesktopsResponseBody {
	s.RequestId = &v
	return s
}

func (s *WakeupDesktopsResponseBody) Validate() error {
	return dara.Validate(s)
}
