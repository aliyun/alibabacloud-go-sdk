// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOSSStatisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDivision(v string) *GetOSSStatisRequest
	GetDivision() *string
	SetEndTime(v int64) *GetOSSStatisRequest
	GetEndTime() *int64
	SetEndTimeUTC(v string) *GetOSSStatisRequest
	GetEndTimeUTC() *string
	SetLevel(v string) *GetOSSStatisRequest
	GetLevel() *string
	SetOwnerAccount(v string) *GetOSSStatisRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetOSSStatisRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetOSSStatisRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetOSSStatisRequest
	GetResourceOwnerId() *string
	SetStartTime(v int64) *GetOSSStatisRequest
	GetStartTime() *int64
	SetStartTimeUTC(v string) *GetOSSStatisRequest
	GetStartTimeUTC() *string
}

type GetOSSStatisRequest struct {
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

func (s GetOSSStatisRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOSSStatisRequest) GoString() string {
	return s.String()
}

func (s *GetOSSStatisRequest) GetDivision() *string {
	return s.Division
}

func (s *GetOSSStatisRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetOSSStatisRequest) GetEndTimeUTC() *string {
	return s.EndTimeUTC
}

func (s *GetOSSStatisRequest) GetLevel() *string {
	return s.Level
}

func (s *GetOSSStatisRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetOSSStatisRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetOSSStatisRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetOSSStatisRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetOSSStatisRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetOSSStatisRequest) GetStartTimeUTC() *string {
	return s.StartTimeUTC
}

func (s *GetOSSStatisRequest) SetDivision(v string) *GetOSSStatisRequest {
	s.Division = &v
	return s
}

func (s *GetOSSStatisRequest) SetEndTime(v int64) *GetOSSStatisRequest {
	s.EndTime = &v
	return s
}

func (s *GetOSSStatisRequest) SetEndTimeUTC(v string) *GetOSSStatisRequest {
	s.EndTimeUTC = &v
	return s
}

func (s *GetOSSStatisRequest) SetLevel(v string) *GetOSSStatisRequest {
	s.Level = &v
	return s
}

func (s *GetOSSStatisRequest) SetOwnerAccount(v string) *GetOSSStatisRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetOSSStatisRequest) SetOwnerId(v string) *GetOSSStatisRequest {
	s.OwnerId = &v
	return s
}

func (s *GetOSSStatisRequest) SetResourceOwnerAccount(v string) *GetOSSStatisRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetOSSStatisRequest) SetResourceOwnerId(v string) *GetOSSStatisRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetOSSStatisRequest) SetStartTime(v int64) *GetOSSStatisRequest {
	s.StartTime = &v
	return s
}

func (s *GetOSSStatisRequest) SetStartTimeUTC(v string) *GetOSSStatisRequest {
	s.StartTimeUTC = &v
	return s
}

func (s *GetOSSStatisRequest) Validate() error {
	return dara.Validate(s)
}
