// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendBatchCardSmsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendBatchCardSmsResponseBody
	GetCode() *string
	SetData(v *SendBatchCardSmsResponseBodyData) *SendBatchCardSmsResponseBody
	GetData() *SendBatchCardSmsResponseBodyData
	SetRequestId(v string) *SendBatchCardSmsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendBatchCardSmsResponseBody
	GetSuccess() *bool
}

type SendBatchCardSmsResponseBody struct {
	// The request status code.
	//
	// 	- If **OK*	- is returned, the request is successful.
	//
	// 	- For information about other error codes, see [API error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *SendBatchCardSmsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700A8C4D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - **true**: The call is successful.
	//
	// - **false**: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendBatchCardSmsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendBatchCardSmsResponseBody) GoString() string {
	return s.String()
}

func (s *SendBatchCardSmsResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendBatchCardSmsResponseBody) GetData() *SendBatchCardSmsResponseBodyData {
	return s.Data
}

func (s *SendBatchCardSmsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendBatchCardSmsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendBatchCardSmsResponseBody) SetCode(v string) *SendBatchCardSmsResponseBody {
	s.Code = &v
	return s
}

func (s *SendBatchCardSmsResponseBody) SetData(v *SendBatchCardSmsResponseBodyData) *SendBatchCardSmsResponseBody {
	s.Data = v
	return s
}

func (s *SendBatchCardSmsResponseBody) SetRequestId(v string) *SendBatchCardSmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendBatchCardSmsResponseBody) SetSuccess(v bool) *SendBatchCardSmsResponseBody {
	s.Success = &v
	return s
}

func (s *SendBatchCardSmsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendBatchCardSmsResponseBodyData struct {
	// The ID of the card SMS sending task.
	//
	// example:
	//
	// 123
	BizCardId *string `json:"BizCardId,omitempty" xml:"BizCardId,omitempty"`
	// The ID of the digital SMS sending task.
	//
	// example:
	//
	// 3214
	BizDigitalId *string `json:"BizDigitalId,omitempty" xml:"BizDigitalId,omitempty"`
	// The ID of the text SMS sending task.
	//
	// example:
	//
	// 3256
	BizSmsId *string `json:"BizSmsId,omitempty" xml:"BizSmsId,omitempty"`
	// The review status of the card SMS template. Valid values:
	//
	// - **0**: Under review.
	//
	// - **1**: Approved.
	//
	// - **2**: Rejected.
	//
	// > For SMS messages that are rejected, you can configure the fallback process by using the **FallbackType*	- parameter.
	//
	// example:
	//
	// 0
	CardTmpState *int32 `json:"CardTmpState,omitempty" xml:"CardTmpState,omitempty"`
	// The mobile phone numbers that receive the card SMS messages.
	//
	// example:
	//
	// 1390000****
	MediaMobiles *string `json:"MediaMobiles,omitempty" xml:"MediaMobiles,omitempty"`
	// The fallback phone numbers.
	//
	// example:
	//
	// 1390000****
	NotMediaMobiles *string `json:"NotMediaMobiles,omitempty" xml:"NotMediaMobiles,omitempty"`
}

func (s SendBatchCardSmsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SendBatchCardSmsResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendBatchCardSmsResponseBodyData) GetBizCardId() *string {
	return s.BizCardId
}

func (s *SendBatchCardSmsResponseBodyData) GetBizDigitalId() *string {
	return s.BizDigitalId
}

func (s *SendBatchCardSmsResponseBodyData) GetBizSmsId() *string {
	return s.BizSmsId
}

func (s *SendBatchCardSmsResponseBodyData) GetCardTmpState() *int32 {
	return s.CardTmpState
}

func (s *SendBatchCardSmsResponseBodyData) GetMediaMobiles() *string {
	return s.MediaMobiles
}

func (s *SendBatchCardSmsResponseBodyData) GetNotMediaMobiles() *string {
	return s.NotMediaMobiles
}

func (s *SendBatchCardSmsResponseBodyData) SetBizCardId(v string) *SendBatchCardSmsResponseBodyData {
	s.BizCardId = &v
	return s
}

func (s *SendBatchCardSmsResponseBodyData) SetBizDigitalId(v string) *SendBatchCardSmsResponseBodyData {
	s.BizDigitalId = &v
	return s
}

func (s *SendBatchCardSmsResponseBodyData) SetBizSmsId(v string) *SendBatchCardSmsResponseBodyData {
	s.BizSmsId = &v
	return s
}

func (s *SendBatchCardSmsResponseBodyData) SetCardTmpState(v int32) *SendBatchCardSmsResponseBodyData {
	s.CardTmpState = &v
	return s
}

func (s *SendBatchCardSmsResponseBodyData) SetMediaMobiles(v string) *SendBatchCardSmsResponseBodyData {
	s.MediaMobiles = &v
	return s
}

func (s *SendBatchCardSmsResponseBodyData) SetNotMediaMobiles(v string) *SendBatchCardSmsResponseBodyData {
	s.NotMediaMobiles = &v
	return s
}

func (s *SendBatchCardSmsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
