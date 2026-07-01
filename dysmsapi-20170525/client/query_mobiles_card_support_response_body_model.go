// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMobilesCardSupportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryMobilesCardSupportResponseBody
	GetCode() *string
	SetData(v *QueryMobilesCardSupportResponseBodyData) *QueryMobilesCardSupportResponseBody
	GetData() *QueryMobilesCardSupportResponseBodyData
	SetRequestId(v string) *QueryMobilesCardSupportResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryMobilesCardSupportResponseBody
	GetSuccess() *bool
}

type QueryMobilesCardSupportResponseBody struct {
	// 请求状态码。
	//
	// - 返回OK代表请求成功。
	//
	// - 其他错误码，请参见[错误码列表](https://help.aliyun.com/document_detail/101346.html)。
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据。
	Data *QueryMobilesCardSupportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 阿里云为该请求生成的唯一标识符。
	//
	// example:
	//
	// 08C17DFE-2E10-54F4-BAFB-7180039CC217
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

func (s QueryMobilesCardSupportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMobilesCardSupportResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMobilesCardSupportResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryMobilesCardSupportResponseBody) GetData() *QueryMobilesCardSupportResponseBodyData {
	return s.Data
}

func (s *QueryMobilesCardSupportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMobilesCardSupportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMobilesCardSupportResponseBody) SetCode(v string) *QueryMobilesCardSupportResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMobilesCardSupportResponseBody) SetData(v *QueryMobilesCardSupportResponseBodyData) *QueryMobilesCardSupportResponseBody {
	s.Data = v
	return s
}

func (s *QueryMobilesCardSupportResponseBody) SetRequestId(v string) *QueryMobilesCardSupportResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMobilesCardSupportResponseBody) SetSuccess(v bool) *QueryMobilesCardSupportResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMobilesCardSupportResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMobilesCardSupportResponseBodyData struct {
	// 查询值。
	QueryResult []*QueryMobilesCardSupportResponseBodyDataQueryResult `json:"QueryResult,omitempty" xml:"QueryResult,omitempty" type:"Repeated"`
}

func (s QueryMobilesCardSupportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryMobilesCardSupportResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMobilesCardSupportResponseBodyData) GetQueryResult() []*QueryMobilesCardSupportResponseBodyDataQueryResult {
	return s.QueryResult
}

func (s *QueryMobilesCardSupportResponseBodyData) SetQueryResult(v []*QueryMobilesCardSupportResponseBodyDataQueryResult) *QueryMobilesCardSupportResponseBodyData {
	s.QueryResult = v
	return s
}

func (s *QueryMobilesCardSupportResponseBodyData) Validate() error {
	if s.QueryResult != nil {
		for _, item := range s.QueryResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMobilesCardSupportResponseBodyDataQueryResult struct {
	// 查询的手机号码。
	//
	// example:
	//
	// 1380000****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// 是否支持卡片短信。取值：
	//
	// - **true**：支持。
	//
	// - **false**：不支持。
	//
	// example:
	//
	// true
	Support *bool `json:"Support,omitempty" xml:"Support,omitempty"`
}

func (s QueryMobilesCardSupportResponseBodyDataQueryResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMobilesCardSupportResponseBodyDataQueryResult) GoString() string {
	return s.String()
}

func (s *QueryMobilesCardSupportResponseBodyDataQueryResult) GetMobile() *string {
	return s.Mobile
}

func (s *QueryMobilesCardSupportResponseBodyDataQueryResult) GetSupport() *bool {
	return s.Support
}

func (s *QueryMobilesCardSupportResponseBodyDataQueryResult) SetMobile(v string) *QueryMobilesCardSupportResponseBodyDataQueryResult {
	s.Mobile = &v
	return s
}

func (s *QueryMobilesCardSupportResponseBodyDataQueryResult) SetSupport(v bool) *QueryMobilesCardSupportResponseBodyDataQueryResult {
	s.Support = &v
	return s
}

func (s *QueryMobilesCardSupportResponseBodyDataQueryResult) Validate() error {
	return dara.Validate(s)
}
