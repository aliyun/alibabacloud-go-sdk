// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLakebaseMountParameterGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeLakebaseMountParameterGroupResponseBody
	GetCreateTime() *string
	SetIsApply(v int32) *DescribeLakebaseMountParameterGroupResponseBody
	GetIsApply() *int32
	SetModifyTime(v string) *DescribeLakebaseMountParameterGroupResponseBody
	GetModifyTime() *string
	SetParameterDetail(v []*DescribeLakebaseMountParameterGroupResponseBodyParameterDetail) *DescribeLakebaseMountParameterGroupResponseBody
	GetParameterDetail() []*DescribeLakebaseMountParameterGroupResponseBodyParameterDetail
	SetParameterGroupDesc(v string) *DescribeLakebaseMountParameterGroupResponseBody
	GetParameterGroupDesc() *string
	SetParameterGroupId(v string) *DescribeLakebaseMountParameterGroupResponseBody
	GetParameterGroupId() *string
	SetParameterGroupName(v string) *DescribeLakebaseMountParameterGroupResponseBody
	GetParameterGroupName() *string
	SetParameterGroupType(v int32) *DescribeLakebaseMountParameterGroupResponseBody
	GetParameterGroupType() *int32
	SetRequestId(v string) *DescribeLakebaseMountParameterGroupResponseBody
	GetRequestId() *string
}

type DescribeLakebaseMountParameterGroupResponseBody struct {
	// The creation time in UTC.
	//
	// example:
	//
	// 2026-08-01T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the parameter template is applied. Valid values: 0: No. 1: Yes.
	//
	// example:
	//
	// 0
	IsApply *int32 `json:"IsApply,omitempty" xml:"IsApply,omitempty"`
	// The modification time in UTC.
	//
	// example:
	//
	// 2026-08-10T12:00:00Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The details of the parameters in the template. An empty array is returned if no data is available.
	ParameterDetail []*DescribeLakebaseMountParameterGroupResponseBodyParameterDetail `json:"ParameterDetail,omitempty" xml:"ParameterDetail,omitempty" type:"Repeated"`
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
	// The parameter templatetype. Valid values: 0: System. 1: User.
	//
	// example:
	//
	// 1
	ParameterGroupType *int32 `json:"ParameterGroupType,omitempty" xml:"ParameterGroupType,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F45FFACC-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLakebaseMountParameterGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakebaseMountParameterGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLakebaseMountParameterGroupResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLakebaseMountParameterGroupResponseBody) GetIsApply() *int32 {
	return s.IsApply
}

func (s *DescribeLakebaseMountParameterGroupResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeLakebaseMountParameterGroupResponseBody) GetParameterDetail() []*DescribeLakebaseMountParameterGroupResponseBodyParameterDetail {
	return s.ParameterDetail
}

func (s *DescribeLakebaseMountParameterGroupResponseBody) GetParameterGroupDesc() *string {
	return s.ParameterGroupDesc
}

func (s *DescribeLakebaseMountParameterGroupResponseBody) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *DescribeLakebaseMountParameterGroupResponseBody) GetParameterGroupName() *string {
	return s.ParameterGroupName
}

func (s *DescribeLakebaseMountParameterGroupResponseBody) GetParameterGroupType() *int32 {
	return s.ParameterGroupType
}

func (s *DescribeLakebaseMountParameterGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLakebaseMountParameterGroupResponseBody) SetCreateTime(v string) *DescribeLakebaseMountParameterGroupResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupResponseBody) SetIsApply(v int32) *DescribeLakebaseMountParameterGroupResponseBody {
	s.IsApply = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupResponseBody) SetModifyTime(v string) *DescribeLakebaseMountParameterGroupResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupResponseBody) SetParameterDetail(v []*DescribeLakebaseMountParameterGroupResponseBodyParameterDetail) *DescribeLakebaseMountParameterGroupResponseBody {
	s.ParameterDetail = v
	return s
}

func (s *DescribeLakebaseMountParameterGroupResponseBody) SetParameterGroupDesc(v string) *DescribeLakebaseMountParameterGroupResponseBody {
	s.ParameterGroupDesc = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupResponseBody) SetParameterGroupId(v string) *DescribeLakebaseMountParameterGroupResponseBody {
	s.ParameterGroupId = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupResponseBody) SetParameterGroupName(v string) *DescribeLakebaseMountParameterGroupResponseBody {
	s.ParameterGroupName = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupResponseBody) SetParameterGroupType(v int32) *DescribeLakebaseMountParameterGroupResponseBody {
	s.ParameterGroupType = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupResponseBody) SetRequestId(v string) *DescribeLakebaseMountParameterGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupResponseBody) Validate() error {
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

type DescribeLakebaseMountParameterGroupResponseBodyParameterDetail struct {
	// The parameter name.
	//
	// example:
	//
	// auto_increment_offset
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// 1
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeLakebaseMountParameterGroupResponseBodyParameterDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakebaseMountParameterGroupResponseBodyParameterDetail) GoString() string {
	return s.String()
}

func (s *DescribeLakebaseMountParameterGroupResponseBodyParameterDetail) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeLakebaseMountParameterGroupResponseBodyParameterDetail) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *DescribeLakebaseMountParameterGroupResponseBodyParameterDetail) SetParameterName(v string) *DescribeLakebaseMountParameterGroupResponseBodyParameterDetail {
	s.ParameterName = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupResponseBodyParameterDetail) SetParameterValue(v string) *DescribeLakebaseMountParameterGroupResponseBodyParameterDetail {
	s.ParameterValue = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupResponseBodyParameterDetail) Validate() error {
	return dara.Validate(s)
}
