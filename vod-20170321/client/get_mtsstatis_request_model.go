// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMTSStatisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDivision(v string) *GetMTSStatisRequest
	GetDivision() *string
	SetEndTime(v int64) *GetMTSStatisRequest
	GetEndTime() *int64
	SetEndTimeUTC(v string) *GetMTSStatisRequest
	GetEndTimeUTC() *string
	SetLevel(v string) *GetMTSStatisRequest
	GetLevel() *string
	SetOwnerAccount(v string) *GetMTSStatisRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetMTSStatisRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetMTSStatisRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetMTSStatisRequest
	GetResourceOwnerId() *string
	SetStartTime(v int64) *GetMTSStatisRequest
	GetStartTime() *int64
	SetStartTimeUTC(v string) *GetMTSStatisRequest
	GetStartTimeUTC() *string
}

type GetMTSStatisRequest struct {
	Division             *string `json:"Division,omitempty" xml:"Division,omitempty"`
	EndTime              *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimeUTC           *string `json:"EndTimeUTC,omitempty" xml:"EndTimeUTC,omitempty"`
	Level                *string `json:"Level,omitempty" xml:"Level,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StartTimeUTC         *string `json:"StartTimeUTC,omitempty" xml:"StartTimeUTC,omitempty"`
}

func (s GetMTSStatisRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMTSStatisRequest) GoString() string {
	return s.String()
}

func (s *GetMTSStatisRequest) GetDivision() *string {
	return s.Division
}

func (s *GetMTSStatisRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetMTSStatisRequest) GetEndTimeUTC() *string {
	return s.EndTimeUTC
}

func (s *GetMTSStatisRequest) GetLevel() *string {
	return s.Level
}

func (s *GetMTSStatisRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetMTSStatisRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetMTSStatisRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetMTSStatisRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetMTSStatisRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetMTSStatisRequest) GetStartTimeUTC() *string {
	return s.StartTimeUTC
}

func (s *GetMTSStatisRequest) SetDivision(v string) *GetMTSStatisRequest {
	s.Division = &v
	return s
}

func (s *GetMTSStatisRequest) SetEndTime(v int64) *GetMTSStatisRequest {
	s.EndTime = &v
	return s
}

func (s *GetMTSStatisRequest) SetEndTimeUTC(v string) *GetMTSStatisRequest {
	s.EndTimeUTC = &v
	return s
}

func (s *GetMTSStatisRequest) SetLevel(v string) *GetMTSStatisRequest {
	s.Level = &v
	return s
}

func (s *GetMTSStatisRequest) SetOwnerAccount(v string) *GetMTSStatisRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetMTSStatisRequest) SetOwnerId(v string) *GetMTSStatisRequest {
	s.OwnerId = &v
	return s
}

func (s *GetMTSStatisRequest) SetResourceOwnerAccount(v string) *GetMTSStatisRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetMTSStatisRequest) SetResourceOwnerId(v string) *GetMTSStatisRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetMTSStatisRequest) SetStartTime(v int64) *GetMTSStatisRequest {
	s.StartTime = &v
	return s
}

func (s *GetMTSStatisRequest) SetStartTimeUTC(v string) *GetMTSStatisRequest {
	s.StartTimeUTC = &v
	return s
}

func (s *GetMTSStatisRequest) Validate() error {
	return dara.Validate(s)
}
