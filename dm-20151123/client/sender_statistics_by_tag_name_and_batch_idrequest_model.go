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
	SetDomain(v string) *SenderStatisticsByTagNameAndBatchIDRequest
	GetDomain() *string
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
	// The sender address. If this parameter is not specified, data for all addresses is returned.
	//
	// example:
	//
	// xxx
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// For dedicated IP users, specifies a dedicated IP address to query data for.
	//
	// If this parameter is not specified, data for all dedicated IP addresses is returned.
	//
	// example:
	//
	// xxx.xxx.xxx.xxx
	DedicatedIp *string `json:"DedicatedIp,omitempty" xml:"DedicatedIp,omitempty"`
	// For dedicated IP users, specifies the ID of a dedicated IP pool to query data for.
	//
	// If this parameter is not specified, data for all dedicated IP pools is returned.
	//
	// example:
	//
	// xxx
	DedicatedIpPoolId *string `json:"DedicatedIpPoolId,omitempty" xml:"DedicatedIpPoolId,omitempty"`
	// The sender domain name.
	//
	// example:
	//
	// dmdomain.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The end time. The time span between the start time and end time cannot exceed 31 days. Format: yyyy-MM-dd.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-29
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// For dedicated IP users, specifies a particular ESP to query data for. Valid values:
	//
	// - gmail.com
	//
	// - yahoo.com
	//
	// - outlook.com
	//
	// - icloud.com
	//
	// - others: data for ESPs other than the ones listed above.
	//
	// If this parameter is not specified, data for all ESPs is returned.
	//
	// example:
	//
	// gmail.com
	Esp                  *string `json:"Esp,omitempty" xml:"Esp,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The start time. The time cannot be earlier than 90 days ago. Format: yyyy-MM-dd.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-29
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The email tag. If this parameter is not specified, data for all tags is returned.
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

func (s *SenderStatisticsByTagNameAndBatchIDRequest) GetDomain() *string {
	return s.Domain
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

func (s *SenderStatisticsByTagNameAndBatchIDRequest) SetDomain(v string) *SenderStatisticsByTagNameAndBatchIDRequest {
	s.Domain = &v
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
