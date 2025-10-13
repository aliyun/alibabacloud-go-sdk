// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuotaDetails interface {
	dara.Model
	String() string
	GoString() string
	SetActualMinQuota(v *ResourceAmount) *QuotaDetails
	GetActualMinQuota() *ResourceAmount
	SetAllocatableQuota(v *ResourceAmount) *QuotaDetails
	GetAllocatableQuota() *ResourceAmount
	SetAllocatedQuota(v *ResourceAmount) *QuotaDetails
	GetAllocatedQuota() *ResourceAmount
	SetAncestorsAllocatedQuota(v *ResourceAmount) *QuotaDetails
	GetAncestorsAllocatedQuota() *ResourceAmount
	SetDescendantsAllocatedQuota(v *ResourceAmount) *QuotaDetails
	GetDescendantsAllocatedQuota() *ResourceAmount
	SetDesiredMinQuota(v *ResourceAmount) *QuotaDetails
	GetDesiredMinQuota() *ResourceAmount
	SetRequestedQuota(v *ResourceAmount) *QuotaDetails
	GetRequestedQuota() *ResourceAmount
	SetSelfAllocatedQuota(v *ResourceAmount) *QuotaDetails
	GetSelfAllocatedQuota() *ResourceAmount
	SetSelfSubmittedQuota(v *ResourceAmount) *QuotaDetails
	GetSelfSubmittedQuota() *ResourceAmount
	SetSystemReservedQuota(v *ResourceAmount) *QuotaDetails
	GetSystemReservedQuota() *ResourceAmount
	SetUsedQuota(v *ResourceAmount) *QuotaDetails
	GetUsedQuota() *ResourceAmount
}

type QuotaDetails struct {
	ActualMinQuota            *ResourceAmount `json:"ActualMinQuota,omitempty" xml:"ActualMinQuota,omitempty"`
	AllocatableQuota          *ResourceAmount `json:"AllocatableQuota,omitempty" xml:"AllocatableQuota,omitempty"`
	AllocatedQuota            *ResourceAmount `json:"AllocatedQuota,omitempty" xml:"AllocatedQuota,omitempty"`
	AncestorsAllocatedQuota   *ResourceAmount `json:"AncestorsAllocatedQuota,omitempty" xml:"AncestorsAllocatedQuota,omitempty"`
	DescendantsAllocatedQuota *ResourceAmount `json:"DescendantsAllocatedQuota,omitempty" xml:"DescendantsAllocatedQuota,omitempty"`
	DesiredMinQuota           *ResourceAmount `json:"DesiredMinQuota,omitempty" xml:"DesiredMinQuota,omitempty"`
	RequestedQuota            *ResourceAmount `json:"RequestedQuota,omitempty" xml:"RequestedQuota,omitempty"`
	SelfAllocatedQuota        *ResourceAmount `json:"SelfAllocatedQuota,omitempty" xml:"SelfAllocatedQuota,omitempty"`
	SelfSubmittedQuota        *ResourceAmount `json:"SelfSubmittedQuota,omitempty" xml:"SelfSubmittedQuota,omitempty"`
	SystemReservedQuota       *ResourceAmount `json:"SystemReservedQuota,omitempty" xml:"SystemReservedQuota,omitempty"`
	UsedQuota                 *ResourceAmount `json:"UsedQuota,omitempty" xml:"UsedQuota,omitempty"`
}

func (s QuotaDetails) String() string {
	return dara.Prettify(s)
}

func (s QuotaDetails) GoString() string {
	return s.String()
}

func (s *QuotaDetails) GetActualMinQuota() *ResourceAmount {
	return s.ActualMinQuota
}

func (s *QuotaDetails) GetAllocatableQuota() *ResourceAmount {
	return s.AllocatableQuota
}

func (s *QuotaDetails) GetAllocatedQuota() *ResourceAmount {
	return s.AllocatedQuota
}

func (s *QuotaDetails) GetAncestorsAllocatedQuota() *ResourceAmount {
	return s.AncestorsAllocatedQuota
}

func (s *QuotaDetails) GetDescendantsAllocatedQuota() *ResourceAmount {
	return s.DescendantsAllocatedQuota
}

func (s *QuotaDetails) GetDesiredMinQuota() *ResourceAmount {
	return s.DesiredMinQuota
}

func (s *QuotaDetails) GetRequestedQuota() *ResourceAmount {
	return s.RequestedQuota
}

func (s *QuotaDetails) GetSelfAllocatedQuota() *ResourceAmount {
	return s.SelfAllocatedQuota
}

func (s *QuotaDetails) GetSelfSubmittedQuota() *ResourceAmount {
	return s.SelfSubmittedQuota
}

func (s *QuotaDetails) GetSystemReservedQuota() *ResourceAmount {
	return s.SystemReservedQuota
}

func (s *QuotaDetails) GetUsedQuota() *ResourceAmount {
	return s.UsedQuota
}

func (s *QuotaDetails) SetActualMinQuota(v *ResourceAmount) *QuotaDetails {
	s.ActualMinQuota = v
	return s
}

func (s *QuotaDetails) SetAllocatableQuota(v *ResourceAmount) *QuotaDetails {
	s.AllocatableQuota = v
	return s
}

func (s *QuotaDetails) SetAllocatedQuota(v *ResourceAmount) *QuotaDetails {
	s.AllocatedQuota = v
	return s
}

func (s *QuotaDetails) SetAncestorsAllocatedQuota(v *ResourceAmount) *QuotaDetails {
	s.AncestorsAllocatedQuota = v
	return s
}

func (s *QuotaDetails) SetDescendantsAllocatedQuota(v *ResourceAmount) *QuotaDetails {
	s.DescendantsAllocatedQuota = v
	return s
}

func (s *QuotaDetails) SetDesiredMinQuota(v *ResourceAmount) *QuotaDetails {
	s.DesiredMinQuota = v
	return s
}

func (s *QuotaDetails) SetRequestedQuota(v *ResourceAmount) *QuotaDetails {
	s.RequestedQuota = v
	return s
}

func (s *QuotaDetails) SetSelfAllocatedQuota(v *ResourceAmount) *QuotaDetails {
	s.SelfAllocatedQuota = v
	return s
}

func (s *QuotaDetails) SetSelfSubmittedQuota(v *ResourceAmount) *QuotaDetails {
	s.SelfSubmittedQuota = v
	return s
}

func (s *QuotaDetails) SetSystemReservedQuota(v *ResourceAmount) *QuotaDetails {
	s.SystemReservedQuota = v
	return s
}

func (s *QuotaDetails) SetUsedQuota(v *ResourceAmount) *QuotaDetails {
	s.UsedQuota = v
	return s
}

func (s *QuotaDetails) Validate() error {
	if s.ActualMinQuota != nil {
		if err := s.ActualMinQuota.Validate(); err != nil {
			return err
		}
	}
	if s.AllocatableQuota != nil {
		if err := s.AllocatableQuota.Validate(); err != nil {
			return err
		}
	}
	if s.AllocatedQuota != nil {
		if err := s.AllocatedQuota.Validate(); err != nil {
			return err
		}
	}
	if s.AncestorsAllocatedQuota != nil {
		if err := s.AncestorsAllocatedQuota.Validate(); err != nil {
			return err
		}
	}
	if s.DescendantsAllocatedQuota != nil {
		if err := s.DescendantsAllocatedQuota.Validate(); err != nil {
			return err
		}
	}
	if s.DesiredMinQuota != nil {
		if err := s.DesiredMinQuota.Validate(); err != nil {
			return err
		}
	}
	if s.RequestedQuota != nil {
		if err := s.RequestedQuota.Validate(); err != nil {
			return err
		}
	}
	if s.SelfAllocatedQuota != nil {
		if err := s.SelfAllocatedQuota.Validate(); err != nil {
			return err
		}
	}
	if s.SelfSubmittedQuota != nil {
		if err := s.SelfSubmittedQuota.Validate(); err != nil {
			return err
		}
	}
	if s.SystemReservedQuota != nil {
		if err := s.SystemReservedQuota.Validate(); err != nil {
			return err
		}
	}
	if s.UsedQuota != nil {
		if err := s.UsedQuota.Validate(); err != nil {
			return err
		}
	}
	return nil
}
