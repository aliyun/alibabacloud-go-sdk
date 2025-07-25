// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkloadDetails interface {
	dara.Model
	String() string
	GoString() string
	SetDLC(v *QuotaJob) *WorkloadDetails
	GetDLC() *QuotaJob
	SetDSW(v *QuotaJob) *WorkloadDetails
	GetDSW() *QuotaJob
	SetEAS(v *QuotaJob) *WorkloadDetails
	GetEAS() *QuotaJob
	SetSummary(v *QuotaJob) *WorkloadDetails
	GetSummary() *QuotaJob
}

type WorkloadDetails struct {
	DLC     *QuotaJob `json:"DLC,omitempty" xml:"DLC,omitempty"`
	DSW     *QuotaJob `json:"DSW,omitempty" xml:"DSW,omitempty"`
	EAS     *QuotaJob `json:"EAS,omitempty" xml:"EAS,omitempty"`
	Summary *QuotaJob `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s WorkloadDetails) String() string {
	return dara.Prettify(s)
}

func (s WorkloadDetails) GoString() string {
	return s.String()
}

func (s *WorkloadDetails) GetDLC() *QuotaJob {
	return s.DLC
}

func (s *WorkloadDetails) GetDSW() *QuotaJob {
	return s.DSW
}

func (s *WorkloadDetails) GetEAS() *QuotaJob {
	return s.EAS
}

func (s *WorkloadDetails) GetSummary() *QuotaJob {
	return s.Summary
}

func (s *WorkloadDetails) SetDLC(v *QuotaJob) *WorkloadDetails {
	s.DLC = v
	return s
}

func (s *WorkloadDetails) SetDSW(v *QuotaJob) *WorkloadDetails {
	s.DSW = v
	return s
}

func (s *WorkloadDetails) SetEAS(v *QuotaJob) *WorkloadDetails {
	s.EAS = v
	return s
}

func (s *WorkloadDetails) SetSummary(v *QuotaJob) *WorkloadDetails {
	s.Summary = v
	return s
}

func (s *WorkloadDetails) Validate() error {
	return dara.Validate(s)
}
