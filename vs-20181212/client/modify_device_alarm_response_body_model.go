// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDeviceAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDeviceAlarmResponseBody
	GetRequestId() *string
}

type ModifyDeviceAlarmResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDeviceAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDeviceAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDeviceAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDeviceAlarmResponseBody) SetRequestId(v string) *ModifyDeviceAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDeviceAlarmResponseBody) Validate() error {
	return dara.Validate(s)
}
