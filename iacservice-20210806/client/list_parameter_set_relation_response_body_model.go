// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListParameterSetRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameterSets(v []*ListParameterSetRelationResponseBodyParameterSets) *ListParameterSetRelationResponseBody
	GetParameterSets() []*ListParameterSetRelationResponseBodyParameterSets
	SetRequestId(v string) *ListParameterSetRelationResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListParameterSetRelationResponseBody
	GetTotalCount() *int32
}

type ListParameterSetRelationResponseBody struct {
	ParameterSets []*ListParameterSetRelationResponseBodyParameterSets `json:"parameterSets,omitempty" xml:"parameterSets,omitempty" type:"Repeated"`
	// example:
	//
	// 2F24E990-E8D3-5C18-ABEA-C7A3F1831C57
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListParameterSetRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListParameterSetRelationResponseBody) GoString() string {
	return s.String()
}

func (s *ListParameterSetRelationResponseBody) GetParameterSets() []*ListParameterSetRelationResponseBodyParameterSets {
	return s.ParameterSets
}

func (s *ListParameterSetRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListParameterSetRelationResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListParameterSetRelationResponseBody) SetParameterSets(v []*ListParameterSetRelationResponseBodyParameterSets) *ListParameterSetRelationResponseBody {
	s.ParameterSets = v
	return s
}

func (s *ListParameterSetRelationResponseBody) SetRequestId(v string) *ListParameterSetRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListParameterSetRelationResponseBody) SetTotalCount(v int32) *ListParameterSetRelationResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListParameterSetRelationResponseBody) Validate() error {
	if s.ParameterSets != nil {
		for _, item := range s.ParameterSets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListParameterSetRelationResponseBodyParameterSets struct {
	// example:
	//
	// 2022-05-14T10:05:19Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 123111
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// alb_enable_ipv6_4.2
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// pts-kw1b11jlssrabb638ptums
	ParameterSetId *string            `json:"parameterSetId,omitempty" xml:"parameterSetId,omitempty"`
	Parameters     map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
}

func (s ListParameterSetRelationResponseBodyParameterSets) String() string {
	return dara.Prettify(s)
}

func (s ListParameterSetRelationResponseBodyParameterSets) GoString() string {
	return s.String()
}

func (s *ListParameterSetRelationResponseBodyParameterSets) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListParameterSetRelationResponseBodyParameterSets) GetDescription() *string {
	return s.Description
}

func (s *ListParameterSetRelationResponseBodyParameterSets) GetName() *string {
	return s.Name
}

func (s *ListParameterSetRelationResponseBodyParameterSets) GetParameterSetId() *string {
	return s.ParameterSetId
}

func (s *ListParameterSetRelationResponseBodyParameterSets) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *ListParameterSetRelationResponseBodyParameterSets) SetCreateTime(v string) *ListParameterSetRelationResponseBodyParameterSets {
	s.CreateTime = &v
	return s
}

func (s *ListParameterSetRelationResponseBodyParameterSets) SetDescription(v string) *ListParameterSetRelationResponseBodyParameterSets {
	s.Description = &v
	return s
}

func (s *ListParameterSetRelationResponseBodyParameterSets) SetName(v string) *ListParameterSetRelationResponseBodyParameterSets {
	s.Name = &v
	return s
}

func (s *ListParameterSetRelationResponseBodyParameterSets) SetParameterSetId(v string) *ListParameterSetRelationResponseBodyParameterSets {
	s.ParameterSetId = &v
	return s
}

func (s *ListParameterSetRelationResponseBodyParameterSets) SetParameters(v map[string]*string) *ListParameterSetRelationResponseBodyParameterSets {
	s.Parameters = v
	return s
}

func (s *ListParameterSetRelationResponseBodyParameterSets) Validate() error {
	return dara.Validate(s)
}
