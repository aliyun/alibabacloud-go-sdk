// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatasetProxyConfig interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetType(v string) *DatasetProxyConfig
	GetDatasetType() *string
	SetSource(v string) *DatasetProxyConfig
	GetSource() *string
	SetSourceDatasetId(v string) *DatasetProxyConfig
	GetSourceDatasetId() *string
}

type DatasetProxyConfig struct {
	// Dataset type. Only LABEL is supported.
	//
	// example:
	//
	// LABEL
	DatasetType *string `json:"DatasetType,omitempty" xml:"DatasetType,omitempty"`
	// Dataset source. Only PAI is supported.
	//
	// example:
	//
	// PAI
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Source dataset ID. For information about how to obtain a dataset ID, see [ListDatasets](https://help.aliyun.com/document_detail/457222.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 214***12514
	SourceDatasetId *string `json:"SourceDatasetId,omitempty" xml:"SourceDatasetId,omitempty"`
}

func (s DatasetProxyConfig) String() string {
	return dara.Prettify(s)
}

func (s DatasetProxyConfig) GoString() string {
	return s.String()
}

func (s *DatasetProxyConfig) GetDatasetType() *string {
	return s.DatasetType
}

func (s *DatasetProxyConfig) GetSource() *string {
	return s.Source
}

func (s *DatasetProxyConfig) GetSourceDatasetId() *string {
	return s.SourceDatasetId
}

func (s *DatasetProxyConfig) SetDatasetType(v string) *DatasetProxyConfig {
	s.DatasetType = &v
	return s
}

func (s *DatasetProxyConfig) SetSource(v string) *DatasetProxyConfig {
	s.Source = &v
	return s
}

func (s *DatasetProxyConfig) SetSourceDatasetId(v string) *DatasetProxyConfig {
	s.SourceDatasetId = &v
	return s
}

func (s *DatasetProxyConfig) Validate() error {
	return dara.Validate(s)
}
