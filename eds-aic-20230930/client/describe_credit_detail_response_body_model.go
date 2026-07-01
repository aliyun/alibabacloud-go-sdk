// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreditDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeCreditDetailResponseBodyData) *DescribeCreditDetailResponseBody
	GetData() *DescribeCreditDetailResponseBodyData
	SetRequestId(v string) *DescribeCreditDetailResponseBody
	GetRequestId() *string
}

type DescribeCreditDetailResponseBody struct {
	// The response object.
	Data *DescribeCreditDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCreditDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCreditDetailResponseBody) GetData() *DescribeCreditDetailResponseBodyData {
	return s.Data
}

func (s *DescribeCreditDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCreditDetailResponseBody) SetData(v *DescribeCreditDetailResponseBodyData) *DescribeCreditDetailResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCreditDetailResponseBody) SetRequestId(v string) *DescribeCreditDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCreditDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCreditDetailResponseBodyData struct {
	// The credit change details.
	Details []*DescribeCreditDetailResponseBodyDataDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of detail records.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total credit change.
	//
	// example:
	//
	// 100
	TotalCreditChange *string `json:"TotalCreditChange,omitempty" xml:"TotalCreditChange,omitempty"`
}

func (s DescribeCreditDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCreditDetailResponseBodyData) GetDetails() []*DescribeCreditDetailResponseBodyDataDetails {
	return s.Details
}

func (s *DescribeCreditDetailResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeCreditDetailResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCreditDetailResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCreditDetailResponseBodyData) GetTotalCreditChange() *string {
	return s.TotalCreditChange
}

func (s *DescribeCreditDetailResponseBodyData) SetDetails(v []*DescribeCreditDetailResponseBodyDataDetails) *DescribeCreditDetailResponseBodyData {
	s.Details = v
	return s
}

func (s *DescribeCreditDetailResponseBodyData) SetPageNum(v int32) *DescribeCreditDetailResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *DescribeCreditDetailResponseBodyData) SetPageSize(v int32) *DescribeCreditDetailResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeCreditDetailResponseBodyData) SetTotalCount(v int32) *DescribeCreditDetailResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeCreditDetailResponseBodyData) SetTotalCreditChange(v string) *DescribeCreditDetailResponseBodyData {
	s.TotalCreditChange = &v
	return s
}

func (s *DescribeCreditDetailResponseBodyData) Validate() error {
	if s.Details != nil {
		for _, item := range s.Details {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCreditDetailResponseBodyDataDetails struct {
	// The time when the change occurred.
	//
	// example:
	//
	// 2026-06-30T08:14:02Z
	ChangeTime *string `json:"ChangeTime,omitempty" xml:"ChangeTime,omitempty"`
	// The credit change amount.
	//
	// example:
	//
	// 0.7637
	CreditChange *string `json:"CreditChange,omitempty" xml:"CreditChange,omitempty"`
	// The task description.
	//
	// example:
	//
	// 打开小红书
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// acp-12oe0l75vl7o5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The credit or resource plan ID.
	//
	// example:
	//
	// cmag-0c1g77wjljl9h****
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// The task ID, which is globally unique.
	//
	// example:
	//
	// t-1fr0k51pozyr5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeCreditDetailResponseBodyDataDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditDetailResponseBodyDataDetails) GoString() string {
	return s.String()
}

func (s *DescribeCreditDetailResponseBodyDataDetails) GetChangeTime() *string {
	return s.ChangeTime
}

func (s *DescribeCreditDetailResponseBodyDataDetails) GetCreditChange() *string {
	return s.CreditChange
}

func (s *DescribeCreditDetailResponseBodyDataDetails) GetDescription() *string {
	return s.Description
}

func (s *DescribeCreditDetailResponseBodyDataDetails) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCreditDetailResponseBodyDataDetails) GetPackageId() *string {
	return s.PackageId
}

func (s *DescribeCreditDetailResponseBodyDataDetails) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeCreditDetailResponseBodyDataDetails) SetChangeTime(v string) *DescribeCreditDetailResponseBodyDataDetails {
	s.ChangeTime = &v
	return s
}

func (s *DescribeCreditDetailResponseBodyDataDetails) SetCreditChange(v string) *DescribeCreditDetailResponseBodyDataDetails {
	s.CreditChange = &v
	return s
}

func (s *DescribeCreditDetailResponseBodyDataDetails) SetDescription(v string) *DescribeCreditDetailResponseBodyDataDetails {
	s.Description = &v
	return s
}

func (s *DescribeCreditDetailResponseBodyDataDetails) SetInstanceId(v string) *DescribeCreditDetailResponseBodyDataDetails {
	s.InstanceId = &v
	return s
}

func (s *DescribeCreditDetailResponseBodyDataDetails) SetPackageId(v string) *DescribeCreditDetailResponseBodyDataDetails {
	s.PackageId = &v
	return s
}

func (s *DescribeCreditDetailResponseBodyDataDetails) SetTaskId(v string) *DescribeCreditDetailResponseBodyDataDetails {
	s.TaskId = &v
	return s
}

func (s *DescribeCreditDetailResponseBodyDataDetails) Validate() error {
	return dara.Validate(s)
}
