// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryResourcePackageInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryResourcePackageInstancesResponseBody
	GetCode() *string
	SetData(v *QueryResourcePackageInstancesResponseBodyData) *QueryResourcePackageInstancesResponseBody
	GetData() *QueryResourcePackageInstancesResponseBodyData
	SetMessage(v string) *QueryResourcePackageInstancesResponseBody
	GetMessage() *string
	SetPage(v int32) *QueryResourcePackageInstancesResponseBody
	GetPage() *int32
	SetPageSize(v int32) *QueryResourcePackageInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryResourcePackageInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryResourcePackageInstancesResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *QueryResourcePackageInstancesResponseBody
	GetTotal() *int32
}

type QueryResourcePackageInstancesResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryResourcePackageInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 12
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryResourcePackageInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryResourcePackageInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryResourcePackageInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryResourcePackageInstancesResponseBody) GetData() *QueryResourcePackageInstancesResponseBodyData {
	return s.Data
}

func (s *QueryResourcePackageInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryResourcePackageInstancesResponseBody) GetPage() *int32 {
	return s.Page
}

func (s *QueryResourcePackageInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryResourcePackageInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryResourcePackageInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryResourcePackageInstancesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *QueryResourcePackageInstancesResponseBody) SetCode(v string) *QueryResourcePackageInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBody) SetData(v *QueryResourcePackageInstancesResponseBodyData) *QueryResourcePackageInstancesResponseBody {
	s.Data = v
	return s
}

func (s *QueryResourcePackageInstancesResponseBody) SetMessage(v string) *QueryResourcePackageInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBody) SetPage(v int32) *QueryResourcePackageInstancesResponseBody {
	s.Page = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBody) SetPageSize(v int32) *QueryResourcePackageInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBody) SetRequestId(v string) *QueryResourcePackageInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBody) SetSuccess(v bool) *QueryResourcePackageInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBody) SetTotal(v int32) *QueryResourcePackageInstancesResponseBody {
	s.Total = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryResourcePackageInstancesResponseBodyData struct {
	// The ID of the host.
	//
	// example:
	//
	// cn
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The details of the instances.
	Instances *QueryResourcePackageInstancesResponseBodyDataInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNum *string `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 12
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryResourcePackageInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryResourcePackageInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryResourcePackageInstancesResponseBodyData) GetHostId() *string {
	return s.HostId
}

func (s *QueryResourcePackageInstancesResponseBodyData) GetInstances() *QueryResourcePackageInstancesResponseBodyDataInstances {
	return s.Instances
}

func (s *QueryResourcePackageInstancesResponseBodyData) GetPageNum() *string {
	return s.PageNum
}

func (s *QueryResourcePackageInstancesResponseBodyData) GetPageSize() *string {
	return s.PageSize
}

func (s *QueryResourcePackageInstancesResponseBodyData) GetTotalCount() *string {
	return s.TotalCount
}

func (s *QueryResourcePackageInstancesResponseBodyData) SetHostId(v string) *QueryResourcePackageInstancesResponseBodyData {
	s.HostId = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyData) SetInstances(v *QueryResourcePackageInstancesResponseBodyDataInstances) *QueryResourcePackageInstancesResponseBodyData {
	s.Instances = v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyData) SetPageNum(v string) *QueryResourcePackageInstancesResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyData) SetPageSize(v string) *QueryResourcePackageInstancesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyData) SetTotalCount(v string) *QueryResourcePackageInstancesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyData) Validate() error {
	if s.Instances != nil {
		if err := s.Instances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryResourcePackageInstancesResponseBodyDataInstances struct {
	Instance []*QueryResourcePackageInstancesResponseBodyDataInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s QueryResourcePackageInstancesResponseBodyDataInstances) String() string {
	return dara.Prettify(s)
}

func (s QueryResourcePackageInstancesResponseBodyDataInstances) GoString() string {
	return s.String()
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstances) GetInstance() []*QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	return s.Instance
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstances) SetInstance(v []*QueryResourcePackageInstancesResponseBodyDataInstancesInstance) *QueryResourcePackageInstancesResponseBodyDataInstances {
	s.Instance = v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstances) Validate() error {
	if s.Instance != nil {
		for _, item := range s.Instance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryResourcePackageInstancesResponseBodyDataInstancesInstance struct {
	// The services to which the resource plan is applicable.
	ApplicableProducts *QueryResourcePackageInstancesResponseBodyDataInstancesInstanceApplicableProducts `json:"ApplicableProducts,omitempty" xml:"ApplicableProducts,omitempty" type:"Struct"`
	// The commodity code.
	//
	// example:
	//
	// rds
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The deduction type. Example: Absolute.
	//
	// example:
	//
	// Absolute
	DeductType *string `json:"DeductType,omitempty" xml:"DeductType,omitempty"`
	// The time when the resource plan took effect.
	//
	// example:
	//
	// 2018-09-12T09:51:56Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The time when the resource plan expired.
	//
	// example:
	//
	// 2019-03-12T16:00:00Z
	ExpiryTime *string `json:"ExpiryTime,omitempty" xml:"ExpiryTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// OSSBAG-cn-v0h1s4hma01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the resource plan.
	//
	// example:
	//
	// FPT_ossbag_absolute_Storage_sh
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The remaining quota.
	//
	// example:
	//
	// 40.000
	RemainingAmount *string `json:"RemainingAmount,omitempty" xml:"RemainingAmount,omitempty"`
	// The unit of the remaining quota.
	//
	// example:
	//
	// GB
	RemainingAmountUnit *string `json:"RemainingAmountUnit,omitempty" xml:"RemainingAmountUnit,omitempty"`
	// The remarks on the resource plan. The remarks must be made in Chinese.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The status of the resource plan. Valid values:
	//
	// 	- Available
	//
	// 	- Expired
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total quota of the resource plan.
	//
	// example:
	//
	// 40.000
	TotalAmount *string `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
	// The unit of the total quota.
	//
	// example:
	//
	// GB
	TotalAmountUnit *string `json:"TotalAmountUnit,omitempty" xml:"TotalAmountUnit,omitempty"`
}

func (s QueryResourcePackageInstancesResponseBodyDataInstancesInstance) String() string {
	return dara.Prettify(s)
}

func (s QueryResourcePackageInstancesResponseBodyDataInstancesInstance) GoString() string {
	return s.String()
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) GetApplicableProducts() *QueryResourcePackageInstancesResponseBodyDataInstancesInstanceApplicableProducts {
	return s.ApplicableProducts
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) GetDeductType() *string {
	return s.DeductType
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) GetExpiryTime() *string {
	return s.ExpiryTime
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) GetPackageType() *string {
	return s.PackageType
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) GetRegion() *string {
	return s.Region
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) GetRemainingAmount() *string {
	return s.RemainingAmount
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) GetRemainingAmountUnit() *string {
	return s.RemainingAmountUnit
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) GetRemark() *string {
	return s.Remark
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) GetStatus() *string {
	return s.Status
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) GetTotalAmount() *string {
	return s.TotalAmount
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) GetTotalAmountUnit() *string {
	return s.TotalAmountUnit
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetApplicableProducts(v *QueryResourcePackageInstancesResponseBodyDataInstancesInstanceApplicableProducts) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.ApplicableProducts = v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetCommodityCode(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.CommodityCode = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetDeductType(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.DeductType = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetEffectiveTime(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.EffectiveTime = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetExpiryTime(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.ExpiryTime = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetInstanceId(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetPackageType(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.PackageType = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetRegion(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.Region = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetRemainingAmount(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.RemainingAmount = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetRemainingAmountUnit(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.RemainingAmountUnit = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetRemark(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.Remark = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetStatus(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.Status = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetTotalAmount(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.TotalAmount = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetTotalAmountUnit(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.TotalAmountUnit = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) Validate() error {
	if s.ApplicableProducts != nil {
		if err := s.ApplicableProducts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryResourcePackageInstancesResponseBodyDataInstancesInstanceApplicableProducts struct {
	Product []*string `json:"Product,omitempty" xml:"Product,omitempty" type:"Repeated"`
}

func (s QueryResourcePackageInstancesResponseBodyDataInstancesInstanceApplicableProducts) String() string {
	return dara.Prettify(s)
}

func (s QueryResourcePackageInstancesResponseBodyDataInstancesInstanceApplicableProducts) GoString() string {
	return s.String()
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstanceApplicableProducts) GetProduct() []*string {
	return s.Product
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstanceApplicableProducts) SetProduct(v []*string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstanceApplicableProducts {
	s.Product = v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstanceApplicableProducts) Validate() error {
	return dara.Validate(s)
}
