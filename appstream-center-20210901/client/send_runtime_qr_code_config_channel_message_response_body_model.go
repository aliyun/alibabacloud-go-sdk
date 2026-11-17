// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendRuntimeQrCodeConfigChannelMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*SendRuntimeQrCodeConfigChannelMessageResponseBodyData) *SendRuntimeQrCodeConfigChannelMessageResponseBody
	GetData() []*SendRuntimeQrCodeConfigChannelMessageResponseBodyData
	SetRequestId(v string) *SendRuntimeQrCodeConfigChannelMessageResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *SendRuntimeQrCodeConfigChannelMessageResponseBody
	GetTotalCount() *int32
}

type SendRuntimeQrCodeConfigChannelMessageResponseBody struct {
	// The list of sending results for each runtime. Each runtime corresponds to one result. The return order is not guaranteed to be the same as the order of RuntimeIds in the request. Match the results by RuntimeId.
	Data []*SendRuntimeQrCodeConfigChannelMessageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID, which is a unique identifier for the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned sending results, which is the same as the number of results in Data.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SendRuntimeQrCodeConfigChannelMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendRuntimeQrCodeConfigChannelMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponseBody) GetData() []*SendRuntimeQrCodeConfigChannelMessageResponseBodyData {
	return s.Data
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponseBody) SetData(v []*SendRuntimeQrCodeConfigChannelMessageResponseBodyData) *SendRuntimeQrCodeConfigChannelMessageResponseBody {
	s.Data = v
	return s
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponseBody) SetRequestId(v string) *SendRuntimeQrCodeConfigChannelMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponseBody) SetTotalCount(v int32) *SendRuntimeQrCodeConfigChannelMessageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SendRuntimeQrCodeConfigChannelMessageResponseBodyData struct {
	// The error code when the sending fails, such as Runtime.NotFound (the runtime does not exist) or Runtime.UserNotSupported (the runtime has no authorized user or has multiple authorized users). This parameter is not returned when the sending is successful.
	//
	// example:
	//
	// Runtime.UserNotSupported
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message when the sending fails. This parameter is not returned when the sending is successful.
	//
	// example:
	//
	// Runtime has no authorized user
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The access URL of the QR code configuration channel page. It carries the taskId, channel code, and region parameters for the current process, and has been sent to the authorized user of the runtime by email and SMS. This parameter is returned only when the sending is successful. The taskId is valid for 1 hour.
	//
	// example:
	//
	// https://eds.console.aliyun.com/public/scan-pair?taskId=8f14e45f-ceea-4672-9f6e-3c5b2a1d****&code=feishu&region=cn-shanghai
	QrCodeNotifyUrl *string `json:"QrCodeNotifyUrl,omitempty" xml:"QrCodeNotifyUrl,omitempty"`
	// The ID of the runtime resource.
	//
	// example:
	//
	// ecd-xgads
	RuntimeId *string `json:"RuntimeId,omitempty" xml:"RuntimeId,omitempty"`
	// Indicates whether the QR code configuration channel message for the runtime was initiated successfully. If the value is false, you can check the Code and Message fields for the cause of the failure.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendRuntimeQrCodeConfigChannelMessageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SendRuntimeQrCodeConfigChannelMessageResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponseBodyData) GetQrCodeNotifyUrl() *string {
	return s.QrCodeNotifyUrl
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponseBodyData) GetRuntimeId() *string {
	return s.RuntimeId
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponseBodyData) SetCode(v string) *SendRuntimeQrCodeConfigChannelMessageResponseBodyData {
	s.Code = &v
	return s
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponseBodyData) SetMessage(v string) *SendRuntimeQrCodeConfigChannelMessageResponseBodyData {
	s.Message = &v
	return s
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponseBodyData) SetQrCodeNotifyUrl(v string) *SendRuntimeQrCodeConfigChannelMessageResponseBodyData {
	s.QrCodeNotifyUrl = &v
	return s
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponseBodyData) SetRuntimeId(v string) *SendRuntimeQrCodeConfigChannelMessageResponseBodyData {
	s.RuntimeId = &v
	return s
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponseBodyData) SetSuccess(v bool) *SendRuntimeQrCodeConfigChannelMessageResponseBodyData {
	s.Success = &v
	return s
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
