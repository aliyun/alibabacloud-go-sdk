// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameterGroups(v []*DescribeParameterGroupsResponseBodyParameterGroups) *DescribeParameterGroupsResponseBody
	GetParameterGroups() []*DescribeParameterGroupsResponseBodyParameterGroups
	SetRequestId(v string) *DescribeParameterGroupsResponseBody
	GetRequestId() *string
}

type DescribeParameterGroupsResponseBody struct {
	// The details of parameter templates.
	ParameterGroups []*DescribeParameterGroupsResponseBodyParameterGroups `json:"ParameterGroups,omitempty" xml:"ParameterGroups,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 944CED46-A6F7-40C6-B6DC-C6E5CC******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeParameterGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupsResponseBody) GetParameterGroups() []*DescribeParameterGroupsResponseBodyParameterGroups {
	return s.ParameterGroups
}

func (s *DescribeParameterGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeParameterGroupsResponseBody) SetParameterGroups(v []*DescribeParameterGroupsResponseBodyParameterGroups) *DescribeParameterGroupsResponseBody {
	s.ParameterGroups = v
	return s
}

func (s *DescribeParameterGroupsResponseBody) SetRequestId(v string) *DescribeParameterGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParameterGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeParameterGroupsResponseBodyParameterGroups struct {
	// The time when the parameter template was created. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-03-10T08:40:39Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The type of the engine.
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The version of the database engine
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
	ParameterCounts *int64 `json:"ParameterCounts,omitempty" xml:"ParameterCounts,omitempty"`
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

func (s DescribeParameterGroupsResponseBodyParameterGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupsResponseBodyParameterGroups) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) GetDBType() *string {
	return s.DBType
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) GetForceRestart() *string {
	return s.ForceRestart
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) GetParameterCounts() *int64 {
	return s.ParameterCounts
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) GetParameterGroupDesc() *string {
	return s.ParameterGroupDesc
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) GetParameterGroupName() *string {
	return s.ParameterGroupName
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) GetParameterGroupType() *string {
	return s.ParameterGroupType
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetCreateTime(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.CreateTime = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetDBType(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.DBType = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetDBVersion(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.DBVersion = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetForceRestart(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.ForceRestart = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetParameterCounts(v int64) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.ParameterCounts = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetParameterGroupDesc(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.ParameterGroupDesc = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetParameterGroupId(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.ParameterGroupId = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetParameterGroupName(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.ParameterGroupName = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetParameterGroupType(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.ParameterGroupType = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) Validate() error {
	return dara.Validate(s)
}
