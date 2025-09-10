// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQuotaAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmId(v string) *CreateQuotaAlarmResponseBody
	GetAlarmId() *string
	SetRequestId(v string) *CreateQuotaAlarmResponseBody
	GetRequestId() *string
}

type CreateQuotaAlarmResponseBody struct {
	// Alarm ID.
	//
	// example:
	//
	// 18b3be23-b7b0-4d45-91bc-d0c331aa****
	AlarmId *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BD219E2B-E687-45EE-B5F3-61FB730551B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateQuotaAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQuotaAlarmResponseBody) GetAlarmId() *string {
	return s.AlarmId
}

func (s *CreateQuotaAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQuotaAlarmResponseBody) SetAlarmId(v string) *CreateQuotaAlarmResponseBody {
	s.AlarmId = &v
	return s
}

func (s *CreateQuotaAlarmResponseBody) SetRequestId(v string) *CreateQuotaAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQuotaAlarmResponseBody) Validate() error {
	return dara.Validate(s)
}
