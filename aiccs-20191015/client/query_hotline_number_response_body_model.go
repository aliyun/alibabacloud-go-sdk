// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHotlineNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryHotlineNumberResponseBody
	GetCode() *string
	SetData(v *QueryHotlineNumberResponseBodyData) *QueryHotlineNumberResponseBody
	GetData() *QueryHotlineNumberResponseBodyData
	SetMessage(v string) *QueryHotlineNumberResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryHotlineNumberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryHotlineNumberResponseBody
	GetSuccess() *bool
}

type QueryHotlineNumberResponseBody struct {
	// The status code. A value of Success indicates that the request succeeded.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Hotline number configuration information.
	Data *QueryHotlineNumberResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Status code description.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE339D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryHotlineNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryHotlineNumberResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHotlineNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryHotlineNumberResponseBody) GetData() *QueryHotlineNumberResponseBodyData {
	return s.Data
}

func (s *QueryHotlineNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryHotlineNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryHotlineNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryHotlineNumberResponseBody) SetCode(v string) *QueryHotlineNumberResponseBody {
	s.Code = &v
	return s
}

func (s *QueryHotlineNumberResponseBody) SetData(v *QueryHotlineNumberResponseBodyData) *QueryHotlineNumberResponseBody {
	s.Data = v
	return s
}

func (s *QueryHotlineNumberResponseBody) SetMessage(v string) *QueryHotlineNumberResponseBody {
	s.Message = &v
	return s
}

func (s *QueryHotlineNumberResponseBody) SetRequestId(v string) *QueryHotlineNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHotlineNumberResponseBody) SetSuccess(v bool) *QueryHotlineNumberResponseBody {
	s.Success = &v
	return s
}

func (s *QueryHotlineNumberResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryHotlineNumberResponseBodyData struct {
	// Current page.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Number list.
	HotlineNumList []*QueryHotlineNumberResponseBodyDataHotlineNumList `json:"HotlineNumList,omitempty" xml:"HotlineNumList,omitempty" type:"Repeated"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total amount of data.
	//
	// example:
	//
	// 123
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryHotlineNumberResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryHotlineNumberResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryHotlineNumberResponseBodyData) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *QueryHotlineNumberResponseBodyData) GetHotlineNumList() []*QueryHotlineNumberResponseBodyDataHotlineNumList {
	return s.HotlineNumList
}

func (s *QueryHotlineNumberResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryHotlineNumberResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QueryHotlineNumberResponseBodyData) SetCurrentPage(v int64) *QueryHotlineNumberResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyData) SetHotlineNumList(v []*QueryHotlineNumberResponseBodyDataHotlineNumList) *QueryHotlineNumberResponseBodyData {
	s.HotlineNumList = v
	return s
}

func (s *QueryHotlineNumberResponseBodyData) SetPageSize(v int64) *QueryHotlineNumberResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyData) SetTotalCount(v int64) *QueryHotlineNumberResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyData) Validate() error {
	if s.HotlineNumList != nil {
		for _, item := range s.HotlineNumList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryHotlineNumberResponseBodyDataHotlineNumList struct {
	// Outbound calls apply to all departments.
	//
	// example:
	//
	// true
	CalloutAllDepartment *bool `json:"CalloutAllDepartment,omitempty" xml:"CalloutAllDepartment,omitempty"`
	// List of departments for which outbound calls are effective.
	CalloutRangeList []*QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList `json:"CalloutRangeList,omitempty" xml:"CalloutRangeList,omitempty" type:"Repeated"`
	// Number description.
	//
	// example:
	//
	// 测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Satisfaction status. Valid values:
	//
	// - **0**: Neither inbound nor outbound calls are enabled.
	//
	// - **1**: Inbound calls are enabled.
	//
	// - **2**: Outbound calls are enabled.
	//
	// - **3**: Both inbound and outbound calls are enabled.
	//
	// example:
	//
	// 0
	EvaluationStatus *int32 `json:"EvaluationStatus,omitempty" xml:"EvaluationStatus,omitempty"`
	// Incoming call flow ID.
	//
	// example:
	//
	// 12
	FlowId *int64 `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// Inbound flow name.
	//
	// example:
	//
	// 测试流程
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// Hotline number.
	//
	// example:
	//
	// 0571****2211
	HotlineNumber *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	// Indicates whether the number is used for inbound calls.
	//
	// example:
	//
	// true
	InBoundEnabled *bool `json:"InBoundEnabled,omitempty" xml:"InBoundEnabled,omitempty"`
	// Number location.
	//
	// example:
	//
	// 浙江杭州
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// Indicates whether the number is used for outbound calls.
	//
	// example:
	//
	// true
	OutboundEnabled *bool `json:"OutboundEnabled,omitempty" xml:"OutboundEnabled,omitempty"`
	// Carrier.
	//
	// example:
	//
	// 电信
	Sp *string `json:"Sp,omitempty" xml:"Sp,omitempty"`
}

func (s QueryHotlineNumberResponseBodyDataHotlineNumList) String() string {
	return dara.Prettify(s)
}

func (s QueryHotlineNumberResponseBodyDataHotlineNumList) GoString() string {
	return s.String()
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) GetCalloutAllDepartment() *bool {
	return s.CalloutAllDepartment
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) GetCalloutRangeList() []*QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList {
	return s.CalloutRangeList
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) GetDescription() *string {
	return s.Description
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) GetEvaluationStatus() *int32 {
	return s.EvaluationStatus
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) GetFlowId() *int64 {
	return s.FlowId
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) GetFlowName() *string {
	return s.FlowName
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) GetHotlineNumber() *string {
	return s.HotlineNumber
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) GetInBoundEnabled() *bool {
	return s.InBoundEnabled
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) GetLocation() *string {
	return s.Location
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) GetOutboundEnabled() *bool {
	return s.OutboundEnabled
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) GetSp() *string {
	return s.Sp
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) SetCalloutAllDepartment(v bool) *QueryHotlineNumberResponseBodyDataHotlineNumList {
	s.CalloutAllDepartment = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) SetCalloutRangeList(v []*QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList) *QueryHotlineNumberResponseBodyDataHotlineNumList {
	s.CalloutRangeList = v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) SetDescription(v string) *QueryHotlineNumberResponseBodyDataHotlineNumList {
	s.Description = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) SetEvaluationStatus(v int32) *QueryHotlineNumberResponseBodyDataHotlineNumList {
	s.EvaluationStatus = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) SetFlowId(v int64) *QueryHotlineNumberResponseBodyDataHotlineNumList {
	s.FlowId = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) SetFlowName(v string) *QueryHotlineNumberResponseBodyDataHotlineNumList {
	s.FlowName = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) SetHotlineNumber(v string) *QueryHotlineNumberResponseBodyDataHotlineNumList {
	s.HotlineNumber = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) SetInBoundEnabled(v bool) *QueryHotlineNumberResponseBodyDataHotlineNumList {
	s.InBoundEnabled = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) SetLocation(v string) *QueryHotlineNumberResponseBodyDataHotlineNumList {
	s.Location = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) SetOutboundEnabled(v bool) *QueryHotlineNumberResponseBodyDataHotlineNumList {
	s.OutboundEnabled = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) SetSp(v string) *QueryHotlineNumberResponseBodyDataHotlineNumList {
	s.Sp = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) Validate() error {
	if s.CalloutRangeList != nil {
		for _, item := range s.CalloutRangeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList struct {
	// Department ID.
	//
	// example:
	//
	// 2256****
	DepartmentId *int64 `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// The department name.
	//
	// example:
	//
	// 部门A
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	// The skill group list.
	GroupDOList []*QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeListGroupDOList `json:"GroupDOList,omitempty" xml:"GroupDOList,omitempty" type:"Repeated"`
}

func (s QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList) String() string {
	return dara.Prettify(s)
}

func (s QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList) GoString() string {
	return s.String()
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList) GetDepartmentId() *int64 {
	return s.DepartmentId
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList) GetDepartmentName() *string {
	return s.DepartmentName
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList) GetGroupDOList() []*QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeListGroupDOList {
	return s.GroupDOList
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList) SetDepartmentId(v int64) *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList {
	s.DepartmentId = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList) SetDepartmentName(v string) *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList {
	s.DepartmentName = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList) SetGroupDOList(v []*QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeListGroupDOList) *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList {
	s.GroupDOList = v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList) Validate() error {
	if s.GroupDOList != nil {
		for _, item := range s.GroupDOList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeListGroupDOList struct {
	// The skill group ID.
	//
	// example:
	//
	// 6083****
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The skill group name.
	//
	// example:
	//
	// 技能组A
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeListGroupDOList) String() string {
	return dara.Prettify(s)
}

func (s QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeListGroupDOList) GoString() string {
	return s.String()
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeListGroupDOList) GetGroupId() *int64 {
	return s.GroupId
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeListGroupDOList) GetGroupName() *string {
	return s.GroupName
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeListGroupDOList) SetGroupId(v int64) *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeListGroupDOList {
	s.GroupId = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeListGroupDOList) SetGroupName(v string) *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeListGroupDOList {
	s.GroupName = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeListGroupDOList) Validate() error {
	return dara.Validate(s)
}
