// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountSecurityPracticeReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountSecurityPracticeInfo(v *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo) *GetAccountSecurityPracticeReportResponseBody
	GetAccountSecurityPracticeInfo() *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo
	SetRequestId(v string) *GetAccountSecurityPracticeReportResponseBody
	GetRequestId() *string
}

type GetAccountSecurityPracticeReportResponseBody struct {
	AccountSecurityPracticeInfo *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo `json:"AccountSecurityPracticeInfo,omitempty" xml:"AccountSecurityPracticeInfo,omitempty" type:"Struct"`
	RequestId                   *string                                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccountSecurityPracticeReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccountSecurityPracticeReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountSecurityPracticeReportResponseBody) GetAccountSecurityPracticeInfo() *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo {
	return s.AccountSecurityPracticeInfo
}

func (s *GetAccountSecurityPracticeReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccountSecurityPracticeReportResponseBody) SetAccountSecurityPracticeInfo(v *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo) *GetAccountSecurityPracticeReportResponseBody {
	s.AccountSecurityPracticeInfo = v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBody) SetRequestId(v string) *GetAccountSecurityPracticeReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBody) Validate() error {
	if s.AccountSecurityPracticeInfo != nil {
		if err := s.AccountSecurityPracticeInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo struct {
	AccountSecurityPracticeUserInfo *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo `json:"AccountSecurityPracticeUserInfo,omitempty" xml:"AccountSecurityPracticeUserInfo,omitempty" type:"Struct"`
	Score                           *int32                                                                                                  `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo) GoString() string {
	return s.String()
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo) GetAccountSecurityPracticeUserInfo() *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	return s.AccountSecurityPracticeUserInfo
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo) GetScore() *int32 {
	return s.Score
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo) SetAccountSecurityPracticeUserInfo(v *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo {
	s.AccountSecurityPracticeUserInfo = v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo) SetScore(v int32) *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo {
	s.Score = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo) Validate() error {
	if s.AccountSecurityPracticeUserInfo != nil {
		if err := s.AccountSecurityPracticeUserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo struct {
	BindMfa                    *bool   `json:"BindMfa,omitempty" xml:"BindMfa,omitempty"`
	OldAkNum                   *int32  `json:"OldAkNum,omitempty" xml:"OldAkNum,omitempty"`
	RootWithAccessKey          *int32  `json:"RootWithAccessKey,omitempty" xml:"RootWithAccessKey,omitempty"`
	SubUser                    *int32  `json:"SubUser,omitempty" xml:"SubUser,omitempty"`
	SubUserBindMfa             *int32  `json:"SubUserBindMfa,omitempty" xml:"SubUserBindMfa,omitempty"`
	SubUserPwdLevel            *string `json:"SubUserPwdLevel,omitempty" xml:"SubUserPwdLevel,omitempty"`
	SubUserWithOldAccessKey    *int32  `json:"SubUserWithOldAccessKey,omitempty" xml:"SubUserWithOldAccessKey,omitempty"`
	SubUserWithUnusedAccessKey *int32  `json:"SubUserWithUnusedAccessKey,omitempty" xml:"SubUserWithUnusedAccessKey,omitempty"`
	UnusedAkNum                *int32  `json:"UnusedAkNum,omitempty" xml:"UnusedAkNum,omitempty"`
}

func (s GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) GoString() string {
	return s.String()
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) GetBindMfa() *bool {
	return s.BindMfa
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) GetOldAkNum() *int32 {
	return s.OldAkNum
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) GetRootWithAccessKey() *int32 {
	return s.RootWithAccessKey
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) GetSubUser() *int32 {
	return s.SubUser
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) GetSubUserBindMfa() *int32 {
	return s.SubUserBindMfa
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) GetSubUserPwdLevel() *string {
	return s.SubUserPwdLevel
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) GetSubUserWithOldAccessKey() *int32 {
	return s.SubUserWithOldAccessKey
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) GetSubUserWithUnusedAccessKey() *int32 {
	return s.SubUserWithUnusedAccessKey
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) GetUnusedAkNum() *int32 {
	return s.UnusedAkNum
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetBindMfa(v bool) *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.BindMfa = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetOldAkNum(v int32) *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.OldAkNum = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetRootWithAccessKey(v int32) *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.RootWithAccessKey = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetSubUser(v int32) *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.SubUser = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetSubUserBindMfa(v int32) *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.SubUserBindMfa = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetSubUserPwdLevel(v string) *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.SubUserPwdLevel = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetSubUserWithOldAccessKey(v int32) *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.SubUserWithOldAccessKey = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetSubUserWithUnusedAccessKey(v int32) *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.SubUserWithUnusedAccessKey = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetUnusedAkNum(v int32) *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.UnusedAkNum = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) Validate() error {
	return dara.Validate(s)
}
