// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventLogPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeEventLogPageResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeEventLogPageResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeEventLogPageResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeEventLogPageResponseBodyResultObject) *DescribeEventLogPageResponseBody
	GetResultObject() []*DescribeEventLogPageResponseBodyResultObject
	SetTotalItem(v int32) *DescribeEventLogPageResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeEventLogPageResponseBody
	GetTotalPage() *int32
}

type DescribeEventLogPageResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, default value is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Returned object.
	ResultObject []*DescribeEventLogPageResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items.
	//
	// example:
	//
	// 6
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 9
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeEventLogPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventLogPageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventLogPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventLogPageResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeEventLogPageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEventLogPageResponseBody) GetResultObject() []*DescribeEventLogPageResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeEventLogPageResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeEventLogPageResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeEventLogPageResponseBody) SetRequestId(v string) *DescribeEventLogPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventLogPageResponseBody) SetCurrentPage(v int32) *DescribeEventLogPageResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeEventLogPageResponseBody) SetPageSize(v int32) *DescribeEventLogPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEventLogPageResponseBody) SetResultObject(v []*DescribeEventLogPageResponseBodyResultObject) *DescribeEventLogPageResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeEventLogPageResponseBody) SetTotalItem(v int32) *DescribeEventLogPageResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeEventLogPageResponseBody) SetTotalPage(v int32) *DescribeEventLogPageResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeEventLogPageResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEventLogPageResponseBodyResultObject struct {
	// Account ID.
	//
	// example:
	//
	// 1631801314885832
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// Email.
	//
	// example:
	//
	// xxxx@123.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// IP address.
	//
	// example:
	//
	// 10.200.5.100
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
	// Mobile phone number.
	//
	// example:
	//
	// 13817606333
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// Nickname.
	//
	// example:
	//
	// 昵称
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 546F8063-0104-5271-9EB7-56FB3F375BAD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Score.
	//
	// example:
	//
	// 10.0
	Score *string `json:"score,omitempty" xml:"score,omitempty"`
	// Event name.
	//
	// example:
	//
	// 注册事件
	Service *string `json:"service,omitempty" xml:"service,omitempty"`
	// Tags.
	//
	// example:
	//
	// rm0102
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// Transaction time.
	//
	// example:
	//
	// 1737101348000
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// Device ID.
	//
	// example:
	//
	// 4239
	Umid *string `json:"umid,omitempty" xml:"umid,omitempty"`
}

func (s DescribeEventLogPageResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventLogPageResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeEventLogPageResponseBodyResultObject) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeEventLogPageResponseBodyResultObject) GetEmail() *string {
	return s.Email
}

func (s *DescribeEventLogPageResponseBodyResultObject) GetIp() *string {
	return s.Ip
}

func (s *DescribeEventLogPageResponseBodyResultObject) GetMobile() *string {
	return s.Mobile
}

func (s *DescribeEventLogPageResponseBodyResultObject) GetNickName() *string {
	return s.NickName
}

func (s *DescribeEventLogPageResponseBodyResultObject) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventLogPageResponseBodyResultObject) GetScore() *string {
	return s.Score
}

func (s *DescribeEventLogPageResponseBodyResultObject) GetService() *string {
	return s.Service
}

func (s *DescribeEventLogPageResponseBodyResultObject) GetTags() *string {
	return s.Tags
}

func (s *DescribeEventLogPageResponseBodyResultObject) GetTimestamp() *string {
	return s.Timestamp
}

func (s *DescribeEventLogPageResponseBodyResultObject) GetUmid() *string {
	return s.Umid
}

func (s *DescribeEventLogPageResponseBodyResultObject) SetAccountId(v string) *DescribeEventLogPageResponseBodyResultObject {
	s.AccountId = &v
	return s
}

func (s *DescribeEventLogPageResponseBodyResultObject) SetEmail(v string) *DescribeEventLogPageResponseBodyResultObject {
	s.Email = &v
	return s
}

func (s *DescribeEventLogPageResponseBodyResultObject) SetIp(v string) *DescribeEventLogPageResponseBodyResultObject {
	s.Ip = &v
	return s
}

func (s *DescribeEventLogPageResponseBodyResultObject) SetMobile(v string) *DescribeEventLogPageResponseBodyResultObject {
	s.Mobile = &v
	return s
}

func (s *DescribeEventLogPageResponseBodyResultObject) SetNickName(v string) *DescribeEventLogPageResponseBodyResultObject {
	s.NickName = &v
	return s
}

func (s *DescribeEventLogPageResponseBodyResultObject) SetRequestId(v string) *DescribeEventLogPageResponseBodyResultObject {
	s.RequestId = &v
	return s
}

func (s *DescribeEventLogPageResponseBodyResultObject) SetScore(v string) *DescribeEventLogPageResponseBodyResultObject {
	s.Score = &v
	return s
}

func (s *DescribeEventLogPageResponseBodyResultObject) SetService(v string) *DescribeEventLogPageResponseBodyResultObject {
	s.Service = &v
	return s
}

func (s *DescribeEventLogPageResponseBodyResultObject) SetTags(v string) *DescribeEventLogPageResponseBodyResultObject {
	s.Tags = &v
	return s
}

func (s *DescribeEventLogPageResponseBodyResultObject) SetTimestamp(v string) *DescribeEventLogPageResponseBodyResultObject {
	s.Timestamp = &v
	return s
}

func (s *DescribeEventLogPageResponseBodyResultObject) SetUmid(v string) *DescribeEventLogPageResponseBodyResultObject {
	s.Umid = &v
	return s
}

func (s *DescribeEventLogPageResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
