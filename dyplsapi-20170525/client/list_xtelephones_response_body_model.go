// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListXTelephonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListXTelephonesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListXTelephonesResponseBody
	GetCode() *string
	SetData(v *ListXTelephonesResponseBodyData) *ListXTelephonesResponseBody
	GetData() *ListXTelephonesResponseBodyData
	SetMessage(v string) *ListXTelephonesResponseBody
	GetMessage() *string
	SetSuccess(v bool) *ListXTelephonesResponseBody
	GetSuccess() *bool
}

type ListXTelephonesResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 返回状态码 0000表示成功 其他表示失败
	//
	// example:
	//
	// 0000
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListXTelephonesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 返回信息
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回是否成功 true  表示成功 false表示失败
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListXTelephonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListXTelephonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListXTelephonesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListXTelephonesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListXTelephonesResponseBody) GetData() *ListXTelephonesResponseBodyData {
	return s.Data
}

func (s *ListXTelephonesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListXTelephonesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListXTelephonesResponseBody) SetAccessDeniedDetail(v string) *ListXTelephonesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListXTelephonesResponseBody) SetCode(v string) *ListXTelephonesResponseBody {
	s.Code = &v
	return s
}

func (s *ListXTelephonesResponseBody) SetData(v *ListXTelephonesResponseBodyData) *ListXTelephonesResponseBody {
	s.Data = v
	return s
}

func (s *ListXTelephonesResponseBody) SetMessage(v string) *ListXTelephonesResponseBody {
	s.Message = &v
	return s
}

func (s *ListXTelephonesResponseBody) SetSuccess(v bool) *ListXTelephonesResponseBody {
	s.Success = &v
	return s
}

func (s *ListXTelephonesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListXTelephonesResponseBodyData struct {
	// 数据集合
	List []*ListXTelephonesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// 页码
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// 每页条数
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 符合查询条件的总数量
	//
	// example:
	//
	// 50
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListXTelephonesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListXTelephonesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListXTelephonesResponseBodyData) GetList() []*ListXTelephonesResponseBodyDataList {
	return s.List
}

func (s *ListXTelephonesResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListXTelephonesResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListXTelephonesResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListXTelephonesResponseBodyData) SetList(v []*ListXTelephonesResponseBodyDataList) *ListXTelephonesResponseBodyData {
	s.List = v
	return s
}

func (s *ListXTelephonesResponseBodyData) SetPageNo(v int64) *ListXTelephonesResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListXTelephonesResponseBodyData) SetPageSize(v int64) *ListXTelephonesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListXTelephonesResponseBodyData) SetTotal(v int64) *ListXTelephonesResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListXTelephonesResponseBodyData) Validate() error {
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

type ListXTelephonesResponseBodyDataList struct {
	// 绑定失败原因
	//
	// example:
	//
	// 绑定失败用户身份证黑名单
	AuthMsg *string `json:"AuthMsg,omitempty" xml:"AuthMsg,omitempty"`
	// 绑定时间
	//
	// example:
	//
	// 2024-08-29 17:23:58
	BindTime *string `json:"BindTime,omitempty" xml:"BindTime,omitempty"`
	// 购买时间
	//
	// example:
	//
	// 2024-08-29 17:23:58
	BuyTime *string `json:"BuyTime,omitempty" xml:"BuyTime,omitempty"`
	// 客户号码池key
	//
	// example:
	//
	// FC533e6eeb81f4400c87ef3745a21a1a
	CustomerPoolKey *string `json:"CustomerPoolKey,omitempty" xml:"CustomerPoolKey,omitempty"`
	// 号码池名称
	//
	// example:
	//
	// 测试号码池
	CustomerPoolName *string `json:"CustomerPoolName,omitempty" xml:"CustomerPoolName,omitempty"`
	// 释放时间
	//
	// example:
	//
	// 2024-08-29 17:23:58
	ReleaseTime *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// 短信开通状态：0 未开通 1已开通
	//
	// example:
	//
	// 0
	SmsStatus *string `json:"SmsStatus,omitempty" xml:"SmsStatus,omitempty"`
	// X号码
	//
	// example:
	//
	// 17816876546
	Telephone *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	// 号码状态：0 空闲中 1 调拨完成待购买 2购买完成待认证  3 实名认证中  4 实名认证成功  5 认证失败  6 解绑中 7 解绑失败 8已释放 99 超时释放
	//
	// example:
	//
	// 0
	TelephoneStatus *string `json:"TelephoneStatus,omitempty" xml:"TelephoneStatus,omitempty"`
	// 解绑时间
	//
	// example:
	//
	// 2024-08-29 17:23:58
	UnbindTime *string `json:"UnbindTime,omitempty" xml:"UnbindTime,omitempty"`
}

func (s ListXTelephonesResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListXTelephonesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListXTelephonesResponseBodyDataList) GetAuthMsg() *string {
	return s.AuthMsg
}

func (s *ListXTelephonesResponseBodyDataList) GetBindTime() *string {
	return s.BindTime
}

func (s *ListXTelephonesResponseBodyDataList) GetBuyTime() *string {
	return s.BuyTime
}

func (s *ListXTelephonesResponseBodyDataList) GetCustomerPoolKey() *string {
	return s.CustomerPoolKey
}

func (s *ListXTelephonesResponseBodyDataList) GetCustomerPoolName() *string {
	return s.CustomerPoolName
}

func (s *ListXTelephonesResponseBodyDataList) GetReleaseTime() *string {
	return s.ReleaseTime
}

func (s *ListXTelephonesResponseBodyDataList) GetSmsStatus() *string {
	return s.SmsStatus
}

func (s *ListXTelephonesResponseBodyDataList) GetTelephone() *string {
	return s.Telephone
}

func (s *ListXTelephonesResponseBodyDataList) GetTelephoneStatus() *string {
	return s.TelephoneStatus
}

func (s *ListXTelephonesResponseBodyDataList) GetUnbindTime() *string {
	return s.UnbindTime
}

func (s *ListXTelephonesResponseBodyDataList) SetAuthMsg(v string) *ListXTelephonesResponseBodyDataList {
	s.AuthMsg = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetBindTime(v string) *ListXTelephonesResponseBodyDataList {
	s.BindTime = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetBuyTime(v string) *ListXTelephonesResponseBodyDataList {
	s.BuyTime = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetCustomerPoolKey(v string) *ListXTelephonesResponseBodyDataList {
	s.CustomerPoolKey = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetCustomerPoolName(v string) *ListXTelephonesResponseBodyDataList {
	s.CustomerPoolName = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetReleaseTime(v string) *ListXTelephonesResponseBodyDataList {
	s.ReleaseTime = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetSmsStatus(v string) *ListXTelephonesResponseBodyDataList {
	s.SmsStatus = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetTelephone(v string) *ListXTelephonesResponseBodyDataList {
	s.Telephone = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetTelephoneStatus(v string) *ListXTelephonesResponseBodyDataList {
	s.TelephoneStatus = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetUnbindTime(v string) *ListXTelephonesResponseBodyDataList {
	s.UnbindTime = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
