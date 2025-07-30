// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionInstanceAlertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDelayAlertPhone(v string) *DescribeSubscriptionInstanceAlertResponseBody
	GetDelayAlertPhone() *string
	SetDelayAlertStatus(v string) *DescribeSubscriptionInstanceAlertResponseBody
	GetDelayAlertStatus() *string
	SetDelayOverSeconds(v string) *DescribeSubscriptionInstanceAlertResponseBody
	GetDelayOverSeconds() *string
	SetErrCode(v string) *DescribeSubscriptionInstanceAlertResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeSubscriptionInstanceAlertResponseBody
	GetErrMessage() *string
	SetErrorAlertPhone(v string) *DescribeSubscriptionInstanceAlertResponseBody
	GetErrorAlertPhone() *string
	SetErrorAlertStatus(v string) *DescribeSubscriptionInstanceAlertResponseBody
	GetErrorAlertStatus() *string
	SetRequestId(v string) *DescribeSubscriptionInstanceAlertResponseBody
	GetRequestId() *string
	SetSubscriptionInstanceID(v string) *DescribeSubscriptionInstanceAlertResponseBody
	GetSubscriptionInstanceID() *string
	SetSubscriptionInstanceName(v string) *DescribeSubscriptionInstanceAlertResponseBody
	GetSubscriptionInstanceName() *string
	SetSuccess(v string) *DescribeSubscriptionInstanceAlertResponseBody
	GetSuccess() *string
}

type DescribeSubscriptionInstanceAlertResponseBody struct {
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
	// The threshold for triggering latency alerts. The unit is seconds and the value is an integer. The recommended value is 10 seconds.
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
	// 210ec2e116055198849072222d****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the change tracking instance.
	//
	// example:
	//
	// dtsl8zl9ek6292****
	SubscriptionInstanceID *string `json:"SubscriptionInstanceID,omitempty" xml:"SubscriptionInstanceID,omitempty"`
	// The name of the change tracking instance.
	//
	// example:
	//
	// test
	SubscriptionInstanceName *string `json:"SubscriptionInstanceName,omitempty" xml:"SubscriptionInstanceName,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSubscriptionInstanceAlertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInstanceAlertResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) GetDelayAlertPhone() *string {
	return s.DelayAlertPhone
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) GetDelayAlertStatus() *string {
	return s.DelayAlertStatus
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) GetDelayOverSeconds() *string {
	return s.DelayOverSeconds
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) GetErrorAlertPhone() *string {
	return s.ErrorAlertPhone
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) GetErrorAlertStatus() *string {
	return s.ErrorAlertStatus
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) GetSubscriptionInstanceID() *string {
	return s.SubscriptionInstanceID
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) GetSubscriptionInstanceName() *string {
	return s.SubscriptionInstanceName
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetDelayAlertPhone(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.DelayAlertPhone = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetDelayAlertStatus(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.DelayAlertStatus = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetDelayOverSeconds(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.DelayOverSeconds = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetErrCode(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetErrMessage(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetErrorAlertPhone(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.ErrorAlertPhone = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetErrorAlertStatus(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.ErrorAlertStatus = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetRequestId(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetSubscriptionInstanceID(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.SubscriptionInstanceID = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetSubscriptionInstanceName(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.SubscriptionInstanceName = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetSuccess(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) Validate() error {
	return dara.Validate(s)
}
