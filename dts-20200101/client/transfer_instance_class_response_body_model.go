// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferInstanceClassResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *TransferInstanceClassResponseBody
	GetChargeType() *string
	SetCode(v string) *TransferInstanceClassResponseBody
	GetCode() *string
	SetDtsJobId(v string) *TransferInstanceClassResponseBody
	GetDtsJobId() *string
	SetDynamicMessage(v string) *TransferInstanceClassResponseBody
	GetDynamicMessage() *string
	SetEndTime(v string) *TransferInstanceClassResponseBody
	GetEndTime() *string
	SetErrCode(v string) *TransferInstanceClassResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *TransferInstanceClassResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *TransferInstanceClassResponseBody
	GetHttpStatusCode() *int32
	SetInstanceId(v string) *TransferInstanceClassResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *TransferInstanceClassResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TransferInstanceClassResponseBody
	GetSuccess() *bool
}

type TransferInstanceClassResponseBody struct {
	// The billing method of the DTS instance. Valid values:
	//
	// 	- **POSTPAY**: pay-as-you-go.
	//
	// 	- **PREPAY**: subscription.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The error code that is returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the data migration or data synchronization task.
	//
	// example:
	//
	// r4yr723m199****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the **%s*	- variable in the value of the **ErrMessage*	- parameter.
	//
	// > For example, if the return value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the return value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified value of the **DtsJobId*	- parameter is invalid.
	//
	// example:
	//
	// DtsJobId
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The expiration time of the DTS instance.
	//
	// > This parameter is returned only if the value of the ChargeType parameter is **PREPAY**.
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
	// The ID of the DTS instance.
	//
	// example:
	//
	// dtsr4yr723m199****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 716DB03E-2D0B-4DC3-BC4C-F7A9EE21****
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

func (s TransferInstanceClassResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferInstanceClassResponseBody) GoString() string {
	return s.String()
}

func (s *TransferInstanceClassResponseBody) GetChargeType() *string {
	return s.ChargeType
}

func (s *TransferInstanceClassResponseBody) GetCode() *string {
	return s.Code
}

func (s *TransferInstanceClassResponseBody) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *TransferInstanceClassResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *TransferInstanceClassResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *TransferInstanceClassResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *TransferInstanceClassResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *TransferInstanceClassResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *TransferInstanceClassResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TransferInstanceClassResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferInstanceClassResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TransferInstanceClassResponseBody) SetChargeType(v string) *TransferInstanceClassResponseBody {
	s.ChargeType = &v
	return s
}

func (s *TransferInstanceClassResponseBody) SetCode(v string) *TransferInstanceClassResponseBody {
	s.Code = &v
	return s
}

func (s *TransferInstanceClassResponseBody) SetDtsJobId(v string) *TransferInstanceClassResponseBody {
	s.DtsJobId = &v
	return s
}

func (s *TransferInstanceClassResponseBody) SetDynamicMessage(v string) *TransferInstanceClassResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *TransferInstanceClassResponseBody) SetEndTime(v string) *TransferInstanceClassResponseBody {
	s.EndTime = &v
	return s
}

func (s *TransferInstanceClassResponseBody) SetErrCode(v string) *TransferInstanceClassResponseBody {
	s.ErrCode = &v
	return s
}

func (s *TransferInstanceClassResponseBody) SetErrMessage(v string) *TransferInstanceClassResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *TransferInstanceClassResponseBody) SetHttpStatusCode(v int32) *TransferInstanceClassResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TransferInstanceClassResponseBody) SetInstanceId(v string) *TransferInstanceClassResponseBody {
	s.InstanceId = &v
	return s
}

func (s *TransferInstanceClassResponseBody) SetRequestId(v string) *TransferInstanceClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferInstanceClassResponseBody) SetSuccess(v bool) *TransferInstanceClassResponseBody {
	s.Success = &v
	return s
}

func (s *TransferInstanceClassResponseBody) Validate() error {
	return dara.Validate(s)
}
