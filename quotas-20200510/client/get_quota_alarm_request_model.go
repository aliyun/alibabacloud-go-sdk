// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmId(v string) *GetQuotaAlarmRequest
	GetAlarmId() *string
}

type GetQuotaAlarmRequest struct {
	// The ID of the quota alert.
	//
	// For more information about how to query the ID of a quota alert, see [ListQuotaAlarms](https://help.aliyun.com/document_detail/184348.html).
	//
	// example:
	//
	// 78d7e436-4b25-4897-84b5-d7b656bb****
	AlarmId *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
}

func (s GetQuotaAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaAlarmRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaAlarmRequest) GetAlarmId() *string {
	return s.AlarmId
}

func (s *GetQuotaAlarmRequest) SetAlarmId(v string) *GetQuotaAlarmRequest {
	s.AlarmId = &v
	return s
}

func (s *GetQuotaAlarmRequest) Validate() error {
	return dara.Validate(s)
}
