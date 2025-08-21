// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindParentPlatformDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindParentPlatformDeviceResponseBody
	GetRequestId() *string
}

type UnbindParentPlatformDeviceResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindParentPlatformDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindParentPlatformDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindParentPlatformDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindParentPlatformDeviceResponseBody) SetRequestId(v string) *UnbindParentPlatformDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindParentPlatformDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
