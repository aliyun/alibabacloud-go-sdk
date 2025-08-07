// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameterGroup(v []*DescribeParameterGroupResponseBodyParameterGroup) *DescribeParameterGroupResponseBody
	GetParameterGroup() []*DescribeParameterGroupResponseBodyParameterGroup
	SetRequestId(v string) *DescribeParameterGroupResponseBody
	GetRequestId() *string
}

type DescribeParameterGroupResponseBody struct {
	// Details about the parameter templates.
	ParameterGroup []*DescribeParameterGroupResponseBodyParameterGroup `json:"ParameterGroup,omitempty" xml:"ParameterGroup,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// F1F16757-D31B-49CA-9BF4-305BAF******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeParameterGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupResponseBody) GetParameterGroup() []*DescribeParameterGroupResponseBodyParameterGroup {
	return s.ParameterGroup
}

func (s *DescribeParameterGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeParameterGroupResponseBody) SetParameterGroup(v []*DescribeParameterGroupResponseBodyParameterGroup) *DescribeParameterGroupResponseBody {
	s.ParameterGroup = v
	return s
}

func (s *DescribeParameterGroupResponseBody) SetRequestId(v string) *DescribeParameterGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParameterGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeParameterGroupResponseBodyParameterGroup struct {
	// The time when the parameter template was created. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-03-10T08:40:39Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The type of the database engine.
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The version of the database engine.
	//
	// example:
	//
	// 8.0
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// Indicates whether to restart the cluster when this parameter template is applied. Valid values:
	//
	// 	- **0**: A restart is not required.
	//
	// 	- **1**: A restart is required.
	//
	// example:
	//
	// 1
	ForceRestart *string `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	// The number of parameters in the parameter template.
	//
	// example:
	//
	// 2
	ParameterCounts *int32 `json:"ParameterCounts,omitempty" xml:"ParameterCounts,omitempty"`
	// Details about the parameters.
	ParameterDetail []*DescribeParameterGroupResponseBodyParameterGroupParameterDetail `json:"ParameterDetail,omitempty" xml:"ParameterDetail,omitempty" type:"Repeated"`
	// The description of the parameter template.
	//
	// example:
	//
	// testgroup
	ParameterGroupDesc *string `json:"ParameterGroupDesc,omitempty" xml:"ParameterGroupDesc,omitempty"`
	// The ID of the parameter template.
	//
	// example:
	//
	// pcpg-**************
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The name of the parameter template.
	//
	// example:
	//
	// test
	ParameterGroupName *string `json:"ParameterGroupName,omitempty" xml:"ParameterGroupName,omitempty"`
	// The type of the parameter template. Valid values:
	//
	// 	- **0**: the default parameter template.
	//
	// 	- **1**: a custom parameter template.
	//
	// 	- **2**: an automatic backup parameter template. After you apply this type of template, the system automatically backs up the original parameter settings and saves the backup as a template.
	//
	// example:
	//
	// 1
	ParameterGroupType *string `json:"ParameterGroupType,omitempty" xml:"ParameterGroupType,omitempty"`
}

func (s DescribeParameterGroupResponseBodyParameterGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupResponseBodyParameterGroup) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) GetDBType() *string {
	return s.DBType
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) GetForceRestart() *string {
	return s.ForceRestart
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) GetParameterCounts() *int32 {
	return s.ParameterCounts
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) GetParameterDetail() []*DescribeParameterGroupResponseBodyParameterGroupParameterDetail {
	return s.ParameterDetail
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) GetParameterGroupDesc() *string {
	return s.ParameterGroupDesc
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) GetParameterGroupName() *string {
	return s.ParameterGroupName
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) GetParameterGroupType() *string {
	return s.ParameterGroupType
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetCreateTime(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.CreateTime = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetDBType(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.DBType = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetDBVersion(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.DBVersion = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetForceRestart(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.ForceRestart = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetParameterCounts(v int32) *DescribeParameterGroupResponseBodyParameterGroup {
	s.ParameterCounts = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetParameterDetail(v []*DescribeParameterGroupResponseBodyParameterGroupParameterDetail) *DescribeParameterGroupResponseBodyParameterGroup {
	s.ParameterDetail = v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetParameterGroupDesc(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.ParameterGroupDesc = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetParameterGroupId(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.ParameterGroupId = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetParameterGroupName(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.ParameterGroupName = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetParameterGroupType(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.ParameterGroupType = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) Validate() error {
	return dara.Validate(s)
}

type DescribeParameterGroupResponseBodyParameterGroupParameterDetail struct {
	// The name of the parameter.
	//
	// example:
	//
	// back_log
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// 3000
	ParamValue *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
}

func (s DescribeParameterGroupResponseBodyParameterGroupParameterDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupResponseBodyParameterGroupParameterDetail) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupResponseBodyParameterGroupParameterDetail) GetParamName() *string {
	return s.ParamName
}

func (s *DescribeParameterGroupResponseBodyParameterGroupParameterDetail) GetParamValue() *string {
	return s.ParamValue
}

func (s *DescribeParameterGroupResponseBodyParameterGroupParameterDetail) SetParamName(v string) *DescribeParameterGroupResponseBodyParameterGroupParameterDetail {
	s.ParamName = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroupParameterDetail) SetParamValue(v string) *DescribeParameterGroupResponseBodyParameterGroupParameterDetail {
	s.ParamValue = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroupParameterDetail) Validate() error {
	return dara.Validate(s)
}
