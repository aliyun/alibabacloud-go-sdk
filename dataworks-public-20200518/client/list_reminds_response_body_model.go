// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemindsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListRemindsResponseBodyData) *ListRemindsResponseBody
	GetData() *ListRemindsResponseBodyData
	SetErrorCode(v string) *ListRemindsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListRemindsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListRemindsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListRemindsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListRemindsResponseBody
	GetSuccess() *bool
}

type ListRemindsResponseBody struct {
	// The data returned.
	Data *ListRemindsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned.
	//
	// example:
	//
	// 1031203110005
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The specified parameters are invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRemindsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRemindsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRemindsResponseBody) GetData() *ListRemindsResponseBodyData {
	return s.Data
}

func (s *ListRemindsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListRemindsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListRemindsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListRemindsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRemindsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRemindsResponseBody) SetData(v *ListRemindsResponseBodyData) *ListRemindsResponseBody {
	s.Data = v
	return s
}

func (s *ListRemindsResponseBody) SetErrorCode(v string) *ListRemindsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRemindsResponseBody) SetErrorMessage(v string) *ListRemindsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRemindsResponseBody) SetHttpStatusCode(v int32) *ListRemindsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRemindsResponseBody) SetRequestId(v string) *ListRemindsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRemindsResponseBody) SetSuccess(v bool) *ListRemindsResponseBody {
	s.Success = &v
	return s
}

func (s *ListRemindsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRemindsResponseBodyData struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of custom alert rules.
	Reminds []*ListRemindsResponseBodyDataReminds `json:"Reminds,omitempty" xml:"Reminds,omitempty" type:"Repeated"`
	// The total number of custom alert rules returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRemindsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRemindsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRemindsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRemindsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRemindsResponseBodyData) GetReminds() []*ListRemindsResponseBodyDataReminds {
	return s.Reminds
}

func (s *ListRemindsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRemindsResponseBodyData) SetPageNumber(v int32) *ListRemindsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListRemindsResponseBodyData) SetPageSize(v int32) *ListRemindsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListRemindsResponseBodyData) SetReminds(v []*ListRemindsResponseBodyDataReminds) *ListRemindsResponseBodyData {
	s.Reminds = v
	return s
}

func (s *ListRemindsResponseBodyData) SetTotalCount(v int32) *ListRemindsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListRemindsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListRemindsResponseBodyDataReminds struct {
	// The notification method. Valid values: MAIL, SMS, and PHONE. The value MAIL indicates that the notification is sent by email. Only DataWorks Professional Edition and more advanced editions support the PHONE notification method.
	AlertMethods []*string `json:"AlertMethods,omitempty" xml:"AlertMethods,omitempty" type:"Repeated"`
	// The IDs of the Alibaba Cloud accounts used by alert recipients.
	AlertTargets []*string `json:"AlertTargets,omitempty" xml:"AlertTargets,omitempty" type:"Repeated"`
	// The alert recipient. Valid values: OWNER and OTHER. The value OWNER indicates the node owner. The value OTHER indicates a specified user.
	//
	// example:
	//
	// OWNER
	AlertUnit *string `json:"AlertUnit,omitempty" xml:"AlertUnit,omitempty"`
	// The IDs of the baselines to which the custom alert rule is applied. This parameter is returned if the value of the RemindUnit parameter is BASELINE.
	BaselineIds []*int64 `json:"BaselineIds,omitempty" xml:"BaselineIds,omitempty" type:"Repeated"`
	// The IDs of the workflows to which the custom alert rule is applied. This parameter is returned if the value of the RemindUnit parameter is BIZPROCESS.
	BizProcessIds []*int64 `json:"BizProcessIds,omitempty" xml:"BizProcessIds,omitempty" type:"Repeated"`
	// The end time of the quiet hours. The time is in the hh:mm format. Valid values of hh: [0,23]. Valid values of mm: [0,59].
	//
	// example:
	//
	// 08:00
	DndEnd *string `json:"DndEnd,omitempty" xml:"DndEnd,omitempty"`
	// The start time of the quiet hours. The time is in the hh:mm format. Valid values of hh: [0,23]. Valid values of mm: [0,59].
	//
	// example:
	//
	// 00:00
	DndStart *string `json:"DndStart,omitempty" xml:"DndStart,omitempty"`
	// The ID of the Alibaba Cloud account used by the rule creator.
	//
	// example:
	//
	// 952795****
	Founder *string `json:"Founder,omitempty" xml:"Founder,omitempty"`
	// The IDs of the nodes to which the custom alert rule is applied. This parameter is returned if the value of the RemindUnit parameter is NODE.
	NodeIds []*int64 `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// The IDs of the workspaces to which the custom alert rule is applied. This parameter is returned if the value of the RemindUnit parameter is PROJECT.
	ProjectIds []*int64 `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The custom alert rule ID.
	//
	// example:
	//
	// 1234
	RemindId *int64 `json:"RemindId,omitempty" xml:"RemindId,omitempty"`
	// The name of the custom alert rule.
	//
	// example:
	//
	// Alert Rule
	RemindName *string `json:"RemindName,omitempty" xml:"RemindName,omitempty"`
	// The condition that triggers an alert. Valid values: FINISHED, UNFINISHED, ERROR, CYCLE_UNFINISHED, and TIMEOUT.
	//
	// example:
	//
	// FINISHED
	RemindType *string `json:"RemindType,omitempty" xml:"RemindType,omitempty"`
	// The type of the object to which the custom alert rule is applied. Valid values: NODE, BASELINE, PROJECT, and BIZPROCESS. The value NODE indicates a node. The value BASELINE indicates a baseline. The value PROJECT indicates a workspace. The value BIZPROCESS indicates a workflow.
	//
	// example:
	//
	// NODE
	RemindUnit *string `json:"RemindUnit,omitempty" xml:"RemindUnit,omitempty"`
	// Indicates whether the custom alert rule is enabled. Valid values: true and false.
	//
	// example:
	//
	// true
	Useflag *bool `json:"Useflag,omitempty" xml:"Useflag,omitempty"`
}

func (s ListRemindsResponseBodyDataReminds) String() string {
	return dara.Prettify(s)
}

func (s ListRemindsResponseBodyDataReminds) GoString() string {
	return s.String()
}

func (s *ListRemindsResponseBodyDataReminds) GetAlertMethods() []*string {
	return s.AlertMethods
}

func (s *ListRemindsResponseBodyDataReminds) GetAlertTargets() []*string {
	return s.AlertTargets
}

func (s *ListRemindsResponseBodyDataReminds) GetAlertUnit() *string {
	return s.AlertUnit
}

func (s *ListRemindsResponseBodyDataReminds) GetBaselineIds() []*int64 {
	return s.BaselineIds
}

func (s *ListRemindsResponseBodyDataReminds) GetBizProcessIds() []*int64 {
	return s.BizProcessIds
}

func (s *ListRemindsResponseBodyDataReminds) GetDndEnd() *string {
	return s.DndEnd
}

func (s *ListRemindsResponseBodyDataReminds) GetDndStart() *string {
	return s.DndStart
}

func (s *ListRemindsResponseBodyDataReminds) GetFounder() *string {
	return s.Founder
}

func (s *ListRemindsResponseBodyDataReminds) GetNodeIds() []*int64 {
	return s.NodeIds
}

func (s *ListRemindsResponseBodyDataReminds) GetProjectIds() []*int64 {
	return s.ProjectIds
}

func (s *ListRemindsResponseBodyDataReminds) GetRemindId() *int64 {
	return s.RemindId
}

func (s *ListRemindsResponseBodyDataReminds) GetRemindName() *string {
	return s.RemindName
}

func (s *ListRemindsResponseBodyDataReminds) GetRemindType() *string {
	return s.RemindType
}

func (s *ListRemindsResponseBodyDataReminds) GetRemindUnit() *string {
	return s.RemindUnit
}

func (s *ListRemindsResponseBodyDataReminds) GetUseflag() *bool {
	return s.Useflag
}

func (s *ListRemindsResponseBodyDataReminds) SetAlertMethods(v []*string) *ListRemindsResponseBodyDataReminds {
	s.AlertMethods = v
	return s
}

func (s *ListRemindsResponseBodyDataReminds) SetAlertTargets(v []*string) *ListRemindsResponseBodyDataReminds {
	s.AlertTargets = v
	return s
}

func (s *ListRemindsResponseBodyDataReminds) SetAlertUnit(v string) *ListRemindsResponseBodyDataReminds {
	s.AlertUnit = &v
	return s
}

func (s *ListRemindsResponseBodyDataReminds) SetBaselineIds(v []*int64) *ListRemindsResponseBodyDataReminds {
	s.BaselineIds = v
	return s
}

func (s *ListRemindsResponseBodyDataReminds) SetBizProcessIds(v []*int64) *ListRemindsResponseBodyDataReminds {
	s.BizProcessIds = v
	return s
}

func (s *ListRemindsResponseBodyDataReminds) SetDndEnd(v string) *ListRemindsResponseBodyDataReminds {
	s.DndEnd = &v
	return s
}

func (s *ListRemindsResponseBodyDataReminds) SetDndStart(v string) *ListRemindsResponseBodyDataReminds {
	s.DndStart = &v
	return s
}

func (s *ListRemindsResponseBodyDataReminds) SetFounder(v string) *ListRemindsResponseBodyDataReminds {
	s.Founder = &v
	return s
}

func (s *ListRemindsResponseBodyDataReminds) SetNodeIds(v []*int64) *ListRemindsResponseBodyDataReminds {
	s.NodeIds = v
	return s
}

func (s *ListRemindsResponseBodyDataReminds) SetProjectIds(v []*int64) *ListRemindsResponseBodyDataReminds {
	s.ProjectIds = v
	return s
}

func (s *ListRemindsResponseBodyDataReminds) SetRemindId(v int64) *ListRemindsResponseBodyDataReminds {
	s.RemindId = &v
	return s
}

func (s *ListRemindsResponseBodyDataReminds) SetRemindName(v string) *ListRemindsResponseBodyDataReminds {
	s.RemindName = &v
	return s
}

func (s *ListRemindsResponseBodyDataReminds) SetRemindType(v string) *ListRemindsResponseBodyDataReminds {
	s.RemindType = &v
	return s
}

func (s *ListRemindsResponseBodyDataReminds) SetRemindUnit(v string) *ListRemindsResponseBodyDataReminds {
	s.RemindUnit = &v
	return s
}

func (s *ListRemindsResponseBodyDataReminds) SetUseflag(v bool) *ListRemindsResponseBodyDataReminds {
	s.Useflag = &v
	return s
}

func (s *ListRemindsResponseBodyDataReminds) Validate() error {
	return dara.Validate(s)
}
