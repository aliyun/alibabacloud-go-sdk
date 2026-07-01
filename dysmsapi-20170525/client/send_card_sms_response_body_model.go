// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendCardSmsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendCardSmsResponseBody
	GetCode() *string
	SetData(v *SendCardSmsResponseBodyData) *SendCardSmsResponseBody
	GetData() *SendCardSmsResponseBodyData
	SetRequestId(v string) *SendCardSmsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendCardSmsResponseBody
	GetSuccess() *bool
}

type SendCardSmsResponseBody struct {
	// 请求状态码。
	//
	// 	- 返回OK代表请求成功。
	//
	// 	- 其他错误码，请参见[错误码列表](https://help.aliyun.com/document_detail/101346.html)。
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据。
	Data *SendCardSmsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 请求ID。
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8D28D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用接口是否成功。取值：
	//
	// - **true**：调用成功。
	//
	// - **false**：调用失败。
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendCardSmsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendCardSmsResponseBody) GoString() string {
	return s.String()
}

func (s *SendCardSmsResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendCardSmsResponseBody) GetData() *SendCardSmsResponseBodyData {
	return s.Data
}

func (s *SendCardSmsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendCardSmsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendCardSmsResponseBody) SetCode(v string) *SendCardSmsResponseBody {
	s.Code = &v
	return s
}

func (s *SendCardSmsResponseBody) SetData(v *SendCardSmsResponseBodyData) *SendCardSmsResponseBody {
	s.Data = v
	return s
}

func (s *SendCardSmsResponseBody) SetRequestId(v string) *SendCardSmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendCardSmsResponseBody) SetSuccess(v bool) *SendCardSmsResponseBody {
	s.Success = &v
	return s
}

func (s *SendCardSmsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendCardSmsResponseBodyData struct {
	// 卡片短信发送ID。
	//
	// example:
	//
	// 123
	BizCardId *string `json:"BizCardId,omitempty" xml:"BizCardId,omitempty"`
	// 数字短信发送ID。
	//
	// example:
	//
	// 232
	BizDigitalId *string `json:"BizDigitalId,omitempty" xml:"BizDigitalId,omitempty"`
	// 文本短信发送ID。
	//
	// example:
	//
	// 524
	BizSmsId *string `json:"BizSmsId,omitempty" xml:"BizSmsId,omitempty"`
	// 卡片短信模板审核状态。取值：
	//
	// - **0**：审核中。
	//
	// - **1**：审核通过。
	//
	// - **2**：审核不通过。
	//
	// >  审核不通过的短信可通过**FallbackType**字段设置回落流程。
	//
	// example:
	//
	// 0
	CardTmpState *int32 `json:"CardTmpState,omitempty" xml:"CardTmpState,omitempty"`
	// 接收卡片短信的手机号。
	//
	// example:
	//
	// 1390000****
	MediaMobiles *string `json:"MediaMobiles,omitempty" xml:"MediaMobiles,omitempty"`
	// 回落的手机号。
	//
	// example:
	//
	// 1390000****
	NotMediaMobiles *string `json:"NotMediaMobiles,omitempty" xml:"NotMediaMobiles,omitempty"`
}

func (s SendCardSmsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SendCardSmsResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendCardSmsResponseBodyData) GetBizCardId() *string {
	return s.BizCardId
}

func (s *SendCardSmsResponseBodyData) GetBizDigitalId() *string {
	return s.BizDigitalId
}

func (s *SendCardSmsResponseBodyData) GetBizSmsId() *string {
	return s.BizSmsId
}

func (s *SendCardSmsResponseBodyData) GetCardTmpState() *int32 {
	return s.CardTmpState
}

func (s *SendCardSmsResponseBodyData) GetMediaMobiles() *string {
	return s.MediaMobiles
}

func (s *SendCardSmsResponseBodyData) GetNotMediaMobiles() *string {
	return s.NotMediaMobiles
}

func (s *SendCardSmsResponseBodyData) SetBizCardId(v string) *SendCardSmsResponseBodyData {
	s.BizCardId = &v
	return s
}

func (s *SendCardSmsResponseBodyData) SetBizDigitalId(v string) *SendCardSmsResponseBodyData {
	s.BizDigitalId = &v
	return s
}

func (s *SendCardSmsResponseBodyData) SetBizSmsId(v string) *SendCardSmsResponseBodyData {
	s.BizSmsId = &v
	return s
}

func (s *SendCardSmsResponseBodyData) SetCardTmpState(v int32) *SendCardSmsResponseBodyData {
	s.CardTmpState = &v
	return s
}

func (s *SendCardSmsResponseBodyData) SetMediaMobiles(v string) *SendCardSmsResponseBodyData {
	s.MediaMobiles = &v
	return s
}

func (s *SendCardSmsResponseBodyData) SetNotMediaMobiles(v string) *SendCardSmsResponseBodyData {
	s.NotMediaMobiles = &v
	return s
}

func (s *SendCardSmsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
