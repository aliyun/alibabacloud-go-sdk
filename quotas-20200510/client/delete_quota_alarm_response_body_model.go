// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQuotaAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteQuotaAlarmResponseBody
	GetRequestId() *string
}

type DeleteQuotaAlarmResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A95C65B3-7CF4-469E-B1D5-1CA0628A6411
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteQuotaAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQuotaAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQuotaAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQuotaAlarmResponseBody) SetRequestId(v string) *DeleteQuotaAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQuotaAlarmResponseBody) Validate() error {
	return dara.Validate(s)
}
