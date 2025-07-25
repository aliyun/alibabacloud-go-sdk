// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuotaTopo interface {
	dara.Model
	String() string
	GoString() string
	SetDepth(v string) *QuotaTopo
	GetDepth() *string
	SetParentQuotaId(v string) *QuotaTopo
	GetParentQuotaId() *string
	SetQuotaDetails(v *QuotaDetails) *QuotaTopo
	GetQuotaDetails() *QuotaDetails
	SetQuotaId(v string) *QuotaTopo
	GetQuotaId() *string
	SetQuotaName(v string) *QuotaTopo
	GetQuotaName() *string
	SetResourceType(v string) *QuotaTopo
	GetResourceType() *string
	SetWorkloadDetails(v *WorkloadDetails) *QuotaTopo
	GetWorkloadDetails() *WorkloadDetails
}

type QuotaTopo struct {
	Depth           *string          `json:"Depth,omitempty" xml:"Depth,omitempty"`
	ParentQuotaId   *string          `json:"ParentQuotaId,omitempty" xml:"ParentQuotaId,omitempty"`
	QuotaDetails    *QuotaDetails    `json:"QuotaDetails,omitempty" xml:"QuotaDetails,omitempty"`
	QuotaId         *string          `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	QuotaName       *string          `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	ResourceType    *string          `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	WorkloadDetails *WorkloadDetails `json:"WorkloadDetails,omitempty" xml:"WorkloadDetails,omitempty"`
}

func (s QuotaTopo) String() string {
	return dara.Prettify(s)
}

func (s QuotaTopo) GoString() string {
	return s.String()
}

func (s *QuotaTopo) GetDepth() *string {
	return s.Depth
}

func (s *QuotaTopo) GetParentQuotaId() *string {
	return s.ParentQuotaId
}

func (s *QuotaTopo) GetQuotaDetails() *QuotaDetails {
	return s.QuotaDetails
}

func (s *QuotaTopo) GetQuotaId() *string {
	return s.QuotaId
}

func (s *QuotaTopo) GetQuotaName() *string {
	return s.QuotaName
}

func (s *QuotaTopo) GetResourceType() *string {
	return s.ResourceType
}

func (s *QuotaTopo) GetWorkloadDetails() *WorkloadDetails {
	return s.WorkloadDetails
}

func (s *QuotaTopo) SetDepth(v string) *QuotaTopo {
	s.Depth = &v
	return s
}

func (s *QuotaTopo) SetParentQuotaId(v string) *QuotaTopo {
	s.ParentQuotaId = &v
	return s
}

func (s *QuotaTopo) SetQuotaDetails(v *QuotaDetails) *QuotaTopo {
	s.QuotaDetails = v
	return s
}

func (s *QuotaTopo) SetQuotaId(v string) *QuotaTopo {
	s.QuotaId = &v
	return s
}

func (s *QuotaTopo) SetQuotaName(v string) *QuotaTopo {
	s.QuotaName = &v
	return s
}

func (s *QuotaTopo) SetResourceType(v string) *QuotaTopo {
	s.ResourceType = &v
	return s
}

func (s *QuotaTopo) SetWorkloadDetails(v *WorkloadDetails) *QuotaTopo {
	s.WorkloadDetails = v
	return s
}

func (s *QuotaTopo) Validate() error {
	return dara.Validate(s)
}
