// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrialInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTrialInfoResponseBody
	GetCode() *string
	SetMessage(v string) *GetTrialInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTrialInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTrialInfoResponseBody
	GetSuccess() *bool
	SetTrialInfo(v *GetTrialInfoResponseBodyTrialInfo) *GetTrialInfoResponseBody
	GetTrialInfo() *GetTrialInfoResponseBodyTrialInfo
}

type GetTrialInfoResponseBody struct {
	// The return code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. The value "successful" is returned on success. An error message is returned on failure.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The free trial properties. This parameter is returned only when a free trial record exists for the data source.
	TrialInfo *GetTrialInfoResponseBodyTrialInfo `json:"TrialInfo,omitempty" xml:"TrialInfo,omitempty" type:"Struct"`
}

func (s GetTrialInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTrialInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrialInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTrialInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTrialInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTrialInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTrialInfoResponseBody) GetTrialInfo() *GetTrialInfoResponseBodyTrialInfo {
	return s.TrialInfo
}

func (s *GetTrialInfoResponseBody) SetCode(v string) *GetTrialInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetTrialInfoResponseBody) SetMessage(v string) *GetTrialInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetTrialInfoResponseBody) SetRequestId(v string) *GetTrialInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrialInfoResponseBody) SetSuccess(v bool) *GetTrialInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetTrialInfoResponseBody) SetTrialInfo(v *GetTrialInfoResponseBodyTrialInfo) *GetTrialInfoResponseBody {
	s.TrialInfo = v
	return s
}

func (s *GetTrialInfoResponseBody) Validate() error {
	if s.TrialInfo != nil {
		if err := s.TrialInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTrialInfoResponseBodyTrialInfo struct {
	// Indicates whether the service is converted to pay-as-you-go after the free trial expires.
	KeepAfterTrialExpiration *bool `json:"KeepAfterTrialExpiration,omitempty" xml:"KeepAfterTrialExpiration,omitempty"`
	// The remaining number of free trial backup plans that can be created. The value 1 is returned if the free trial has not expired and no free trial backup plan has been created. Otherwise, the value 0 is returned.
	//
	// example:
	//
	// 1
	TrialBackupPlanQuota *int64 `json:"TrialBackupPlanQuota,omitempty" xml:"TrialBackupPlanQuota,omitempty"`
	// The expiration time of the free trial. UNIX timestamp, in seconds.
	//
	// example:
	//
	// 1584597600
	TrialExpireTime *int64 `json:"TrialExpireTime,omitempty" xml:"TrialExpireTime,omitempty"`
	// The start time of the free trial. UNIX timestamp, in seconds.
	//
	// example:
	//
	// 1579413159
	TrialStartTime *int64 `json:"TrialStartTime,omitempty" xml:"TrialStartTime,omitempty"`
	// The remaining number of free trial backup vaults that can be created. The value 1 is returned if the free trial has not expired and no free trial backup vault has been created. Otherwise, the value 0 is returned.
	//
	// example:
	//
	// 1
	TrialVaultQuota *int64 `json:"TrialVaultQuota,omitempty" xml:"TrialVaultQuota,omitempty"`
	// The release time of the free trial backup vault. UNIX timestamp, in seconds.
	//
	// example:
	//
	// 1594965600
	TrialVaultReleaseTime *int64 `json:"TrialVaultReleaseTime,omitempty" xml:"TrialVaultReleaseTime,omitempty"`
}

func (s GetTrialInfoResponseBodyTrialInfo) String() string {
	return dara.Prettify(s)
}

func (s GetTrialInfoResponseBodyTrialInfo) GoString() string {
	return s.String()
}

func (s *GetTrialInfoResponseBodyTrialInfo) GetKeepAfterTrialExpiration() *bool {
	return s.KeepAfterTrialExpiration
}

func (s *GetTrialInfoResponseBodyTrialInfo) GetTrialBackupPlanQuota() *int64 {
	return s.TrialBackupPlanQuota
}

func (s *GetTrialInfoResponseBodyTrialInfo) GetTrialExpireTime() *int64 {
	return s.TrialExpireTime
}

func (s *GetTrialInfoResponseBodyTrialInfo) GetTrialStartTime() *int64 {
	return s.TrialStartTime
}

func (s *GetTrialInfoResponseBodyTrialInfo) GetTrialVaultQuota() *int64 {
	return s.TrialVaultQuota
}

func (s *GetTrialInfoResponseBodyTrialInfo) GetTrialVaultReleaseTime() *int64 {
	return s.TrialVaultReleaseTime
}

func (s *GetTrialInfoResponseBodyTrialInfo) SetKeepAfterTrialExpiration(v bool) *GetTrialInfoResponseBodyTrialInfo {
	s.KeepAfterTrialExpiration = &v
	return s
}

func (s *GetTrialInfoResponseBodyTrialInfo) SetTrialBackupPlanQuota(v int64) *GetTrialInfoResponseBodyTrialInfo {
	s.TrialBackupPlanQuota = &v
	return s
}

func (s *GetTrialInfoResponseBodyTrialInfo) SetTrialExpireTime(v int64) *GetTrialInfoResponseBodyTrialInfo {
	s.TrialExpireTime = &v
	return s
}

func (s *GetTrialInfoResponseBodyTrialInfo) SetTrialStartTime(v int64) *GetTrialInfoResponseBodyTrialInfo {
	s.TrialStartTime = &v
	return s
}

func (s *GetTrialInfoResponseBodyTrialInfo) SetTrialVaultQuota(v int64) *GetTrialInfoResponseBodyTrialInfo {
	s.TrialVaultQuota = &v
	return s
}

func (s *GetTrialInfoResponseBodyTrialInfo) SetTrialVaultReleaseTime(v int64) *GetTrialInfoResponseBodyTrialInfo {
	s.TrialVaultReleaseTime = &v
	return s
}

func (s *GetTrialInfoResponseBodyTrialInfo) Validate() error {
	return dara.Validate(s)
}
