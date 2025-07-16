// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCatalogSummary interface {
	dara.Model
	String() string
	GoString() string
	SetApiVisitCountMonthly(v int64) *CatalogSummary
	GetApiVisitCountMonthly() *int64
	SetDatabaseCount(v *MoMValues) *CatalogSummary
	GetDatabaseCount() *MoMValues
	SetFileAccessCountMonthly(v int64) *CatalogSummary
	GetFileAccessCountMonthly() *int64
	SetGeneratedDate(v string) *CatalogSummary
	GetGeneratedDate() *string
	SetPartitionCount(v *MoMValues) *CatalogSummary
	GetPartitionCount() *MoMValues
	SetTableCount(v *MoMValues) *CatalogSummary
	GetTableCount() *MoMValues
	SetThroughputMonthly(v int64) *CatalogSummary
	GetThroughputMonthly() *int64
	SetTotalFileCount(v *MoMValues) *CatalogSummary
	GetTotalFileCount() *MoMValues
	SetTotalFileSizeInBytes(v *MoMValues) *CatalogSummary
	GetTotalFileSizeInBytes() *MoMValues
}

type CatalogSummary struct {
	ApiVisitCountMonthly   *int64     `json:"apiVisitCountMonthly,omitempty" xml:"apiVisitCountMonthly,omitempty"`
	DatabaseCount          *MoMValues `json:"databaseCount,omitempty" xml:"databaseCount,omitempty"`
	FileAccessCountMonthly *int64     `json:"fileAccessCountMonthly,omitempty" xml:"fileAccessCountMonthly,omitempty"`
	// Update date of the statistics
	GeneratedDate        *string    `json:"generatedDate,omitempty" xml:"generatedDate,omitempty"`
	PartitionCount       *MoMValues `json:"partitionCount,omitempty" xml:"partitionCount,omitempty"`
	TableCount           *MoMValues `json:"tableCount,omitempty" xml:"tableCount,omitempty"`
	ThroughputMonthly    *int64     `json:"throughputMonthly,omitempty" xml:"throughputMonthly,omitempty"`
	TotalFileCount       *MoMValues `json:"totalFileCount,omitempty" xml:"totalFileCount,omitempty"`
	TotalFileSizeInBytes *MoMValues `json:"totalFileSizeInBytes,omitempty" xml:"totalFileSizeInBytes,omitempty"`
}

func (s CatalogSummary) String() string {
	return dara.Prettify(s)
}

func (s CatalogSummary) GoString() string {
	return s.String()
}

func (s *CatalogSummary) GetApiVisitCountMonthly() *int64 {
	return s.ApiVisitCountMonthly
}

func (s *CatalogSummary) GetDatabaseCount() *MoMValues {
	return s.DatabaseCount
}

func (s *CatalogSummary) GetFileAccessCountMonthly() *int64 {
	return s.FileAccessCountMonthly
}

func (s *CatalogSummary) GetGeneratedDate() *string {
	return s.GeneratedDate
}

func (s *CatalogSummary) GetPartitionCount() *MoMValues {
	return s.PartitionCount
}

func (s *CatalogSummary) GetTableCount() *MoMValues {
	return s.TableCount
}

func (s *CatalogSummary) GetThroughputMonthly() *int64 {
	return s.ThroughputMonthly
}

func (s *CatalogSummary) GetTotalFileCount() *MoMValues {
	return s.TotalFileCount
}

func (s *CatalogSummary) GetTotalFileSizeInBytes() *MoMValues {
	return s.TotalFileSizeInBytes
}

func (s *CatalogSummary) SetApiVisitCountMonthly(v int64) *CatalogSummary {
	s.ApiVisitCountMonthly = &v
	return s
}

func (s *CatalogSummary) SetDatabaseCount(v *MoMValues) *CatalogSummary {
	s.DatabaseCount = v
	return s
}

func (s *CatalogSummary) SetFileAccessCountMonthly(v int64) *CatalogSummary {
	s.FileAccessCountMonthly = &v
	return s
}

func (s *CatalogSummary) SetGeneratedDate(v string) *CatalogSummary {
	s.GeneratedDate = &v
	return s
}

func (s *CatalogSummary) SetPartitionCount(v *MoMValues) *CatalogSummary {
	s.PartitionCount = v
	return s
}

func (s *CatalogSummary) SetTableCount(v *MoMValues) *CatalogSummary {
	s.TableCount = v
	return s
}

func (s *CatalogSummary) SetThroughputMonthly(v int64) *CatalogSummary {
	s.ThroughputMonthly = &v
	return s
}

func (s *CatalogSummary) SetTotalFileCount(v *MoMValues) *CatalogSummary {
	s.TotalFileCount = v
	return s
}

func (s *CatalogSummary) SetTotalFileSizeInBytes(v *MoMValues) *CatalogSummary {
	s.TotalFileSizeInBytes = v
	return s
}

func (s *CatalogSummary) Validate() error {
	return dara.Validate(s)
}
