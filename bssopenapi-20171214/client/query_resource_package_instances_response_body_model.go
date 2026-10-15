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
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QueryResourcePackageInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
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
	// The total number of records.
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
	// The host ID.
	//
	// example:
	//
	// cn
	HostId    *string                                                 `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Instances *QueryResourcePackageInstancesResponseBodyDataInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *string `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
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
	ApplicableProducts  *QueryResourcePackageInstancesResponseBodyDataInstancesInstanceApplicableProducts `json:"ApplicableProducts,omitempty" xml:"ApplicableProducts,omitempty" type:"Struct"`
	CommodityCode       *string                                                                           `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	DeductType          *string                                                                           `json:"DeductType,omitempty" xml:"DeductType,omitempty"`
	EffectiveTime       *string                                                                           `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	ExpiryTime          *string                                                                           `json:"ExpiryTime,omitempty" xml:"ExpiryTime,omitempty"`
	InstanceId          *string                                                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PackageType         *string                                                                           `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	Region              *string                                                                           `json:"Region,omitempty" xml:"Region,omitempty"`
	RemainingAmount     *string                                                                           `json:"RemainingAmount,omitempty" xml:"RemainingAmount,omitempty"`
	RemainingAmountUnit *string                                                                           `json:"RemainingAmountUnit,omitempty" xml:"RemainingAmountUnit,omitempty"`
	Remark              *string                                                                           `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status              *string                                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalAmount         *string                                                                           `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
	TotalAmountUnit     *string                                                                           `json:"TotalAmountUnit,omitempty" xml:"TotalAmountUnit,omitempty"`
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
