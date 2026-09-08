// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTotalSensitiveInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataCountDOList(v []*ListTotalSensitiveInfoResponseBodyDataCountDOList) *ListTotalSensitiveInfoResponseBody
	GetDataCountDOList() []*ListTotalSensitiveInfoResponseBodyDataCountDOList
	SetDbCount(v int64) *ListTotalSensitiveInfoResponseBody
	GetDbCount() *int64
	SetInstanceCount(v int64) *ListTotalSensitiveInfoResponseBody
	GetInstanceCount() *int64
	SetRequestId(v string) *ListTotalSensitiveInfoResponseBody
	GetRequestId() *string
	SetRuleInfoList(v []*ListTotalSensitiveInfoResponseBodyRuleInfoList) *ListTotalSensitiveInfoResponseBody
	GetRuleInfoList() []*ListTotalSensitiveInfoResponseBodyRuleInfoList
	SetS0Count(v int64) *ListTotalSensitiveInfoResponseBody
	GetS0Count() *int64
	SetS10Count(v int64) *ListTotalSensitiveInfoResponseBody
	GetS10Count() *int64
	SetS1Count(v int64) *ListTotalSensitiveInfoResponseBody
	GetS1Count() *int64
	SetS2Count(v int64) *ListTotalSensitiveInfoResponseBody
	GetS2Count() *int64
	SetS3Count(v int64) *ListTotalSensitiveInfoResponseBody
	GetS3Count() *int64
	SetS4Count(v int64) *ListTotalSensitiveInfoResponseBody
	GetS4Count() *int64
	SetS5Count(v int64) *ListTotalSensitiveInfoResponseBody
	GetS5Count() *int64
	SetS6Count(v int64) *ListTotalSensitiveInfoResponseBody
	GetS6Count() *int64
	SetS7Count(v int64) *ListTotalSensitiveInfoResponseBody
	GetS7Count() *int64
	SetS8Count(v int64) *ListTotalSensitiveInfoResponseBody
	GetS8Count() *int64
	SetS9Count(v int64) *ListTotalSensitiveInfoResponseBody
	GetS9Count() *int64
	SetSensitiveCount(v int64) *ListTotalSensitiveInfoResponseBody
	GetSensitiveCount() *int64
	SetSensitiveDbCount(v int64) *ListTotalSensitiveInfoResponseBody
	GetSensitiveDbCount() *int64
	SetSensitiveInstanceCount(v int64) *ListTotalSensitiveInfoResponseBody
	GetSensitiveInstanceCount() *int64
	SetSensitiveUnStructSize(v int64) *ListTotalSensitiveInfoResponseBody
	GetSensitiveUnStructSize() *int64
	SetSubSensitiveCount(v int64) *ListTotalSensitiveInfoResponseBody
	GetSubSensitiveCount() *int64
	SetSubTotalCount(v int64) *ListTotalSensitiveInfoResponseBody
	GetSubTotalCount() *int64
	SetTotalCount(v int64) *ListTotalSensitiveInfoResponseBody
	GetTotalCount() *int64
	SetUnStructSize(v int64) *ListTotalSensitiveInfoResponseBody
	GetUnStructSize() *int64
}

type ListTotalSensitiveInfoResponseBody struct {
	DataCountDOList []*ListTotalSensitiveInfoResponseBodyDataCountDOList `json:"DataCountDOList,omitempty" xml:"DataCountDOList,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	DbCount *int64 `json:"DbCount,omitempty" xml:"DbCount,omitempty"`
	// example:
	//
	// 5
	InstanceCount *int64 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// example:
	//
	// 7C6D8E9F-1234-5678-ABCD-0123456789AB
	RequestId    *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleInfoList []*ListTotalSensitiveInfoResponseBodyRuleInfoList `json:"RuleInfoList,omitempty" xml:"RuleInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	S0Count *int64 `json:"S0Count,omitempty" xml:"S0Count,omitempty"`
	// example:
	//
	// 0
	S10Count *int64 `json:"S10Count,omitempty" xml:"S10Count,omitempty"`
	// example:
	//
	// 0
	S1Count *int64 `json:"S1Count,omitempty" xml:"S1Count,omitempty"`
	// example:
	//
	// 0
	S2Count *int64 `json:"S2Count,omitempty" xml:"S2Count,omitempty"`
	// example:
	//
	// 0
	S3Count *int64 `json:"S3Count,omitempty" xml:"S3Count,omitempty"`
	// example:
	//
	// 0
	S4Count *int64 `json:"S4Count,omitempty" xml:"S4Count,omitempty"`
	// example:
	//
	// 0
	S5Count *int64 `json:"S5Count,omitempty" xml:"S5Count,omitempty"`
	// example:
	//
	// 0
	S6Count *int64 `json:"S6Count,omitempty" xml:"S6Count,omitempty"`
	// example:
	//
	// 0
	S7Count *int64 `json:"S7Count,omitempty" xml:"S7Count,omitempty"`
	// example:
	//
	// 0
	S8Count *int64 `json:"S8Count,omitempty" xml:"S8Count,omitempty"`
	// example:
	//
	// 0
	S9Count *int64 `json:"S9Count,omitempty" xml:"S9Count,omitempty"`
	// example:
	//
	// 20
	SensitiveCount *int64 `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	// example:
	//
	// 5
	SensitiveDbCount *int64 `json:"SensitiveDbCount,omitempty" xml:"SensitiveDbCount,omitempty"`
	// example:
	//
	// 2
	SensitiveInstanceCount *int64 `json:"SensitiveInstanceCount,omitempty" xml:"SensitiveInstanceCount,omitempty"`
	// example:
	//
	// 512
	SensitiveUnStructSize *int64 `json:"SensitiveUnStructSize,omitempty" xml:"SensitiveUnStructSize,omitempty"`
	// example:
	//
	// 20
	SubSensitiveCount *int64 `json:"SubSensitiveCount,omitempty" xml:"SubSensitiveCount,omitempty"`
	// example:
	//
	// 100
	SubTotalCount *int64 `json:"SubTotalCount,omitempty" xml:"SubTotalCount,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 1024
	UnStructSize *int64 `json:"UnStructSize,omitempty" xml:"UnStructSize,omitempty"`
}

func (s ListTotalSensitiveInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTotalSensitiveInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListTotalSensitiveInfoResponseBody) GetDataCountDOList() []*ListTotalSensitiveInfoResponseBodyDataCountDOList {
	return s.DataCountDOList
}

func (s *ListTotalSensitiveInfoResponseBody) GetDbCount() *int64 {
	return s.DbCount
}

func (s *ListTotalSensitiveInfoResponseBody) GetInstanceCount() *int64 {
	return s.InstanceCount
}

func (s *ListTotalSensitiveInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTotalSensitiveInfoResponseBody) GetRuleInfoList() []*ListTotalSensitiveInfoResponseBodyRuleInfoList {
	return s.RuleInfoList
}

func (s *ListTotalSensitiveInfoResponseBody) GetS0Count() *int64 {
	return s.S0Count
}

func (s *ListTotalSensitiveInfoResponseBody) GetS10Count() *int64 {
	return s.S10Count
}

func (s *ListTotalSensitiveInfoResponseBody) GetS1Count() *int64 {
	return s.S1Count
}

func (s *ListTotalSensitiveInfoResponseBody) GetS2Count() *int64 {
	return s.S2Count
}

func (s *ListTotalSensitiveInfoResponseBody) GetS3Count() *int64 {
	return s.S3Count
}

func (s *ListTotalSensitiveInfoResponseBody) GetS4Count() *int64 {
	return s.S4Count
}

func (s *ListTotalSensitiveInfoResponseBody) GetS5Count() *int64 {
	return s.S5Count
}

func (s *ListTotalSensitiveInfoResponseBody) GetS6Count() *int64 {
	return s.S6Count
}

func (s *ListTotalSensitiveInfoResponseBody) GetS7Count() *int64 {
	return s.S7Count
}

func (s *ListTotalSensitiveInfoResponseBody) GetS8Count() *int64 {
	return s.S8Count
}

func (s *ListTotalSensitiveInfoResponseBody) GetS9Count() *int64 {
	return s.S9Count
}

func (s *ListTotalSensitiveInfoResponseBody) GetSensitiveCount() *int64 {
	return s.SensitiveCount
}

func (s *ListTotalSensitiveInfoResponseBody) GetSensitiveDbCount() *int64 {
	return s.SensitiveDbCount
}

func (s *ListTotalSensitiveInfoResponseBody) GetSensitiveInstanceCount() *int64 {
	return s.SensitiveInstanceCount
}

func (s *ListTotalSensitiveInfoResponseBody) GetSensitiveUnStructSize() *int64 {
	return s.SensitiveUnStructSize
}

func (s *ListTotalSensitiveInfoResponseBody) GetSubSensitiveCount() *int64 {
	return s.SubSensitiveCount
}

func (s *ListTotalSensitiveInfoResponseBody) GetSubTotalCount() *int64 {
	return s.SubTotalCount
}

func (s *ListTotalSensitiveInfoResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTotalSensitiveInfoResponseBody) GetUnStructSize() *int64 {
	return s.UnStructSize
}

func (s *ListTotalSensitiveInfoResponseBody) SetDataCountDOList(v []*ListTotalSensitiveInfoResponseBodyDataCountDOList) *ListTotalSensitiveInfoResponseBody {
	s.DataCountDOList = v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetDbCount(v int64) *ListTotalSensitiveInfoResponseBody {
	s.DbCount = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetInstanceCount(v int64) *ListTotalSensitiveInfoResponseBody {
	s.InstanceCount = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetRequestId(v string) *ListTotalSensitiveInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetRuleInfoList(v []*ListTotalSensitiveInfoResponseBodyRuleInfoList) *ListTotalSensitiveInfoResponseBody {
	s.RuleInfoList = v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetS0Count(v int64) *ListTotalSensitiveInfoResponseBody {
	s.S0Count = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetS10Count(v int64) *ListTotalSensitiveInfoResponseBody {
	s.S10Count = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetS1Count(v int64) *ListTotalSensitiveInfoResponseBody {
	s.S1Count = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetS2Count(v int64) *ListTotalSensitiveInfoResponseBody {
	s.S2Count = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetS3Count(v int64) *ListTotalSensitiveInfoResponseBody {
	s.S3Count = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetS4Count(v int64) *ListTotalSensitiveInfoResponseBody {
	s.S4Count = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetS5Count(v int64) *ListTotalSensitiveInfoResponseBody {
	s.S5Count = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetS6Count(v int64) *ListTotalSensitiveInfoResponseBody {
	s.S6Count = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetS7Count(v int64) *ListTotalSensitiveInfoResponseBody {
	s.S7Count = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetS8Count(v int64) *ListTotalSensitiveInfoResponseBody {
	s.S8Count = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetS9Count(v int64) *ListTotalSensitiveInfoResponseBody {
	s.S9Count = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetSensitiveCount(v int64) *ListTotalSensitiveInfoResponseBody {
	s.SensitiveCount = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetSensitiveDbCount(v int64) *ListTotalSensitiveInfoResponseBody {
	s.SensitiveDbCount = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetSensitiveInstanceCount(v int64) *ListTotalSensitiveInfoResponseBody {
	s.SensitiveInstanceCount = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetSensitiveUnStructSize(v int64) *ListTotalSensitiveInfoResponseBody {
	s.SensitiveUnStructSize = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetSubSensitiveCount(v int64) *ListTotalSensitiveInfoResponseBody {
	s.SubSensitiveCount = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetSubTotalCount(v int64) *ListTotalSensitiveInfoResponseBody {
	s.SubTotalCount = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetTotalCount(v int64) *ListTotalSensitiveInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) SetUnStructSize(v int64) *ListTotalSensitiveInfoResponseBody {
	s.UnStructSize = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBody) Validate() error {
	if s.DataCountDOList != nil {
		for _, item := range s.DataCountDOList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RuleInfoList != nil {
		for _, item := range s.RuleInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTotalSensitiveInfoResponseBodyDataCountDOList struct {
	DataCountDOList []*ListTotalSensitiveInfoResponseBodyDataCountDOListDataCountDOList `json:"DataCountDOList,omitempty" xml:"DataCountDOList,omitempty" type:"Repeated"`
	// example:
	//
	// 1788537600000000000
	Date *int64 `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// cn-zhangjiakou
	RegionId     *string                                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RuleInfoList []*ListTotalSensitiveInfoResponseBodyDataCountDOListRuleInfoList `json:"RuleInfoList,omitempty" xml:"RuleInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	S0Count *int64 `json:"S0Count,omitempty" xml:"S0Count,omitempty"`
	// example:
	//
	// 0
	S10Count *int64 `json:"S10Count,omitempty" xml:"S10Count,omitempty"`
	// example:
	//
	// 0
	S1Count *int64 `json:"S1Count,omitempty" xml:"S1Count,omitempty"`
	// example:
	//
	// 0
	S2Count *int64 `json:"S2Count,omitempty" xml:"S2Count,omitempty"`
	// example:
	//
	// 0
	S3Count *int64 `json:"S3Count,omitempty" xml:"S3Count,omitempty"`
	// example:
	//
	// 0
	S4Count *int64 `json:"S4Count,omitempty" xml:"S4Count,omitempty"`
	// example:
	//
	// 0
	S5Count *int64 `json:"S5Count,omitempty" xml:"S5Count,omitempty"`
	// example:
	//
	// 0
	S6Count *int64 `json:"S6Count,omitempty" xml:"S6Count,omitempty"`
	// example:
	//
	// 0
	S7Count *int64 `json:"S7Count,omitempty" xml:"S7Count,omitempty"`
	// example:
	//
	// 0
	S8Count *int64 `json:"S8Count,omitempty" xml:"S8Count,omitempty"`
	// example:
	//
	// 0
	S9Count *int64 `json:"S9Count,omitempty" xml:"S9Count,omitempty"`
	// example:
	//
	// 20
	SensitiveCount *int64 `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	// example:
	//
	// 1
	StructFlag *int32 `json:"StructFlag,omitempty" xml:"StructFlag,omitempty"`
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// 通用分类分级模板
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTotalSensitiveInfoResponseBodyDataCountDOList) String() string {
	return dara.Prettify(s)
}

func (s ListTotalSensitiveInfoResponseBodyDataCountDOList) GoString() string {
	return s.String()
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) GetDataCountDOList() []*ListTotalSensitiveInfoResponseBodyDataCountDOListDataCountDOList {
	return s.DataCountDOList
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) GetDate() *int64 {
	return s.Date
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) GetRuleInfoList() []*ListTotalSensitiveInfoResponseBodyDataCountDOListRuleInfoList {
	return s.RuleInfoList
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) GetS0Count() *int64 {
	return s.S0Count
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) GetS10Count() *int64 {
	return s.S10Count
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) GetS1Count() *int64 {
	return s.S1Count
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) GetS2Count() *int64 {
	return s.S2Count
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) GetS3Count() *int64 {
	return s.S3Count
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) GetS4Count() *int64 {
	return s.S4Count
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) GetS5Count() *int64 {
	return s.S5Count
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) GetS6Count() *int64 {
	return s.S6Count
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) GetS7Count() *int64 {
	return s.S7Count
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) GetS8Count() *int64 {
	return s.S8Count
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) GetS9Count() *int64 {
	return s.S9Count
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) GetSensitiveCount() *int64 {
	return s.SensitiveCount
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) GetStructFlag() *int32 {
	return s.StructFlag
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) SetDataCountDOList(v []*ListTotalSensitiveInfoResponseBodyDataCountDOListDataCountDOList) *ListTotalSensitiveInfoResponseBodyDataCountDOList {
	s.DataCountDOList = v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) SetDate(v int64) *ListTotalSensitiveInfoResponseBodyDataCountDOList {
	s.Date = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) SetRegionId(v string) *ListTotalSensitiveInfoResponseBodyDataCountDOList {
	s.RegionId = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) SetRuleInfoList(v []*ListTotalSensitiveInfoResponseBodyDataCountDOListRuleInfoList) *ListTotalSensitiveInfoResponseBodyDataCountDOList {
	s.RuleInfoList = v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) SetS0Count(v int64) *ListTotalSensitiveInfoResponseBodyDataCountDOList {
	s.S0Count = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) SetS10Count(v int64) *ListTotalSensitiveInfoResponseBodyDataCountDOList {
	s.S10Count = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) SetS1Count(v int64) *ListTotalSensitiveInfoResponseBodyDataCountDOList {
	s.S1Count = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) SetS2Count(v int64) *ListTotalSensitiveInfoResponseBodyDataCountDOList {
	s.S2Count = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) SetS3Count(v int64) *ListTotalSensitiveInfoResponseBodyDataCountDOList {
	s.S3Count = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) SetS4Count(v int64) *ListTotalSensitiveInfoResponseBodyDataCountDOList {
	s.S4Count = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) SetS5Count(v int64) *ListTotalSensitiveInfoResponseBodyDataCountDOList {
	s.S5Count = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) SetS6Count(v int64) *ListTotalSensitiveInfoResponseBodyDataCountDOList {
	s.S6Count = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) SetS7Count(v int64) *ListTotalSensitiveInfoResponseBodyDataCountDOList {
	s.S7Count = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) SetS8Count(v int64) *ListTotalSensitiveInfoResponseBodyDataCountDOList {
	s.S8Count = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) SetS9Count(v int64) *ListTotalSensitiveInfoResponseBodyDataCountDOList {
	s.S9Count = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) SetSensitiveCount(v int64) *ListTotalSensitiveInfoResponseBodyDataCountDOList {
	s.SensitiveCount = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) SetStructFlag(v int32) *ListTotalSensitiveInfoResponseBodyDataCountDOList {
	s.StructFlag = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) SetTemplateId(v int64) *ListTotalSensitiveInfoResponseBodyDataCountDOList {
	s.TemplateId = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) SetTemplateName(v string) *ListTotalSensitiveInfoResponseBodyDataCountDOList {
	s.TemplateName = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) SetTotalCount(v int64) *ListTotalSensitiveInfoResponseBodyDataCountDOList {
	s.TotalCount = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOList) Validate() error {
	if s.DataCountDOList != nil {
		for _, item := range s.DataCountDOList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RuleInfoList != nil {
		for _, item := range s.RuleInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTotalSensitiveInfoResponseBodyDataCountDOListDataCountDOList struct {
	// example:
	//
	// 1788537600000000000
	Date *int64 `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// 20
	SensitiveCount *int64 `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTotalSensitiveInfoResponseBodyDataCountDOListDataCountDOList) String() string {
	return dara.Prettify(s)
}

func (s ListTotalSensitiveInfoResponseBodyDataCountDOListDataCountDOList) GoString() string {
	return s.String()
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOListDataCountDOList) GetDate() *int64 {
	return s.Date
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOListDataCountDOList) GetSensitiveCount() *int64 {
	return s.SensitiveCount
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOListDataCountDOList) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOListDataCountDOList) SetDate(v int64) *ListTotalSensitiveInfoResponseBodyDataCountDOListDataCountDOList {
	s.Date = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOListDataCountDOList) SetSensitiveCount(v int64) *ListTotalSensitiveInfoResponseBodyDataCountDOListDataCountDOList {
	s.SensitiveCount = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOListDataCountDOList) SetTotalCount(v int64) *ListTotalSensitiveInfoResponseBodyDataCountDOListDataCountDOList {
	s.TotalCount = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOListDataCountDOList) Validate() error {
	return dara.Validate(s)
}

type ListTotalSensitiveInfoResponseBodyDataCountDOListRuleInfoList struct {
	// example:
	//
	// 10
	RuleCount *int64 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
	// example:
	//
	// 1001
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// 手机号
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListTotalSensitiveInfoResponseBodyDataCountDOListRuleInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListTotalSensitiveInfoResponseBodyDataCountDOListRuleInfoList) GoString() string {
	return s.String()
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOListRuleInfoList) GetRuleCount() *int64 {
	return s.RuleCount
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOListRuleInfoList) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOListRuleInfoList) GetRuleName() *string {
	return s.RuleName
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOListRuleInfoList) SetRuleCount(v int64) *ListTotalSensitiveInfoResponseBodyDataCountDOListRuleInfoList {
	s.RuleCount = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOListRuleInfoList) SetRuleId(v int64) *ListTotalSensitiveInfoResponseBodyDataCountDOListRuleInfoList {
	s.RuleId = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOListRuleInfoList) SetRuleName(v string) *ListTotalSensitiveInfoResponseBodyDataCountDOListRuleInfoList {
	s.RuleName = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyDataCountDOListRuleInfoList) Validate() error {
	return dara.Validate(s)
}

type ListTotalSensitiveInfoResponseBodyRuleInfoList struct {
	// example:
	//
	// 10
	RuleCount *int64 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
	// example:
	//
	// 1001
	RuleId *int32 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// 手机号
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListTotalSensitiveInfoResponseBodyRuleInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListTotalSensitiveInfoResponseBodyRuleInfoList) GoString() string {
	return s.String()
}

func (s *ListTotalSensitiveInfoResponseBodyRuleInfoList) GetRuleCount() *int64 {
	return s.RuleCount
}

func (s *ListTotalSensitiveInfoResponseBodyRuleInfoList) GetRuleId() *int32 {
	return s.RuleId
}

func (s *ListTotalSensitiveInfoResponseBodyRuleInfoList) GetRuleName() *string {
	return s.RuleName
}

func (s *ListTotalSensitiveInfoResponseBodyRuleInfoList) SetRuleCount(v int64) *ListTotalSensitiveInfoResponseBodyRuleInfoList {
	s.RuleCount = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyRuleInfoList) SetRuleId(v int32) *ListTotalSensitiveInfoResponseBodyRuleInfoList {
	s.RuleId = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyRuleInfoList) SetRuleName(v string) *ListTotalSensitiveInfoResponseBodyRuleInfoList {
	s.RuleName = &v
	return s
}

func (s *ListTotalSensitiveInfoResponseBodyRuleInfoList) Validate() error {
	return dara.Validate(s)
}
