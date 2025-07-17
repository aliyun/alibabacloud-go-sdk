// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableAlarmResponseBody
	GetRequestId() *string
}

type DisableAlarmResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 086EFCD4-C76F-4DC6-9EE9-0D9B711E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableAlarmResponseBody) SetRequestId(v string) *DisableAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableAlarmResponseBody) Validate() error {
	return dara.Validate(s)
}
