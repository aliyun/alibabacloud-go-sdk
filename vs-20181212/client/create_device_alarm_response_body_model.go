// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeviceAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmDelay(v int64) *CreateDeviceAlarmResponseBody
	GetAlarmDelay() *int64
	SetAlarmId(v string) *CreateDeviceAlarmResponseBody
	GetAlarmId() *string
	SetExpire(v int64) *CreateDeviceAlarmResponseBody
	GetExpire() *int64
	SetRequestId(v string) *CreateDeviceAlarmResponseBody
	GetRequestId() *string
	SetUrl(v string) *CreateDeviceAlarmResponseBody
	GetUrl() *string
}

type CreateDeviceAlarmResponseBody struct {
	// example:
	//
	// 60
	AlarmDelay *int64 `json:"AlarmDelay,omitempty" xml:"AlarmDelay,omitempty"`
	// example:
	//
	// 0hyNgTdgv2D000195842
	AlarmId *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// example:
	//
	// 3600
	Expire *int64 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// example:
	//
	// 76E11E6A-4441-51C9-AF60-D354362257A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rtmp://demo.aliyundoc.com/live/310101*****7542007?auth_key=1639130258-0-0-b2b04fe85ece6*****a6b1a42bc7e
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateDeviceAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDeviceAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeviceAlarmResponseBody) GetAlarmDelay() *int64 {
	return s.AlarmDelay
}

func (s *CreateDeviceAlarmResponseBody) GetAlarmId() *string {
	return s.AlarmId
}

func (s *CreateDeviceAlarmResponseBody) GetExpire() *int64 {
	return s.Expire
}

func (s *CreateDeviceAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDeviceAlarmResponseBody) GetUrl() *string {
	return s.Url
}

func (s *CreateDeviceAlarmResponseBody) SetAlarmDelay(v int64) *CreateDeviceAlarmResponseBody {
	s.AlarmDelay = &v
	return s
}

func (s *CreateDeviceAlarmResponseBody) SetAlarmId(v string) *CreateDeviceAlarmResponseBody {
	s.AlarmId = &v
	return s
}

func (s *CreateDeviceAlarmResponseBody) SetExpire(v int64) *CreateDeviceAlarmResponseBody {
	s.Expire = &v
	return s
}

func (s *CreateDeviceAlarmResponseBody) SetRequestId(v string) *CreateDeviceAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeviceAlarmResponseBody) SetUrl(v string) *CreateDeviceAlarmResponseBody {
	s.Url = &v
	return s
}

func (s *CreateDeviceAlarmResponseBody) Validate() error {
	return dara.Validate(s)
}
