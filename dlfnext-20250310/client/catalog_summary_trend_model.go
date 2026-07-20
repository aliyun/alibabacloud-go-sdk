// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCatalogSummaryTrend interface {
	dara.Model
	String() string
	GoString() string
	SetApiVisitCount(v []*DateSummary) *CatalogSummaryTrend
	GetApiVisitCount() []*DateSummary
	SetFileAccessCount(v []*DateSummary) *CatalogSummaryTrend
	GetFileAccessCount() []*DateSummary
	SetThroughput(v []*DateSummary) *CatalogSummaryTrend
	GetThroughput() []*DateSummary
	SetTotalFileCount(v []*DateSummary) *CatalogSummaryTrend
	GetTotalFileCount() []*DateSummary
	SetTotalFileSizeInBytes(v []*DateSummary) *CatalogSummaryTrend
	GetTotalFileSizeInBytes() []*DateSummary
	SetTotalMetaCount(v []*DateSummary) *CatalogSummaryTrend
	GetTotalMetaCount() []*DateSummary
}

type CatalogSummaryTrend struct {
	// API visit count trends
	ApiVisitCount []*DateSummary `json:"apiVisitCount,omitempty" xml:"apiVisitCount,omitempty" type:"Repeated"`
	// file access count trends
	FileAccessCount []*DateSummary `json:"fileAccessCount,omitempty" xml:"fileAccessCount,omitempty" type:"Repeated"`
	// Table count trends
	Throughput []*DateSummary `json:"throughput,omitempty" xml:"throughput,omitempty" type:"Repeated"`
	// Historical total file count
	TotalFileCount []*DateSummary `json:"totalFileCount,omitempty" xml:"totalFileCount,omitempty" type:"Repeated"`
	// Database count trends
	TotalFileSizeInBytes []*DateSummary `json:"totalFileSizeInBytes,omitempty" xml:"totalFileSizeInBytes,omitempty" type:"Repeated"`
	// Latest snapshot file count
	TotalMetaCount []*DateSummary `json:"totalMetaCount,omitempty" xml:"totalMetaCount,omitempty" type:"Repeated"`
}

func (s CatalogSummaryTrend) String() string {
	return dara.Prettify(s)
}

func (s CatalogSummaryTrend) GoString() string {
	return s.String()
}

func (s *CatalogSummaryTrend) GetApiVisitCount() []*DateSummary {
	return s.ApiVisitCount
}

func (s *CatalogSummaryTrend) GetFileAccessCount() []*DateSummary {
	return s.FileAccessCount
}

func (s *CatalogSummaryTrend) GetThroughput() []*DateSummary {
	return s.Throughput
}

func (s *CatalogSummaryTrend) GetTotalFileCount() []*DateSummary {
	return s.TotalFileCount
}

func (s *CatalogSummaryTrend) GetTotalFileSizeInBytes() []*DateSummary {
	return s.TotalFileSizeInBytes
}

func (s *CatalogSummaryTrend) GetTotalMetaCount() []*DateSummary {
	return s.TotalMetaCount
}

func (s *CatalogSummaryTrend) SetApiVisitCount(v []*DateSummary) *CatalogSummaryTrend {
	s.ApiVisitCount = v
	return s
}

func (s *CatalogSummaryTrend) SetFileAccessCount(v []*DateSummary) *CatalogSummaryTrend {
	s.FileAccessCount = v
	return s
}

func (s *CatalogSummaryTrend) SetThroughput(v []*DateSummary) *CatalogSummaryTrend {
	s.Throughput = v
	return s
}

func (s *CatalogSummaryTrend) SetTotalFileCount(v []*DateSummary) *CatalogSummaryTrend {
	s.TotalFileCount = v
	return s
}

func (s *CatalogSummaryTrend) SetTotalFileSizeInBytes(v []*DateSummary) *CatalogSummaryTrend {
	s.TotalFileSizeInBytes = v
	return s
}

func (s *CatalogSummaryTrend) SetTotalMetaCount(v []*DateSummary) *CatalogSummaryTrend {
	s.TotalMetaCount = v
	return s
}

func (s *CatalogSummaryTrend) Validate() error {
	if s.ApiVisitCount != nil {
		for _, item := range s.ApiVisitCount {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FileAccessCount != nil {
		for _, item := range s.FileAccessCount {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Throughput != nil {
		for _, item := range s.Throughput {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TotalFileCount != nil {
		for _, item := range s.TotalFileCount {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TotalFileSizeInBytes != nil {
		for _, item := range s.TotalFileSizeInBytes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TotalMetaCount != nil {
		for _, item := range s.TotalMetaCount {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
