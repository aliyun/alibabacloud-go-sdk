// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsoleScoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetConsoleScoreResponseBody
	GetCode() *string
	SetData(v interface{}) *GetConsoleScoreResponseBody
	GetData() interface{}
	SetHttpStatusCode(v int32) *GetConsoleScoreResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetConsoleScoreResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetConsoleScoreResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetConsoleScoreResponseBody
	GetSuccess() *bool
}

type GetConsoleScoreResponseBody struct {
	// Interface response code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	//
	// example:
	//
	// {
	//
	//     "score": "94.00",
	//
	//     "consoleScoreTrendDTOS": [
	//
	//         {
	//
	//             "date": "20241009",
	//
	//             "score": "100.0"
	//
	//         }
	//
	//     ],
	//
	//     "cyclicYearOverYear": "-6.00",
	//
	//     "recordDate": "20241209",
	//
	//     "weeklyYearOverYear": "1.62",
	//
	//     "aboveWholeNetworkUserRatio": "6.25",
	//
	//     "aliUid": "1601097845544644",
	//
	//     "detailJson": "[{\\"detailDTO\\":[{\\"count\\":0,\\"itemName\\":\\"应用漏洞POC验证\\",\\"mark\\":\\"1\\"},{\\"count\\":0,\\"itemName\\":\\"未授权访问漏洞（公网暴露）\\",\\"mark\\":\\"1\\"},{\\"count\\":0,\\"itemName\\":\\"后台弱口令漏洞（公网暴露）\\",\\"mark\\":\\"1\\"},{\\"count\\":0,\\"itemName\\":\\"文件上传漏洞（公网暴露）\\",\\"mark\\":\\"1\\"}],\\"markRate\\":\\"0.5\\",\\"markType\\":\\"vul\\"},{\\"detailDTO\\":[{\\"count\\":12,\\"itemName\\":\\"WAF3.0回源配置不正确\\",\\"mark\\":\\"1\\"},{\\"count\\":0,\\"itemName\\":\\"AK泄露检查未开启\\",\\"mark\\":\\"1\\"},{\\"count\\":0,\\"itemName\\":\\"DNAT管理端口开放\\",\\"mark\\":\\"1\\"},{\\"count\\":0,\\"itemName\\":\\"高危端口暴露\\",\\"mark\\":\\"0.5\\"}],\\"markRate\\":\\"0.5\\",\\"markType\\":\\"risk\\"}]"
	//
	// }
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message for the result returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// D0937B0F-9180-5F70-B6ED-0BA22591627F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. true means success, false means failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetConsoleScoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConsoleScoreResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsoleScoreResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetConsoleScoreResponseBody) GetData() interface{} {
	return s.Data
}

func (s *GetConsoleScoreResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetConsoleScoreResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetConsoleScoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConsoleScoreResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetConsoleScoreResponseBody) SetCode(v string) *GetConsoleScoreResponseBody {
	s.Code = &v
	return s
}

func (s *GetConsoleScoreResponseBody) SetData(v interface{}) *GetConsoleScoreResponseBody {
	s.Data = v
	return s
}

func (s *GetConsoleScoreResponseBody) SetHttpStatusCode(v int32) *GetConsoleScoreResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetConsoleScoreResponseBody) SetMessage(v string) *GetConsoleScoreResponseBody {
	s.Message = &v
	return s
}

func (s *GetConsoleScoreResponseBody) SetRequestId(v string) *GetConsoleScoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsoleScoreResponseBody) SetSuccess(v bool) *GetConsoleScoreResponseBody {
	s.Success = &v
	return s
}

func (s *GetConsoleScoreResponseBody) Validate() error {
	return dara.Validate(s)
}
