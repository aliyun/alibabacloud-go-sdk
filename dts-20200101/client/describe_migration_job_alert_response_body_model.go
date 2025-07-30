// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMigrationJobAlertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDelayAlertPhone(v string) *DescribeMigrationJobAlertResponseBody
	GetDelayAlertPhone() *string
	SetDelayAlertStatus(v string) *DescribeMigrationJobAlertResponseBody
	GetDelayAlertStatus() *string
	SetDelayOverSeconds(v string) *DescribeMigrationJobAlertResponseBody
	GetDelayOverSeconds() *string
	SetErrCode(v string) *DescribeMigrationJobAlertResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeMigrationJobAlertResponseBody
	GetErrMessage() *string
	SetErrorAlertPhone(v string) *DescribeMigrationJobAlertResponseBody
	GetErrorAlertPhone() *string
	SetErrorAlertStatus(v string) *DescribeMigrationJobAlertResponseBody
	GetErrorAlertStatus() *string
	SetMigrationJobId(v string) *DescribeMigrationJobAlertResponseBody
	GetMigrationJobId() *string
	SetMigrationJobName(v string) *DescribeMigrationJobAlertResponseBody
	GetMigrationJobName() *string
	SetRequestId(v string) *DescribeMigrationJobAlertResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeMigrationJobAlertResponseBody
	GetSuccess() *string
}

type DescribeMigrationJobAlertResponseBody struct {
	// Phone number of the contact for delay alarm.
	//
	// example:
	//
	// 1361234****,1371234****
	DelayAlertPhone *string `json:"DelayAlertPhone,omitempty" xml:"DelayAlertPhone,omitempty"`
	// Whether to monitor the delay status, return values:
	//
	// - **enable**: Yes. - **disable**: No.
	//
	// example:
	//
	// enable
	DelayAlertStatus *string `json:"DelayAlertStatus,omitempty" xml:"DelayAlertStatus,omitempty"`
	// The threshold for triggering a delayed alarm, in seconds.
	//
	// example:
	//
	// 0
	DelayOverSeconds *string `json:"DelayOverSeconds,omitempty" xml:"DelayOverSeconds,omitempty"`
	// Error code returned when the call fails.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// Phone number of the contact for abnormal alarm notifications
	//
	// example:
	//
	// 1361234****,1371234****
	ErrorAlertPhone *string `json:"ErrorAlertPhone,omitempty" xml:"ErrorAlertPhone,omitempty"`
	// Whether to monitor for abnormal status, return values:
	//
	// - **enable**: Yes. - **disable**: No.
	//
	// example:
	//
	// enable
	ErrorAlertStatus *string `json:"ErrorAlertStatus,omitempty" xml:"ErrorAlertStatus,omitempty"`
	// Data migration instance ID.
	//
	// example:
	//
	// dtslb9113qq11n****
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	// Name of the data migration task.
	//
	// example:
	//
	// zwy-test5
	MigrationJobName *string `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CFB89C51-6F03-519C-A921-AAE28D50AEFB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMigrationJobAlertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobAlertResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobAlertResponseBody) GetDelayAlertPhone() *string {
	return s.DelayAlertPhone
}

func (s *DescribeMigrationJobAlertResponseBody) GetDelayAlertStatus() *string {
	return s.DelayAlertStatus
}

func (s *DescribeMigrationJobAlertResponseBody) GetDelayOverSeconds() *string {
	return s.DelayOverSeconds
}

func (s *DescribeMigrationJobAlertResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeMigrationJobAlertResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeMigrationJobAlertResponseBody) GetErrorAlertPhone() *string {
	return s.ErrorAlertPhone
}

func (s *DescribeMigrationJobAlertResponseBody) GetErrorAlertStatus() *string {
	return s.ErrorAlertStatus
}

func (s *DescribeMigrationJobAlertResponseBody) GetMigrationJobId() *string {
	return s.MigrationJobId
}

func (s *DescribeMigrationJobAlertResponseBody) GetMigrationJobName() *string {
	return s.MigrationJobName
}

func (s *DescribeMigrationJobAlertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMigrationJobAlertResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeMigrationJobAlertResponseBody) SetDelayAlertPhone(v string) *DescribeMigrationJobAlertResponseBody {
	s.DelayAlertPhone = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetDelayAlertStatus(v string) *DescribeMigrationJobAlertResponseBody {
	s.DelayAlertStatus = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetDelayOverSeconds(v string) *DescribeMigrationJobAlertResponseBody {
	s.DelayOverSeconds = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetErrCode(v string) *DescribeMigrationJobAlertResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetErrMessage(v string) *DescribeMigrationJobAlertResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetErrorAlertPhone(v string) *DescribeMigrationJobAlertResponseBody {
	s.ErrorAlertPhone = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetErrorAlertStatus(v string) *DescribeMigrationJobAlertResponseBody {
	s.ErrorAlertStatus = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetMigrationJobId(v string) *DescribeMigrationJobAlertResponseBody {
	s.MigrationJobId = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetMigrationJobName(v string) *DescribeMigrationJobAlertResponseBody {
	s.MigrationJobName = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetRequestId(v string) *DescribeMigrationJobAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetSuccess(v string) *DescribeMigrationJobAlertResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) Validate() error {
	return dara.Validate(s)
}
