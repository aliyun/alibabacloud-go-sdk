// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeUserDeviceSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeUserDeviceSessionResponseBody
	GetRequestId() *string
}

type RevokeUserDeviceSessionResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 102350E7-1A20-58F5-9D63-ABEA820AE6E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeUserDeviceSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeUserDeviceSessionResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeUserDeviceSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeUserDeviceSessionResponseBody) SetRequestId(v string) *RevokeUserDeviceSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeUserDeviceSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
