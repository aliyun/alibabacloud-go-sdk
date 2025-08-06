// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOSSFlowStatisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDivision(v string) *GetOSSFlowStatisRequest
	GetDivision() *string
	SetEndTime(v int64) *GetOSSFlowStatisRequest
	GetEndTime() *int64
	SetEndTimeUTC(v string) *GetOSSFlowStatisRequest
	GetEndTimeUTC() *string
	SetLevel(v string) *GetOSSFlowStatisRequest
	GetLevel() *string
	SetOwnerAccount(v string) *GetOSSFlowStatisRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetOSSFlowStatisRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetOSSFlowStatisRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetOSSFlowStatisRequest
	GetResourceOwnerId() *string
	SetStartTime(v int64) *GetOSSFlowStatisRequest
	GetStartTime() *int64
	SetStartTimeUTC(v string) *GetOSSFlowStatisRequest
	GetStartTimeUTC() *string
}

type GetOSSFlowStatisRequest struct {
	Division   *string `json:"Division,omitempty" xml:"Division,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimeUTC *string `json:"EndTimeUTC,omitempty" xml:"EndTimeUTC,omitempty"`
	// This parameter is required.
	Level                *string `json:"Level,omitempty" xml:"Level,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StartTimeUTC         *string `json:"StartTimeUTC,omitempty" xml:"StartTimeUTC,omitempty"`
}

func (s GetOSSFlowStatisRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOSSFlowStatisRequest) GoString() string {
	return s.String()
}

func (s *GetOSSFlowStatisRequest) GetDivision() *string {
	return s.Division
}

func (s *GetOSSFlowStatisRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetOSSFlowStatisRequest) GetEndTimeUTC() *string {
	return s.EndTimeUTC
}

func (s *GetOSSFlowStatisRequest) GetLevel() *string {
	return s.Level
}

func (s *GetOSSFlowStatisRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetOSSFlowStatisRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetOSSFlowStatisRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetOSSFlowStatisRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetOSSFlowStatisRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetOSSFlowStatisRequest) GetStartTimeUTC() *string {
	return s.StartTimeUTC
}

func (s *GetOSSFlowStatisRequest) SetDivision(v string) *GetOSSFlowStatisRequest {
	s.Division = &v
	return s
}

func (s *GetOSSFlowStatisRequest) SetEndTime(v int64) *GetOSSFlowStatisRequest {
	s.EndTime = &v
	return s
}

func (s *GetOSSFlowStatisRequest) SetEndTimeUTC(v string) *GetOSSFlowStatisRequest {
	s.EndTimeUTC = &v
	return s
}

func (s *GetOSSFlowStatisRequest) SetLevel(v string) *GetOSSFlowStatisRequest {
	s.Level = &v
	return s
}

func (s *GetOSSFlowStatisRequest) SetOwnerAccount(v string) *GetOSSFlowStatisRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetOSSFlowStatisRequest) SetOwnerId(v string) *GetOSSFlowStatisRequest {
	s.OwnerId = &v
	return s
}

func (s *GetOSSFlowStatisRequest) SetResourceOwnerAccount(v string) *GetOSSFlowStatisRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetOSSFlowStatisRequest) SetResourceOwnerId(v string) *GetOSSFlowStatisRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetOSSFlowStatisRequest) SetStartTime(v int64) *GetOSSFlowStatisRequest {
	s.StartTime = &v
	return s
}

func (s *GetOSSFlowStatisRequest) SetStartTimeUTC(v string) *GetOSSFlowStatisRequest {
	s.StartTimeUTC = &v
	return s
}

func (s *GetOSSFlowStatisRequest) Validate() error {
	return dara.Validate(s)
}
