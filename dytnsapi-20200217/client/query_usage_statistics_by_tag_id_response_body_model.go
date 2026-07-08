// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUsageStatisticsByTagIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryUsageStatisticsByTagIdResponseBody
	GetCode() *string
	SetData(v []*QueryUsageStatisticsByTagIdResponseBodyData) *QueryUsageStatisticsByTagIdResponseBody
	GetData() []*QueryUsageStatisticsByTagIdResponseBodyData
	SetMessage(v string) *QueryUsageStatisticsByTagIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryUsageStatisticsByTagIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryUsageStatisticsByTagIdResponseBody
	GetSuccess() *bool
}

type QueryUsageStatisticsByTagIdResponseBody struct {
	// 请求状态码。 **OK**为成功。
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回的数据内容。
	Data []*QueryUsageStatisticsByTagIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// 返回状态码描述。
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求ID。
	//
	// example:
	//
	// D45CC751-34DF-5797-81FB-9A2ED6DC024B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 表示是否调用成功。
	//
	// - true：调用成功。
	//
	//  - false：调用失败。
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUsageStatisticsByTagIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUsageStatisticsByTagIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUsageStatisticsByTagIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryUsageStatisticsByTagIdResponseBody) GetData() []*QueryUsageStatisticsByTagIdResponseBodyData {
	return s.Data
}

func (s *QueryUsageStatisticsByTagIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryUsageStatisticsByTagIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUsageStatisticsByTagIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryUsageStatisticsByTagIdResponseBody) SetCode(v string) *QueryUsageStatisticsByTagIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBody) SetData(v []*QueryUsageStatisticsByTagIdResponseBodyData) *QueryUsageStatisticsByTagIdResponseBody {
	s.Data = v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBody) SetMessage(v string) *QueryUsageStatisticsByTagIdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBody) SetRequestId(v string) *QueryUsageStatisticsByTagIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBody) SetSuccess(v bool) *QueryUsageStatisticsByTagIdResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBody) Validate() error {
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

type QueryUsageStatisticsByTagIdResponseBodyData struct {
	// 授权码。
	//
	// example:
	//
	// g61I8UV5zd
	AuthorizationCode *string `json:"AuthorizationCode,omitempty" xml:"AuthorizationCode,omitempty"`
	// 查询失败号码数。
	//
	// example:
	//
	// 71
	FailTotal *int64 `json:"FailTotal,omitempty" xml:"FailTotal,omitempty"`
	// 创建时间。
	//
	// example:
	//
	// 20230312
	GmtDateStr *string `json:"GmtDateStr,omitempty" xml:"GmtDateStr,omitempty"`
	// 授权码使用记录ID。
	//
	// example:
	//
	// 4d32c9ef-c8b0-4d69-a437-cca5****185d
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 行业名称。
	//
	// example:
	//
	// 家政安保
	IndustryName *string `json:"IndustryName,omitempty" xml:"IndustryName,omitempty"`
	// 客户pid。
	//
	// example:
	//
	// 10000013****071
	PartnerId *int64 `json:"PartnerId,omitempty" xml:"PartnerId,omitempty"`
	// 场景名称。
	//
	// example:
	//
	// 电话回访1
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// 查询成功号码数。
	//
	// example:
	//
	// 93
	SuccessTotal *int64 `json:"SuccessTotal,omitempty" xml:"SuccessTotal,omitempty"`
	// 标签名称。
	//
	// example:
	//
	// 69
	TagId *int64 `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// 标签名称。
	//
	// example:
	//
	// 阿里云查询
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// 查询总号码数。
	//
	// example:
	//
	// 41
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryUsageStatisticsByTagIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryUsageStatisticsByTagIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) GetAuthorizationCode() *string {
	return s.AuthorizationCode
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) GetFailTotal() *int64 {
	return s.FailTotal
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) GetGmtDateStr() *string {
	return s.GmtDateStr
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) GetIndustryName() *string {
	return s.IndustryName
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) GetPartnerId() *int64 {
	return s.PartnerId
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) GetSceneName() *string {
	return s.SceneName
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) GetSuccessTotal() *int64 {
	return s.SuccessTotal
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) GetTagId() *int64 {
	return s.TagId
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) GetTagName() *string {
	return s.TagName
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) SetAuthorizationCode(v string) *QueryUsageStatisticsByTagIdResponseBodyData {
	s.AuthorizationCode = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) SetFailTotal(v int64) *QueryUsageStatisticsByTagIdResponseBodyData {
	s.FailTotal = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) SetGmtDateStr(v string) *QueryUsageStatisticsByTagIdResponseBodyData {
	s.GmtDateStr = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) SetId(v int64) *QueryUsageStatisticsByTagIdResponseBodyData {
	s.Id = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) SetIndustryName(v string) *QueryUsageStatisticsByTagIdResponseBodyData {
	s.IndustryName = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) SetPartnerId(v int64) *QueryUsageStatisticsByTagIdResponseBodyData {
	s.PartnerId = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) SetSceneName(v string) *QueryUsageStatisticsByTagIdResponseBodyData {
	s.SceneName = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) SetSuccessTotal(v int64) *QueryUsageStatisticsByTagIdResponseBodyData {
	s.SuccessTotal = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) SetTagId(v int64) *QueryUsageStatisticsByTagIdResponseBodyData {
	s.TagId = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) SetTagName(v string) *QueryUsageStatisticsByTagIdResponseBodyData {
	s.TagName = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) SetTotal(v int64) *QueryUsageStatisticsByTagIdResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
