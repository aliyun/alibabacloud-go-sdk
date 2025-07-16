// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTableCompaction interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogId(v string) *TableCompaction
	GetCatalogId() *string
	SetCuUsage(v float64) *TableCompaction
	GetCuUsage() *float64
	SetLastCompactedFileTime(v int64) *TableCompaction
	GetLastCompactedFileTime() *int64
	SetMaxLevel0FileCount(v string) *TableCompaction
	GetMaxLevel0FileCount() *string
	SetTableId(v string) *TableCompaction
	GetTableId() *string
}

type TableCompaction struct {
	CatalogId             *string  `json:"catalogId,omitempty" xml:"catalogId,omitempty"`
	CuUsage               *float64 `json:"cuUsage,omitempty" xml:"cuUsage,omitempty"`
	LastCompactedFileTime *int64   `json:"lastCompactedFileTime,omitempty" xml:"lastCompactedFileTime,omitempty"`
	MaxLevel0FileCount    *string  `json:"maxLevel0FileCount,omitempty" xml:"maxLevel0FileCount,omitempty"`
	TableId               *string  `json:"tableId,omitempty" xml:"tableId,omitempty"`
}

func (s TableCompaction) String() string {
	return dara.Prettify(s)
}

func (s TableCompaction) GoString() string {
	return s.String()
}

func (s *TableCompaction) GetCatalogId() *string {
	return s.CatalogId
}

func (s *TableCompaction) GetCuUsage() *float64 {
	return s.CuUsage
}

func (s *TableCompaction) GetLastCompactedFileTime() *int64 {
	return s.LastCompactedFileTime
}

func (s *TableCompaction) GetMaxLevel0FileCount() *string {
	return s.MaxLevel0FileCount
}

func (s *TableCompaction) GetTableId() *string {
	return s.TableId
}

func (s *TableCompaction) SetCatalogId(v string) *TableCompaction {
	s.CatalogId = &v
	return s
}

func (s *TableCompaction) SetCuUsage(v float64) *TableCompaction {
	s.CuUsage = &v
	return s
}

func (s *TableCompaction) SetLastCompactedFileTime(v int64) *TableCompaction {
	s.LastCompactedFileTime = &v
	return s
}

func (s *TableCompaction) SetMaxLevel0FileCount(v string) *TableCompaction {
	s.MaxLevel0FileCount = &v
	return s
}

func (s *TableCompaction) SetTableId(v string) *TableCompaction {
	s.TableId = &v
	return s
}

func (s *TableCompaction) Validate() error {
	return dara.Validate(s)
}
