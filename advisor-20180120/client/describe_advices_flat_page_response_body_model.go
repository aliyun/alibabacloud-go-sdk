// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvicesFlatPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeAdvicesFlatPageResponseBodyData) *DescribeAdvicesFlatPageResponseBody
	GetData() *DescribeAdvicesFlatPageResponseBodyData
	SetRequestId(v string) *DescribeAdvicesFlatPageResponseBody
	GetRequestId() *string
}

type DescribeAdvicesFlatPageResponseBody struct {
	Data *DescribeAdvicesFlatPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 566331F9-5AB3-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAdvicesFlatPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvicesFlatPageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesFlatPageResponseBody) GetData() *DescribeAdvicesFlatPageResponseBodyData {
	return s.Data
}

func (s *DescribeAdvicesFlatPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAdvicesFlatPageResponseBody) SetData(v *DescribeAdvicesFlatPageResponseBodyData) *DescribeAdvicesFlatPageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBody) SetRequestId(v string) *DescribeAdvicesFlatPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAdvicesFlatPageResponseBodyData struct {
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int64                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result   []*DescribeAdvicesFlatPageResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeAdvicesFlatPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvicesFlatPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesFlatPageResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribeAdvicesFlatPageResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAdvicesFlatPageResponseBodyData) GetResult() []*DescribeAdvicesFlatPageResponseBodyDataResult {
	return s.Result
}

func (s *DescribeAdvicesFlatPageResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeAdvicesFlatPageResponseBodyData) SetPageNo(v int64) *DescribeAdvicesFlatPageResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyData) SetPageSize(v int64) *DescribeAdvicesFlatPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyData) SetResult(v []*DescribeAdvicesFlatPageResponseBodyDataResult) *DescribeAdvicesFlatPageResponseBodyData {
	s.Result = v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyData) SetTotal(v int64) *DescribeAdvicesFlatPageResponseBodyData {
	s.Total = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAdvicesFlatPageResponseBodyDataResult struct {
	// example:
	//
	// 192895059480****
	AliyunId *int64 `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	// example:
	//
	// EcsHighCpuUtilization
	CheckId   *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// example:
	//
	// [{"key":"EcsHighCpuUtilization_xxxx", "value":"xxx"}, {"key":"EcsHighCpuUtilization_xxxx", "value":"xxx"}, {"key":"EcsHighCpuUtilization_xxxx", "value":"xxx"}, ]
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2023-07-01 00:00:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 2023-07-01 00:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 40200899
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// false
	IsExpired *bool `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// {"resourceId": "i-2zecnwitr2s7aca6****","resourceName": "ecs-20230701","regionId": "cn-hangzhou",...}
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// example:
	//
	// i-bp67acfmxazb4p****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// 1
	Severity *int64 `json:"Severity,omitempty" xml:"Severity,omitempty"`
}

func (s DescribeAdvicesFlatPageResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvicesFlatPageResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) GetAliyunId() *int64 {
	return s.AliyunId
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) GetCheckId() *string {
	return s.CheckId
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) GetCheckName() *string {
	return s.CheckName
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) GetContent() *string {
	return s.Content
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) GetDescription() *string {
	return s.Description
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) GetId() *int64 {
	return s.Id
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) GetIsExpired() *bool {
	return s.IsExpired
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) GetProduct() *string {
	return s.Product
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) GetResource() *string {
	return s.Resource
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) GetSeverity() *int64 {
	return s.Severity
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetAliyunId(v int64) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.AliyunId = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetCheckId(v string) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.CheckId = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetCheckName(v string) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.CheckName = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetContent(v string) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.Content = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetDescription(v string) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetGmtCreated(v string) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.GmtCreated = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetGmtModified(v string) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.GmtModified = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetId(v int64) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.Id = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetIsExpired(v bool) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.IsExpired = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetProduct(v string) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.Product = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetResource(v string) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.Resource = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetResourceId(v string) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.ResourceId = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetSeverity(v int64) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.Severity = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
