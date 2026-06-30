// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListApiKeysResponseBodyData) *ListApiKeysResponseBody
	GetData() *ListApiKeysResponseBodyData
	SetMessage(v string) *ListApiKeysResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListApiKeysResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListApiKeysResponseBody
	GetSuccess() *bool
}

type ListApiKeysResponseBody struct {
	// The response data.
	Data *ListApiKeysResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates if the request succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListApiKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApiKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListApiKeysResponseBody) GetData() *ListApiKeysResponseBodyData {
	return s.Data
}

func (s *ListApiKeysResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListApiKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApiKeysResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListApiKeysResponseBody) SetData(v *ListApiKeysResponseBodyData) *ListApiKeysResponseBody {
	s.Data = v
	return s
}

func (s *ListApiKeysResponseBody) SetMessage(v string) *ListApiKeysResponseBody {
	s.Message = &v
	return s
}

func (s *ListApiKeysResponseBody) SetRequestId(v string) *ListApiKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApiKeysResponseBody) SetSuccess(v bool) *ListApiKeysResponseBody {
	s.Success = &v
	return s
}

func (s *ListApiKeysResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApiKeysResponseBodyData struct {
	// The base URL for model calls.
	//
	// example:
	//
	// http://xxx.yy/v1
	BaseUrl *string `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
	// The custom API key list.
	CustomKeyList []*ListApiKeysResponseBodyDataCustomKeyList `json:"CustomKeyList,omitempty" xml:"CustomKeyList,omitempty" type:"Repeated"`
	// example:
	//
	// 100000000
	DailyTokenQuota *int64 `json:"DailyTokenQuota,omitempty" xml:"DailyTokenQuota,omitempty"`
	// Specifies if the system-generated key is rate-limited.
	//
	// example:
	//
	// false
	IsRateLimited *bool `json:"IsRateLimited,omitempty" xml:"IsRateLimited,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The system-generated key.
	//
	// example:
	//
	// sk-rds-xxx
	SystemApiKey *string `json:"SystemApiKey,omitempty" xml:"SystemApiKey,omitempty"`
	// The alarm threshold percentage for the SystemApiKey. For example, a value of 80 indicates that an alarm is triggered when usage reaches 80% of the quota. The alarm clears when usage falls below this threshold.
	//
	// example:
	//
	// 80
	ThresholdPercent *int32 `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 138
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListApiKeysResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListApiKeysResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListApiKeysResponseBodyData) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *ListApiKeysResponseBodyData) GetCustomKeyList() []*ListApiKeysResponseBodyDataCustomKeyList {
	return s.CustomKeyList
}

func (s *ListApiKeysResponseBodyData) GetDailyTokenQuota() *int64 {
	return s.DailyTokenQuota
}

func (s *ListApiKeysResponseBodyData) GetIsRateLimited() *bool {
	return s.IsRateLimited
}

func (s *ListApiKeysResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ListApiKeysResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApiKeysResponseBodyData) GetSystemApiKey() *string {
	return s.SystemApiKey
}

func (s *ListApiKeysResponseBodyData) GetThresholdPercent() *int32 {
	return s.ThresholdPercent
}

func (s *ListApiKeysResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListApiKeysResponseBodyData) SetBaseUrl(v string) *ListApiKeysResponseBodyData {
	s.BaseUrl = &v
	return s
}

func (s *ListApiKeysResponseBodyData) SetCustomKeyList(v []*ListApiKeysResponseBodyDataCustomKeyList) *ListApiKeysResponseBodyData {
	s.CustomKeyList = v
	return s
}

func (s *ListApiKeysResponseBodyData) SetDailyTokenQuota(v int64) *ListApiKeysResponseBodyData {
	s.DailyTokenQuota = &v
	return s
}

func (s *ListApiKeysResponseBodyData) SetIsRateLimited(v bool) *ListApiKeysResponseBodyData {
	s.IsRateLimited = &v
	return s
}

func (s *ListApiKeysResponseBodyData) SetPage(v int32) *ListApiKeysResponseBodyData {
	s.Page = &v
	return s
}

func (s *ListApiKeysResponseBodyData) SetPageSize(v int32) *ListApiKeysResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListApiKeysResponseBodyData) SetSystemApiKey(v string) *ListApiKeysResponseBodyData {
	s.SystemApiKey = &v
	return s
}

func (s *ListApiKeysResponseBodyData) SetThresholdPercent(v int32) *ListApiKeysResponseBodyData {
	s.ThresholdPercent = &v
	return s
}

func (s *ListApiKeysResponseBodyData) SetTotal(v int32) *ListApiKeysResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListApiKeysResponseBodyData) Validate() error {
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

type ListApiKeysResponseBodyDataCustomKeyList struct {
	// The API key.
	//
	// example:
	//
	// sk-rds-*****
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// 100000000
	DailyTokenQuota *int64 `json:"DailyTokenQuota,omitempty" xml:"DailyTokenQuota,omitempty"`
	// Specifies if the API key is rate-limited.
	//
	// example:
	//
	// false
	IsRateLimited *bool `json:"IsRateLimited,omitempty" xml:"IsRateLimited,omitempty"`
	// The key name.
	//
	// example:
	//
	// api-*****
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// The limit, specified as a ratio in decimal format. This parameter is used when LimitType is ratio.
	//
	// example:
	//
	// 0.2
	LimitRate *float32 `json:"LimitRate,omitempty" xml:"LimitRate,omitempty"`
	// The limit type. Valid values:
	//
	// - **fixed**: A fixed value.
	//
	// - **ratio**: A percentage of the total quota.
	//
	// - **auto**: The quota is allocated automatically.
	//
	// example:
	//
	// fixed
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
	// The alarm threshold percentage. For example, a value of 80 indicates that an alarm is triggered when usage reaches 80% of the quota. The alarm clears when usage falls below this threshold.
	//
	// example:
	//
	// 80
	ThresholdPercent *int32 `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
	// The token quota.
	//
	// example:
	//
	// 10000
	TokenQuota *int64 `json:"TokenQuota,omitempty" xml:"TokenQuota,omitempty"`
}

func (s ListApiKeysResponseBodyDataCustomKeyList) String() string {
	return dara.Prettify(s)
}

func (s ListApiKeysResponseBodyDataCustomKeyList) GoString() string {
	return s.String()
}

func (s *ListApiKeysResponseBodyDataCustomKeyList) GetApiKey() *string {
	return s.ApiKey
}

func (s *ListApiKeysResponseBodyDataCustomKeyList) GetDailyTokenQuota() *int64 {
	return s.DailyTokenQuota
}

func (s *ListApiKeysResponseBodyDataCustomKeyList) GetIsRateLimited() *bool {
	return s.IsRateLimited
}

func (s *ListApiKeysResponseBodyDataCustomKeyList) GetKeyName() *string {
	return s.KeyName
}

func (s *ListApiKeysResponseBodyDataCustomKeyList) GetLimitRate() *float32 {
	return s.LimitRate
}

func (s *ListApiKeysResponseBodyDataCustomKeyList) GetLimitType() *string {
	return s.LimitType
}

func (s *ListApiKeysResponseBodyDataCustomKeyList) GetThresholdPercent() *int32 {
	return s.ThresholdPercent
}

func (s *ListApiKeysResponseBodyDataCustomKeyList) GetTokenQuota() *int64 {
	return s.TokenQuota
}

func (s *ListApiKeysResponseBodyDataCustomKeyList) SetApiKey(v string) *ListApiKeysResponseBodyDataCustomKeyList {
	s.ApiKey = &v
	return s
}

func (s *ListApiKeysResponseBodyDataCustomKeyList) SetDailyTokenQuota(v int64) *ListApiKeysResponseBodyDataCustomKeyList {
	s.DailyTokenQuota = &v
	return s
}

func (s *ListApiKeysResponseBodyDataCustomKeyList) SetIsRateLimited(v bool) *ListApiKeysResponseBodyDataCustomKeyList {
	s.IsRateLimited = &v
	return s
}

func (s *ListApiKeysResponseBodyDataCustomKeyList) SetKeyName(v string) *ListApiKeysResponseBodyDataCustomKeyList {
	s.KeyName = &v
	return s
}

func (s *ListApiKeysResponseBodyDataCustomKeyList) SetLimitRate(v float32) *ListApiKeysResponseBodyDataCustomKeyList {
	s.LimitRate = &v
	return s
}

func (s *ListApiKeysResponseBodyDataCustomKeyList) SetLimitType(v string) *ListApiKeysResponseBodyDataCustomKeyList {
	s.LimitType = &v
	return s
}

func (s *ListApiKeysResponseBodyDataCustomKeyList) SetThresholdPercent(v int32) *ListApiKeysResponseBodyDataCustomKeyList {
	s.ThresholdPercent = &v
	return s
}

func (s *ListApiKeysResponseBodyDataCustomKeyList) SetTokenQuota(v int64) *ListApiKeysResponseBodyDataCustomKeyList {
	s.TokenQuota = &v
	return s
}

func (s *ListApiKeysResponseBodyDataCustomKeyList) Validate() error {
	return dara.Validate(s)
}
