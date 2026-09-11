// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferPayTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *TransferPayTypeResponseBody
	GetChargeType() *string
	SetCode(v string) *TransferPayTypeResponseBody
	GetCode() *string
	SetDtsJobId(v string) *TransferPayTypeResponseBody
	GetDtsJobId() *string
	SetDynamicMessage(v string) *TransferPayTypeResponseBody
	GetDynamicMessage() *string
	SetEndTime(v string) *TransferPayTypeResponseBody
	GetEndTime() *string
	SetErrCode(v string) *TransferPayTypeResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *TransferPayTypeResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *TransferPayTypeResponseBody
	GetHttpStatusCode() *int32
	SetInstanceId(v string) *TransferPayTypeResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *TransferPayTypeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TransferPayTypeResponseBody
	GetSuccess() *bool
}

type TransferPayTypeResponseBody struct {
	// The billing method after conversion. Valid values:
	//
	// - **PrePaid**: subscription.
	//
	// - **PostPaid**: pay-as-you-go.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The error code. This parameter will be deprecated.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The task ID.
	//
	// example:
	//
	// o4nh3g7jg56****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The dynamic error message used to replace the **%s*	- variable in the **ErrMessage*	- parameter.
	//
	// > For example, if **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// DtsJobId
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The time when the subscription instance expires, in UNIX timestamp format.
	//
	// > - If the instance is converted to pay-as-you-go, this value is empty.
	//
	// - You can use a search engine to find a UNIX timestamp conversion tool.
	//
	// example:
	//
	// 1614916318
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
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
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// dtso4nh3g7jg56****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 601B6F25-21E7-4484-99D5-3EF2625C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TransferPayTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferPayTypeResponseBody) GoString() string {
	return s.String()
}

func (s *TransferPayTypeResponseBody) GetChargeType() *string {
	return s.ChargeType
}

func (s *TransferPayTypeResponseBody) GetCode() *string {
	return s.Code
}

func (s *TransferPayTypeResponseBody) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *TransferPayTypeResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *TransferPayTypeResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *TransferPayTypeResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *TransferPayTypeResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *TransferPayTypeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *TransferPayTypeResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TransferPayTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferPayTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TransferPayTypeResponseBody) SetChargeType(v string) *TransferPayTypeResponseBody {
	s.ChargeType = &v
	return s
}

func (s *TransferPayTypeResponseBody) SetCode(v string) *TransferPayTypeResponseBody {
	s.Code = &v
	return s
}

func (s *TransferPayTypeResponseBody) SetDtsJobId(v string) *TransferPayTypeResponseBody {
	s.DtsJobId = &v
	return s
}

func (s *TransferPayTypeResponseBody) SetDynamicMessage(v string) *TransferPayTypeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *TransferPayTypeResponseBody) SetEndTime(v string) *TransferPayTypeResponseBody {
	s.EndTime = &v
	return s
}

func (s *TransferPayTypeResponseBody) SetErrCode(v string) *TransferPayTypeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *TransferPayTypeResponseBody) SetErrMessage(v string) *TransferPayTypeResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *TransferPayTypeResponseBody) SetHttpStatusCode(v int32) *TransferPayTypeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TransferPayTypeResponseBody) SetInstanceId(v string) *TransferPayTypeResponseBody {
	s.InstanceId = &v
	return s
}

func (s *TransferPayTypeResponseBody) SetRequestId(v string) *TransferPayTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferPayTypeResponseBody) SetSuccess(v bool) *TransferPayTypeResponseBody {
	s.Success = &v
	return s
}

func (s *TransferPayTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
