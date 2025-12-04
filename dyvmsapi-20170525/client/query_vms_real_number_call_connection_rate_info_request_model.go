// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVmsRealNumberCallConnectionRateInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QueryVmsRealNumberCallConnectionRateInfoRequest
	GetOwnerId() *int64
	SetRealNumber(v string) *QueryVmsRealNumberCallConnectionRateInfoRequest
	GetRealNumber() *string
	SetResourceOwnerAccount(v string) *QueryVmsRealNumberCallConnectionRateInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryVmsRealNumberCallConnectionRateInfoRequest
	GetResourceOwnerId() *int64
	SetTimePeriod(v string) *QueryVmsRealNumberCallConnectionRateInfoRequest
	GetTimePeriod() *string
	SetVirtualNumber(v string) *QueryVmsRealNumberCallConnectionRateInfoRequest
	GetVirtualNumber() *string
}

type QueryVmsRealNumberCallConnectionRateInfoRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 真实号码
	//
	// example:
	//
	// 131***1234
	RealNumber           *string `json:"RealNumber,omitempty" xml:"RealNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 时间段
	//
	// example:
	//
	// 1764666240123-1764686240567
	TimePeriod *string `json:"TimePeriod,omitempty" xml:"TimePeriod,omitempty"`
	// 虚拟号码
	//
	// example:
	//
	// 0571***1234
	VirtualNumber *string `json:"VirtualNumber,omitempty" xml:"VirtualNumber,omitempty"`
}

func (s QueryVmsRealNumberCallConnectionRateInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryVmsRealNumberCallConnectionRateInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryVmsRealNumberCallConnectionRateInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryVmsRealNumberCallConnectionRateInfoRequest) GetRealNumber() *string {
	return s.RealNumber
}

func (s *QueryVmsRealNumberCallConnectionRateInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryVmsRealNumberCallConnectionRateInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryVmsRealNumberCallConnectionRateInfoRequest) GetTimePeriod() *string {
	return s.TimePeriod
}

func (s *QueryVmsRealNumberCallConnectionRateInfoRequest) GetVirtualNumber() *string {
	return s.VirtualNumber
}

func (s *QueryVmsRealNumberCallConnectionRateInfoRequest) SetOwnerId(v int64) *QueryVmsRealNumberCallConnectionRateInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoRequest) SetRealNumber(v string) *QueryVmsRealNumberCallConnectionRateInfoRequest {
	s.RealNumber = &v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoRequest) SetResourceOwnerAccount(v string) *QueryVmsRealNumberCallConnectionRateInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoRequest) SetResourceOwnerId(v int64) *QueryVmsRealNumberCallConnectionRateInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoRequest) SetTimePeriod(v string) *QueryVmsRealNumberCallConnectionRateInfoRequest {
	s.TimePeriod = &v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoRequest) SetVirtualNumber(v string) *QueryVmsRealNumberCallConnectionRateInfoRequest {
	s.VirtualNumber = &v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoRequest) Validate() error {
	return dara.Validate(s)
}
