// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UnlockDeviceResponseBody
	GetId() *string
	SetRequestId(v string) *UnlockDeviceResponseBody
	GetRequestId() *string
}

type UnlockDeviceResponseBody struct {
	// example:
	//
	// 323884****9092996
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnlockDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnlockDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockDeviceResponseBody) GetId() *string {
	return s.Id
}

func (s *UnlockDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnlockDeviceResponseBody) SetId(v string) *UnlockDeviceResponseBody {
	s.Id = &v
	return s
}

func (s *UnlockDeviceResponseBody) SetRequestId(v string) *UnlockDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnlockDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
