// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLakebaseMountParameterGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameterGroups(v []*DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups) *DescribeLakebaseMountParameterGroupsResponseBody
	GetParameterGroups() []*DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups
	SetRequestId(v string) *DescribeLakebaseMountParameterGroupsResponseBody
	GetRequestId() *string
}

type DescribeLakebaseMountParameterGroupsResponseBody struct {
	// The parameter templates.
	ParameterGroups []*DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups `json:"ParameterGroups,omitempty" xml:"ParameterGroups,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F45FFACC-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLakebaseMountParameterGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakebaseMountParameterGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLakebaseMountParameterGroupsResponseBody) GetParameterGroups() []*DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups {
	return s.ParameterGroups
}

func (s *DescribeLakebaseMountParameterGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLakebaseMountParameterGroupsResponseBody) SetParameterGroups(v []*DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups) *DescribeLakebaseMountParameterGroupsResponseBody {
	s.ParameterGroups = v
	return s
}

func (s *DescribeLakebaseMountParameterGroupsResponseBody) SetRequestId(v string) *DescribeLakebaseMountParameterGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupsResponseBody) Validate() error {
	if s.ParameterGroups != nil {
		for _, item := range s.ParameterGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups struct {
	// The creation time in UTC.
	//
	// example:
	//
	// 2026-08-07T12:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the parameter template is applied. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	IsApply *string `json:"IsApply,omitempty" xml:"IsApply,omitempty"`
	// The modification time in UTC.
	//
	// example:
	//
	// 2026-08-10T12:00:00Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The number of parameters.
	//
	// example:
	//
	// 10
	ParameterCounts *int32 `json:"ParameterCounts,omitempty" xml:"ParameterCounts,omitempty"`
	// The description of the parameter template.
	//
	// example:
	//
	// Default template
	ParameterGroupDesc *string `json:"ParameterGroupDesc,omitempty" xml:"ParameterGroupDesc,omitempty"`
	// The parameter template ID.
	//
	// example:
	//
	// pcpg-lakebase-xxx
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The name of the parameter template.
	//
	// example:
	//
	// default-template
	ParameterGroupName *string `json:"ParameterGroupName,omitempty" xml:"ParameterGroupName,omitempty"`
	// The templatetype. Valid values: 0 (system) and 1 (user).
	//
	// example:
	//
	// 0
	ParameterGroupType *string `json:"ParameterGroupType,omitempty" xml:"ParameterGroupType,omitempty"`
}

func (s DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups) GoString() string {
	return s.String()
}

func (s *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups) GetIsApply() *string {
	return s.IsApply
}

func (s *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups) GetParameterCounts() *int32 {
	return s.ParameterCounts
}

func (s *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups) GetParameterGroupDesc() *string {
	return s.ParameterGroupDesc
}

func (s *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups) GetParameterGroupName() *string {
	return s.ParameterGroupName
}

func (s *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups) GetParameterGroupType() *string {
	return s.ParameterGroupType
}

func (s *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups) SetCreateTime(v string) *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups {
	s.CreateTime = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups) SetIsApply(v string) *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups {
	s.IsApply = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups) SetModifyTime(v string) *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups {
	s.ModifyTime = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups) SetParameterCounts(v int32) *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups {
	s.ParameterCounts = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups) SetParameterGroupDesc(v string) *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups {
	s.ParameterGroupDesc = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups) SetParameterGroupId(v string) *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups {
	s.ParameterGroupId = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups) SetParameterGroupName(v string) *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups {
	s.ParameterGroupName = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups) SetParameterGroupType(v string) *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups {
	s.ParameterGroupType = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupsResponseBodyParameterGroups) Validate() error {
	return dara.Validate(s)
}
