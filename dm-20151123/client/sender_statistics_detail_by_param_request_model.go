// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSenderStatisticsDetailByParamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *SenderStatisticsDetailByParamRequest
	GetAccountName() *string
	SetConfigSetId(v string) *SenderStatisticsDetailByParamRequest
	GetConfigSetId() *string
	SetEndTime(v string) *SenderStatisticsDetailByParamRequest
	GetEndTime() *string
	SetIpPoolId(v string) *SenderStatisticsDetailByParamRequest
	GetIpPoolId() *string
	SetLength(v int32) *SenderStatisticsDetailByParamRequest
	GetLength() *int32
	SetNextStart(v string) *SenderStatisticsDetailByParamRequest
	GetNextStart() *string
	SetOwnerId(v int64) *SenderStatisticsDetailByParamRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SenderStatisticsDetailByParamRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SenderStatisticsDetailByParamRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *SenderStatisticsDetailByParamRequest
	GetStartTime() *string
	SetStatus(v int32) *SenderStatisticsDetailByParamRequest
	GetStatus() *int32
	SetTagName(v string) *SenderStatisticsDetailByParamRequest
	GetTagName() *string
	SetToAddress(v string) *SenderStatisticsDetailByParamRequest
	GetToAddress() *string
}

type SenderStatisticsDetailByParamRequest struct {
	// The sender address. If you leave this parameter empty, emails from all sender addresses are queried.
	//
	// > You can leave **AccountName**, **TagName**, and **ToAddress*	- empty. You can specify a value for only one of these parameters at a time.
	//
	// example:
	//
	// s***@example.net
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the configuration set.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 95437e39-5433-4c6c-8b66-0c4eac2f4d97
	ConfigSetId *string `json:"ConfigSetId,omitempty" xml:"ConfigSetId,omitempty"`
	// The end time. The time range between the start time and the end time cannot exceed 30 days. The format is yyyy-MM-dd HH:mm.
	//
	// example:
	//
	// 2021-04-29 00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the independent IP address pool.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// e42a4d8e-66e0-4ea3-88aa-74c66c661cef
	IpPoolId *string `json:"IpPoolId,omitempty" xml:"IpPoolId,omitempty"`
	// The number of results to return. The value can range from 1 to 100.
	//
	// example:
	//
	// 5
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// The offset for this request, used for paging. If more results are available, set the \\`NextStart\\` parameter in your next request to this return value.
	//
	// example:
	//
	// 90f0243616#203#a***@example.net-1658817837#a***@example.net.247475288187
	NextStart            *string `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The start time. The time range between the start time and the end time cannot exceed 30 days. The format is yyyy-MM-dd HH:mm.
	//
	// example:
	//
	// 2021-04-28 00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The delivery status. If you leave this parameter empty, emails in all delivery statuses are queried. Valid values:
	//
	// - 0: delivered successfully
	//
	// - 2: invalid address
	//
	// - 3: spam
	//
	// - 4: failed
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The email tag. If you leave this parameter empty, emails with all tags are queried.
	//
	// When you send an email using SMTP, specify the \\`TagName\\` and its value for the \\`X-AliDM-Trace\\` field. For more information, see the examples for email tracking.
	//
	// example:
	//
	// EmailQuestionnaireHelioscam
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The recipient address. If you leave this parameter empty, emails sent to all recipient addresses are queried.
	//
	// example:
	//
	// b***@example.net
	ToAddress *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
}

func (s SenderStatisticsDetailByParamRequest) String() string {
	return dara.Prettify(s)
}

func (s SenderStatisticsDetailByParamRequest) GoString() string {
	return s.String()
}

func (s *SenderStatisticsDetailByParamRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *SenderStatisticsDetailByParamRequest) GetConfigSetId() *string {
	return s.ConfigSetId
}

func (s *SenderStatisticsDetailByParamRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *SenderStatisticsDetailByParamRequest) GetIpPoolId() *string {
	return s.IpPoolId
}

func (s *SenderStatisticsDetailByParamRequest) GetLength() *int32 {
	return s.Length
}

func (s *SenderStatisticsDetailByParamRequest) GetNextStart() *string {
	return s.NextStart
}

func (s *SenderStatisticsDetailByParamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SenderStatisticsDetailByParamRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SenderStatisticsDetailByParamRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SenderStatisticsDetailByParamRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *SenderStatisticsDetailByParamRequest) GetStatus() *int32 {
	return s.Status
}

func (s *SenderStatisticsDetailByParamRequest) GetTagName() *string {
	return s.TagName
}

func (s *SenderStatisticsDetailByParamRequest) GetToAddress() *string {
	return s.ToAddress
}

func (s *SenderStatisticsDetailByParamRequest) SetAccountName(v string) *SenderStatisticsDetailByParamRequest {
	s.AccountName = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetConfigSetId(v string) *SenderStatisticsDetailByParamRequest {
	s.ConfigSetId = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetEndTime(v string) *SenderStatisticsDetailByParamRequest {
	s.EndTime = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetIpPoolId(v string) *SenderStatisticsDetailByParamRequest {
	s.IpPoolId = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetLength(v int32) *SenderStatisticsDetailByParamRequest {
	s.Length = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetNextStart(v string) *SenderStatisticsDetailByParamRequest {
	s.NextStart = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetOwnerId(v int64) *SenderStatisticsDetailByParamRequest {
	s.OwnerId = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetResourceOwnerAccount(v string) *SenderStatisticsDetailByParamRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetResourceOwnerId(v int64) *SenderStatisticsDetailByParamRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetStartTime(v string) *SenderStatisticsDetailByParamRequest {
	s.StartTime = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetStatus(v int32) *SenderStatisticsDetailByParamRequest {
	s.Status = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetTagName(v string) *SenderStatisticsDetailByParamRequest {
	s.TagName = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetToAddress(v string) *SenderStatisticsDetailByParamRequest {
	s.ToAddress = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) Validate() error {
	return dara.Validate(s)
}
