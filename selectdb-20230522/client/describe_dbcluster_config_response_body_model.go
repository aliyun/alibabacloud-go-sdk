// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeDBClusterConfigResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *DescribeDBClusterConfigResponseBodyData) *DescribeDBClusterConfigResponseBody
	GetData() *DescribeDBClusterConfigResponseBodyData
	SetDynamicCode(v string) *DescribeDBClusterConfigResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeDBClusterConfigResponseBody
	GetDynamicMessage() *string
	SetRequestId(v string) *DescribeDBClusterConfigResponseBody
	GetRequestId() *string
}

type DescribeDBClusterConfigResponseBody struct {
	// The details about the access denial. This parameter is returned only if Resource Access Management (RAM) authentication failed.
	//
	// example:
	//
	// failed
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The information returned.
	Data *DescribeDBClusterConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// The request ID.
	//
	// example:
	//
	// ADF42B18-43FD-5100-83A9-BE81AB70C863
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeDBClusterConfigResponseBody) GetData() *DescribeDBClusterConfigResponseBodyData {
	return s.Data
}

func (s *DescribeDBClusterConfigResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeDBClusterConfigResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeDBClusterConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterConfigResponseBody) SetAccessDeniedDetail(v string) *DescribeDBClusterConfigResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBody) SetData(v *DescribeDBClusterConfigResponseBodyData) *DescribeDBClusterConfigResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBClusterConfigResponseBody) SetDynamicCode(v string) *DescribeDBClusterConfigResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBody) SetDynamicMessage(v string) *DescribeDBClusterConfigResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBody) SetRequestId(v string) *DescribeDBClusterConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterConfigResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// selectdb-cn-wny3li00g02-be
	DbClusterId *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	// The numeric ID of the instance.
	//
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
	// The details about each parameter returned.
	Params []*DescribeDBClusterConfigResponseBodyDataParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The task ID.
	//
	// example:
	//
	// 107841167
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeDBClusterConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigResponseBodyData) GetDbClusterId() *string {
	return s.DbClusterId
}

func (s *DescribeDBClusterConfigResponseBodyData) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *DescribeDBClusterConfigResponseBodyData) GetDbInstanceName() *string {
	return s.DbInstanceName
}

func (s *DescribeDBClusterConfigResponseBodyData) GetParams() []*DescribeDBClusterConfigResponseBodyDataParams {
	return s.Params
}

func (s *DescribeDBClusterConfigResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DescribeDBClusterConfigResponseBodyData) SetDbClusterId(v string) *DescribeDBClusterConfigResponseBodyData {
	s.DbClusterId = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyData) SetDbInstanceId(v string) *DescribeDBClusterConfigResponseBodyData {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyData) SetDbInstanceName(v string) *DescribeDBClusterConfigResponseBodyData {
	s.DbInstanceName = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyData) SetParams(v []*DescribeDBClusterConfigResponseBodyDataParams) *DescribeDBClusterConfigResponseBodyData {
	s.Params = v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyData) SetTaskId(v int32) *DescribeDBClusterConfigResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterConfigResponseBodyDataParams struct {
	// The comments on the parameter.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The default value of the parameter.
	//
	// example:
	//
	// 15
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// Indicates whether the parameter immediately takes effect without requiring a restart.
	//
	// example:
	//
	// true
	IsDynamic *int32 `json:"IsDynamic,omitempty" xml:"IsDynamic,omitempty"`
	// Indicates whether the parameter is modifiable.
	//
	// example:
	//
	// true
	IsUserModifiable *int32 `json:"IsUserModifiable,omitempty" xml:"IsUserModifiable,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// doris_scanner_thread_pool_thread_num
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value range of the parameter.
	//
	// example:
	//
	// [0-20000]
	Optional *string `json:"Optional,omitempty" xml:"Optional,omitempty"`
	// The category of the parameter.
	ParamCategory *string `json:"ParamCategory,omitempty" xml:"ParamCategory,omitempty"`
	// The current value of the parameter.
	//
	// example:
	//
	// 10
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClusterConfigResponseBodyDataParams) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterConfigResponseBodyDataParams) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) GetComment() *string {
	return s.Comment
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) GetIsDynamic() *int32 {
	return s.IsDynamic
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) GetIsUserModifiable() *int32 {
	return s.IsUserModifiable
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) GetName() *string {
	return s.Name
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) GetOptional() *string {
	return s.Optional
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) GetParamCategory() *string {
	return s.ParamCategory
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) GetValue() *string {
	return s.Value
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) SetComment(v string) *DescribeDBClusterConfigResponseBodyDataParams {
	s.Comment = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) SetDefaultValue(v string) *DescribeDBClusterConfigResponseBodyDataParams {
	s.DefaultValue = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) SetIsDynamic(v int32) *DescribeDBClusterConfigResponseBodyDataParams {
	s.IsDynamic = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) SetIsUserModifiable(v int32) *DescribeDBClusterConfigResponseBodyDataParams {
	s.IsUserModifiable = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) SetName(v string) *DescribeDBClusterConfigResponseBodyDataParams {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) SetOptional(v string) *DescribeDBClusterConfigResponseBodyDataParams {
	s.Optional = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) SetParamCategory(v string) *DescribeDBClusterConfigResponseBodyDataParams {
	s.ParamCategory = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) SetValue(v string) *DescribeDBClusterConfigResponseBodyDataParams {
	s.Value = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) Validate() error {
	return dara.Validate(s)
}
