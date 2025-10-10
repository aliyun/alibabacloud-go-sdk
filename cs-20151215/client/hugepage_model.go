// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHugepage interface {
	dara.Model
	String() string
	GoString() string
	SetKhugepagedAllocSleepMillisecs(v int64) *Hugepage
	GetKhugepagedAllocSleepMillisecs() *int64
	SetKhugepagedDefrag(v int64) *Hugepage
	GetKhugepagedDefrag() *int64
	SetKhugepagedPagesToScan(v int64) *Hugepage
	GetKhugepagedPagesToScan() *int64
	SetKhugepagedScanSleepMillisecs(v int64) *Hugepage
	GetKhugepagedScanSleepMillisecs() *int64
	SetTransparentDefrag(v string) *Hugepage
	GetTransparentDefrag() *string
	SetTransparentEnabled(v string) *Hugepage
	GetTransparentEnabled() *string
}

type Hugepage struct {
	KhugepagedAllocSleepMillisecs *int64  `json:"khugepagedAllocSleepMillisecs,omitempty" xml:"khugepagedAllocSleepMillisecs,omitempty"`
	KhugepagedDefrag              *int64  `json:"khugepagedDefrag,omitempty" xml:"khugepagedDefrag,omitempty"`
	KhugepagedPagesToScan         *int64  `json:"khugepagedPagesToScan,omitempty" xml:"khugepagedPagesToScan,omitempty"`
	KhugepagedScanSleepMillisecs  *int64  `json:"khugepagedScanSleepMillisecs,omitempty" xml:"khugepagedScanSleepMillisecs,omitempty"`
	TransparentDefrag             *string `json:"transparentDefrag,omitempty" xml:"transparentDefrag,omitempty"`
	TransparentEnabled            *string `json:"transparentEnabled,omitempty" xml:"transparentEnabled,omitempty"`
}

func (s Hugepage) String() string {
	return dara.Prettify(s)
}

func (s Hugepage) GoString() string {
	return s.String()
}

func (s *Hugepage) GetKhugepagedAllocSleepMillisecs() *int64 {
	return s.KhugepagedAllocSleepMillisecs
}

func (s *Hugepage) GetKhugepagedDefrag() *int64 {
	return s.KhugepagedDefrag
}

func (s *Hugepage) GetKhugepagedPagesToScan() *int64 {
	return s.KhugepagedPagesToScan
}

func (s *Hugepage) GetKhugepagedScanSleepMillisecs() *int64 {
	return s.KhugepagedScanSleepMillisecs
}

func (s *Hugepage) GetTransparentDefrag() *string {
	return s.TransparentDefrag
}

func (s *Hugepage) GetTransparentEnabled() *string {
	return s.TransparentEnabled
}

func (s *Hugepage) SetKhugepagedAllocSleepMillisecs(v int64) *Hugepage {
	s.KhugepagedAllocSleepMillisecs = &v
	return s
}

func (s *Hugepage) SetKhugepagedDefrag(v int64) *Hugepage {
	s.KhugepagedDefrag = &v
	return s
}

func (s *Hugepage) SetKhugepagedPagesToScan(v int64) *Hugepage {
	s.KhugepagedPagesToScan = &v
	return s
}

func (s *Hugepage) SetKhugepagedScanSleepMillisecs(v int64) *Hugepage {
	s.KhugepagedScanSleepMillisecs = &v
	return s
}

func (s *Hugepage) SetTransparentDefrag(v string) *Hugepage {
	s.TransparentDefrag = &v
	return s
}

func (s *Hugepage) SetTransparentEnabled(v string) *Hugepage {
	s.TransparentEnabled = &v
	return s
}

func (s *Hugepage) Validate() error {
	return dara.Validate(s)
}
