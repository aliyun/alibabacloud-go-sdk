// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFootprintListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetFootprintListResponseBodyData) *GetFootprintListResponseBody
	GetData() *GetFootprintListResponseBodyData
	SetRequestId(v string) *GetFootprintListResponseBody
	GetRequestId() *string
}

type GetFootprintListResponseBody struct {
	// The response parameters.
	Data *GetFootprintListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// 06DA2909-7736-5738-AA31-ACD617D828F1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetFootprintListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFootprintListResponseBody) GoString() string {
	return s.String()
}

func (s *GetFootprintListResponseBody) GetData() *GetFootprintListResponseBodyData {
	return s.Data
}

func (s *GetFootprintListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFootprintListResponseBody) SetData(v *GetFootprintListResponseBodyData) *GetFootprintListResponseBody {
	s.Data = v
	return s
}

func (s *GetFootprintListResponseBody) SetRequestId(v string) *GetFootprintListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFootprintListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFootprintListResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Product Detail List.
	Records []*GetFootprintListResponseBodyDataRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 21
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
	// Total Pages
	//
	// example:
	//
	// 3
	TotalPage *int64 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s GetFootprintListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFootprintListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFootprintListResponseBodyData) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *GetFootprintListResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetFootprintListResponseBodyData) GetRecords() []*GetFootprintListResponseBodyDataRecords {
	return s.Records
}

func (s *GetFootprintListResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetFootprintListResponseBodyData) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *GetFootprintListResponseBodyData) SetCurrentPage(v int64) *GetFootprintListResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetFootprintListResponseBodyData) SetPageSize(v int64) *GetFootprintListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetFootprintListResponseBodyData) SetRecords(v []*GetFootprintListResponseBodyDataRecords) *GetFootprintListResponseBodyData {
	s.Records = v
	return s
}

func (s *GetFootprintListResponseBodyData) SetTotal(v int64) *GetFootprintListResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetFootprintListResponseBodyData) SetTotalPage(v int64) *GetFootprintListResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *GetFootprintListResponseBodyData) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetFootprintListResponseBodyDataRecords struct {
	// The amount of the product.
	//
	// example:
	//
	// 100.0000000000000000000000000
	Amount *float64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// Authentication status enumeration value, please refer to [link](https://carbon-doc.oss-cn-hangzhou.aliyuncs.com/CarbonFootprintAppendices-en.pdf).
	//
	// example:
	//
	// 1
	AuthStatus *int64 `json:"authStatus,omitempty" xml:"authStatus,omitempty"`
	// Calculation end time.
	//
	// example:
	//
	// 2024/01/31
	CheckEndTime *string `json:"checkEndTime,omitempty" xml:"checkEndTime,omitempty"`
	// Calculation start time.
	//
	// example:
	//
	// 2024/01/01
	CheckStartTime *string `json:"checkStartTime,omitempty" xml:"checkStartTime,omitempty"`
	// The enterprise code.
	//
	// example:
	//
	// C-20080808-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The user who created it.
	//
	// example:
	//
	// Energy Expert
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// The product ID.
	//
	// example:
	//
	// 1024
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Indicates whether the demo model is a built-in model. A value of 1 indicates a true value and a value of 0 indicates a false value.
	//
	// example:
	//
	// 1
	IsDemoModel *int64 `json:"isDemoModel,omitempty" xml:"isDemoModel,omitempty"`
	// The literal expression corresponding to lifeCycleType, `From Cradle to Gate` and `From Cradle to Grave`.
	//
	// example:
	//
	// From Cradle to Gate
	LifeCycle *string `json:"lifeCycle,omitempty" xml:"lifeCycle,omitempty"`
	// 1 is `From Cradle to Gate`, and 2 is `From Cradle to Grave`.
	//
	// example:
	//
	// 1
	LifeCycleType *int64 `json:"lifeCycleType,omitempty" xml:"lifeCycleType,omitempty"`
	// The product name.
	//
	// example:
	//
	// Carbon-footprint-demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Product category enumeration value, please refer to [Carbon Footprint Appendices](https://carbon-doc.oss-cn-hangzhou.aliyuncs.com/CarbonFootprintAppendices-en.pdf).
	//
	// example:
	//
	// 158-159
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Unit enumeration value. Please refer to [Carbon Footprint Appendices](https://carbon-doc.oss-cn-hangzhou.aliyuncs.com/CarbonFootprintAppendices-en.pdf).
	//
	// example:
	//
	// 1-4
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GetFootprintListResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s GetFootprintListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetFootprintListResponseBodyDataRecords) GetAmount() *float64 {
	return s.Amount
}

func (s *GetFootprintListResponseBodyDataRecords) GetAuthStatus() *int64 {
	return s.AuthStatus
}

func (s *GetFootprintListResponseBodyDataRecords) GetCheckEndTime() *string {
	return s.CheckEndTime
}

func (s *GetFootprintListResponseBodyDataRecords) GetCheckStartTime() *string {
	return s.CheckStartTime
}

func (s *GetFootprintListResponseBodyDataRecords) GetCode() *string {
	return s.Code
}

func (s *GetFootprintListResponseBodyDataRecords) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *GetFootprintListResponseBodyDataRecords) GetId() *int64 {
	return s.Id
}

func (s *GetFootprintListResponseBodyDataRecords) GetIsDemoModel() *int64 {
	return s.IsDemoModel
}

func (s *GetFootprintListResponseBodyDataRecords) GetLifeCycle() *string {
	return s.LifeCycle
}

func (s *GetFootprintListResponseBodyDataRecords) GetLifeCycleType() *int64 {
	return s.LifeCycleType
}

func (s *GetFootprintListResponseBodyDataRecords) GetName() *string {
	return s.Name
}

func (s *GetFootprintListResponseBodyDataRecords) GetType() *string {
	return s.Type
}

func (s *GetFootprintListResponseBodyDataRecords) GetUnit() *string {
	return s.Unit
}

func (s *GetFootprintListResponseBodyDataRecords) SetAmount(v float64) *GetFootprintListResponseBodyDataRecords {
	s.Amount = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetAuthStatus(v int64) *GetFootprintListResponseBodyDataRecords {
	s.AuthStatus = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetCheckEndTime(v string) *GetFootprintListResponseBodyDataRecords {
	s.CheckEndTime = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetCheckStartTime(v string) *GetFootprintListResponseBodyDataRecords {
	s.CheckStartTime = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetCode(v string) *GetFootprintListResponseBodyDataRecords {
	s.Code = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetCreatedBy(v string) *GetFootprintListResponseBodyDataRecords {
	s.CreatedBy = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetId(v int64) *GetFootprintListResponseBodyDataRecords {
	s.Id = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetIsDemoModel(v int64) *GetFootprintListResponseBodyDataRecords {
	s.IsDemoModel = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetLifeCycle(v string) *GetFootprintListResponseBodyDataRecords {
	s.LifeCycle = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetLifeCycleType(v int64) *GetFootprintListResponseBodyDataRecords {
	s.LifeCycleType = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetName(v string) *GetFootprintListResponseBodyDataRecords {
	s.Name = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetType(v string) *GetFootprintListResponseBodyDataRecords {
	s.Type = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetUnit(v string) *GetFootprintListResponseBodyDataRecords {
	s.Unit = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
