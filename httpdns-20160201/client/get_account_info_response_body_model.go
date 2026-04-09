// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountInfo(v *GetAccountInfoResponseBodyAccountInfo) *GetAccountInfoResponseBody
	GetAccountInfo() *GetAccountInfoResponseBodyAccountInfo
	SetRequestId(v string) *GetAccountInfoResponseBody
	GetRequestId() *string
}

type GetAccountInfoResponseBody struct {
	AccountInfo *GetAccountInfoResponseBodyAccountInfo `json:"AccountInfo,omitempty" xml:"AccountInfo,omitempty" type:"Struct"`
	// example:
	//
	// 50F9C40E-188D-4208-BE2C-74271337****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccountInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccountInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountInfoResponseBody) GetAccountInfo() *GetAccountInfoResponseBodyAccountInfo {
	return s.AccountInfo
}

func (s *GetAccountInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccountInfoResponseBody) SetAccountInfo(v *GetAccountInfoResponseBodyAccountInfo) *GetAccountInfoResponseBody {
	s.AccountInfo = v
	return s
}

func (s *GetAccountInfoResponseBody) SetRequestId(v string) *GetAccountInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountInfoResponseBody) Validate() error {
	if s.AccountInfo != nil {
		if err := s.AccountInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAccountInfoResponseBodyAccountInfo struct {
	// example:
	//
	// 1337****
	AccountId            *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	DohEnabled           *bool   `json:"DohEnabled,omitempty" xml:"DohEnabled,omitempty"`
	DohResolveAllEnabled *bool   `json:"DohResolveAllEnabled,omitempty" xml:"DohResolveAllEnabled,omitempty"`
	MonthDohResolveCount *int64  `json:"MonthDohResolveCount,omitempty" xml:"MonthDohResolveCount,omitempty"`
	// example:
	//
	// 1500000
	MonthFreeCount            *int32 `json:"MonthFreeCount,omitempty" xml:"MonthFreeCount,omitempty"`
	MonthHttpAesResolveCount  *int64 `json:"MonthHttpAesResolveCount,omitempty" xml:"MonthHttpAesResolveCount,omitempty"`
	MonthHttpsAesResolveCount *int64 `json:"MonthHttpsAesResolveCount,omitempty" xml:"MonthHttpsAesResolveCount,omitempty"`
	// example:
	//
	// 3
	MonthHttpsResolveCount *int32 `json:"MonthHttpsResolveCount,omitempty" xml:"MonthHttpsResolveCount,omitempty"`
	// example:
	//
	// 9927326
	MonthResolveCount *int32 `json:"MonthResolveCount,omitempty" xml:"MonthResolveCount,omitempty"`
	// example:
	//
	// 0
	PackageCount *int32 `json:"PackageCount,omitempty" xml:"PackageCount,omitempty"`
	// example:
	//
	// 50F9C40E****
	SignSecret *string `json:"SignSecret,omitempty" xml:"SignSecret,omitempty"`
	// example:
	//
	// 611523
	SignedCount *int64 `json:"SignedCount,omitempty" xml:"SignedCount,omitempty"`
	// example:
	//
	// 1523
	UnsignedCount *int64 `json:"UnsignedCount,omitempty" xml:"UnsignedCount,omitempty"`
	// example:
	//
	// true
	UnsignedEnabled *bool `json:"UnsignedEnabled,omitempty" xml:"UnsignedEnabled,omitempty"`
	// example:
	//
	// 1
	UserStatus *int32 `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s GetAccountInfoResponseBodyAccountInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAccountInfoResponseBodyAccountInfo) GoString() string {
	return s.String()
}

func (s *GetAccountInfoResponseBodyAccountInfo) GetAccountId() *string {
	return s.AccountId
}

func (s *GetAccountInfoResponseBodyAccountInfo) GetDohEnabled() *bool {
	return s.DohEnabled
}

func (s *GetAccountInfoResponseBodyAccountInfo) GetDohResolveAllEnabled() *bool {
	return s.DohResolveAllEnabled
}

func (s *GetAccountInfoResponseBodyAccountInfo) GetMonthDohResolveCount() *int64 {
	return s.MonthDohResolveCount
}

func (s *GetAccountInfoResponseBodyAccountInfo) GetMonthFreeCount() *int32 {
	return s.MonthFreeCount
}

func (s *GetAccountInfoResponseBodyAccountInfo) GetMonthHttpAesResolveCount() *int64 {
	return s.MonthHttpAesResolveCount
}

func (s *GetAccountInfoResponseBodyAccountInfo) GetMonthHttpsAesResolveCount() *int64 {
	return s.MonthHttpsAesResolveCount
}

func (s *GetAccountInfoResponseBodyAccountInfo) GetMonthHttpsResolveCount() *int32 {
	return s.MonthHttpsResolveCount
}

func (s *GetAccountInfoResponseBodyAccountInfo) GetMonthResolveCount() *int32 {
	return s.MonthResolveCount
}

func (s *GetAccountInfoResponseBodyAccountInfo) GetPackageCount() *int32 {
	return s.PackageCount
}

func (s *GetAccountInfoResponseBodyAccountInfo) GetSignSecret() *string {
	return s.SignSecret
}

func (s *GetAccountInfoResponseBodyAccountInfo) GetSignedCount() *int64 {
	return s.SignedCount
}

func (s *GetAccountInfoResponseBodyAccountInfo) GetUnsignedCount() *int64 {
	return s.UnsignedCount
}

func (s *GetAccountInfoResponseBodyAccountInfo) GetUnsignedEnabled() *bool {
	return s.UnsignedEnabled
}

func (s *GetAccountInfoResponseBodyAccountInfo) GetUserStatus() *int32 {
	return s.UserStatus
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetAccountId(v string) *GetAccountInfoResponseBodyAccountInfo {
	s.AccountId = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetDohEnabled(v bool) *GetAccountInfoResponseBodyAccountInfo {
	s.DohEnabled = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetDohResolveAllEnabled(v bool) *GetAccountInfoResponseBodyAccountInfo {
	s.DohResolveAllEnabled = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetMonthDohResolveCount(v int64) *GetAccountInfoResponseBodyAccountInfo {
	s.MonthDohResolveCount = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetMonthFreeCount(v int32) *GetAccountInfoResponseBodyAccountInfo {
	s.MonthFreeCount = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetMonthHttpAesResolveCount(v int64) *GetAccountInfoResponseBodyAccountInfo {
	s.MonthHttpAesResolveCount = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetMonthHttpsAesResolveCount(v int64) *GetAccountInfoResponseBodyAccountInfo {
	s.MonthHttpsAesResolveCount = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetMonthHttpsResolveCount(v int32) *GetAccountInfoResponseBodyAccountInfo {
	s.MonthHttpsResolveCount = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetMonthResolveCount(v int32) *GetAccountInfoResponseBodyAccountInfo {
	s.MonthResolveCount = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetPackageCount(v int32) *GetAccountInfoResponseBodyAccountInfo {
	s.PackageCount = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetSignSecret(v string) *GetAccountInfoResponseBodyAccountInfo {
	s.SignSecret = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetSignedCount(v int64) *GetAccountInfoResponseBodyAccountInfo {
	s.SignedCount = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetUnsignedCount(v int64) *GetAccountInfoResponseBodyAccountInfo {
	s.UnsignedCount = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetUnsignedEnabled(v bool) *GetAccountInfoResponseBodyAccountInfo {
	s.UnsignedEnabled = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetUserStatus(v int32) *GetAccountInfoResponseBodyAccountInfo {
	s.UserStatus = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) Validate() error {
	return dara.Validate(s)
}
