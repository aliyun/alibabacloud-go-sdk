// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJVSInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeJVSInstanceResponseBodyData) *DescribeJVSInstanceResponseBody
	GetData() []*DescribeJVSInstanceResponseBodyData
	SetMaxResults(v int32) *DescribeJVSInstanceResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeJVSInstanceResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeJVSInstanceResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeJVSInstanceResponseBody
	GetTotalCount() *int32
}

type DescribeJVSInstanceResponseBody struct {
	Data []*DescribeJVSInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// 当前页实际返回条数
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 下一页游标，末页不返回
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 符合条件的总记录数
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeJVSInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeJVSInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJVSInstanceResponseBody) GetData() []*DescribeJVSInstanceResponseBodyData {
	return s.Data
}

func (s *DescribeJVSInstanceResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeJVSInstanceResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeJVSInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeJVSInstanceResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeJVSInstanceResponseBody) SetData(v []*DescribeJVSInstanceResponseBodyData) *DescribeJVSInstanceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeJVSInstanceResponseBody) SetMaxResults(v int32) *DescribeJVSInstanceResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeJVSInstanceResponseBody) SetNextToken(v string) *DescribeJVSInstanceResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeJVSInstanceResponseBody) SetRequestId(v string) *DescribeJVSInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJVSInstanceResponseBody) SetTotalCount(v int32) *DescribeJVSInstanceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeJVSInstanceResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeJVSInstanceResponseBodyData struct {
	// example:
	//
	// 2026-04-10T01:31:32Z
	CreateTime   *string                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreditConfig []*DescribeJVSInstanceResponseBodyDataCreditConfig `json:"CreditConfig,omitempty" xml:"CreditConfig,omitempty" type:"Repeated"`
	// example:
	//
	// 2026-04-10T01:31:32Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// acp-uto81vfd8t8z****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// null
	JvsPackageId *string `json:"JvsPackageId,omitempty" xml:"JvsPackageId,omitempty"`
	// example:
	//
	// 2026-04-10T01:31:32Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// RUNNNING
	Status     *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	UsedCredit []*DescribeJVSInstanceResponseBodyDataUsedCredit `json:"UsedCredit,omitempty" xml:"UsedCredit,omitempty" type:"Repeated"`
}

func (s DescribeJVSInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeJVSInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeJVSInstanceResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeJVSInstanceResponseBodyData) GetCreditConfig() []*DescribeJVSInstanceResponseBodyDataCreditConfig {
	return s.CreditConfig
}

func (s *DescribeJVSInstanceResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeJVSInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeJVSInstanceResponseBodyData) GetJvsPackageId() *string {
	return s.JvsPackageId
}

func (s *DescribeJVSInstanceResponseBodyData) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeJVSInstanceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeJVSInstanceResponseBodyData) GetUsedCredit() []*DescribeJVSInstanceResponseBodyDataUsedCredit {
	return s.UsedCredit
}

func (s *DescribeJVSInstanceResponseBodyData) SetCreateTime(v string) *DescribeJVSInstanceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyData) SetCreditConfig(v []*DescribeJVSInstanceResponseBodyDataCreditConfig) *DescribeJVSInstanceResponseBodyData {
	s.CreditConfig = v
	return s
}

func (s *DescribeJVSInstanceResponseBodyData) SetExpireTime(v string) *DescribeJVSInstanceResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyData) SetInstanceId(v string) *DescribeJVSInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyData) SetJvsPackageId(v string) *DescribeJVSInstanceResponseBodyData {
	s.JvsPackageId = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyData) SetModifyTime(v string) *DescribeJVSInstanceResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyData) SetStatus(v string) *DescribeJVSInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyData) SetUsedCredit(v []*DescribeJVSInstanceResponseBodyDataUsedCredit) *DescribeJVSInstanceResponseBodyData {
	s.UsedCredit = v
	return s
}

func (s *DescribeJVSInstanceResponseBodyData) Validate() error {
	if s.CreditConfig != nil {
		for _, item := range s.CreditConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UsedCredit != nil {
		for _, item := range s.UsedCredit {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeJVSInstanceResponseBodyDataCreditConfig struct {
	// example:
	//
	// -1
	CreditLimit *int64 `json:"CreditLimit,omitempty" xml:"CreditLimit,omitempty"`
	// example:
	//
	// day
	LimitPeriod *string `json:"LimitPeriod,omitempty" xml:"LimitPeriod,omitempty"`
}

func (s DescribeJVSInstanceResponseBodyDataCreditConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeJVSInstanceResponseBodyDataCreditConfig) GoString() string {
	return s.String()
}

func (s *DescribeJVSInstanceResponseBodyDataCreditConfig) GetCreditLimit() *int64 {
	return s.CreditLimit
}

func (s *DescribeJVSInstanceResponseBodyDataCreditConfig) GetLimitPeriod() *string {
	return s.LimitPeriod
}

func (s *DescribeJVSInstanceResponseBodyDataCreditConfig) SetCreditLimit(v int64) *DescribeJVSInstanceResponseBodyDataCreditConfig {
	s.CreditLimit = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyDataCreditConfig) SetLimitPeriod(v string) *DescribeJVSInstanceResponseBodyDataCreditConfig {
	s.LimitPeriod = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyDataCreditConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeJVSInstanceResponseBodyDataUsedCredit struct {
	// example:
	//
	// 5
	Credit *int64 `json:"Credit,omitempty" xml:"Credit,omitempty"`
	// example:
	//
	// day
	LimitPeriod *string `json:"LimitPeriod,omitempty" xml:"LimitPeriod,omitempty"`
}

func (s DescribeJVSInstanceResponseBodyDataUsedCredit) String() string {
	return dara.Prettify(s)
}

func (s DescribeJVSInstanceResponseBodyDataUsedCredit) GoString() string {
	return s.String()
}

func (s *DescribeJVSInstanceResponseBodyDataUsedCredit) GetCredit() *int64 {
	return s.Credit
}

func (s *DescribeJVSInstanceResponseBodyDataUsedCredit) GetLimitPeriod() *string {
	return s.LimitPeriod
}

func (s *DescribeJVSInstanceResponseBodyDataUsedCredit) SetCredit(v int64) *DescribeJVSInstanceResponseBodyDataUsedCredit {
	s.Credit = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyDataUsedCredit) SetLimitPeriod(v string) *DescribeJVSInstanceResponseBodyDataUsedCredit {
	s.LimitPeriod = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyDataUsedCredit) Validate() error {
	return dara.Validate(s)
}
