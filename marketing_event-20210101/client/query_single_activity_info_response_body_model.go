// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySingleActivityInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QuerySingleActivityInfoResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QuerySingleActivityInfoResponseBody
	GetCode() *string
	SetData(v []*QuerySingleActivityInfoResponseBodyData) *QuerySingleActivityInfoResponseBody
	GetData() []*QuerySingleActivityInfoResponseBodyData
	SetHttpStatusCode(v string) *QuerySingleActivityInfoResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *QuerySingleActivityInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySingleActivityInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySingleActivityInfoResponseBody
	GetSuccess() *bool
}

type QuerySingleActivityInfoResponseBody struct {
	AccessDeniedDetail *string                                    `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               []*QuerySingleActivityInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode     *string                                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySingleActivityInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySingleActivityInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySingleActivityInfoResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QuerySingleActivityInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySingleActivityInfoResponseBody) GetData() []*QuerySingleActivityInfoResponseBodyData {
	return s.Data
}

func (s *QuerySingleActivityInfoResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *QuerySingleActivityInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySingleActivityInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySingleActivityInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySingleActivityInfoResponseBody) SetAccessDeniedDetail(v string) *QuerySingleActivityInfoResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBody) SetCode(v string) *QuerySingleActivityInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBody) SetData(v []*QuerySingleActivityInfoResponseBodyData) *QuerySingleActivityInfoResponseBody {
	s.Data = v
	return s
}

func (s *QuerySingleActivityInfoResponseBody) SetHttpStatusCode(v string) *QuerySingleActivityInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBody) SetMessage(v string) *QuerySingleActivityInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBody) SetRequestId(v string) *QuerySingleActivityInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBody) SetSuccess(v bool) *QuerySingleActivityInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBody) Validate() error {
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

type QuerySingleActivityInfoResponseBodyData struct {
	ActivityId    *int64  `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	ChannelName   *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	CompanyName   *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	CustomerName  *string `json:"CustomerName,omitempty" xml:"CustomerName,omitempty"`
	Email         *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IsVipCustomer *string `json:"IsVipCustomer,omitempty" xml:"IsVipCustomer,omitempty"`
	Mobile        *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	QRCode        *string `json:"QRCode,omitempty" xml:"QRCode,omitempty"`
	ReportFields  *string `json:"ReportFields,omitempty" xml:"ReportFields,omitempty"`
}

func (s QuerySingleActivityInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySingleActivityInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySingleActivityInfoResponseBodyData) GetActivityId() *int64 {
	return s.ActivityId
}

func (s *QuerySingleActivityInfoResponseBodyData) GetChannelName() *string {
	return s.ChannelName
}

func (s *QuerySingleActivityInfoResponseBodyData) GetCompanyName() *string {
	return s.CompanyName
}

func (s *QuerySingleActivityInfoResponseBodyData) GetCustomerName() *string {
	return s.CustomerName
}

func (s *QuerySingleActivityInfoResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *QuerySingleActivityInfoResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *QuerySingleActivityInfoResponseBodyData) GetIsVipCustomer() *string {
	return s.IsVipCustomer
}

func (s *QuerySingleActivityInfoResponseBodyData) GetMobile() *string {
	return s.Mobile
}

func (s *QuerySingleActivityInfoResponseBodyData) GetQRCode() *string {
	return s.QRCode
}

func (s *QuerySingleActivityInfoResponseBodyData) GetReportFields() *string {
	return s.ReportFields
}

func (s *QuerySingleActivityInfoResponseBodyData) SetActivityId(v int64) *QuerySingleActivityInfoResponseBodyData {
	s.ActivityId = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBodyData) SetChannelName(v string) *QuerySingleActivityInfoResponseBodyData {
	s.ChannelName = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBodyData) SetCompanyName(v string) *QuerySingleActivityInfoResponseBodyData {
	s.CompanyName = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBodyData) SetCustomerName(v string) *QuerySingleActivityInfoResponseBodyData {
	s.CustomerName = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBodyData) SetEmail(v string) *QuerySingleActivityInfoResponseBodyData {
	s.Email = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBodyData) SetId(v int64) *QuerySingleActivityInfoResponseBodyData {
	s.Id = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBodyData) SetIsVipCustomer(v string) *QuerySingleActivityInfoResponseBodyData {
	s.IsVipCustomer = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBodyData) SetMobile(v string) *QuerySingleActivityInfoResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBodyData) SetQRCode(v string) *QuerySingleActivityInfoResponseBodyData {
	s.QRCode = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBodyData) SetReportFields(v string) *QuerySingleActivityInfoResponseBodyData {
	s.ReportFields = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
