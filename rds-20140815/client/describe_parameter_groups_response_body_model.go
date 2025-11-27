// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameterGroups(v *DescribeParameterGroupsResponseBodyParameterGroups) *DescribeParameterGroupsResponseBody
	GetParameterGroups() *DescribeParameterGroupsResponseBodyParameterGroups
	SetRequestId(v string) *DescribeParameterGroupsResponseBody
	GetRequestId() *string
	SetSignalForOptimizeParams(v bool) *DescribeParameterGroupsResponseBody
	GetSignalForOptimizeParams() *bool
}

type DescribeParameterGroupsResponseBody struct {
	// The details of the parameter templates.
	ParameterGroups *DescribeParameterGroupsResponseBodyParameterGroups `json:"ParameterGroups,omitempty" xml:"ParameterGroups,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D4A23265-C5B6-42E1-98A0-EFA1EB42E723
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether parameter templates exist in the specified region. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// 	Notice: This parameter is deprecated.
	//
	// example:
	//
	// false
	SignalForOptimizeParams *bool `json:"SignalForOptimizeParams,omitempty" xml:"SignalForOptimizeParams,omitempty"`
}

func (s DescribeParameterGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupsResponseBody) GetParameterGroups() *DescribeParameterGroupsResponseBodyParameterGroups {
	return s.ParameterGroups
}

func (s *DescribeParameterGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeParameterGroupsResponseBody) GetSignalForOptimizeParams() *bool {
	return s.SignalForOptimizeParams
}

func (s *DescribeParameterGroupsResponseBody) SetParameterGroups(v *DescribeParameterGroupsResponseBodyParameterGroups) *DescribeParameterGroupsResponseBody {
	s.ParameterGroups = v
	return s
}

func (s *DescribeParameterGroupsResponseBody) SetRequestId(v string) *DescribeParameterGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParameterGroupsResponseBody) SetSignalForOptimizeParams(v bool) *DescribeParameterGroupsResponseBody {
	s.SignalForOptimizeParams = &v
	return s
}

func (s *DescribeParameterGroupsResponseBody) Validate() error {
	if s.ParameterGroups != nil {
		if err := s.ParameterGroups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeParameterGroupsResponseBodyParameterGroups struct {
	ParameterGroup []*DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup `json:"ParameterGroup,omitempty" xml:"ParameterGroup,omitempty" type:"Repeated"`
}

func (s DescribeParameterGroupsResponseBodyParameterGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupsResponseBodyParameterGroups) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) GetParameterGroup() []*DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup {
	return s.ParameterGroup
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetParameterGroup(v []*DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.ParameterGroup = v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) Validate() error {
	if s.ParameterGroup != nil {
		for _, item := range s.ParameterGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup struct {
	// The time when the parameter template was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-11-21T01:48:39Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The database engine of the instance.
	//
	// example:
	//
	// mysql
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// Indicates whether the restart of an instance is required for the parameter template to take effect. Valid values:
	//
	// 	- 0: A restart is not required.
	//
	// 	- 1: A restart is required.
	//
	// example:
	//
	// 1
	ForceRestart *int32 `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	// The number of parameters in the parameter template.
	//
	// example:
	//
	// 2
	ParamCounts *int32 `json:"ParamCounts,omitempty" xml:"ParamCounts,omitempty"`
	// The type of the parameter template. Valid values:
	//
	// 	- 0: the default parameter template.
	//
	// 	- 1: a custom parameter template.
	//
	// 	- 2: an automatic backup parameter template. After you apply this type of template, the system automatically backs up the original parameter settings and saves the backup as a template.
	//
	// example:
	//
	// 1
	ParameterGroupDesc *string `json:"ParameterGroupDesc,omitempty" xml:"ParameterGroupDesc,omitempty"`
	// The ID of the parameter template.
	//
	// example:
	//
	// rpg-****
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The name of the parameter template.
	//
	// example:
	//
	// test1234
	ParameterGroupName *string `json:"ParameterGroupName,omitempty" xml:"ParameterGroupName,omitempty"`
	// The type of the parameter template. Valid values:
	//
	// 	- 0: the default parameter template.
	//
	// 	- 1: a custom parameter template.
	//
	// 	- 2: an automatic backup parameter template. After you apply this type of template, the system automatically backs up the original parameter settings and saves the backup as a template.
	//
	// example:
	//
	// 1
	ParameterGroupType *int32 `json:"ParameterGroupType,omitempty" xml:"ParameterGroupType,omitempty"`
	// The time when the parameter template was last updated. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-11-21T02:21:35Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) GetEngine() *string {
	return s.Engine
}

func (s *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) GetForceRestart() *int32 {
	return s.ForceRestart
}

func (s *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) GetParamCounts() *int32 {
	return s.ParamCounts
}

func (s *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) GetParameterGroupDesc() *string {
	return s.ParameterGroupDesc
}

func (s *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) GetParameterGroupName() *string {
	return s.ParameterGroupName
}

func (s *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) GetParameterGroupType() *int32 {
	return s.ParameterGroupType
}

func (s *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) SetCreateTime(v string) *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup {
	s.CreateTime = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) SetEngine(v string) *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup {
	s.Engine = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) SetEngineVersion(v string) *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup {
	s.EngineVersion = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) SetForceRestart(v int32) *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup {
	s.ForceRestart = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) SetParamCounts(v int32) *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup {
	s.ParamCounts = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) SetParameterGroupDesc(v string) *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup {
	s.ParameterGroupDesc = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) SetParameterGroupId(v string) *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup {
	s.ParameterGroupId = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) SetParameterGroupName(v string) *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup {
	s.ParameterGroupName = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) SetParameterGroupType(v int32) *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup {
	s.ParameterGroupType = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) SetUpdateTime(v string) *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup {
	s.UpdateTime = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroupsParameterGroup) Validate() error {
	return dara.Validate(s)
}
