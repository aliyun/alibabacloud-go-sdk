// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountSufficient(v bool) *GetUserConfigResponseBody
	GetAccountSufficient() *bool
	SetCode(v string) *GetUserConfigResponseBody
	GetCode() *string
	SetEnableEciDisk(v bool) *GetUserConfigResponseBody
	GetEnableEciDisk() *bool
	SetFreeTier(v *GetUserConfigResponseBodyFreeTier) *GetUserConfigResponseBody
	GetFreeTier() *GetUserConfigResponseBodyFreeTier
	SetFreeTierSpecAvailable(v bool) *GetUserConfigResponseBody
	GetFreeTierSpecAvailable() *bool
	SetHttpStatusCode(v int32) *GetUserConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetUserConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUserConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUserConfigResponseBody
	GetSuccess() *bool
}

type GetUserConfigResponseBody struct {
	// example:
	//
	// false
	AccountSufficient *bool `json:"AccountSufficient,omitempty" xml:"AccountSufficient,omitempty"`
	// example:
	//
	// null
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	EnableEciDisk *bool                              `json:"EnableEciDisk,omitempty" xml:"EnableEciDisk,omitempty"`
	FreeTier      *GetUserConfigResponseBodyFreeTier `json:"FreeTier,omitempty" xml:"FreeTier,omitempty" type:"Struct"`
	// example:
	//
	// true
	FreeTierSpecAvailable *bool `json:"FreeTierSpecAvailable,omitempty" xml:"FreeTierSpecAvailable,omitempty"`
	// example:
	//
	// null
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// "Access denied"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserConfigResponseBody) GetAccountSufficient() *bool {
	return s.AccountSufficient
}

func (s *GetUserConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUserConfigResponseBody) GetEnableEciDisk() *bool {
	return s.EnableEciDisk
}

func (s *GetUserConfigResponseBody) GetFreeTier() *GetUserConfigResponseBodyFreeTier {
	return s.FreeTier
}

func (s *GetUserConfigResponseBody) GetFreeTierSpecAvailable() *bool {
	return s.FreeTierSpecAvailable
}

func (s *GetUserConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetUserConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUserConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserConfigResponseBody) SetAccountSufficient(v bool) *GetUserConfigResponseBody {
	s.AccountSufficient = &v
	return s
}

func (s *GetUserConfigResponseBody) SetCode(v string) *GetUserConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserConfigResponseBody) SetEnableEciDisk(v bool) *GetUserConfigResponseBody {
	s.EnableEciDisk = &v
	return s
}

func (s *GetUserConfigResponseBody) SetFreeTier(v *GetUserConfigResponseBodyFreeTier) *GetUserConfigResponseBody {
	s.FreeTier = v
	return s
}

func (s *GetUserConfigResponseBody) SetFreeTierSpecAvailable(v bool) *GetUserConfigResponseBody {
	s.FreeTierSpecAvailable = &v
	return s
}

func (s *GetUserConfigResponseBody) SetHttpStatusCode(v int32) *GetUserConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUserConfigResponseBody) SetMessage(v string) *GetUserConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserConfigResponseBody) SetRequestId(v string) *GetUserConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserConfigResponseBody) SetSuccess(v bool) *GetUserConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUserConfigResponseBodyFreeTier struct {
	// example:
	//
	// 2020-11-08T15:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 个
	InitBaseUnit *string `json:"InitBaseUnit,omitempty" xml:"InitBaseUnit,omitempty"`
	// example:
	//
	// 5000
	InitBaseValue *float64 `json:"InitBaseValue,omitempty" xml:"InitBaseValue,omitempty"`
	// example:
	//
	// 个
	InitShowUnit *string `json:"InitShowUnit,omitempty" xml:"InitShowUnit,omitempty"`
	// example:
	//
	// 5000
	InitShowValue *string `json:"InitShowValue,omitempty" xml:"InitShowValue,omitempty"`
	// example:
	//
	// true
	IsFreeTierUser *bool `json:"IsFreeTierUser,omitempty" xml:"IsFreeTierUser,omitempty"`
	// example:
	//
	// 个
	PeriodBaseUnit *string `json:"PeriodBaseUnit,omitempty" xml:"PeriodBaseUnit,omitempty"`
	// example:
	//
	// 3000
	PeriodBaseValue *float64 `json:"PeriodBaseValue,omitempty" xml:"PeriodBaseValue,omitempty"`
	// example:
	//
	// 个
	PeriodShowUnit *string `json:"PeriodShowUnit,omitempty" xml:"PeriodShowUnit,omitempty"`
	// example:
	//
	// 3000
	PeriodShowValue *string `json:"PeriodShowValue,omitempty" xml:"PeriodShowValue,omitempty"`
	// example:
	//
	// 2020-11-08T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// valid
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetUserConfigResponseBodyFreeTier) String() string {
	return dara.Prettify(s)
}

func (s GetUserConfigResponseBodyFreeTier) GoString() string {
	return s.String()
}

func (s *GetUserConfigResponseBodyFreeTier) GetEndTime() *string {
	return s.EndTime
}

func (s *GetUserConfigResponseBodyFreeTier) GetInitBaseUnit() *string {
	return s.InitBaseUnit
}

func (s *GetUserConfigResponseBodyFreeTier) GetInitBaseValue() *float64 {
	return s.InitBaseValue
}

func (s *GetUserConfigResponseBodyFreeTier) GetInitShowUnit() *string {
	return s.InitShowUnit
}

func (s *GetUserConfigResponseBodyFreeTier) GetInitShowValue() *string {
	return s.InitShowValue
}

func (s *GetUserConfigResponseBodyFreeTier) GetIsFreeTierUser() *bool {
	return s.IsFreeTierUser
}

func (s *GetUserConfigResponseBodyFreeTier) GetPeriodBaseUnit() *string {
	return s.PeriodBaseUnit
}

func (s *GetUserConfigResponseBodyFreeTier) GetPeriodBaseValue() *float64 {
	return s.PeriodBaseValue
}

func (s *GetUserConfigResponseBodyFreeTier) GetPeriodShowUnit() *string {
	return s.PeriodShowUnit
}

func (s *GetUserConfigResponseBodyFreeTier) GetPeriodShowValue() *string {
	return s.PeriodShowValue
}

func (s *GetUserConfigResponseBodyFreeTier) GetStartTime() *string {
	return s.StartTime
}

func (s *GetUserConfigResponseBodyFreeTier) GetStatus() *string {
	return s.Status
}

func (s *GetUserConfigResponseBodyFreeTier) SetEndTime(v string) *GetUserConfigResponseBodyFreeTier {
	s.EndTime = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) SetInitBaseUnit(v string) *GetUserConfigResponseBodyFreeTier {
	s.InitBaseUnit = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) SetInitBaseValue(v float64) *GetUserConfigResponseBodyFreeTier {
	s.InitBaseValue = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) SetInitShowUnit(v string) *GetUserConfigResponseBodyFreeTier {
	s.InitShowUnit = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) SetInitShowValue(v string) *GetUserConfigResponseBodyFreeTier {
	s.InitShowValue = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) SetIsFreeTierUser(v bool) *GetUserConfigResponseBodyFreeTier {
	s.IsFreeTierUser = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) SetPeriodBaseUnit(v string) *GetUserConfigResponseBodyFreeTier {
	s.PeriodBaseUnit = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) SetPeriodBaseValue(v float64) *GetUserConfigResponseBodyFreeTier {
	s.PeriodBaseValue = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) SetPeriodShowUnit(v string) *GetUserConfigResponseBodyFreeTier {
	s.PeriodShowUnit = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) SetPeriodShowValue(v string) *GetUserConfigResponseBodyFreeTier {
	s.PeriodShowValue = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) SetStartTime(v string) *GetUserConfigResponseBodyFreeTier {
	s.StartTime = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) SetStatus(v string) *GetUserConfigResponseBodyFreeTier {
	s.Status = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) Validate() error {
	return dara.Validate(s)
}
