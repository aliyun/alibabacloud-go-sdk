// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListExtenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudListExtenResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudListExtenResponseBody
	GetCode() *string
	SetData(v *CloudListExtenResponseBodyData) *CloudListExtenResponseBody
	GetData() *CloudListExtenResponseBodyData
	SetMessage(v string) *CloudListExtenResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudListExtenResponseBody
	GetRequestId() *string
}

type CloudListExtenResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudListExtenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudListExtenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudListExtenResponseBody) GoString() string {
	return s.String()
}

func (s *CloudListExtenResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudListExtenResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudListExtenResponseBody) GetData() *CloudListExtenResponseBodyData {
	return s.Data
}

func (s *CloudListExtenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudListExtenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudListExtenResponseBody) SetAccessDeniedDetail(v string) *CloudListExtenResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudListExtenResponseBody) SetCode(v string) *CloudListExtenResponseBody {
	s.Code = &v
	return s
}

func (s *CloudListExtenResponseBody) SetData(v *CloudListExtenResponseBodyData) *CloudListExtenResponseBody {
	s.Data = v
	return s
}

func (s *CloudListExtenResponseBody) SetMessage(v string) *CloudListExtenResponseBody {
	s.Message = &v
	return s
}

func (s *CloudListExtenResponseBody) SetRequestId(v string) *CloudListExtenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudListExtenResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudListExtenResponseBodyData struct {
	// example:
	//
	// 3
	EndRow *string `json:"EndRow,omitempty" xml:"EndRow,omitempty"`
	// example:
	//
	// 1
	FirstPage *string `json:"FirstPage,omitempty" xml:"FirstPage,omitempty"`
	// example:
	//
	// false
	HasNextPage *int64 `json:"HasNextPage,omitempty" xml:"HasNextPage,omitempty"`
	// example:
	//
	// true
	HasPreviousPage *int64 `json:"HasPreviousPage,omitempty" xml:"HasPreviousPage,omitempty"`
	// example:
	//
	// true
	IsFirstPage *int64 `json:"IsFirstPage,omitempty" xml:"IsFirstPage,omitempty"`
	// example:
	//
	// true
	IsLastPage *int64 `json:"IsLastPage,omitempty" xml:"IsLastPage,omitempty"`
	// example:
	//
	// 1
	LastPage *string                               `json:"LastPage,omitempty" xml:"LastPage,omitempty"`
	List     []*CloudListExtenResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值示例值
	NavigatePages    *string   `json:"NavigatePages,omitempty" xml:"NavigatePages,omitempty"`
	NavigatepageNums []*string `json:"NavigatepageNums,omitempty" xml:"NavigatepageNums,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	NextPage *string `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// example:
	//
	// exten
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// 1
	PageNum *string `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	Pages *string `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// example:
	//
	// 0
	PrePage *string `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// example:
	//
	// 3
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 1
	StartRow *string `json:"StartRow,omitempty" xml:"StartRow,omitempty"`
	// example:
	//
	// 3
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s CloudListExtenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudListExtenResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudListExtenResponseBodyData) GetEndRow() *string {
	return s.EndRow
}

func (s *CloudListExtenResponseBodyData) GetFirstPage() *string {
	return s.FirstPage
}

func (s *CloudListExtenResponseBodyData) GetHasNextPage() *int64 {
	return s.HasNextPage
}

func (s *CloudListExtenResponseBodyData) GetHasPreviousPage() *int64 {
	return s.HasPreviousPage
}

func (s *CloudListExtenResponseBodyData) GetIsFirstPage() *int64 {
	return s.IsFirstPage
}

func (s *CloudListExtenResponseBodyData) GetIsLastPage() *int64 {
	return s.IsLastPage
}

func (s *CloudListExtenResponseBodyData) GetLastPage() *string {
	return s.LastPage
}

func (s *CloudListExtenResponseBodyData) GetList() []*CloudListExtenResponseBodyDataList {
	return s.List
}

func (s *CloudListExtenResponseBodyData) GetNavigatePages() *string {
	return s.NavigatePages
}

func (s *CloudListExtenResponseBodyData) GetNavigatepageNums() []*string {
	return s.NavigatepageNums
}

func (s *CloudListExtenResponseBodyData) GetNextPage() *string {
	return s.NextPage
}

func (s *CloudListExtenResponseBodyData) GetOrderBy() *string {
	return s.OrderBy
}

func (s *CloudListExtenResponseBodyData) GetPageNum() *string {
	return s.PageNum
}

func (s *CloudListExtenResponseBodyData) GetPageSize() *string {
	return s.PageSize
}

func (s *CloudListExtenResponseBodyData) GetPages() *string {
	return s.Pages
}

func (s *CloudListExtenResponseBodyData) GetPrePage() *string {
	return s.PrePage
}

func (s *CloudListExtenResponseBodyData) GetSize() *string {
	return s.Size
}

func (s *CloudListExtenResponseBodyData) GetStartRow() *string {
	return s.StartRow
}

func (s *CloudListExtenResponseBodyData) GetTotal() *string {
	return s.Total
}

func (s *CloudListExtenResponseBodyData) SetEndRow(v string) *CloudListExtenResponseBodyData {
	s.EndRow = &v
	return s
}

func (s *CloudListExtenResponseBodyData) SetFirstPage(v string) *CloudListExtenResponseBodyData {
	s.FirstPage = &v
	return s
}

func (s *CloudListExtenResponseBodyData) SetHasNextPage(v int64) *CloudListExtenResponseBodyData {
	s.HasNextPage = &v
	return s
}

func (s *CloudListExtenResponseBodyData) SetHasPreviousPage(v int64) *CloudListExtenResponseBodyData {
	s.HasPreviousPage = &v
	return s
}

func (s *CloudListExtenResponseBodyData) SetIsFirstPage(v int64) *CloudListExtenResponseBodyData {
	s.IsFirstPage = &v
	return s
}

func (s *CloudListExtenResponseBodyData) SetIsLastPage(v int64) *CloudListExtenResponseBodyData {
	s.IsLastPage = &v
	return s
}

func (s *CloudListExtenResponseBodyData) SetLastPage(v string) *CloudListExtenResponseBodyData {
	s.LastPage = &v
	return s
}

func (s *CloudListExtenResponseBodyData) SetList(v []*CloudListExtenResponseBodyDataList) *CloudListExtenResponseBodyData {
	s.List = v
	return s
}

func (s *CloudListExtenResponseBodyData) SetNavigatePages(v string) *CloudListExtenResponseBodyData {
	s.NavigatePages = &v
	return s
}

func (s *CloudListExtenResponseBodyData) SetNavigatepageNums(v []*string) *CloudListExtenResponseBodyData {
	s.NavigatepageNums = v
	return s
}

func (s *CloudListExtenResponseBodyData) SetNextPage(v string) *CloudListExtenResponseBodyData {
	s.NextPage = &v
	return s
}

func (s *CloudListExtenResponseBodyData) SetOrderBy(v string) *CloudListExtenResponseBodyData {
	s.OrderBy = &v
	return s
}

func (s *CloudListExtenResponseBodyData) SetPageNum(v string) *CloudListExtenResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *CloudListExtenResponseBodyData) SetPageSize(v string) *CloudListExtenResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *CloudListExtenResponseBodyData) SetPages(v string) *CloudListExtenResponseBodyData {
	s.Pages = &v
	return s
}

func (s *CloudListExtenResponseBodyData) SetPrePage(v string) *CloudListExtenResponseBodyData {
	s.PrePage = &v
	return s
}

func (s *CloudListExtenResponseBodyData) SetSize(v string) *CloudListExtenResponseBodyData {
	s.Size = &v
	return s
}

func (s *CloudListExtenResponseBodyData) SetStartRow(v string) *CloudListExtenResponseBodyData {
	s.StartRow = &v
	return s
}

func (s *CloudListExtenResponseBodyData) SetTotal(v string) *CloudListExtenResponseBodyData {
	s.Total = &v
	return s
}

func (s *CloudListExtenResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CloudListExtenResponseBodyDataList struct {
	// 允许的语音编码,支持格式: 1. g729 2. g729,alaw,ulaw 3. alaw,ulaw,g729 4. alaw,ulaw 5. myopus,alaw,ulaw 公网+远程话机支持配置1/2/3;专线+远程话机支持配置1/2;公网+软电话支持配置4;专线+软电话支持配置5;
	//
	// example:
	//
	// alaw,ulaw
	Allow *string `json:"Allow,omitempty" xml:"Allow,omitempty"`
	// 区号
	//
	// example:
	//
	// 010
	AreaCode *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	// 呼叫权限，0：无限制，1：国内长途，2：国内本市，3：内部呼叫，默认无限制
	//
	// example:
	//
	// 0
	CallPower *string `json:"CallPower,omitempty" xml:"CallPower,omitempty"`
	// 创建时间
	//
	// example:
	//
	// 2026-03-30 06:09:02
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 降噪开关 0:关闭 1:开启 默认关闭(该参数只有在类型为注册到webrtc才有效)
	//
	// example:
	//
	// 0
	Denoise *int64 `json:"Denoise,omitempty" xml:"Denoise,omitempty"`
	// 企业编号
	//
	// example:
	//
	// 7000002
	EnterpriseId *string `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 分机号,3-11位
	//
	// example:
	//
	// 66749
	Exten *string `json:"Exten,omitempty" xml:"Exten,omitempty"`
	// 呼入是否录音，0：不录用，1：录音，默认录音
	//
	// example:
	//
	// 1
	IbRecord *int64 `json:"IbRecord,omitempty" xml:"IbRecord,omitempty"`
	// 分机号id
	//
	// example:
	//
	// 59
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 是否允许摘机外呼，0：不允许，1：可以，默认允许
	//
	// example:
	//
	// 1
	IsDirect *int64 `json:"IsDirect,omitempty" xml:"IsDirect,omitempty"`
	// 是否允许外呼，0：不允许，1：可以，默认允许
	//
	// example:
	//
	// 1
	IsOb *int64 `json:"IsOb,omitempty" xml:"IsOb,omitempty"`
	// 网络防抖
	//
	// example:
	//
	// 0
	JitterBuffer *int64 `json:"JitterBuffer,omitempty" xml:"JitterBuffer,omitempty"`
	// 外呼是否录音，0：不录音，1：录音，默认录音
	//
	// example:
	//
	// 1
	ObRecord *int64 `json:"ObRecord,omitempty" xml:"ObRecord,omitempty"`
	// 密码
	//
	// example:
	//
	// Aa626670
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// 类型，1. 注册到IAD分机 2.注册到webrtc 3.注册到远程话机
	//
	// example:
	//
	// 2
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CloudListExtenResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s CloudListExtenResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *CloudListExtenResponseBodyDataList) GetAllow() *string {
	return s.Allow
}

func (s *CloudListExtenResponseBodyDataList) GetAreaCode() *string {
	return s.AreaCode
}

func (s *CloudListExtenResponseBodyDataList) GetCallPower() *string {
	return s.CallPower
}

func (s *CloudListExtenResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudListExtenResponseBodyDataList) GetDenoise() *int64 {
	return s.Denoise
}

func (s *CloudListExtenResponseBodyDataList) GetEnterpriseId() *string {
	return s.EnterpriseId
}

func (s *CloudListExtenResponseBodyDataList) GetExten() *string {
	return s.Exten
}

func (s *CloudListExtenResponseBodyDataList) GetIbRecord() *int64 {
	return s.IbRecord
}

func (s *CloudListExtenResponseBodyDataList) GetId() *int64 {
	return s.Id
}

func (s *CloudListExtenResponseBodyDataList) GetIsDirect() *int64 {
	return s.IsDirect
}

func (s *CloudListExtenResponseBodyDataList) GetIsOb() *int64 {
	return s.IsOb
}

func (s *CloudListExtenResponseBodyDataList) GetJitterBuffer() *int64 {
	return s.JitterBuffer
}

func (s *CloudListExtenResponseBodyDataList) GetObRecord() *int64 {
	return s.ObRecord
}

func (s *CloudListExtenResponseBodyDataList) GetPassword() *string {
	return s.Password
}

func (s *CloudListExtenResponseBodyDataList) GetType() *int64 {
	return s.Type
}

func (s *CloudListExtenResponseBodyDataList) SetAllow(v string) *CloudListExtenResponseBodyDataList {
	s.Allow = &v
	return s
}

func (s *CloudListExtenResponseBodyDataList) SetAreaCode(v string) *CloudListExtenResponseBodyDataList {
	s.AreaCode = &v
	return s
}

func (s *CloudListExtenResponseBodyDataList) SetCallPower(v string) *CloudListExtenResponseBodyDataList {
	s.CallPower = &v
	return s
}

func (s *CloudListExtenResponseBodyDataList) SetCreateTime(v string) *CloudListExtenResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *CloudListExtenResponseBodyDataList) SetDenoise(v int64) *CloudListExtenResponseBodyDataList {
	s.Denoise = &v
	return s
}

func (s *CloudListExtenResponseBodyDataList) SetEnterpriseId(v string) *CloudListExtenResponseBodyDataList {
	s.EnterpriseId = &v
	return s
}

func (s *CloudListExtenResponseBodyDataList) SetExten(v string) *CloudListExtenResponseBodyDataList {
	s.Exten = &v
	return s
}

func (s *CloudListExtenResponseBodyDataList) SetIbRecord(v int64) *CloudListExtenResponseBodyDataList {
	s.IbRecord = &v
	return s
}

func (s *CloudListExtenResponseBodyDataList) SetId(v int64) *CloudListExtenResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *CloudListExtenResponseBodyDataList) SetIsDirect(v int64) *CloudListExtenResponseBodyDataList {
	s.IsDirect = &v
	return s
}

func (s *CloudListExtenResponseBodyDataList) SetIsOb(v int64) *CloudListExtenResponseBodyDataList {
	s.IsOb = &v
	return s
}

func (s *CloudListExtenResponseBodyDataList) SetJitterBuffer(v int64) *CloudListExtenResponseBodyDataList {
	s.JitterBuffer = &v
	return s
}

func (s *CloudListExtenResponseBodyDataList) SetObRecord(v int64) *CloudListExtenResponseBodyDataList {
	s.ObRecord = &v
	return s
}

func (s *CloudListExtenResponseBodyDataList) SetPassword(v string) *CloudListExtenResponseBodyDataList {
	s.Password = &v
	return s
}

func (s *CloudListExtenResponseBodyDataList) SetType(v int64) *CloudListExtenResponseBodyDataList {
	s.Type = &v
	return s
}

func (s *CloudListExtenResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
