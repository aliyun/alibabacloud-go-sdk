// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSenderStatisticsByTagNameAndBatchIDRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *SenderStatisticsByTagNameAndBatchIDRequest
	GetAccountName() *string
	SetDedicatedIp(v string) *SenderStatisticsByTagNameAndBatchIDRequest
	GetDedicatedIp() *string
	SetDedicatedIpPoolId(v string) *SenderStatisticsByTagNameAndBatchIDRequest
	GetDedicatedIpPoolId() *string
	SetEndTime(v string) *SenderStatisticsByTagNameAndBatchIDRequest
	GetEndTime() *string
	SetEsp(v string) *SenderStatisticsByTagNameAndBatchIDRequest
	GetEsp() *string
	SetOwnerId(v int64) *SenderStatisticsByTagNameAndBatchIDRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SenderStatisticsByTagNameAndBatchIDRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SenderStatisticsByTagNameAndBatchIDRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *SenderStatisticsByTagNameAndBatchIDRequest
	GetStartTime() *string
	SetTagName(v string) *SenderStatisticsByTagNameAndBatchIDRequest
	GetTagName() *string
}

type SenderStatisticsByTagNameAndBatchIDRequest struct {
	// Sending address. If not filled, it represents all addresses.
	//
	// example:
	//
	// xxx
	AccountName       *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DedicatedIp       *string `json:"DedicatedIp,omitempty" xml:"DedicatedIp,omitempty"`
	DedicatedIpPoolId *string `json:"DedicatedIpPoolId,omitempty" xml:"DedicatedIpPoolId,omitempty"`
	// End time, which cannot exceed 7 days from the start time, in the format yyyy-MM-dd.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-29
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Esp                  *string `json:"Esp,omitempty" xml:"Esp,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Start time, in the format yyyy-MM-dd.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-29
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Email tag. If not filled, it represents all tags.
	//
	// example:
	//
	// xxx
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s SenderStatisticsByTagNameAndBatchIDRequest) String() string {
	return dara.Prettify(s)
}

func (s SenderStatisticsByTagNameAndBatchIDRequest) GoString() string {
	return s.String()
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) GetDedicatedIp() *string {
	return s.DedicatedIp
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) GetDedicatedIpPoolId() *string {
	return s.DedicatedIpPoolId
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) GetEsp() *string {
	return s.Esp
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) GetTagName() *string {
	return s.TagName
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) SetAccountName(v string) *SenderStatisticsByTagNameAndBatchIDRequest {
	s.AccountName = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) SetDedicatedIp(v string) *SenderStatisticsByTagNameAndBatchIDRequest {
	s.DedicatedIp = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) SetDedicatedIpPoolId(v string) *SenderStatisticsByTagNameAndBatchIDRequest {
	s.DedicatedIpPoolId = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) SetEndTime(v string) *SenderStatisticsByTagNameAndBatchIDRequest {
	s.EndTime = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) SetEsp(v string) *SenderStatisticsByTagNameAndBatchIDRequest {
	s.Esp = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) SetOwnerId(v int64) *SenderStatisticsByTagNameAndBatchIDRequest {
	s.OwnerId = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) SetResourceOwnerAccount(v string) *SenderStatisticsByTagNameAndBatchIDRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) SetResourceOwnerId(v int64) *SenderStatisticsByTagNameAndBatchIDRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) SetStartTime(v string) *SenderStatisticsByTagNameAndBatchIDRequest {
	s.StartTime = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) SetTagName(v string) *SenderStatisticsByTagNameAndBatchIDRequest {
	s.TagName = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) Validate() error {
	return dara.Validate(s)
}
