// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuota interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *Quota
	GetClusterId() *string
	SetClusterName(v string) *Quota
	GetClusterName() *string
	SetQuotaConfig(v *QuotaConfig) *Quota
	GetQuotaConfig() *QuotaConfig
	SetQuotaId(v string) *Quota
	GetQuotaId() *string
	SetQuotaName(v string) *Quota
	GetQuotaName() *string
	SetQuotaType(v string) *Quota
	GetQuotaType() *string
	SetTotalQuota(v *QuotaDetail) *Quota
	GetTotalQuota() *QuotaDetail
	SetTotalTideQuota(v *QuotaDetail) *Quota
	GetTotalTideQuota() *QuotaDetail
	SetUsedQuota(v *QuotaDetail) *Quota
	GetUsedQuota() *QuotaDetail
	SetUsedTideQuota(v *QuotaDetail) *Quota
	GetUsedTideQuota() *QuotaDetail
}

type Quota struct {
	ClusterId   *string      `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName *string      `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	QuotaConfig *QuotaConfig `json:"QuotaConfig,omitempty" xml:"QuotaConfig,omitempty"`
	// example:
	//
	// quotamtl37ge7gkvdz
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// example:
	//
	// dlc-quota
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// example:
	//
	// asiquota
	QuotaType      *string      `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	TotalQuota     *QuotaDetail `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	TotalTideQuota *QuotaDetail `json:"TotalTideQuota,omitempty" xml:"TotalTideQuota,omitempty"`
	UsedQuota      *QuotaDetail `json:"UsedQuota,omitempty" xml:"UsedQuota,omitempty"`
	UsedTideQuota  *QuotaDetail `json:"UsedTideQuota,omitempty" xml:"UsedTideQuota,omitempty"`
}

func (s Quota) String() string {
	return dara.Prettify(s)
}

func (s Quota) GoString() string {
	return s.String()
}

func (s *Quota) GetClusterId() *string {
	return s.ClusterId
}

func (s *Quota) GetClusterName() *string {
	return s.ClusterName
}

func (s *Quota) GetQuotaConfig() *QuotaConfig {
	return s.QuotaConfig
}

func (s *Quota) GetQuotaId() *string {
	return s.QuotaId
}

func (s *Quota) GetQuotaName() *string {
	return s.QuotaName
}

func (s *Quota) GetQuotaType() *string {
	return s.QuotaType
}

func (s *Quota) GetTotalQuota() *QuotaDetail {
	return s.TotalQuota
}

func (s *Quota) GetTotalTideQuota() *QuotaDetail {
	return s.TotalTideQuota
}

func (s *Quota) GetUsedQuota() *QuotaDetail {
	return s.UsedQuota
}

func (s *Quota) GetUsedTideQuota() *QuotaDetail {
	return s.UsedTideQuota
}

func (s *Quota) SetClusterId(v string) *Quota {
	s.ClusterId = &v
	return s
}

func (s *Quota) SetClusterName(v string) *Quota {
	s.ClusterName = &v
	return s
}

func (s *Quota) SetQuotaConfig(v *QuotaConfig) *Quota {
	s.QuotaConfig = v
	return s
}

func (s *Quota) SetQuotaId(v string) *Quota {
	s.QuotaId = &v
	return s
}

func (s *Quota) SetQuotaName(v string) *Quota {
	s.QuotaName = &v
	return s
}

func (s *Quota) SetQuotaType(v string) *Quota {
	s.QuotaType = &v
	return s
}

func (s *Quota) SetTotalQuota(v *QuotaDetail) *Quota {
	s.TotalQuota = v
	return s
}

func (s *Quota) SetTotalTideQuota(v *QuotaDetail) *Quota {
	s.TotalTideQuota = v
	return s
}

func (s *Quota) SetUsedQuota(v *QuotaDetail) *Quota {
	s.UsedQuota = v
	return s
}

func (s *Quota) SetUsedTideQuota(v *QuotaDetail) *Quota {
	s.UsedTideQuota = v
	return s
}

func (s *Quota) Validate() error {
	if s.QuotaConfig != nil {
		if err := s.QuotaConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TotalQuota != nil {
		if err := s.TotalQuota.Validate(); err != nil {
			return err
		}
	}
	if s.TotalTideQuota != nil {
		if err := s.TotalTideQuota.Validate(); err != nil {
			return err
		}
	}
	if s.UsedQuota != nil {
		if err := s.UsedQuota.Validate(); err != nil {
			return err
		}
	}
	if s.UsedTideQuota != nil {
		if err := s.UsedTideQuota.Validate(); err != nil {
			return err
		}
	}
	return nil
}
