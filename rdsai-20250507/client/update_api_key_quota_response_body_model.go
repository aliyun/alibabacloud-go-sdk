// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiKeyQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateApiKeyQuotaResponseBodyData) *UpdateApiKeyQuotaResponseBody
	GetData() *UpdateApiKeyQuotaResponseBodyData
	SetMessage(v string) *UpdateApiKeyQuotaResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateApiKeyQuotaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateApiKeyQuotaResponseBody
	GetSuccess() *bool
}

type UpdateApiKeyQuotaResponseBody struct {
	// The returned data.
	Data *UpdateApiKeyQuotaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateApiKeyQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiKeyQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApiKeyQuotaResponseBody) GetData() *UpdateApiKeyQuotaResponseBodyData {
	return s.Data
}

func (s *UpdateApiKeyQuotaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateApiKeyQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApiKeyQuotaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateApiKeyQuotaResponseBody) SetData(v *UpdateApiKeyQuotaResponseBodyData) *UpdateApiKeyQuotaResponseBody {
	s.Data = v
	return s
}

func (s *UpdateApiKeyQuotaResponseBody) SetMessage(v string) *UpdateApiKeyQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateApiKeyQuotaResponseBody) SetRequestId(v string) *UpdateApiKeyQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApiKeyQuotaResponseBody) SetSuccess(v bool) *UpdateApiKeyQuotaResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateApiKeyQuotaResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApiKeyQuotaResponseBodyData struct {
	// The list of custom API keys.
	CustomKeyList []*UpdateApiKeyQuotaResponseBodyDataCustomKeyList `json:"CustomKeyList,omitempty" xml:"CustomKeyList,omitempty" type:"Repeated"`
}

func (s UpdateApiKeyQuotaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiKeyQuotaResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateApiKeyQuotaResponseBodyData) GetCustomKeyList() []*UpdateApiKeyQuotaResponseBodyDataCustomKeyList {
	return s.CustomKeyList
}

func (s *UpdateApiKeyQuotaResponseBodyData) SetCustomKeyList(v []*UpdateApiKeyQuotaResponseBodyDataCustomKeyList) *UpdateApiKeyQuotaResponseBodyData {
	s.CustomKeyList = v
	return s
}

func (s *UpdateApiKeyQuotaResponseBodyData) Validate() error {
	if s.CustomKeyList != nil {
		for _, item := range s.CustomKeyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateApiKeyQuotaResponseBodyDataCustomKeyList struct {
	// The API key.
	//
	// example:
	//
	// sk-rds-xxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// 1000000000
	DailyTokenQuota *int64 `json:"DailyTokenQuota,omitempty" xml:"DailyTokenQuota,omitempty"`
	// The limit rate.
	//
	// example:
	//
	// 0.2
	LimitRate *float32 `json:"LimitRate,omitempty" xml:"LimitRate,omitempty"`
	// The quota limiting method. Valid values:
	//
	// - `ratio`: Sets the limit based on a ratio.
	//
	// - `fixed`: Sets the limit to a fixed value.
	//
	// - `auto`: Allocates the limit automatically.
	//
	// example:
	//
	// fixed
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
	// The token quota for the API key.
	//
	// example:
	//
	// 10000
	TokenQuota *int64 `json:"TokenQuota,omitempty" xml:"TokenQuota,omitempty"`
}

func (s UpdateApiKeyQuotaResponseBodyDataCustomKeyList) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiKeyQuotaResponseBodyDataCustomKeyList) GoString() string {
	return s.String()
}

func (s *UpdateApiKeyQuotaResponseBodyDataCustomKeyList) GetApiKey() *string {
	return s.ApiKey
}

func (s *UpdateApiKeyQuotaResponseBodyDataCustomKeyList) GetDailyTokenQuota() *int64 {
	return s.DailyTokenQuota
}

func (s *UpdateApiKeyQuotaResponseBodyDataCustomKeyList) GetLimitRate() *float32 {
	return s.LimitRate
}

func (s *UpdateApiKeyQuotaResponseBodyDataCustomKeyList) GetLimitType() *string {
	return s.LimitType
}

func (s *UpdateApiKeyQuotaResponseBodyDataCustomKeyList) GetTokenQuota() *int64 {
	return s.TokenQuota
}

func (s *UpdateApiKeyQuotaResponseBodyDataCustomKeyList) SetApiKey(v string) *UpdateApiKeyQuotaResponseBodyDataCustomKeyList {
	s.ApiKey = &v
	return s
}

func (s *UpdateApiKeyQuotaResponseBodyDataCustomKeyList) SetDailyTokenQuota(v int64) *UpdateApiKeyQuotaResponseBodyDataCustomKeyList {
	s.DailyTokenQuota = &v
	return s
}

func (s *UpdateApiKeyQuotaResponseBodyDataCustomKeyList) SetLimitRate(v float32) *UpdateApiKeyQuotaResponseBodyDataCustomKeyList {
	s.LimitRate = &v
	return s
}

func (s *UpdateApiKeyQuotaResponseBodyDataCustomKeyList) SetLimitType(v string) *UpdateApiKeyQuotaResponseBodyDataCustomKeyList {
	s.LimitType = &v
	return s
}

func (s *UpdateApiKeyQuotaResponseBodyDataCustomKeyList) SetTokenQuota(v int64) *UpdateApiKeyQuotaResponseBodyDataCustomKeyList {
	s.TokenQuota = &v
	return s
}

func (s *UpdateApiKeyQuotaResponseBodyDataCustomKeyList) Validate() error {
	return dara.Validate(s)
}
