// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQuotaAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateQuotaAlarmResponseBody
	GetRequestId() *string
}

type UpdateQuotaAlarmResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A95C65B3-7CF4-469E-B1D5-1CA0628A6411
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateQuotaAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateQuotaAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQuotaAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateQuotaAlarmResponseBody) SetRequestId(v string) *UpdateQuotaAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateQuotaAlarmResponseBody) Validate() error {
	return dara.Validate(s)
}
