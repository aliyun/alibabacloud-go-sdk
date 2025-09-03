// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSenderStatisticsDetailByParamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextStart(v string) *SenderStatisticsDetailByParamResponseBody
	GetNextStart() *string
	SetRequestId(v string) *SenderStatisticsDetailByParamResponseBody
	GetRequestId() *string
	SetData(v *SenderStatisticsDetailByParamResponseBodyData) *SenderStatisticsDetailByParamResponseBody
	GetData() *SenderStatisticsDetailByParamResponseBodyData
}

type SenderStatisticsDetailByParamResponseBody struct {
	// Used for pagination. If there are more results, set this returned value to the NextStart in the next request.
	//
	// example:
	//
	// 90f0243616#203#a***@example.net-1658817689#a***@example.net.247141122178
	NextStart *string `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
	// Request ID
	//
	// example:
	//
	// B5AB8EBB-EE64-4BB2-B085-B92CC5DEDC41
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Detailed records
	Data *SenderStatisticsDetailByParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s SenderStatisticsDetailByParamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SenderStatisticsDetailByParamResponseBody) GoString() string {
	return s.String()
}

func (s *SenderStatisticsDetailByParamResponseBody) GetNextStart() *string {
	return s.NextStart
}

func (s *SenderStatisticsDetailByParamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SenderStatisticsDetailByParamResponseBody) GetData() *SenderStatisticsDetailByParamResponseBodyData {
	return s.Data
}

func (s *SenderStatisticsDetailByParamResponseBody) SetNextStart(v string) *SenderStatisticsDetailByParamResponseBody {
	s.NextStart = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBody) SetRequestId(v string) *SenderStatisticsDetailByParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBody) SetData(v *SenderStatisticsDetailByParamResponseBodyData) *SenderStatisticsDetailByParamResponseBody {
	s.Data = v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBody) Validate() error {
	return dara.Validate(s)
}

type SenderStatisticsDetailByParamResponseBodyData struct {
	MailDetail []*SenderStatisticsDetailByParamResponseBodyDataMailDetail `json:"mailDetail,omitempty" xml:"mailDetail,omitempty" type:"Repeated"`
}

func (s SenderStatisticsDetailByParamResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SenderStatisticsDetailByParamResponseBodyData) GoString() string {
	return s.String()
}

func (s *SenderStatisticsDetailByParamResponseBodyData) GetMailDetail() []*SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	return s.MailDetail
}

func (s *SenderStatisticsDetailByParamResponseBodyData) SetMailDetail(v []*SenderStatisticsDetailByParamResponseBodyDataMailDetail) *SenderStatisticsDetailByParamResponseBodyData {
	s.MailDetail = v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type SenderStatisticsDetailByParamResponseBodyDataMailDetail struct {
	// Sending address
	//
	// example:
	//
	// s***@example.net
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// if can be null:
	// true
	ConfigSetId *string `json:"ConfigSetId,omitempty" xml:"ConfigSetId,omitempty"`
	// if can be null:
	// true
	ConfigSetName *string `json:"ConfigSetName,omitempty" xml:"ConfigSetName,omitempty"`
	// Detailed classification of error reasons: - SendOk - SmtpNxBox
	//
	// etc.
	//
	// example:
	//
	// SendOk
	ErrorClassification *string `json:"ErrorClassification,omitempty" xml:"ErrorClassification,omitempty"`
	// if can be null:
	// true
	IpPoolId *string `json:"IpPoolId,omitempty" xml:"IpPoolId,omitempty"`
	// if can be null:
	// true
	IpPoolName *string `json:"IpPoolName,omitempty" xml:"IpPoolName,omitempty"`
	// Update time
	//
	// example:
	//
	// 2021-04-28T17:11Z
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	// Delivery detail information
	//
	// example:
	//
	// 250 Send Mail OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Delivery status: 0 Success, 2 Invalid Address, 3 Spam, 4 Other Failures
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Email subject
	//
	// example:
	//
	// test subject
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// Recipient address
	//
	// example:
	//
	// b***@example.net
	ToAddress *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
	// UTC formatted update time
	//
	// example:
	//
	// 1619601108
	UtcLastUpdateTime *string `json:"UtcLastUpdateTime,omitempty" xml:"UtcLastUpdateTime,omitempty"`
}

func (s SenderStatisticsDetailByParamResponseBodyDataMailDetail) String() string {
	return dara.Prettify(s)
}

func (s SenderStatisticsDetailByParamResponseBodyDataMailDetail) GoString() string {
	return s.String()
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) GetAccountName() *string {
	return s.AccountName
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) GetConfigSetId() *string {
	return s.ConfigSetId
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) GetConfigSetName() *string {
	return s.ConfigSetName
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) GetErrorClassification() *string {
	return s.ErrorClassification
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) GetIpPoolId() *string {
	return s.IpPoolId
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) GetIpPoolName() *string {
	return s.IpPoolName
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) GetLastUpdateTime() *string {
	return s.LastUpdateTime
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) GetMessage() *string {
	return s.Message
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) GetStatus() *int32 {
	return s.Status
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) GetSubject() *string {
	return s.Subject
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) GetToAddress() *string {
	return s.ToAddress
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) GetUtcLastUpdateTime() *string {
	return s.UtcLastUpdateTime
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetAccountName(v string) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.AccountName = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetConfigSetId(v string) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.ConfigSetId = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetConfigSetName(v string) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.ConfigSetName = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetErrorClassification(v string) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.ErrorClassification = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetIpPoolId(v string) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.IpPoolId = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetIpPoolName(v string) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.IpPoolName = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetLastUpdateTime(v string) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.LastUpdateTime = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetMessage(v string) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.Message = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetStatus(v int32) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.Status = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetSubject(v string) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.Subject = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetToAddress(v string) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.ToAddress = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetUtcLastUpdateTime(v string) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.UtcLastUpdateTime = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) Validate() error {
	return dara.Validate(s)
}
