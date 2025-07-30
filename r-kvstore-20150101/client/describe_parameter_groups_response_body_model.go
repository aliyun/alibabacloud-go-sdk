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
	// The list of parameter templates.
	ParameterGroups []*DescribeParameterGroupsResponseBodyParameterGroups `json:"ParameterGroups,omitempty" xml:"ParameterGroups,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 686BB8A6-BBA5-47E5-8A75-D2ADE433****
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
	// The service category. Valid values:
	//
	// 	- **0**: Redis Open-Source Edition
	//
	// 	- **1**: Tair (Enterprise Edition)
	//
	// example:
	//
	// 0
	Category *int64 `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the parameter template was created.
	//
	// example:
	//
	// 2023-04-18 16:32:45
	Created *string `json:"Created,omitempty" xml:"Created,omitempty"`
	// The engine type. Valid values:
	//
	// 	- **redis**: Redis Open-Source Edition or Tair (In-Memory)
	//
	// 	- **tair_pena**: Tair (On NVM)
	//
	// 	- **tair_pdb**: Tair (On Disk)
	//
	// example:
	//
	// redis
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The compatible engine version.
	//
	// example:
	//
	// 5
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The time when the parameter template was last modified.
	//
	// example:
	//
	// 2023-04-18 16:32:45
	Modified *string `json:"Modified,omitempty" xml:"Modified,omitempty"`
	// The parameter template ID.
	//
	// example:
	//
	// test01
	ParamGroupId *string `json:"ParamGroupId,omitempty" xml:"ParamGroupId,omitempty"`
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

func (s DescribeParameterGroupsResponseBodyParameterGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupsResponseBodyParameterGroups) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) GetCategory() *int64 {
	return s.Category
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) GetCreated() *string {
	return s.Created
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) GetEngine() *string {
	return s.Engine
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) GetModified() *string {
	return s.Modified
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) GetParamGroupId() *string {
	return s.ParamGroupId
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) GetParameterGroupDesc() *string {
	return s.ParameterGroupDesc
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) GetParameterGroupName() *string {
	return s.ParameterGroupName
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetCategory(v int64) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.Category = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetCreated(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.Created = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetEngine(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.Engine = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetEngineVersion(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.EngineVersion = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetModified(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.Modified = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetParamGroupId(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.ParamGroupId = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetParameterGroupDesc(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.ParameterGroupDesc = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetParameterGroupName(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.ParameterGroupName = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) Validate() error {
	return dara.Validate(s)
}
