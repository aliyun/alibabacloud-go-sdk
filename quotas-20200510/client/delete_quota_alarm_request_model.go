// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQuotaAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmId(v string) *DeleteQuotaAlarmRequest
	GetAlarmId() *string
}

type DeleteQuotaAlarmRequest struct {
	// The ID of the quota alert.
	//
	// >  You can call the [ListQuotaAlarms](https://help.aliyun.com/document_detail/440561.html) operation to obtain the ID of a quota alert.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6b512ab7-da3a-4142-b529-2b2a9294****
	AlarmId *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
}

func (s DeleteQuotaAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQuotaAlarmRequest) GoString() string {
	return s.String()
}

func (s *DeleteQuotaAlarmRequest) GetAlarmId() *string {
	return s.AlarmId
}

func (s *DeleteQuotaAlarmRequest) SetAlarmId(v string) *DeleteQuotaAlarmRequest {
	s.AlarmId = &v
	return s
}

func (s *DeleteQuotaAlarmRequest) Validate() error {
	return dara.Validate(s)
}
