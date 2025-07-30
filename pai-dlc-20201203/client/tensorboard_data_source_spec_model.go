// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTensorboardDataSourceSpec interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceType(v string) *TensorboardDataSourceSpec
	GetDataSourceType() *string
	SetDirectoryName(v string) *TensorboardDataSourceSpec
	GetDirectoryName() *string
	SetFullSummaryPath(v string) *TensorboardDataSourceSpec
	GetFullSummaryPath() *string
	SetId(v string) *TensorboardDataSourceSpec
	GetId() *string
	SetName(v string) *TensorboardDataSourceSpec
	GetName() *string
	SetSourceType(v string) *TensorboardDataSourceSpec
	GetSourceType() *string
	SetSummaryPath(v string) *TensorboardDataSourceSpec
	GetSummaryPath() *string
	SetUri(v string) *TensorboardDataSourceSpec
	GetUri() *string
}

type TensorboardDataSourceSpec struct {
	// example:
	//
	// OSS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// example:
	//
	// dlcJobName
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	// example:
	//
	// oss://xxxxx/tensorboard/run1
	FullSummaryPath *string `json:"FullSummaryPath,omitempty" xml:"FullSummaryPath,omitempty"`
	// example:
	//
	// d-vf2fdhxxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// dlcJobName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// datasource
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// /tensorboard/run1
	SummaryPath *string `json:"SummaryPath,omitempty" xml:"SummaryPath,omitempty"`
	// example:
	//
	// oss://.oss-cn-shanghai-finance-1.aliyuncs.com/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s TensorboardDataSourceSpec) String() string {
	return dara.Prettify(s)
}

func (s TensorboardDataSourceSpec) GoString() string {
	return s.String()
}

func (s *TensorboardDataSourceSpec) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *TensorboardDataSourceSpec) GetDirectoryName() *string {
	return s.DirectoryName
}

func (s *TensorboardDataSourceSpec) GetFullSummaryPath() *string {
	return s.FullSummaryPath
}

func (s *TensorboardDataSourceSpec) GetId() *string {
	return s.Id
}

func (s *TensorboardDataSourceSpec) GetName() *string {
	return s.Name
}

func (s *TensorboardDataSourceSpec) GetSourceType() *string {
	return s.SourceType
}

func (s *TensorboardDataSourceSpec) GetSummaryPath() *string {
	return s.SummaryPath
}

func (s *TensorboardDataSourceSpec) GetUri() *string {
	return s.Uri
}

func (s *TensorboardDataSourceSpec) SetDataSourceType(v string) *TensorboardDataSourceSpec {
	s.DataSourceType = &v
	return s
}

func (s *TensorboardDataSourceSpec) SetDirectoryName(v string) *TensorboardDataSourceSpec {
	s.DirectoryName = &v
	return s
}

func (s *TensorboardDataSourceSpec) SetFullSummaryPath(v string) *TensorboardDataSourceSpec {
	s.FullSummaryPath = &v
	return s
}

func (s *TensorboardDataSourceSpec) SetId(v string) *TensorboardDataSourceSpec {
	s.Id = &v
	return s
}

func (s *TensorboardDataSourceSpec) SetName(v string) *TensorboardDataSourceSpec {
	s.Name = &v
	return s
}

func (s *TensorboardDataSourceSpec) SetSourceType(v string) *TensorboardDataSourceSpec {
	s.SourceType = &v
	return s
}

func (s *TensorboardDataSourceSpec) SetSummaryPath(v string) *TensorboardDataSourceSpec {
	s.SummaryPath = &v
	return s
}

func (s *TensorboardDataSourceSpec) SetUri(v string) *TensorboardDataSourceSpec {
	s.Uri = &v
	return s
}

func (s *TensorboardDataSourceSpec) Validate() error {
	return dara.Validate(s)
}
