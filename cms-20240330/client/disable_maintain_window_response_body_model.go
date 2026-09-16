// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableMaintainWindowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaintainWindowId(v string) *DisableMaintainWindowResponseBody
	GetMaintainWindowId() *string
	SetRequestId(v string) *DisableMaintainWindowResponseBody
	GetRequestId() *string
}

type DisableMaintainWindowResponseBody struct {
	// The ID of the paused silence policy.
	//
	// example:
	//
	// 3ff3fbd0-8a0b-4b31-9b1c-8e3f0a2c5d71
	MaintainWindowId *string `json:"maintainWindowId,omitempty" xml:"maintainWindowId,omitempty"`
	// The unique ID of the request. You can use this ID for troubleshooting and ticket tracking.
	//
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-A01D6CC3F4B8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DisableMaintainWindowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableMaintainWindowResponseBody) GoString() string {
	return s.String()
}

func (s *DisableMaintainWindowResponseBody) GetMaintainWindowId() *string {
	return s.MaintainWindowId
}

func (s *DisableMaintainWindowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableMaintainWindowResponseBody) SetMaintainWindowId(v string) *DisableMaintainWindowResponseBody {
	s.MaintainWindowId = &v
	return s
}

func (s *DisableMaintainWindowResponseBody) SetRequestId(v string) *DisableMaintainWindowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableMaintainWindowResponseBody) Validate() error {
	return dara.Validate(s)
}
