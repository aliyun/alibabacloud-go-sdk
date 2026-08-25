// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMFADeviceForUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMFADeviceForUserResponseBody
	GetRequestId() *string
}

type DeleteMFADeviceForUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8B9982ED-FD0D-5622-8EA0-7B768685DCE7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMFADeviceForUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMFADeviceForUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMFADeviceForUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMFADeviceForUserResponseBody) SetRequestId(v string) *DeleteMFADeviceForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMFADeviceForUserResponseBody) Validate() error {
	return dara.Validate(s)
}
