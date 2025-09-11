// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryExtCodeSignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryExtCodeSignResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryExtCodeSignResponseBody
	GetCode() *string
	SetData(v *QueryExtCodeSignResponseBodyData) *QueryExtCodeSignResponseBody
	GetData() *QueryExtCodeSignResponseBodyData
	SetMessage(v string) *QueryExtCodeSignResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryExtCodeSignResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryExtCodeSignResponseBody
	GetSuccess() *bool
}

type QueryExtCodeSignResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryExtCodeSignResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryExtCodeSignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryExtCodeSignResponseBody) GoString() string {
	return s.String()
}

func (s *QueryExtCodeSignResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryExtCodeSignResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryExtCodeSignResponseBody) GetData() *QueryExtCodeSignResponseBodyData {
	return s.Data
}

func (s *QueryExtCodeSignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryExtCodeSignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryExtCodeSignResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryExtCodeSignResponseBody) SetAccessDeniedDetail(v string) *QueryExtCodeSignResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryExtCodeSignResponseBody) SetCode(v string) *QueryExtCodeSignResponseBody {
	s.Code = &v
	return s
}

func (s *QueryExtCodeSignResponseBody) SetData(v *QueryExtCodeSignResponseBodyData) *QueryExtCodeSignResponseBody {
	s.Data = v
	return s
}

func (s *QueryExtCodeSignResponseBody) SetMessage(v string) *QueryExtCodeSignResponseBody {
	s.Message = &v
	return s
}

func (s *QueryExtCodeSignResponseBody) SetRequestId(v string) *QueryExtCodeSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryExtCodeSignResponseBody) SetSuccess(v bool) *QueryExtCodeSignResponseBody {
	s.Success = &v
	return s
}

func (s *QueryExtCodeSignResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryExtCodeSignResponseBodyData struct {
	List []*QueryExtCodeSignResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 5
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryExtCodeSignResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryExtCodeSignResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryExtCodeSignResponseBodyData) GetList() []*QueryExtCodeSignResponseBodyDataList {
	return s.List
}

func (s *QueryExtCodeSignResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QueryExtCodeSignResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryExtCodeSignResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *QueryExtCodeSignResponseBodyData) SetList(v []*QueryExtCodeSignResponseBodyDataList) *QueryExtCodeSignResponseBodyData {
	s.List = v
	return s
}

func (s *QueryExtCodeSignResponseBodyData) SetPageNo(v int64) *QueryExtCodeSignResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryExtCodeSignResponseBodyData) SetPageSize(v int64) *QueryExtCodeSignResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryExtCodeSignResponseBodyData) SetTotal(v int64) *QueryExtCodeSignResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryExtCodeSignResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryExtCodeSignResponseBodyDataList struct {
	// 是否可回收
	//
	// example:
	//
	// 1
	Active *int64 `json:"Active,omitempty" xml:"Active,omitempty"`
	// 扩展码A3
	//
	// example:
	//
	// 01
	ExtCode *string `json:"ExtCode,omitempty" xml:"ExtCode,omitempty"`
	// 近1个月发送成功条数（只读）
	//
	// example:
	//
	// 69
	SendCount *int64 `json:"SendCount,omitempty" xml:"SendCount,omitempty"`
	// 签名
	//
	// example:
	//
	// 示例值示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// 来源
	//
	// example:
	//
	// 示例值示例值示例值
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s QueryExtCodeSignResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s QueryExtCodeSignResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryExtCodeSignResponseBodyDataList) GetActive() *int64 {
	return s.Active
}

func (s *QueryExtCodeSignResponseBodyDataList) GetExtCode() *string {
	return s.ExtCode
}

func (s *QueryExtCodeSignResponseBodyDataList) GetSendCount() *int64 {
	return s.SendCount
}

func (s *QueryExtCodeSignResponseBodyDataList) GetSignName() *string {
	return s.SignName
}

func (s *QueryExtCodeSignResponseBodyDataList) GetSource() *string {
	return s.Source
}

func (s *QueryExtCodeSignResponseBodyDataList) SetActive(v int64) *QueryExtCodeSignResponseBodyDataList {
	s.Active = &v
	return s
}

func (s *QueryExtCodeSignResponseBodyDataList) SetExtCode(v string) *QueryExtCodeSignResponseBodyDataList {
	s.ExtCode = &v
	return s
}

func (s *QueryExtCodeSignResponseBodyDataList) SetSendCount(v int64) *QueryExtCodeSignResponseBodyDataList {
	s.SendCount = &v
	return s
}

func (s *QueryExtCodeSignResponseBodyDataList) SetSignName(v string) *QueryExtCodeSignResponseBodyDataList {
	s.SignName = &v
	return s
}

func (s *QueryExtCodeSignResponseBodyDataList) SetSource(v string) *QueryExtCodeSignResponseBodyDataList {
	s.Source = &v
	return s
}

func (s *QueryExtCodeSignResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
