// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterConfigChangeLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeDBClusterConfigChangeLogsResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *DescribeDBClusterConfigChangeLogsResponseBodyData) *DescribeDBClusterConfigChangeLogsResponseBody
	GetData() *DescribeDBClusterConfigChangeLogsResponseBodyData
	SetDynamicCode(v string) *DescribeDBClusterConfigChangeLogsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeDBClusterConfigChangeLogsResponseBody
	GetDynamicMessage() *string
	SetRequestId(v string) *DescribeDBClusterConfigChangeLogsResponseBody
	GetRequestId() *string
}

type DescribeDBClusterConfigChangeLogsResponseBody struct {
	// example:
	//
	// failed
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The information returned.
	Data *DescribeDBClusterConfigChangeLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The dynamic code. This parameter is not returned.
	//
	// example:
	//
	// 0
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic message. This parameter is not returned.
	//
	// example:
	//
	// An error occurred while processing your request.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// F8900A96-67F7-5274-A41B-7722E1ECF8C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterConfigChangeLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterConfigChangeLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigChangeLogsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeDBClusterConfigChangeLogsResponseBody) GetData() *DescribeDBClusterConfigChangeLogsResponseBodyData {
	return s.Data
}

func (s *DescribeDBClusterConfigChangeLogsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeDBClusterConfigChangeLogsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeDBClusterConfigChangeLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterConfigChangeLogsResponseBody) SetAccessDeniedDetail(v string) *DescribeDBClusterConfigChangeLogsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBody) SetData(v *DescribeDBClusterConfigChangeLogsResponseBodyData) *DescribeDBClusterConfigChangeLogsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBody) SetDynamicCode(v string) *DescribeDBClusterConfigChangeLogsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBody) SetDynamicMessage(v string) *DescribeDBClusterConfigChangeLogsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBody) SetRequestId(v string) *DescribeDBClusterConfigChangeLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterConfigChangeLogsResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// selectdb-cn-wny3li00g02-be
	DbClusterId *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	// example:
	//
	// 6585
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// selectdb-cn-wny3li00g02
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	// The parameter change logs.
	ParamChangeLogs []*DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs `json:"ParamChangeLogs,omitempty" xml:"ParamChangeLogs,omitempty" type:"Repeated"`
	// The task ID.
	//
	// example:
	//
	// 107841167
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeDBClusterConfigChangeLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterConfigChangeLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyData) GetDbClusterId() *string {
	return s.DbClusterId
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyData) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyData) GetDbInstanceName() *string {
	return s.DbInstanceName
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyData) GetParamChangeLogs() []*DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs {
	return s.ParamChangeLogs
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyData) SetDbClusterId(v string) *DescribeDBClusterConfigChangeLogsResponseBodyData {
	s.DbClusterId = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyData) SetDbInstanceId(v string) *DescribeDBClusterConfigChangeLogsResponseBodyData {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyData) SetDbInstanceName(v string) *DescribeDBClusterConfigChangeLogsResponseBodyData {
	s.DbInstanceName = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyData) SetParamChangeLogs(v []*DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) *DescribeDBClusterConfigChangeLogsResponseBodyData {
	s.ParamChangeLogs = v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyData) SetTaskId(v int32) *DescribeDBClusterConfigChangeLogsResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs struct {
	// example:
	//
	// 2022-10-11T08:53:32Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 2024-03-08T10:08Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the change log.
	//
	// example:
	//
	// 617975
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the modification has taken effect.
	//
	// example:
	//
	// false
	IsApplied *bool `json:"IsApplied,omitempty" xml:"IsApplied,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// cumulative_compaction_rounds_for_each_base_compaction_round
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 12
	NewValue *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
	// example:
	//
	// 10
	OldValue *string `json:"OldValue,omitempty" xml:"OldValue,omitempty"`
}

func (s DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) GetId() *int64 {
	return s.Id
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) GetIsApplied() *bool {
	return s.IsApplied
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) GetName() *string {
	return s.Name
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) GetNewValue() *string {
	return s.NewValue
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) GetOldValue() *string {
	return s.OldValue
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) SetGmtCreated(v string) *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs {
	s.GmtCreated = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) SetGmtModified(v string) *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs {
	s.GmtModified = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) SetId(v int64) *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs {
	s.Id = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) SetIsApplied(v bool) *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs {
	s.IsApplied = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) SetName(v string) *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) SetNewValue(v string) *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs {
	s.NewValue = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) SetOldValue(v string) *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs {
	s.OldValue = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) Validate() error {
	return dara.Validate(s)
}
