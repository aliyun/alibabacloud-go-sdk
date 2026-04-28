// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvisorChecksFoPagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeAdvisorChecksFoPagesResponseBody
	GetCode() *string
	SetData(v []*DescribeAdvisorChecksFoPagesResponseBodyData) *DescribeAdvisorChecksFoPagesResponseBody
	GetData() []*DescribeAdvisorChecksFoPagesResponseBodyData
	SetMessage(v string) *DescribeAdvisorChecksFoPagesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAdvisorChecksFoPagesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAdvisorChecksFoPagesResponseBody
	GetSuccess() *bool
}

type DescribeAdvisorChecksFoPagesResponseBody struct {
	// example:
	//
	// 200
	Code *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*DescribeAdvisorChecksFoPagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 566331F9-5AB3-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAdvisorChecksFoPagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvisorChecksFoPagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksFoPagesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAdvisorChecksFoPagesResponseBody) GetData() []*DescribeAdvisorChecksFoPagesResponseBodyData {
	return s.Data
}

func (s *DescribeAdvisorChecksFoPagesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAdvisorChecksFoPagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAdvisorChecksFoPagesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAdvisorChecksFoPagesResponseBody) SetCode(v string) *DescribeAdvisorChecksFoPagesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBody) SetData(v []*DescribeAdvisorChecksFoPagesResponseBodyData) *DescribeAdvisorChecksFoPagesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBody) SetMessage(v string) *DescribeAdvisorChecksFoPagesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBody) SetRequestId(v string) *DescribeAdvisorChecksFoPagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBody) SetSuccess(v bool) *DescribeAdvisorChecksFoPagesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBody) Validate() error {
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

type DescribeAdvisorChecksFoPagesResponseBodyData struct {
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result   []*DescribeAdvisorChecksFoPagesResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeAdvisorChecksFoPagesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvisorChecksFoPagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyData) GetResult() []*DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	return s.Result
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyData) SetPageNo(v int32) *DescribeAdvisorChecksFoPagesResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyData) SetPageSize(v int32) *DescribeAdvisorChecksFoPagesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyData) SetResult(v []*DescribeAdvisorChecksFoPagesResponseBodyDataResult) *DescribeAdvisorChecksFoPagesResponseBodyData {
	s.Result = v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyData) SetTotal(v int32) *DescribeAdvisorChecksFoPagesResponseBodyData {
	s.Total = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyData) Validate() error {
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

type DescribeAdvisorChecksFoPagesResponseBodyDataResult struct {
	// example:
	//
	// 21
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// EcsCostLowUtilizationCheck
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// false
	ConfigSupport *string `json:"ConfigSupport,omitempty" xml:"ConfigSupport,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	InspectionScope *string `json:"InspectionScope,omitempty" xml:"InspectionScope,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// [{\\"type\\":\\"template\\",\\"value\\":\\"cloudmonitor.console.aliyun.com/index.htm?custom_trace=ecs_console#/hostDetail/chart/instanceId=${Resource.resourceId}&system=Linux&region=${Resource.regionId}&aliyunhost=true\\",\\"key\\":\\"Detail\\"},{\\"type\\":\\"template\\",\\"value\\":\\"/diagnosis?product=${Product}&resourceId=${Resource.resourceId}\\",\\"key\\":\\"Refresh\\"}]
	OperateColumn *string `json:"OperateColumn,omitempty" xml:"OperateColumn,omitempty"`
	// example:
	//
	// ECS
	Product   *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RiskLevel *int64  `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// Advisor
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// enabled
	Status      *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	SubCategory []*int64 `json:"SubCategory,omitempty" xml:"SubCategory,omitempty" type:"Repeated"`
	Tips        *string  `json:"Tips,omitempty" xml:"Tips,omitempty"`
	// example:
	//
	// [{\\"type\\":\\"DEFAULT\\",\\"key\\":\\"EcsCostIdleCheck_resourceId\\"},{\\"type\\":\\"DEFAULT\\",\\"key\\":\\"EcsCostIdleCheck_resourceName\\"},{\\"type\\":\\"DEFAULT\\",\\"key\\":\\"EcsCostIdleCheck_regionId\\"},{\\"type\\":\\"DEFAULT\\",\\"key\\":\\"EcsCostIdleCheck_instanceChargeType\\"},{\\"type\\":\\"DEFAULT\\",\\"key\\":\\"EcsCostIdleCheck_instanceType\\"},{\\"type\\":\\"DEFAULT\\",\\"key\\":\\"EcsCostIdleCheck_severity\\"},{\\"type\\":\\"DEFAULT\\",\\"key\\":\\"EcsCostIdleCheck_costBefore\\"},{\\"type\\":\\"DEFAULT\\",\\"key\\":\\"EcsCostIdleCheck_costAfter\\"},{\\"type\\":\\"DEFAULT\\",\\"key\\":\\"EcsCostIdleCheck_costSavings\\"},{\\"type\\":\\"DEFAULT\\",\\"key\\":\\"First_time\\"},{\\"type\\":\\"DEFAULT\\",\\"key\\":\\"Duration_time\\"}]
	ViewColumn *string `json:"ViewColumn,omitempty" xml:"ViewColumn,omitempty"`
}

func (s DescribeAdvisorChecksFoPagesResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvisorChecksFoPagesResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) GetCategory() *string {
	return s.Category
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) GetCode() *string {
	return s.Code
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) GetConfigSupport() *string {
	return s.ConfigSupport
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) GetDescription() *string {
	return s.Description
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) GetInspectionScope() *string {
	return s.InspectionScope
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) GetName() *string {
	return s.Name
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) GetOperateColumn() *string {
	return s.OperateColumn
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) GetProduct() *string {
	return s.Product
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) GetRiskLevel() *int64 {
	return s.RiskLevel
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) GetSource() *string {
	return s.Source
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) GetStatus() *string {
	return s.Status
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) GetSubCategory() []*int64 {
	return s.SubCategory
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) GetTips() *string {
	return s.Tips
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) GetViewColumn() *string {
	return s.ViewColumn
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetCategory(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.Category = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetCode(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.Code = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetConfigSupport(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.ConfigSupport = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetDescription(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetInspectionScope(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.InspectionScope = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetName(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.Name = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetOperateColumn(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.OperateColumn = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetProduct(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.Product = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetRiskLevel(v int64) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.RiskLevel = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetSource(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.Source = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetStatus(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.Status = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetSubCategory(v []*int64) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.SubCategory = v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetTips(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.Tips = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetViewColumn(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.ViewColumn = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
