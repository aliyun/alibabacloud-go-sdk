// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameterGroup(v *DescribeParameterGroupResponseBodyParameterGroup) *DescribeParameterGroupResponseBody
	GetParameterGroup() *DescribeParameterGroupResponseBodyParameterGroup
	SetRequestId(v string) *DescribeParameterGroupResponseBody
	GetRequestId() *string
}

type DescribeParameterGroupResponseBody struct {
	// The information about the parameter template.
	ParameterGroup *DescribeParameterGroupResponseBodyParameterGroup `json:"ParameterGroup,omitempty" xml:"ParameterGroup,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A501A191-BD70-5E50-98A9-C2A486A82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeParameterGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupResponseBody) GetParameterGroup() *DescribeParameterGroupResponseBodyParameterGroup {
	return s.ParameterGroup
}

func (s *DescribeParameterGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeParameterGroupResponseBody) SetParameterGroup(v *DescribeParameterGroupResponseBodyParameterGroup) *DescribeParameterGroupResponseBody {
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
	// The service category. Valid values:
	//
	// 	- **0**: Redis Open-Source Edition
	//
	// 	- **1**: Tair (Enterprise Edition)
	//
	// example:
	//
	// 1
	Category *int64 `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the parameter template was created.
	//
	// example:
	//
	// 2023-04-18 16:32:45
	Created *string `json:"Created,omitempty" xml:"Created,omitempty"`
	// The engine type. Valid values:
	//
	// 	- *redis*: Redis or Tair DRAM-based instance
	//
	// 	- *tair_pena*: Tair persistent memory-optimized instance
	//
	// 	- *tair_pdb*: Tair ESSD-based instance
	//
	// example:
	//
	// redis
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The compatible engine version.
	//
	// example:
	//
	// 5.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The time when the parameter template was last modified.
	//
	// example:
	//
	// 2023-04-18 16:32:45
	Modified *string `json:"Modified,omitempty" xml:"Modified,omitempty"`
	// The parameter template ID, which is globally unique.
	//
	// example:
	//
	// sys-001*****
	ParamGroupId *string `json:"ParamGroupId,omitempty" xml:"ParamGroupId,omitempty"`
	// The parameter details of the parameter template.
	ParamGroupsDetails []*DescribeParameterGroupResponseBodyParameterGroupParamGroupsDetails `json:"ParamGroupsDetails,omitempty" xml:"ParamGroupsDetails,omitempty" type:"Repeated"`
	// The description of the parameter template.
	//
	// example:
	//
	// test
	ParameterGroupDesc *string `json:"ParameterGroupDesc,omitempty" xml:"ParameterGroupDesc,omitempty"`
	// The name of the parameter template.
	//
	// example:
	//
	// testGroupName
	ParameterGroupName *string `json:"ParameterGroupName,omitempty" xml:"ParameterGroupName,omitempty"`
}

func (s DescribeParameterGroupResponseBodyParameterGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupResponseBodyParameterGroup) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) GetCategory() *int64 {
	return s.Category
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) GetCreated() *string {
	return s.Created
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) GetEngine() *string {
	return s.Engine
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) GetModified() *string {
	return s.Modified
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) GetParamGroupId() *string {
	return s.ParamGroupId
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) GetParamGroupsDetails() []*DescribeParameterGroupResponseBodyParameterGroupParamGroupsDetails {
	return s.ParamGroupsDetails
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) GetParameterGroupDesc() *string {
	return s.ParameterGroupDesc
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) GetParameterGroupName() *string {
	return s.ParameterGroupName
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetCategory(v int64) *DescribeParameterGroupResponseBodyParameterGroup {
	s.Category = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetCreated(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.Created = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetEngine(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.Engine = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetEngineVersion(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.EngineVersion = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetModified(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.Modified = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetParamGroupId(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.ParamGroupId = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetParamGroupsDetails(v []*DescribeParameterGroupResponseBodyParameterGroupParamGroupsDetails) *DescribeParameterGroupResponseBodyParameterGroup {
	s.ParamGroupsDetails = v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetParameterGroupDesc(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.ParameterGroupDesc = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetParameterGroupName(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.ParameterGroupName = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) Validate() error {
	return dara.Validate(s)
}

type DescribeParameterGroupResponseBodyParameterGroupParamGroupsDetails struct {
	// The name of the parameter.
	//
	// example:
	//
	// timeout
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// 1000
	ParamValue *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
}

func (s DescribeParameterGroupResponseBodyParameterGroupParamGroupsDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupResponseBodyParameterGroupParamGroupsDetails) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupResponseBodyParameterGroupParamGroupsDetails) GetParamName() *string {
	return s.ParamName
}

func (s *DescribeParameterGroupResponseBodyParameterGroupParamGroupsDetails) GetParamValue() *string {
	return s.ParamValue
}

func (s *DescribeParameterGroupResponseBodyParameterGroupParamGroupsDetails) SetParamName(v string) *DescribeParameterGroupResponseBodyParameterGroupParamGroupsDetails {
	s.ParamName = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroupParamGroupsDetails) SetParamValue(v string) *DescribeParameterGroupResponseBodyParameterGroupParamGroupsDetails {
	s.ParamValue = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroupParamGroupsDetails) Validate() error {
	return dara.Validate(s)
}
