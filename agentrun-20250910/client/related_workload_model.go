// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRelatedWorkload interface {
	dara.Model
	String() string
	GoString() string
	SetWorkloadId(v string) *RelatedWorkload
	GetWorkloadId() *string
	SetWorkloadName(v string) *RelatedWorkload
	GetWorkloadName() *string
	SetWorkloadType(v string) *RelatedWorkload
	GetWorkloadType() *string
}

type RelatedWorkload struct {
	WorkloadId   *string `json:"workloadId,omitempty" xml:"workloadId,omitempty"`
	WorkloadName *string `json:"workloadName,omitempty" xml:"workloadName,omitempty"`
	WorkloadType *string `json:"workloadType,omitempty" xml:"workloadType,omitempty"`
}

func (s RelatedWorkload) String() string {
	return dara.Prettify(s)
}

func (s RelatedWorkload) GoString() string {
	return s.String()
}

func (s *RelatedWorkload) GetWorkloadId() *string {
	return s.WorkloadId
}

func (s *RelatedWorkload) GetWorkloadName() *string {
	return s.WorkloadName
}

func (s *RelatedWorkload) GetWorkloadType() *string {
	return s.WorkloadType
}

func (s *RelatedWorkload) SetWorkloadId(v string) *RelatedWorkload {
	s.WorkloadId = &v
	return s
}

func (s *RelatedWorkload) SetWorkloadName(v string) *RelatedWorkload {
	s.WorkloadName = &v
	return s
}

func (s *RelatedWorkload) SetWorkloadType(v string) *RelatedWorkload {
	s.WorkloadType = &v
	return s
}

func (s *RelatedWorkload) Validate() error {
	return dara.Validate(s)
}
