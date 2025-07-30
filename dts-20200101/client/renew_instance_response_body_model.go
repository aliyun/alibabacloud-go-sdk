// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *RenewInstanceResponseBody
	GetChargeType() *string
	SetCode(v string) *RenewInstanceResponseBody
	GetCode() *string
	SetDtsJobId(v string) *RenewInstanceResponseBody
	GetDtsJobId() *string
	SetDynamicMessage(v string) *RenewInstanceResponseBody
	GetDynamicMessage() *string
	SetEndTime(v string) *RenewInstanceResponseBody
	GetEndTime() *string
	SetErrCode(v string) *RenewInstanceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *RenewInstanceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *RenewInstanceResponseBody
	GetHttpStatusCode() *int32
	SetInstanceId(v string) *RenewInstanceResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *RenewInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RenewInstanceResponseBody
	GetSuccess() *bool
}

type RenewInstanceResponseBody struct {
	// The billing method of the DTS instance. Only **PREPAY*	- may be returned, which indicates the subscription billing method.
	//
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The error code returned if the request failed.
	//
	// > This parameter will be removed in the future.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the data synchronization or change tracking task.
	//
	// example:
	//
	// qi0r643lc31****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the **%s*	- variable in the value of **ErrMessage**.
	//
	// > If the return value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the return value of **DynamicMessage*	- is **DtsJobId**, the specified value of **DtsJobId*	- is invalid.
	//
	// example:
	//
	// DtsJobId
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The time when the DTS instance expires after renewal. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ssZZZ	- format. The time is displayed in UTC.
	//
	// > **ZZZ*	- indicates the offset of the time zone, which is displayed in the format of a plus sign (+) or a minus sign (-) followed by hours and minutes, such as **+00:00**.
	//
	// example:
	//
	// 2021-08-04T16:00:00.000+00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// dtsqi0r643lc31****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1112E255-0C38-4970-8159-1D54AD92****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenewInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBody) GetChargeType() *string {
	return s.ChargeType
}

func (s *RenewInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *RenewInstanceResponseBody) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *RenewInstanceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *RenewInstanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *RenewInstanceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *RenewInstanceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *RenewInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RenewInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RenewInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RenewInstanceResponseBody) SetChargeType(v string) *RenewInstanceResponseBody {
	s.ChargeType = &v
	return s
}

func (s *RenewInstanceResponseBody) SetCode(v string) *RenewInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *RenewInstanceResponseBody) SetDtsJobId(v string) *RenewInstanceResponseBody {
	s.DtsJobId = &v
	return s
}

func (s *RenewInstanceResponseBody) SetDynamicMessage(v string) *RenewInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *RenewInstanceResponseBody) SetEndTime(v string) *RenewInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *RenewInstanceResponseBody) SetErrCode(v string) *RenewInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RenewInstanceResponseBody) SetErrMessage(v string) *RenewInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *RenewInstanceResponseBody) SetHttpStatusCode(v int32) *RenewInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RenewInstanceResponseBody) SetInstanceId(v string) *RenewInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *RenewInstanceResponseBody) SetRequestId(v string) *RenewInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewInstanceResponseBody) SetSuccess(v bool) *RenewInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *RenewInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
