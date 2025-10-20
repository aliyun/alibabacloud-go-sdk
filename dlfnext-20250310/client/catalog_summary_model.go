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
	SetObjTypeArchiveSize(v int64) *CatalogSummary
	GetObjTypeArchiveSize() *int64
	SetObjTypeColdArchiveSize(v int64) *CatalogSummary
	GetObjTypeColdArchiveSize() *int64
	SetObjTypeIaSize(v int64) *CatalogSummary
	GetObjTypeIaSize() *int64
	SetObjTypeStandardSize(v int64) *CatalogSummary
	GetObjTypeStandardSize() *int64
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
	SetTotalMetaSizeInBytes(v *MoMValues) *CatalogSummary
	GetTotalMetaSizeInBytes() *MoMValues
}

type CatalogSummary struct {
	ApiVisitCountMonthly   *int64     `json:"apiVisitCountMonthly,omitempty" xml:"apiVisitCountMonthly,omitempty"`
	DatabaseCount          *MoMValues `json:"databaseCount,omitempty" xml:"databaseCount,omitempty"`
	FileAccessCountMonthly *int64     `json:"fileAccessCountMonthly,omitempty" xml:"fileAccessCountMonthly,omitempty"`
	// Update date of the statistics
	GeneratedDate          *string    `json:"generatedDate,omitempty" xml:"generatedDate,omitempty"`
	ObjTypeArchiveSize     *int64     `json:"objTypeArchiveSize,omitempty" xml:"objTypeArchiveSize,omitempty"`
	ObjTypeColdArchiveSize *int64     `json:"objTypeColdArchiveSize,omitempty" xml:"objTypeColdArchiveSize,omitempty"`
	ObjTypeIaSize          *int64     `json:"objTypeIaSize,omitempty" xml:"objTypeIaSize,omitempty"`
	ObjTypeStandardSize    *int64     `json:"objTypeStandardSize,omitempty" xml:"objTypeStandardSize,omitempty"`
	PartitionCount         *MoMValues `json:"partitionCount,omitempty" xml:"partitionCount,omitempty"`
	TableCount             *MoMValues `json:"tableCount,omitempty" xml:"tableCount,omitempty"`
	ThroughputMonthly      *int64     `json:"throughputMonthly,omitempty" xml:"throughputMonthly,omitempty"`
	TotalFileCount         *MoMValues `json:"totalFileCount,omitempty" xml:"totalFileCount,omitempty"`
	TotalFileSizeInBytes   *MoMValues `json:"totalFileSizeInBytes,omitempty" xml:"totalFileSizeInBytes,omitempty"`
	TotalMetaSizeInBytes   *MoMValues `json:"totalMetaSizeInBytes,omitempty" xml:"totalMetaSizeInBytes,omitempty"`
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

func (s *CatalogSummary) GetObjTypeArchiveSize() *int64 {
	return s.ObjTypeArchiveSize
}

func (s *CatalogSummary) GetObjTypeColdArchiveSize() *int64 {
	return s.ObjTypeColdArchiveSize
}

func (s *CatalogSummary) GetObjTypeIaSize() *int64 {
	return s.ObjTypeIaSize
}

func (s *CatalogSummary) GetObjTypeStandardSize() *int64 {
	return s.ObjTypeStandardSize
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

func (s *CatalogSummary) GetTotalMetaSizeInBytes() *MoMValues {
	return s.TotalMetaSizeInBytes
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

func (s *CatalogSummary) SetObjTypeArchiveSize(v int64) *CatalogSummary {
	s.ObjTypeArchiveSize = &v
	return s
}

func (s *CatalogSummary) SetObjTypeColdArchiveSize(v int64) *CatalogSummary {
	s.ObjTypeColdArchiveSize = &v
	return s
}

func (s *CatalogSummary) SetObjTypeIaSize(v int64) *CatalogSummary {
	s.ObjTypeIaSize = &v
	return s
}

func (s *CatalogSummary) SetObjTypeStandardSize(v int64) *CatalogSummary {
	s.ObjTypeStandardSize = &v
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

func (s *CatalogSummary) SetTotalMetaSizeInBytes(v *MoMValues) *CatalogSummary {
	s.TotalMetaSizeInBytes = v
	return s
}

func (s *CatalogSummary) Validate() error {
	if s.DatabaseCount != nil {
		if err := s.DatabaseCount.Validate(); err != nil {
			return err
		}
	}
	if s.PartitionCount != nil {
		if err := s.PartitionCount.Validate(); err != nil {
			return err
		}
	}
	if s.TableCount != nil {
		if err := s.TableCount.Validate(); err != nil {
			return err
		}
	}
	if s.TotalFileCount != nil {
		if err := s.TotalFileCount.Validate(); err != nil {
			return err
		}
	}
	if s.TotalFileSizeInBytes != nil {
		if err := s.TotalFileSizeInBytes.Validate(); err != nil {
			return err
		}
	}
	if s.TotalMetaSizeInBytes != nil {
		if err := s.TotalMetaSizeInBytes.Validate(); err != nil {
			return err
		}
	}
	return nil
}
