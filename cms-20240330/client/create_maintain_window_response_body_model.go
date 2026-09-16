// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMaintainWindowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaintainWindowId(v string) *CreateMaintainWindowResponseBody
	GetMaintainWindowId() *string
	SetRequestId(v string) *CreateMaintainWindowResponseBody
	GetRequestId() *string
}

type CreateMaintainWindowResponseBody struct {
	// The ID of the created silence policy.
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

func (s CreateMaintainWindowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMaintainWindowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMaintainWindowResponseBody) GetMaintainWindowId() *string {
	return s.MaintainWindowId
}

func (s *CreateMaintainWindowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMaintainWindowResponseBody) SetMaintainWindowId(v string) *CreateMaintainWindowResponseBody {
	s.MaintainWindowId = &v
	return s
}

func (s *CreateMaintainWindowResponseBody) SetRequestId(v string) *CreateMaintainWindowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMaintainWindowResponseBody) Validate() error {
	return dara.Validate(s)
}
