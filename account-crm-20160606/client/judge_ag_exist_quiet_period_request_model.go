// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJudgeAgExistQuietPeriodRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgAccountType(v string) *JudgeAgExistQuietPeriodRequest
	GetAgAccountType() *string
	SetAppName(v string) *JudgeAgExistQuietPeriodRequest
	GetAppName() *string
	SetMpk(v string) *JudgeAgExistQuietPeriodRequest
	GetMpk() *string
	SetPk(v string) *JudgeAgExistQuietPeriodRequest
	GetPk() *string
}

type JudgeAgExistQuietPeriodRequest struct {
	// This parameter is required.
	AgAccountType *string `json:"AgAccountType,omitempty" xml:"AgAccountType,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s JudgeAgExistQuietPeriodRequest) String() string {
	return dara.Prettify(s)
}

func (s JudgeAgExistQuietPeriodRequest) GoString() string {
	return s.String()
}

func (s *JudgeAgExistQuietPeriodRequest) GetAgAccountType() *string {
	return s.AgAccountType
}

func (s *JudgeAgExistQuietPeriodRequest) GetAppName() *string {
	return s.AppName
}

func (s *JudgeAgExistQuietPeriodRequest) GetMpk() *string {
	return s.Mpk
}

func (s *JudgeAgExistQuietPeriodRequest) GetPk() *string {
	return s.Pk
}

func (s *JudgeAgExistQuietPeriodRequest) SetAgAccountType(v string) *JudgeAgExistQuietPeriodRequest {
	s.AgAccountType = &v
	return s
}

func (s *JudgeAgExistQuietPeriodRequest) SetAppName(v string) *JudgeAgExistQuietPeriodRequest {
	s.AppName = &v
	return s
}

func (s *JudgeAgExistQuietPeriodRequest) SetMpk(v string) *JudgeAgExistQuietPeriodRequest {
	s.Mpk = &v
	return s
}

func (s *JudgeAgExistQuietPeriodRequest) SetPk(v string) *JudgeAgExistQuietPeriodRequest {
	s.Pk = &v
	return s
}

func (s *JudgeAgExistQuietPeriodRequest) Validate() error {
	return dara.Validate(s)
}
