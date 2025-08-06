// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIStatisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDivision(v string) *GetAIStatisRequest
	GetDivision() *string
	SetEndTime(v int64) *GetAIStatisRequest
	GetEndTime() *int64
	SetEndTimeUTC(v string) *GetAIStatisRequest
	GetEndTimeUTC() *string
	SetLevel(v string) *GetAIStatisRequest
	GetLevel() *string
	SetOwnerAccount(v string) *GetAIStatisRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetAIStatisRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetAIStatisRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetAIStatisRequest
	GetResourceOwnerId() *string
	SetStartTime(v int64) *GetAIStatisRequest
	GetStartTime() *int64
	SetStartTimeUTC(v string) *GetAIStatisRequest
	GetStartTimeUTC() *string
	SetType(v string) *GetAIStatisRequest
	GetType() *string
}

type GetAIStatisRequest struct {
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
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAIStatisRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAIStatisRequest) GoString() string {
	return s.String()
}

func (s *GetAIStatisRequest) GetDivision() *string {
	return s.Division
}

func (s *GetAIStatisRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetAIStatisRequest) GetEndTimeUTC() *string {
	return s.EndTimeUTC
}

func (s *GetAIStatisRequest) GetLevel() *string {
	return s.Level
}

func (s *GetAIStatisRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetAIStatisRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetAIStatisRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetAIStatisRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetAIStatisRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetAIStatisRequest) GetStartTimeUTC() *string {
	return s.StartTimeUTC
}

func (s *GetAIStatisRequest) GetType() *string {
	return s.Type
}

func (s *GetAIStatisRequest) SetDivision(v string) *GetAIStatisRequest {
	s.Division = &v
	return s
}

func (s *GetAIStatisRequest) SetEndTime(v int64) *GetAIStatisRequest {
	s.EndTime = &v
	return s
}

func (s *GetAIStatisRequest) SetEndTimeUTC(v string) *GetAIStatisRequest {
	s.EndTimeUTC = &v
	return s
}

func (s *GetAIStatisRequest) SetLevel(v string) *GetAIStatisRequest {
	s.Level = &v
	return s
}

func (s *GetAIStatisRequest) SetOwnerAccount(v string) *GetAIStatisRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetAIStatisRequest) SetOwnerId(v string) *GetAIStatisRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAIStatisRequest) SetResourceOwnerAccount(v string) *GetAIStatisRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetAIStatisRequest) SetResourceOwnerId(v string) *GetAIStatisRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAIStatisRequest) SetStartTime(v int64) *GetAIStatisRequest {
	s.StartTime = &v
	return s
}

func (s *GetAIStatisRequest) SetStartTimeUTC(v string) *GetAIStatisRequest {
	s.StartTimeUTC = &v
	return s
}

func (s *GetAIStatisRequest) SetType(v string) *GetAIStatisRequest {
	s.Type = &v
	return s
}

func (s *GetAIStatisRequest) Validate() error {
	return dara.Validate(s)
}
