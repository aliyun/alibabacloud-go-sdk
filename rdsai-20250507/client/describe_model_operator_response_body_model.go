// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelOperatorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeModelOperatorResponseBodyData) *DescribeModelOperatorResponseBody
	GetData() *DescribeModelOperatorResponseBodyData
	SetMessage(v string) *DescribeModelOperatorResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeModelOperatorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeModelOperatorResponseBody
	GetSuccess() *bool
}

type DescribeModelOperatorResponseBody struct {
	Data *DescribeModelOperatorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeModelOperatorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOperatorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModelOperatorResponseBody) GetData() *DescribeModelOperatorResponseBodyData {
	return s.Data
}

func (s *DescribeModelOperatorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeModelOperatorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeModelOperatorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeModelOperatorResponseBody) SetData(v *DescribeModelOperatorResponseBodyData) *DescribeModelOperatorResponseBody {
	s.Data = v
	return s
}

func (s *DescribeModelOperatorResponseBody) SetMessage(v string) *DescribeModelOperatorResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeModelOperatorResponseBody) SetRequestId(v string) *DescribeModelOperatorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModelOperatorResponseBody) SetSuccess(v bool) *DescribeModelOperatorResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeModelOperatorResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeModelOperatorResponseBodyData struct {
	// example:
	//
	// sk-rds-xxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// False
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// http://xxx.yy/v1
	BaseUrl *string `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
	// example:
	//
	// PREPAY / POSTPAY
	ChargeType *string                                            `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DailyUsage []*DescribeModelOperatorResponseBodyDataDailyUsage `json:"DailyUsage,omitempty" xml:"DailyUsage,omitempty" type:"Repeated"`
	// example:
	//
	// 1775145600000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// xlarge
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// example:
	//
	// rds_copilot***_public_cn-*********6
	InstanceId   *string                                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	KeyUsageList []*DescribeModelOperatorResponseBodyDataKeyUsageList `json:"KeyUsageList,omitempty" xml:"KeyUsageList,omitempty" type:"Repeated"`
	// example:
	//
	// 1772439028000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// active/creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 200000000
	TotalQuota *int64 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// example:
	//
	// 1000000
	UsedQuota *int64 `json:"UsedQuota,omitempty" xml:"UsedQuota,omitempty"`
}

func (s DescribeModelOperatorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOperatorResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeModelOperatorResponseBodyData) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeModelOperatorResponseBodyData) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *DescribeModelOperatorResponseBodyData) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *DescribeModelOperatorResponseBodyData) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeModelOperatorResponseBodyData) GetDailyUsage() []*DescribeModelOperatorResponseBodyDataDailyUsage {
	return s.DailyUsage
}

func (s *DescribeModelOperatorResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeModelOperatorResponseBodyData) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeModelOperatorResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeModelOperatorResponseBodyData) GetKeyUsageList() []*DescribeModelOperatorResponseBodyDataKeyUsageList {
	return s.KeyUsageList
}

func (s *DescribeModelOperatorResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeModelOperatorResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeModelOperatorResponseBodyData) GetTotalQuota() *int64 {
	return s.TotalQuota
}

func (s *DescribeModelOperatorResponseBodyData) GetUsedQuota() *int64 {
	return s.UsedQuota
}

func (s *DescribeModelOperatorResponseBodyData) SetApiKey(v string) *DescribeModelOperatorResponseBodyData {
	s.ApiKey = &v
	return s
}

func (s *DescribeModelOperatorResponseBodyData) SetAutoRenew(v bool) *DescribeModelOperatorResponseBodyData {
	s.AutoRenew = &v
	return s
}

func (s *DescribeModelOperatorResponseBodyData) SetBaseUrl(v string) *DescribeModelOperatorResponseBodyData {
	s.BaseUrl = &v
	return s
}

func (s *DescribeModelOperatorResponseBodyData) SetChargeType(v string) *DescribeModelOperatorResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *DescribeModelOperatorResponseBodyData) SetDailyUsage(v []*DescribeModelOperatorResponseBodyDataDailyUsage) *DescribeModelOperatorResponseBodyData {
	s.DailyUsage = v
	return s
}

func (s *DescribeModelOperatorResponseBodyData) SetEndTime(v int64) *DescribeModelOperatorResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeModelOperatorResponseBodyData) SetInstanceClass(v string) *DescribeModelOperatorResponseBodyData {
	s.InstanceClass = &v
	return s
}

func (s *DescribeModelOperatorResponseBodyData) SetInstanceId(v string) *DescribeModelOperatorResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeModelOperatorResponseBodyData) SetKeyUsageList(v []*DescribeModelOperatorResponseBodyDataKeyUsageList) *DescribeModelOperatorResponseBodyData {
	s.KeyUsageList = v
	return s
}

func (s *DescribeModelOperatorResponseBodyData) SetStartTime(v int64) *DescribeModelOperatorResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeModelOperatorResponseBodyData) SetStatus(v string) *DescribeModelOperatorResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeModelOperatorResponseBodyData) SetTotalQuota(v int64) *DescribeModelOperatorResponseBodyData {
	s.TotalQuota = &v
	return s
}

func (s *DescribeModelOperatorResponseBodyData) SetUsedQuota(v int64) *DescribeModelOperatorResponseBodyData {
	s.UsedQuota = &v
	return s
}

func (s *DescribeModelOperatorResponseBodyData) Validate() error {
	if s.DailyUsage != nil {
		for _, item := range s.DailyUsage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.KeyUsageList != nil {
		for _, item := range s.KeyUsageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeModelOperatorResponseBodyDataDailyUsage struct {
	// example:
	//
	// 2026-03-31
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// 100000
	Usage *int64 `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s DescribeModelOperatorResponseBodyDataDailyUsage) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOperatorResponseBodyDataDailyUsage) GoString() string {
	return s.String()
}

func (s *DescribeModelOperatorResponseBodyDataDailyUsage) GetDate() *string {
	return s.Date
}

func (s *DescribeModelOperatorResponseBodyDataDailyUsage) GetUsage() *int64 {
	return s.Usage
}

func (s *DescribeModelOperatorResponseBodyDataDailyUsage) SetDate(v string) *DescribeModelOperatorResponseBodyDataDailyUsage {
	s.Date = &v
	return s
}

func (s *DescribeModelOperatorResponseBodyDataDailyUsage) SetUsage(v int64) *DescribeModelOperatorResponseBodyDataDailyUsage {
	s.Usage = &v
	return s
}

func (s *DescribeModelOperatorResponseBodyDataDailyUsage) Validate() error {
	return dara.Validate(s)
}

type DescribeModelOperatorResponseBodyDataKeyUsageList struct {
	// API Key
	//
	// example:
	//
	// sk-rds-*****
	ApiKey     *string                                                        `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	DailyUsage []*DescribeModelOperatorResponseBodyDataKeyUsageListDailyUsage `json:"DailyUsage,omitempty" xml:"DailyUsage,omitempty" type:"Repeated"`
	// example:
	//
	// fase
	Deleted *bool `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// example:
	//
	// api-*****
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// example:
	//
	// fixed
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	// example:
	//
	// 100000
	KeyUsed *string `json:"KeyUsed,omitempty" xml:"KeyUsed,omitempty"`
	// example:
	//
	// 2000000
	UsedQuota *string `json:"UsedQuota,omitempty" xml:"UsedQuota,omitempty"`
}

func (s DescribeModelOperatorResponseBodyDataKeyUsageList) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOperatorResponseBodyDataKeyUsageList) GoString() string {
	return s.String()
}

func (s *DescribeModelOperatorResponseBodyDataKeyUsageList) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeModelOperatorResponseBodyDataKeyUsageList) GetDailyUsage() []*DescribeModelOperatorResponseBodyDataKeyUsageListDailyUsage {
	return s.DailyUsage
}

func (s *DescribeModelOperatorResponseBodyDataKeyUsageList) GetDeleted() *bool {
	return s.Deleted
}

func (s *DescribeModelOperatorResponseBodyDataKeyUsageList) GetKeyName() *string {
	return s.KeyName
}

func (s *DescribeModelOperatorResponseBodyDataKeyUsageList) GetKeyType() *string {
	return s.KeyType
}

func (s *DescribeModelOperatorResponseBodyDataKeyUsageList) GetKeyUsed() *string {
	return s.KeyUsed
}

func (s *DescribeModelOperatorResponseBodyDataKeyUsageList) GetUsedQuota() *string {
	return s.UsedQuota
}

func (s *DescribeModelOperatorResponseBodyDataKeyUsageList) SetApiKey(v string) *DescribeModelOperatorResponseBodyDataKeyUsageList {
	s.ApiKey = &v
	return s
}

func (s *DescribeModelOperatorResponseBodyDataKeyUsageList) SetDailyUsage(v []*DescribeModelOperatorResponseBodyDataKeyUsageListDailyUsage) *DescribeModelOperatorResponseBodyDataKeyUsageList {
	s.DailyUsage = v
	return s
}

func (s *DescribeModelOperatorResponseBodyDataKeyUsageList) SetDeleted(v bool) *DescribeModelOperatorResponseBodyDataKeyUsageList {
	s.Deleted = &v
	return s
}

func (s *DescribeModelOperatorResponseBodyDataKeyUsageList) SetKeyName(v string) *DescribeModelOperatorResponseBodyDataKeyUsageList {
	s.KeyName = &v
	return s
}

func (s *DescribeModelOperatorResponseBodyDataKeyUsageList) SetKeyType(v string) *DescribeModelOperatorResponseBodyDataKeyUsageList {
	s.KeyType = &v
	return s
}

func (s *DescribeModelOperatorResponseBodyDataKeyUsageList) SetKeyUsed(v string) *DescribeModelOperatorResponseBodyDataKeyUsageList {
	s.KeyUsed = &v
	return s
}

func (s *DescribeModelOperatorResponseBodyDataKeyUsageList) SetUsedQuota(v string) *DescribeModelOperatorResponseBodyDataKeyUsageList {
	s.UsedQuota = &v
	return s
}

func (s *DescribeModelOperatorResponseBodyDataKeyUsageList) Validate() error {
	if s.DailyUsage != nil {
		for _, item := range s.DailyUsage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeModelOperatorResponseBodyDataKeyUsageListDailyUsage struct {
	// example:
	//
	// 2026-03-31
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// 2000
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s DescribeModelOperatorResponseBodyDataKeyUsageListDailyUsage) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOperatorResponseBodyDataKeyUsageListDailyUsage) GoString() string {
	return s.String()
}

func (s *DescribeModelOperatorResponseBodyDataKeyUsageListDailyUsage) GetDate() *string {
	return s.Date
}

func (s *DescribeModelOperatorResponseBodyDataKeyUsageListDailyUsage) GetUsage() *string {
	return s.Usage
}

func (s *DescribeModelOperatorResponseBodyDataKeyUsageListDailyUsage) SetDate(v string) *DescribeModelOperatorResponseBodyDataKeyUsageListDailyUsage {
	s.Date = &v
	return s
}

func (s *DescribeModelOperatorResponseBodyDataKeyUsageListDailyUsage) SetUsage(v string) *DescribeModelOperatorResponseBodyDataKeyUsageListDailyUsage {
	s.Usage = &v
	return s
}

func (s *DescribeModelOperatorResponseBodyDataKeyUsageListDailyUsage) Validate() error {
	return dara.Validate(s)
}
