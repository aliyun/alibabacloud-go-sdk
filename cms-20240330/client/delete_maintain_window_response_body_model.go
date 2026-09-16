// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaintainWindowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaintainWindowId(v string) *DeleteMaintainWindowResponseBody
	GetMaintainWindowId() *string
	SetRequestId(v string) *DeleteMaintainWindowResponseBody
	GetRequestId() *string
}

type DeleteMaintainWindowResponseBody struct {
	// The ID of the deleted silence policy.
	//
	// example:
	//
	// 3ff3fbd0-8a0b-4b31-9b1c-8e3f0a2c5d71
	MaintainWindowId *string `json:"maintainWindowId,omitempty" xml:"maintainWindowId,omitempty"`
	// The unique ID of the request. You can use this ID for troubleshooting and ticket tracking.
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A0D1C36
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteMaintainWindowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaintainWindowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMaintainWindowResponseBody) GetMaintainWindowId() *string {
	return s.MaintainWindowId
}

func (s *DeleteMaintainWindowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMaintainWindowResponseBody) SetMaintainWindowId(v string) *DeleteMaintainWindowResponseBody {
	s.MaintainWindowId = &v
	return s
}

func (s *DeleteMaintainWindowResponseBody) SetRequestId(v string) *DeleteMaintainWindowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMaintainWindowResponseBody) Validate() error {
	return dara.Validate(s)
}
