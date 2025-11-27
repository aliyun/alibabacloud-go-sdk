// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParamGroup(v *DescribeParameterGroupResponseBodyParamGroup) *DescribeParameterGroupResponseBody
	GetParamGroup() *DescribeParameterGroupResponseBodyParamGroup
	SetRelatedCustinsInfo(v *DescribeParameterGroupResponseBodyRelatedCustinsInfo) *DescribeParameterGroupResponseBody
	GetRelatedCustinsInfo() *DescribeParameterGroupResponseBodyRelatedCustinsInfo
	SetRequestId(v string) *DescribeParameterGroupResponseBody
	GetRequestId() *string
}

type DescribeParameterGroupResponseBody struct {
	// The information about the parameter template.
	ParamGroup *DescribeParameterGroupResponseBodyParamGroup `json:"ParamGroup,omitempty" xml:"ParamGroup,omitempty" type:"Struct"`
	// The information about the instance to which the parameter template is applied.
	//
	// >  This parameter is available only for ApsaraDB RDS for PostgreSQL instances.
	RelatedCustinsInfo *DescribeParameterGroupResponseBodyRelatedCustinsInfo `json:"RelatedCustinsInfo,omitempty" xml:"RelatedCustinsInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 498AE8CA-8C81-4A01-AF37-2B902014ED30
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeParameterGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupResponseBody) GetParamGroup() *DescribeParameterGroupResponseBodyParamGroup {
	return s.ParamGroup
}

func (s *DescribeParameterGroupResponseBody) GetRelatedCustinsInfo() *DescribeParameterGroupResponseBodyRelatedCustinsInfo {
	return s.RelatedCustinsInfo
}

func (s *DescribeParameterGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeParameterGroupResponseBody) SetParamGroup(v *DescribeParameterGroupResponseBodyParamGroup) *DescribeParameterGroupResponseBody {
	s.ParamGroup = v
	return s
}

func (s *DescribeParameterGroupResponseBody) SetRelatedCustinsInfo(v *DescribeParameterGroupResponseBodyRelatedCustinsInfo) *DescribeParameterGroupResponseBody {
	s.RelatedCustinsInfo = v
	return s
}

func (s *DescribeParameterGroupResponseBody) SetRequestId(v string) *DescribeParameterGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParameterGroupResponseBody) Validate() error {
	if s.ParamGroup != nil {
		if err := s.ParamGroup.Validate(); err != nil {
			return err
		}
	}
	if s.RelatedCustinsInfo != nil {
		if err := s.RelatedCustinsInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeParameterGroupResponseBodyParamGroup struct {
	ParameterGroup []*DescribeParameterGroupResponseBodyParamGroupParameterGroup `json:"ParameterGroup,omitempty" xml:"ParameterGroup,omitempty" type:"Repeated"`
}

func (s DescribeParameterGroupResponseBodyParamGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupResponseBodyParamGroup) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupResponseBodyParamGroup) GetParameterGroup() []*DescribeParameterGroupResponseBodyParamGroupParameterGroup {
	return s.ParameterGroup
}

func (s *DescribeParameterGroupResponseBodyParamGroup) SetParameterGroup(v []*DescribeParameterGroupResponseBodyParamGroupParameterGroup) *DescribeParameterGroupResponseBodyParamGroup {
	s.ParameterGroup = v
	return s
}

func (s *DescribeParameterGroupResponseBodyParamGroup) Validate() error {
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

type DescribeParameterGroupResponseBodyParamGroupParameterGroup struct {
	// The time when the parameter template was created.
	//
	// example:
	//
	// 2019-10-22T06:02:53Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The database engine of the instance.
	//
	// example:
	//
	// mysql
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version of the instance.
	//
	// example:
	//
	// 5.6
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// Indicates whether the restart of an instance is required for the parameter template to take effect. Valid values:
	//
	// 	- **0**: A restart is not required.
	//
	// 	- **1**: A restart is required.
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
	// The details of the parameters.
	ParamDetail *DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetail `json:"ParamDetail,omitempty" xml:"ParamDetail,omitempty" type:"Struct"`
	// The description of the parameter template.
	//
	// example:
	//
	// testGroup1
	ParameterGroupDesc *string `json:"ParameterGroupDesc,omitempty" xml:"ParameterGroupDesc,omitempty"`
	// The ID of the parameter template.
	//
	// example:
	//
	// rpg-dp****
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The name of the parameter template.
	//
	// example:
	//
	// test123456
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
	ParameterGroupType *int32 `json:"ParameterGroupType,omitempty" xml:"ParameterGroupType,omitempty"`
	// The time when the parameter template was last updated. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-10-22T06:07:54Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeParameterGroupResponseBodyParamGroupParameterGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupResponseBodyParamGroupParameterGroup) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) GetEngine() *string {
	return s.Engine
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) GetForceRestart() *int32 {
	return s.ForceRestart
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) GetParamCounts() *int32 {
	return s.ParamCounts
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) GetParamDetail() *DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetail {
	return s.ParamDetail
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) GetParameterGroupDesc() *string {
	return s.ParameterGroupDesc
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) GetParameterGroupName() *string {
	return s.ParameterGroupName
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) GetParameterGroupType() *int32 {
	return s.ParameterGroupType
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) SetCreateTime(v string) *DescribeParameterGroupResponseBodyParamGroupParameterGroup {
	s.CreateTime = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) SetEngine(v string) *DescribeParameterGroupResponseBodyParamGroupParameterGroup {
	s.Engine = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) SetEngineVersion(v string) *DescribeParameterGroupResponseBodyParamGroupParameterGroup {
	s.EngineVersion = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) SetForceRestart(v int32) *DescribeParameterGroupResponseBodyParamGroupParameterGroup {
	s.ForceRestart = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) SetParamCounts(v int32) *DescribeParameterGroupResponseBodyParamGroupParameterGroup {
	s.ParamCounts = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) SetParamDetail(v *DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetail) *DescribeParameterGroupResponseBodyParamGroupParameterGroup {
	s.ParamDetail = v
	return s
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) SetParameterGroupDesc(v string) *DescribeParameterGroupResponseBodyParamGroupParameterGroup {
	s.ParameterGroupDesc = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) SetParameterGroupId(v string) *DescribeParameterGroupResponseBodyParamGroupParameterGroup {
	s.ParameterGroupId = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) SetParameterGroupName(v string) *DescribeParameterGroupResponseBodyParamGroupParameterGroup {
	s.ParameterGroupName = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) SetParameterGroupType(v int32) *DescribeParameterGroupResponseBodyParamGroupParameterGroup {
	s.ParameterGroupType = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) SetUpdateTime(v string) *DescribeParameterGroupResponseBodyParamGroupParameterGroup {
	s.UpdateTime = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroup) Validate() error {
	if s.ParamDetail != nil {
		if err := s.ParamDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetail struct {
	ParameterDetail []*DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetailParameterDetail `json:"ParameterDetail,omitempty" xml:"ParameterDetail,omitempty" type:"Repeated"`
}

func (s DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetail) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetail) GetParameterDetail() []*DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetailParameterDetail {
	return s.ParameterDetail
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetail) SetParameterDetail(v []*DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetailParameterDetail) *DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetail {
	s.ParameterDetail = v
	return s
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetail) Validate() error {
	if s.ParameterDetail != nil {
		for _, item := range s.ParameterDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetailParameterDetail struct {
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
	// 2000
	ParamValue *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
}

func (s DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetailParameterDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetailParameterDetail) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetailParameterDetail) GetParamName() *string {
	return s.ParamName
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetailParameterDetail) GetParamValue() *string {
	return s.ParamValue
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetailParameterDetail) SetParamName(v string) *DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetailParameterDetail {
	s.ParamName = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetailParameterDetail) SetParamValue(v string) *DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetailParameterDetail {
	s.ParamValue = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParamGroupParameterGroupParamDetailParameterDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeParameterGroupResponseBodyRelatedCustinsInfo struct {
	RelatedCustinsInfo []*DescribeParameterGroupResponseBodyRelatedCustinsInfoRelatedCustinsInfo `json:"RelatedCustinsInfo,omitempty" xml:"RelatedCustinsInfo,omitempty" type:"Repeated"`
}

func (s DescribeParameterGroupResponseBodyRelatedCustinsInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupResponseBodyRelatedCustinsInfo) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupResponseBodyRelatedCustinsInfo) GetRelatedCustinsInfo() []*DescribeParameterGroupResponseBodyRelatedCustinsInfoRelatedCustinsInfo {
	return s.RelatedCustinsInfo
}

func (s *DescribeParameterGroupResponseBodyRelatedCustinsInfo) SetRelatedCustinsInfo(v []*DescribeParameterGroupResponseBodyRelatedCustinsInfoRelatedCustinsInfo) *DescribeParameterGroupResponseBodyRelatedCustinsInfo {
	s.RelatedCustinsInfo = v
	return s
}

func (s *DescribeParameterGroupResponseBodyRelatedCustinsInfo) Validate() error {
	if s.RelatedCustinsInfo != nil {
		for _, item := range s.RelatedCustinsInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeParameterGroupResponseBodyRelatedCustinsInfoRelatedCustinsInfo struct {
	// The time when the parameter template was applied.
	//
	// example:
	//
	// 2022-10-17T03:19:02Z
	AppliedTime *string `json:"AppliedTime,omitempty" xml:"AppliedTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-bp170****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
}

func (s DescribeParameterGroupResponseBodyRelatedCustinsInfoRelatedCustinsInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupResponseBodyRelatedCustinsInfoRelatedCustinsInfo) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupResponseBodyRelatedCustinsInfoRelatedCustinsInfo) GetAppliedTime() *string {
	return s.AppliedTime
}

func (s *DescribeParameterGroupResponseBodyRelatedCustinsInfoRelatedCustinsInfo) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeParameterGroupResponseBodyRelatedCustinsInfoRelatedCustinsInfo) SetAppliedTime(v string) *DescribeParameterGroupResponseBodyRelatedCustinsInfoRelatedCustinsInfo {
	s.AppliedTime = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyRelatedCustinsInfoRelatedCustinsInfo) SetDBInstanceName(v string) *DescribeParameterGroupResponseBodyRelatedCustinsInfoRelatedCustinsInfo {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyRelatedCustinsInfoRelatedCustinsInfo) Validate() error {
	return dara.Validate(s)
}
