// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserDevicesResponseBody
	GetRequestId() *string
}

type DeleteUserDevicesResponseBody struct {
	// example:
	//
	// 5FEF5CFA-14CC-5DE5-BD1F-AFFE0996E71D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserDevicesResponseBody) SetRequestId(v string) *DeleteUserDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}
