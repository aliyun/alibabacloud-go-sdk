// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynchronizationJobAlertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDelayAlertPhone(v string) *DescribeSynchronizationJobAlertResponseBody
	GetDelayAlertPhone() *string
	SetDelayAlertStatus(v string) *DescribeSynchronizationJobAlertResponseBody
	GetDelayAlertStatus() *string
	SetDelayOverSeconds(v string) *DescribeSynchronizationJobAlertResponseBody
	GetDelayOverSeconds() *string
	SetErrCode(v string) *DescribeSynchronizationJobAlertResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeSynchronizationJobAlertResponseBody
	GetErrMessage() *string
	SetErrorAlertPhone(v string) *DescribeSynchronizationJobAlertResponseBody
	GetErrorAlertPhone() *string
	SetErrorAlertStatus(v string) *DescribeSynchronizationJobAlertResponseBody
	GetErrorAlertStatus() *string
	SetRequestId(v string) *DescribeSynchronizationJobAlertResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSynchronizationJobAlertResponseBody
	GetSuccess() *string
	SetSynchronizationDirection(v string) *DescribeSynchronizationJobAlertResponseBody
	GetSynchronizationDirection() *string
	SetSynchronizationJobId(v string) *DescribeSynchronizationJobAlertResponseBody
	GetSynchronizationJobId() *string
	SetSynchronizationJobName(v string) *DescribeSynchronizationJobAlertResponseBody
	GetSynchronizationJobName() *string
}

type DescribeSynchronizationJobAlertResponseBody struct {
	// The mobile phone numbers that receive latency-related alerts.
	//
	// example:
	//
	// 1361234****,1371234****
	DelayAlertPhone *string `json:"DelayAlertPhone,omitempty" xml:"DelayAlertPhone,omitempty"`
	// Indicates whether task latency is monitored. Valid values:
	//
	// 	- **enable**: yes
	//
	// 	- **disable**: no
	//
	// example:
	//
	// enable
	DelayAlertStatus *string `json:"DelayAlertStatus,omitempty" xml:"DelayAlertStatus,omitempty"`
	// The threshold for triggering latency alerts. Unit: seconds.
	//
	// example:
	//
	// 10
	DelayOverSeconds *string `json:"DelayOverSeconds,omitempty" xml:"DelayOverSeconds,omitempty"`
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The mobile phone numbers that receive status-related alerts.
	//
	// example:
	//
	// 1361234****,1371234****
	ErrorAlertPhone *string `json:"ErrorAlertPhone,omitempty" xml:"ErrorAlertPhone,omitempty"`
	// Indicates whether task status is monitored. Valid values:
	//
	// 	- **enable**: yes
	//
	// 	- **disable**: no
	//
	// example:
	//
	// enable
	ErrorAlertStatus *string `json:"ErrorAlertStatus,omitempty" xml:"ErrorAlertStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 210ec20e16055205968635339d****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The synchronization direction. Valid values:
	//
	// 	- **Forward**
	//
	// 	- **Reverse**
	//
	// > This parameter is returned only when the topology of data synchronization is two-way synchronization.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// The ID of the data synchronization instance.
	//
	// example:
	//
	// kxz1170c10p****
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	// The name of the data synchronization task.
	//
	// example:
	//
	// Polar MySQL_TO_RDS MySQL
	SynchronizationJobName *string `json:"SynchronizationJobName,omitempty" xml:"SynchronizationJobName,omitempty"`
}

func (s DescribeSynchronizationJobAlertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobAlertResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobAlertResponseBody) GetDelayAlertPhone() *string {
	return s.DelayAlertPhone
}

func (s *DescribeSynchronizationJobAlertResponseBody) GetDelayAlertStatus() *string {
	return s.DelayAlertStatus
}

func (s *DescribeSynchronizationJobAlertResponseBody) GetDelayOverSeconds() *string {
	return s.DelayOverSeconds
}

func (s *DescribeSynchronizationJobAlertResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeSynchronizationJobAlertResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeSynchronizationJobAlertResponseBody) GetErrorAlertPhone() *string {
	return s.ErrorAlertPhone
}

func (s *DescribeSynchronizationJobAlertResponseBody) GetErrorAlertStatus() *string {
	return s.ErrorAlertStatus
}

func (s *DescribeSynchronizationJobAlertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSynchronizationJobAlertResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSynchronizationJobAlertResponseBody) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *DescribeSynchronizationJobAlertResponseBody) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *DescribeSynchronizationJobAlertResponseBody) GetSynchronizationJobName() *string {
	return s.SynchronizationJobName
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetDelayAlertPhone(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.DelayAlertPhone = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetDelayAlertStatus(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.DelayAlertStatus = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetDelayOverSeconds(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.DelayOverSeconds = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetErrCode(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetErrMessage(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetErrorAlertPhone(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.ErrorAlertPhone = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetErrorAlertStatus(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.ErrorAlertStatus = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetRequestId(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetSuccess(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetSynchronizationDirection(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetSynchronizationJobId(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.SynchronizationJobId = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetSynchronizationJobName(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.SynchronizationJobName = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) Validate() error {
	return dara.Validate(s)
}
