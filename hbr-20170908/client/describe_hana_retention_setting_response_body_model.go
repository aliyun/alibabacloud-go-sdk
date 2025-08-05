// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHanaRetentionSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeHanaRetentionSettingResponseBody
	GetClusterId() *string
	SetCode(v string) *DescribeHanaRetentionSettingResponseBody
	GetCode() *string
	SetDatabaseName(v string) *DescribeHanaRetentionSettingResponseBody
	GetDatabaseName() *string
	SetDisabled(v bool) *DescribeHanaRetentionSettingResponseBody
	GetDisabled() *bool
	SetMessage(v string) *DescribeHanaRetentionSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeHanaRetentionSettingResponseBody
	GetRequestId() *string
	SetRetentionDays(v int64) *DescribeHanaRetentionSettingResponseBody
	GetRetentionDays() *int64
	SetSchedule(v string) *DescribeHanaRetentionSettingResponseBody
	GetSchedule() *string
	SetSuccess(v bool) *DescribeHanaRetentionSettingResponseBody
	GetSuccess() *bool
	SetVaultId(v string) *DescribeHanaRetentionSettingResponseBody
	GetVaultId() *string
}

type DescribeHanaRetentionSettingResponseBody struct {
	// The ID of the SAP HANA instance.
	//
	// example:
	//
	// cl-0003jyv******fsku5m
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The database name.
	//
	// example:
	//
	// Q01
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// Indicates whether the backup is permanently retained. Valid values:
	//
	// 	- true: The backup is permanently retained.
	//
	// 	- false: The backup is retained for the specified number of days.
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 280DD872-EE25-52E8-9CB4-491067173DD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of days for which the backup is retained. If the value of the Disabled parameter is false, the backup is retained for the number of days specified by this parameter.
	//
	// example:
	//
	// 3650
	RetentionDays *int64 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The policy to update the retention period. Format: `I|{startTime}|{interval}`, which indicates that the retention period is updated at an interval of {interval} starting from {startTime}.
	//
	// 	- startTime: the time at which the system starts to run a backup job. The time follows the UNIX time format. Unit: seconds.
	//
	// 	- interval: the interval at which the system runs a backup job. The interval follows the ISO 8601 standard. For example, PT1H indicates an interval of 1 hour. P1D indicates an interval of one day.
	//
	// example:
	//
	// I|0|P1D
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-0006wkn7******zkn
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeHanaRetentionSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaRetentionSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHanaRetentionSettingResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeHanaRetentionSettingResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeHanaRetentionSettingResponseBody) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeHanaRetentionSettingResponseBody) GetDisabled() *bool {
	return s.Disabled
}

func (s *DescribeHanaRetentionSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeHanaRetentionSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHanaRetentionSettingResponseBody) GetRetentionDays() *int64 {
	return s.RetentionDays
}

func (s *DescribeHanaRetentionSettingResponseBody) GetSchedule() *string {
	return s.Schedule
}

func (s *DescribeHanaRetentionSettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeHanaRetentionSettingResponseBody) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeHanaRetentionSettingResponseBody) SetClusterId(v string) *DescribeHanaRetentionSettingResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaRetentionSettingResponseBody) SetCode(v string) *DescribeHanaRetentionSettingResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHanaRetentionSettingResponseBody) SetDatabaseName(v string) *DescribeHanaRetentionSettingResponseBody {
	s.DatabaseName = &v
	return s
}

func (s *DescribeHanaRetentionSettingResponseBody) SetDisabled(v bool) *DescribeHanaRetentionSettingResponseBody {
	s.Disabled = &v
	return s
}

func (s *DescribeHanaRetentionSettingResponseBody) SetMessage(v string) *DescribeHanaRetentionSettingResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHanaRetentionSettingResponseBody) SetRequestId(v string) *DescribeHanaRetentionSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHanaRetentionSettingResponseBody) SetRetentionDays(v int64) *DescribeHanaRetentionSettingResponseBody {
	s.RetentionDays = &v
	return s
}

func (s *DescribeHanaRetentionSettingResponseBody) SetSchedule(v string) *DescribeHanaRetentionSettingResponseBody {
	s.Schedule = &v
	return s
}

func (s *DescribeHanaRetentionSettingResponseBody) SetSuccess(v bool) *DescribeHanaRetentionSettingResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHanaRetentionSettingResponseBody) SetVaultId(v string) *DescribeHanaRetentionSettingResponseBody {
	s.VaultId = &v
	return s
}

func (s *DescribeHanaRetentionSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
