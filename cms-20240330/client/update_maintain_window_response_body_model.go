// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMaintainWindowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaintainWindowId(v string) *UpdateMaintainWindowResponseBody
	GetMaintainWindowId() *string
	SetRequestId(v string) *UpdateMaintainWindowResponseBody
	GetRequestId() *string
}

type UpdateMaintainWindowResponseBody struct {
	// The ID of the updated silence policy.
	//
	// example:
	//
	// 3ff3fbd0-8a0b-4b31-9b1c-8e3f0a2c5d71
	MaintainWindowId *string `json:"maintainWindowId,omitempty" xml:"maintainWindowId,omitempty"`
	// The unique ID of the request. You can use this ID for troubleshooting and ticket tracking.
	//
	// example:
	//
	// 0CEC5375-C554-562B-A65F-9A629907C1F0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateMaintainWindowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMaintainWindowResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMaintainWindowResponseBody) GetMaintainWindowId() *string {
	return s.MaintainWindowId
}

func (s *UpdateMaintainWindowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMaintainWindowResponseBody) SetMaintainWindowId(v string) *UpdateMaintainWindowResponseBody {
	s.MaintainWindowId = &v
	return s
}

func (s *UpdateMaintainWindowResponseBody) SetRequestId(v string) *UpdateMaintainWindowResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMaintainWindowResponseBody) Validate() error {
	return dara.Validate(s)
}
