// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAllActivityInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryAllActivityInfoResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryAllActivityInfoResponseBody
	GetCode() *string
	SetData(v []*QueryAllActivityInfoResponseBodyData) *QueryAllActivityInfoResponseBody
	GetData() []*QueryAllActivityInfoResponseBodyData
	SetHttpStatusCode(v int32) *QueryAllActivityInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryAllActivityInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAllActivityInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAllActivityInfoResponseBody
	GetSuccess() *bool
}

type QueryAllActivityInfoResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code           *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*QueryAllActivityInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// null
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAllActivityInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAllActivityInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllActivityInfoResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryAllActivityInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAllActivityInfoResponseBody) GetData() []*QueryAllActivityInfoResponseBodyData {
	return s.Data
}

func (s *QueryAllActivityInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryAllActivityInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAllActivityInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAllActivityInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAllActivityInfoResponseBody) SetAccessDeniedDetail(v string) *QueryAllActivityInfoResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryAllActivityInfoResponseBody) SetCode(v string) *QueryAllActivityInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAllActivityInfoResponseBody) SetData(v []*QueryAllActivityInfoResponseBodyData) *QueryAllActivityInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryAllActivityInfoResponseBody) SetHttpStatusCode(v int32) *QueryAllActivityInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryAllActivityInfoResponseBody) SetMessage(v string) *QueryAllActivityInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAllActivityInfoResponseBody) SetRequestId(v string) *QueryAllActivityInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAllActivityInfoResponseBody) SetSuccess(v bool) *QueryAllActivityInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAllActivityInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryAllActivityInfoResponseBodyData struct {
	// example:
	//
	// 1234
	ActivityId *int64 `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// vip
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// example:
	//
	// test
	CustomerName *string `json:"CustomerName,omitempty" xml:"CustomerName,omitempty"`
	// example:
	//
	// xx@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Id    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true
	IsVipCustomer *string `json:"IsVipCustomer,omitempty" xml:"IsVipCustomer,omitempty"`
	// example:
	//
	// 12123455
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	QRCode *string `json:"QRCode,omitempty" xml:"QRCode,omitempty"`
	// example:
	//
	// {}
	ReportFields *string `json:"ReportFields,omitempty" xml:"ReportFields,omitempty"`
}

func (s QueryAllActivityInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAllActivityInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAllActivityInfoResponseBodyData) GetActivityId() *int64 {
	return s.ActivityId
}

func (s *QueryAllActivityInfoResponseBodyData) GetChannelName() *string {
	return s.ChannelName
}

func (s *QueryAllActivityInfoResponseBodyData) GetCompanyName() *string {
	return s.CompanyName
}

func (s *QueryAllActivityInfoResponseBodyData) GetCustomerName() *string {
	return s.CustomerName
}

func (s *QueryAllActivityInfoResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *QueryAllActivityInfoResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *QueryAllActivityInfoResponseBodyData) GetIsVipCustomer() *string {
	return s.IsVipCustomer
}

func (s *QueryAllActivityInfoResponseBodyData) GetMobile() *string {
	return s.Mobile
}

func (s *QueryAllActivityInfoResponseBodyData) GetQRCode() *string {
	return s.QRCode
}

func (s *QueryAllActivityInfoResponseBodyData) GetReportFields() *string {
	return s.ReportFields
}

func (s *QueryAllActivityInfoResponseBodyData) SetActivityId(v int64) *QueryAllActivityInfoResponseBodyData {
	s.ActivityId = &v
	return s
}

func (s *QueryAllActivityInfoResponseBodyData) SetChannelName(v string) *QueryAllActivityInfoResponseBodyData {
	s.ChannelName = &v
	return s
}

func (s *QueryAllActivityInfoResponseBodyData) SetCompanyName(v string) *QueryAllActivityInfoResponseBodyData {
	s.CompanyName = &v
	return s
}

func (s *QueryAllActivityInfoResponseBodyData) SetCustomerName(v string) *QueryAllActivityInfoResponseBodyData {
	s.CustomerName = &v
	return s
}

func (s *QueryAllActivityInfoResponseBodyData) SetEmail(v string) *QueryAllActivityInfoResponseBodyData {
	s.Email = &v
	return s
}

func (s *QueryAllActivityInfoResponseBodyData) SetId(v int64) *QueryAllActivityInfoResponseBodyData {
	s.Id = &v
	return s
}

func (s *QueryAllActivityInfoResponseBodyData) SetIsVipCustomer(v string) *QueryAllActivityInfoResponseBodyData {
	s.IsVipCustomer = &v
	return s
}

func (s *QueryAllActivityInfoResponseBodyData) SetMobile(v string) *QueryAllActivityInfoResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *QueryAllActivityInfoResponseBodyData) SetQRCode(v string) *QueryAllActivityInfoResponseBodyData {
	s.QRCode = &v
	return s
}

func (s *QueryAllActivityInfoResponseBodyData) SetReportFields(v string) *QueryAllActivityInfoResponseBodyData {
	s.ReportFields = &v
	return s
}

func (s *QueryAllActivityInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
