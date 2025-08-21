// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindParentPlatformDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindParentPlatformDeviceResponseBody
	GetRequestId() *string
}

type BindParentPlatformDeviceResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindParentPlatformDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindParentPlatformDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *BindParentPlatformDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindParentPlatformDeviceResponseBody) SetRequestId(v string) *BindParentPlatformDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindParentPlatformDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
