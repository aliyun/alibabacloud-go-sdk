// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuotaJob interface {
	dara.Model
	String() string
	GoString() string
	SetQueuing(v int64) *QuotaJob
	GetQueuing() *int64
	SetRunning(v int64) *QuotaJob
	GetRunning() *int64
	SetTotal(v int64) *QuotaJob
	GetTotal() *int64
}

type QuotaJob struct {
	Queuing *int64 `json:"Queuing,omitempty" xml:"Queuing,omitempty"`
	Running *int64 `json:"Running,omitempty" xml:"Running,omitempty"`
	Total   *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QuotaJob) String() string {
	return dara.Prettify(s)
}

func (s QuotaJob) GoString() string {
	return s.String()
}

func (s *QuotaJob) GetQueuing() *int64 {
	return s.Queuing
}

func (s *QuotaJob) GetRunning() *int64 {
	return s.Running
}

func (s *QuotaJob) GetTotal() *int64 {
	return s.Total
}

func (s *QuotaJob) SetQueuing(v int64) *QuotaJob {
	s.Queuing = &v
	return s
}

func (s *QuotaJob) SetRunning(v int64) *QuotaJob {
	s.Running = &v
	return s
}

func (s *QuotaJob) SetTotal(v int64) *QuotaJob {
	s.Total = &v
	return s
}

func (s *QuotaJob) Validate() error {
	return dara.Validate(s)
}
