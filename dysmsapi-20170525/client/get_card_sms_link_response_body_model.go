// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCardSmsLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCardSmsLinkResponseBody
	GetCode() *string
	SetData(v *GetCardSmsLinkResponseBodyData) *GetCardSmsLinkResponseBody
	GetData() *GetCardSmsLinkResponseBodyData
	SetRequestId(v string) *GetCardSmsLinkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCardSmsLinkResponseBody
	GetSuccess() *bool
}

type GetCardSmsLinkResponseBody struct {
	// 请求状态码。取值：
	//
	// - OK：代表请求成功。
	//
	// - 其他错误码，请参见[错误码列表](https://help.aliyun.com/document_detail/101346.html)。
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据。
	Data *GetCardSmsLinkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 请求ID。
	//
	// example:
	//
	// CC89A90C-978F-46AC-B80D-54738371E7CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 接口调用是否成功。取值：
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

func (s GetCardSmsLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCardSmsLinkResponseBody) GoString() string {
	return s.String()
}

func (s *GetCardSmsLinkResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCardSmsLinkResponseBody) GetData() *GetCardSmsLinkResponseBodyData {
	return s.Data
}

func (s *GetCardSmsLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCardSmsLinkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCardSmsLinkResponseBody) SetCode(v string) *GetCardSmsLinkResponseBody {
	s.Code = &v
	return s
}

func (s *GetCardSmsLinkResponseBody) SetData(v *GetCardSmsLinkResponseBodyData) *GetCardSmsLinkResponseBody {
	s.Data = v
	return s
}

func (s *GetCardSmsLinkResponseBody) SetRequestId(v string) *GetCardSmsLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCardSmsLinkResponseBody) SetSuccess(v bool) *GetCardSmsLinkResponseBody {
	s.Success = &v
	return s
}

func (s *GetCardSmsLinkResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCardSmsLinkResponseBodyData struct {
	// 支持卡片短信的手机号码。
	//
	// example:
	//
	// [\\"1390000****\\",\\"1370000****\\"]
	CardPhoneNumbers *string `json:"CardPhoneNumbers,omitempty" xml:"CardPhoneNumbers,omitempty"`
	// 用于申请卡片短信短链的短信签名，在发送时签名、接收号码、卡片短信短链要一一对应。
	//
	// example:
	//
	// ["阿里云","阿里云2"]
	CardSignNames *string `json:"CardSignNames,omitempty" xml:"CardSignNames,omitempty"`
	// 卡片短信短链。
	//
	// example:
	//
	// [\\"mw2m.cn/LAaGGa\\",\\"mw2m.cn/LAAaes\\"]
	CardSmsLinks *string `json:"CardSmsLinks,omitempty" xml:"CardSmsLinks,omitempty"`
	// 卡片短信模板审核状态。取值：
	//
	// - **0**：审核中。
	//
	// - **1**：审核通过。
	//
	// - **2**：审核不通过。
	//
	// > 未审核通过的短信走回落流程。
	//
	// example:
	//
	// 0
	CardTmpState *int32 `json:"CardTmpState,omitempty" xml:"CardTmpState,omitempty"`
	// 不支持卡片短信的手机号。
	//
	// example:
	//
	// 1390000****
	NotMediaMobiles *string `json:"NotMediaMobiles,omitempty" xml:"NotMediaMobiles,omitempty"`
}

func (s GetCardSmsLinkResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCardSmsLinkResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCardSmsLinkResponseBodyData) GetCardPhoneNumbers() *string {
	return s.CardPhoneNumbers
}

func (s *GetCardSmsLinkResponseBodyData) GetCardSignNames() *string {
	return s.CardSignNames
}

func (s *GetCardSmsLinkResponseBodyData) GetCardSmsLinks() *string {
	return s.CardSmsLinks
}

func (s *GetCardSmsLinkResponseBodyData) GetCardTmpState() *int32 {
	return s.CardTmpState
}

func (s *GetCardSmsLinkResponseBodyData) GetNotMediaMobiles() *string {
	return s.NotMediaMobiles
}

func (s *GetCardSmsLinkResponseBodyData) SetCardPhoneNumbers(v string) *GetCardSmsLinkResponseBodyData {
	s.CardPhoneNumbers = &v
	return s
}

func (s *GetCardSmsLinkResponseBodyData) SetCardSignNames(v string) *GetCardSmsLinkResponseBodyData {
	s.CardSignNames = &v
	return s
}

func (s *GetCardSmsLinkResponseBodyData) SetCardSmsLinks(v string) *GetCardSmsLinkResponseBodyData {
	s.CardSmsLinks = &v
	return s
}

func (s *GetCardSmsLinkResponseBodyData) SetCardTmpState(v int32) *GetCardSmsLinkResponseBodyData {
	s.CardTmpState = &v
	return s
}

func (s *GetCardSmsLinkResponseBodyData) SetNotMediaMobiles(v string) *GetCardSmsLinkResponseBodyData {
	s.NotMediaMobiles = &v
	return s
}

func (s *GetCardSmsLinkResponseBodyData) Validate() error {
	return dara.Validate(s)
}
