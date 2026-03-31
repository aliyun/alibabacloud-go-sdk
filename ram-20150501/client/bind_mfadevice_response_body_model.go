// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindMFADeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindMFADeviceResponseBody
	GetRequestId() *string
}

type BindMFADeviceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindMFADeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindMFADeviceResponseBody) GoString() string {
	return s.String()
}

func (s *BindMFADeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindMFADeviceResponseBody) SetRequestId(v string) *BindMFADeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindMFADeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
