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
	// B9AF80E4-1565-42D9-9256-0B8B0D9FD3EC
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
