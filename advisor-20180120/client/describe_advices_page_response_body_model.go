// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvicesPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeAdvicesPageResponseBodyData) *DescribeAdvicesPageResponseBody
	GetData() *DescribeAdvicesPageResponseBodyData
	SetRequestId(v string) *DescribeAdvicesPageResponseBody
	GetRequestId() *string
}

type DescribeAdvicesPageResponseBody struct {
	Data *DescribeAdvicesPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 566331F9-5AB3-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAdvicesPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvicesPageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesPageResponseBody) GetData() *DescribeAdvicesPageResponseBodyData {
	return s.Data
}

func (s *DescribeAdvicesPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAdvicesPageResponseBody) SetData(v *DescribeAdvicesPageResponseBodyData) *DescribeAdvicesPageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAdvicesPageResponseBody) SetRequestId(v string) *DescribeAdvicesPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdvicesPageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAdvicesPageResponseBodyData struct {
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int64                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result   []*DescribeAdvicesPageResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeAdvicesPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvicesPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesPageResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribeAdvicesPageResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAdvicesPageResponseBodyData) GetResult() []*DescribeAdvicesPageResponseBodyDataResult {
	return s.Result
}

func (s *DescribeAdvicesPageResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeAdvicesPageResponseBodyData) SetPageNo(v int64) *DescribeAdvicesPageResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyData) SetPageSize(v int64) *DescribeAdvicesPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyData) SetResult(v []*DescribeAdvicesPageResponseBodyDataResult) *DescribeAdvicesPageResponseBodyData {
	s.Result = v
	return s
}

func (s *DescribeAdvicesPageResponseBodyData) SetTotal(v int64) *DescribeAdvicesPageResponseBodyData {
	s.Total = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyData) Validate() error {
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

type DescribeAdvicesPageResponseBodyDataResult struct {
	// example:
	//
	// 1234567891234567
	AliyunId *int64 `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	// example:
	//
	// EcsHighCpuUtilization
	CheckId   *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// example:
	//
	// [
	//
	// 	{
	//
	// 		"key":"EcsHighCpuUtilization_xxxx",
	//
	// 		"value":xxx
	//
	// 	},
	//
	// 	{
	//
	// 		"key":"EcsHighCpuUtilization_xxxx",
	//
	// 		"value":xxx
	//
	// 	},
	//
	// 	{
	//
	// 		"key":"EcsHighCpuUtilization_xxxx",
	//
	// 		"value":xxx
	//
	// 	},
	//
	// ]
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
	// ID
	//
	// example:
	//
	// 123
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
	// {
	//
	// 	"resourceId": xxxx,
	//
	// 	"resourceName": xxxxxx,
	//
	// 	"regionId": xxxx,
	//
	// 	...
	//
	// }
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

func (s DescribeAdvicesPageResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvicesPageResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesPageResponseBodyDataResult) GetAliyunId() *int64 {
	return s.AliyunId
}

func (s *DescribeAdvicesPageResponseBodyDataResult) GetCheckId() *string {
	return s.CheckId
}

func (s *DescribeAdvicesPageResponseBodyDataResult) GetCheckName() *string {
	return s.CheckName
}

func (s *DescribeAdvicesPageResponseBodyDataResult) GetContent() *string {
	return s.Content
}

func (s *DescribeAdvicesPageResponseBodyDataResult) GetDescription() *string {
	return s.Description
}

func (s *DescribeAdvicesPageResponseBodyDataResult) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeAdvicesPageResponseBodyDataResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeAdvicesPageResponseBodyDataResult) GetId() *int64 {
	return s.Id
}

func (s *DescribeAdvicesPageResponseBodyDataResult) GetIsExpired() *bool {
	return s.IsExpired
}

func (s *DescribeAdvicesPageResponseBodyDataResult) GetProduct() *string {
	return s.Product
}

func (s *DescribeAdvicesPageResponseBodyDataResult) GetResource() *string {
	return s.Resource
}

func (s *DescribeAdvicesPageResponseBodyDataResult) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeAdvicesPageResponseBodyDataResult) GetSeverity() *int64 {
	return s.Severity
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetAliyunId(v int64) *DescribeAdvicesPageResponseBodyDataResult {
	s.AliyunId = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetCheckId(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.CheckId = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetCheckName(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.CheckName = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetContent(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.Content = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetDescription(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetGmtCreated(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.GmtCreated = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetGmtModified(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.GmtModified = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetId(v int64) *DescribeAdvicesPageResponseBodyDataResult {
	s.Id = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetIsExpired(v bool) *DescribeAdvicesPageResponseBodyDataResult {
	s.IsExpired = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetProduct(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.Product = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetResource(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.Resource = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetResourceId(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.ResourceId = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetSeverity(v int64) *DescribeAdvicesPageResponseBodyDataResult {
	s.Severity = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
