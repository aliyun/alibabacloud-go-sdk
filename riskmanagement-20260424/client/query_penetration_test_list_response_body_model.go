// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPenetrationTestListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryPenetrationTestListResponseBody
	GetCode() *string
	SetData(v *QueryPenetrationTestListResponseBodyData) *QueryPenetrationTestListResponseBody
	GetData() *QueryPenetrationTestListResponseBodyData
	SetMessage(v string) *QueryPenetrationTestListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryPenetrationTestListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryPenetrationTestListResponseBody
	GetSuccess() *bool
}

type QueryPenetrationTestListResponseBody struct {
	// The status code.
	//
	// - **200**: Success.
	//
	// - **Other values (such as 400 and 500)**: Failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query results.
	Data *QueryPenetrationTestListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 99D93ED4-D462-5FC5-8518-9BC1C49C7B6C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - **true**: The call is successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPenetrationTestListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPenetrationTestListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPenetrationTestListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryPenetrationTestListResponseBody) GetData() *QueryPenetrationTestListResponseBodyData {
	return s.Data
}

func (s *QueryPenetrationTestListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryPenetrationTestListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPenetrationTestListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryPenetrationTestListResponseBody) SetCode(v string) *QueryPenetrationTestListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPenetrationTestListResponseBody) SetData(v *QueryPenetrationTestListResponseBodyData) *QueryPenetrationTestListResponseBody {
	s.Data = v
	return s
}

func (s *QueryPenetrationTestListResponseBody) SetMessage(v string) *QueryPenetrationTestListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPenetrationTestListResponseBody) SetRequestId(v string) *QueryPenetrationTestListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPenetrationTestListResponseBody) SetSuccess(v bool) *QueryPenetrationTestListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryPenetrationTestListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPenetrationTestListResponseBodyData struct {
	// The event data.
	List []*QueryPenetrationTestListResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *QueryPenetrationTestListResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
}

func (s QueryPenetrationTestListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryPenetrationTestListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPenetrationTestListResponseBodyData) GetList() []*QueryPenetrationTestListResponseBodyDataList {
	return s.List
}

func (s *QueryPenetrationTestListResponseBodyData) GetPageInfo() *QueryPenetrationTestListResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *QueryPenetrationTestListResponseBodyData) SetList(v []*QueryPenetrationTestListResponseBodyDataList) *QueryPenetrationTestListResponseBodyData {
	s.List = v
	return s
}

func (s *QueryPenetrationTestListResponseBodyData) SetPageInfo(v *QueryPenetrationTestListResponseBodyDataPageInfo) *QueryPenetrationTestListResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *QueryPenetrationTestListResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPenetrationTestListResponseBodyDataList struct {
	// The application time.
	//
	// example:
	//
	// 1788873796000
	ApplyTime *int64 `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	// The type of the asset. Valid values:
	//
	// - **EIP**: The public IP address of a VPC.
	//
	// - **SLB**: Server Load Balancer.
	//
	// - **ECS**: The public IP address of an ECS instance in the classic network.
	//
	// example:
	//
	// EIP
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The unique identifier of the alert event record.
	//
	// > The data type of this field is Long. Precision loss may occur during the serialization or deserialization process. Make sure that the value does not exceed 9007199254740991.
	//
	// example:
	//
	// 710864062
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The IP address of the asset.
	//
	// example:
	//
	// 10.0.158.58
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The status of the task. Valid values:
	//
	// - **early_end**: The user applied for early termination.
	//
	// - **expired**: The authorization expired.
	//
	// - **auditing**: The application is being reviewed.
	//
	// - **approved**: The application is approved.
	//
	// - **rejected**: The application is rejected.
	//
	// example:
	//
	// early_end
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The end time of the test.
	//
	// example:
	//
	// 1788873796000
	TestEndDate *int64 `json:"TestEndDate,omitempty" xml:"TestEndDate,omitempty"`
	// The start date of the test.
	//
	// example:
	//
	// 1788873796000
	TestStartDate *int64 `json:"TestStartDate,omitempty" xml:"TestStartDate,omitempty"`
}

func (s QueryPenetrationTestListResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s QueryPenetrationTestListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryPenetrationTestListResponseBodyDataList) GetApplyTime() *int64 {
	return s.ApplyTime
}

func (s *QueryPenetrationTestListResponseBodyDataList) GetAssetType() *string {
	return s.AssetType
}

func (s *QueryPenetrationTestListResponseBodyDataList) GetId() *int64 {
	return s.Id
}

func (s *QueryPenetrationTestListResponseBodyDataList) GetIp() *string {
	return s.Ip
}

func (s *QueryPenetrationTestListResponseBodyDataList) GetStatus() *string {
	return s.Status
}

func (s *QueryPenetrationTestListResponseBodyDataList) GetTestEndDate() *int64 {
	return s.TestEndDate
}

func (s *QueryPenetrationTestListResponseBodyDataList) GetTestStartDate() *int64 {
	return s.TestStartDate
}

func (s *QueryPenetrationTestListResponseBodyDataList) SetApplyTime(v int64) *QueryPenetrationTestListResponseBodyDataList {
	s.ApplyTime = &v
	return s
}

func (s *QueryPenetrationTestListResponseBodyDataList) SetAssetType(v string) *QueryPenetrationTestListResponseBodyDataList {
	s.AssetType = &v
	return s
}

func (s *QueryPenetrationTestListResponseBodyDataList) SetId(v int64) *QueryPenetrationTestListResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *QueryPenetrationTestListResponseBodyDataList) SetIp(v string) *QueryPenetrationTestListResponseBodyDataList {
	s.Ip = &v
	return s
}

func (s *QueryPenetrationTestListResponseBodyDataList) SetStatus(v string) *QueryPenetrationTestListResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *QueryPenetrationTestListResponseBodyDataList) SetTestEndDate(v int64) *QueryPenetrationTestListResponseBodyDataList {
	s.TestEndDate = &v
	return s
}

func (s *QueryPenetrationTestListResponseBodyDataList) SetTestStartDate(v int64) *QueryPenetrationTestListResponseBodyDataList {
	s.TestStartDate = &v
	return s
}

func (s *QueryPenetrationTestListResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type QueryPenetrationTestListResponseBodyDataPageInfo struct {
	// The page number from which the query results are displayed.
	//
	// > Default value: 1, which indicates that the query results are displayed starting from the first page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 24
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of events.
	//
	// example:
	//
	// 20
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryPenetrationTestListResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryPenetrationTestListResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *QueryPenetrationTestListResponseBodyDataPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryPenetrationTestListResponseBodyDataPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryPenetrationTestListResponseBodyDataPageInfo) GetTotal() *int32 {
	return s.Total
}

func (s *QueryPenetrationTestListResponseBodyDataPageInfo) SetCurrentPage(v int32) *QueryPenetrationTestListResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *QueryPenetrationTestListResponseBodyDataPageInfo) SetPageSize(v int32) *QueryPenetrationTestListResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *QueryPenetrationTestListResponseBodyDataPageInfo) SetTotal(v int32) *QueryPenetrationTestListResponseBodyDataPageInfo {
	s.Total = &v
	return s
}

func (s *QueryPenetrationTestListResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}
