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
	SetDesiredMinQuota(v *ResourceAmount) *QuotaDetails
	GetDesiredMinQuota() *ResourceAmount
	SetRequestedQuota(v *ResourceAmount) *QuotaDetails
	GetRequestedQuota() *ResourceAmount
	SetUsedQuota(v *ResourceAmount) *QuotaDetails
	GetUsedQuota() *ResourceAmount
}

type QuotaDetails struct {
	ActualMinQuota  *ResourceAmount `json:"ActualMinQuota,omitempty" xml:"ActualMinQuota,omitempty"`
	DesiredMinQuota *ResourceAmount `json:"DesiredMinQuota,omitempty" xml:"DesiredMinQuota,omitempty"`
	RequestedQuota  *ResourceAmount `json:"RequestedQuota,omitempty" xml:"RequestedQuota,omitempty"`
	UsedQuota       *ResourceAmount `json:"UsedQuota,omitempty" xml:"UsedQuota,omitempty"`
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

func (s *QuotaDetails) GetDesiredMinQuota() *ResourceAmount {
	return s.DesiredMinQuota
}

func (s *QuotaDetails) GetRequestedQuota() *ResourceAmount {
	return s.RequestedQuota
}

func (s *QuotaDetails) GetUsedQuota() *ResourceAmount {
	return s.UsedQuota
}

func (s *QuotaDetails) SetActualMinQuota(v *ResourceAmount) *QuotaDetails {
	s.ActualMinQuota = v
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
	if s.UsedQuota != nil {
		if err := s.UsedQuota.Validate(); err != nil {
			return err
		}
	}
	return nil
}
