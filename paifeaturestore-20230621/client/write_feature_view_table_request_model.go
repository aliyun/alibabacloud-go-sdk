// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWriteFeatureViewTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMode(v string) *WriteFeatureViewTableRequest
	GetMode() *string
	SetPartitions(v map[string]map[string]interface{}) *WriteFeatureViewTableRequest
	GetPartitions() map[string]map[string]interface{}
	SetUrlDatasource(v *WriteFeatureViewTableRequestUrlDatasource) *WriteFeatureViewTableRequest
	GetUrlDatasource() *WriteFeatureViewTableRequestUrlDatasource
}

type WriteFeatureViewTableRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Merge
	Mode          *string                                    `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Partitions    map[string]map[string]interface{}          `json:"Partitions,omitempty" xml:"Partitions,omitempty"`
	UrlDatasource *WriteFeatureViewTableRequestUrlDatasource `json:"UrlDatasource,omitempty" xml:"UrlDatasource,omitempty" type:"Struct"`
}

func (s WriteFeatureViewTableRequest) String() string {
	return dara.Prettify(s)
}

func (s WriteFeatureViewTableRequest) GoString() string {
	return s.String()
}

func (s *WriteFeatureViewTableRequest) GetMode() *string {
	return s.Mode
}

func (s *WriteFeatureViewTableRequest) GetPartitions() map[string]map[string]interface{} {
	return s.Partitions
}

func (s *WriteFeatureViewTableRequest) GetUrlDatasource() *WriteFeatureViewTableRequestUrlDatasource {
	return s.UrlDatasource
}

func (s *WriteFeatureViewTableRequest) SetMode(v string) *WriteFeatureViewTableRequest {
	s.Mode = &v
	return s
}

func (s *WriteFeatureViewTableRequest) SetPartitions(v map[string]map[string]interface{}) *WriteFeatureViewTableRequest {
	s.Partitions = v
	return s
}

func (s *WriteFeatureViewTableRequest) SetUrlDatasource(v *WriteFeatureViewTableRequestUrlDatasource) *WriteFeatureViewTableRequest {
	s.UrlDatasource = v
	return s
}

func (s *WriteFeatureViewTableRequest) Validate() error {
	return dara.Validate(s)
}

type WriteFeatureViewTableRequestUrlDatasource struct {
	// example:
	//
	// ,
	Delimiter *string `json:"Delimiter,omitempty" xml:"Delimiter,omitempty"`
	// example:
	//
	// true
	OmitHeader *bool `json:"OmitHeader,omitempty" xml:"OmitHeader,omitempty"`
	// example:
	//
	// xxx.xxx.com/file.csv
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s WriteFeatureViewTableRequestUrlDatasource) String() string {
	return dara.Prettify(s)
}

func (s WriteFeatureViewTableRequestUrlDatasource) GoString() string {
	return s.String()
}

func (s *WriteFeatureViewTableRequestUrlDatasource) GetDelimiter() *string {
	return s.Delimiter
}

func (s *WriteFeatureViewTableRequestUrlDatasource) GetOmitHeader() *bool {
	return s.OmitHeader
}

func (s *WriteFeatureViewTableRequestUrlDatasource) GetPath() *string {
	return s.Path
}

func (s *WriteFeatureViewTableRequestUrlDatasource) SetDelimiter(v string) *WriteFeatureViewTableRequestUrlDatasource {
	s.Delimiter = &v
	return s
}

func (s *WriteFeatureViewTableRequestUrlDatasource) SetOmitHeader(v bool) *WriteFeatureViewTableRequestUrlDatasource {
	s.OmitHeader = &v
	return s
}

func (s *WriteFeatureViewTableRequestUrlDatasource) SetPath(v string) *WriteFeatureViewTableRequestUrlDatasource {
	s.Path = &v
	return s
}

func (s *WriteFeatureViewTableRequestUrlDatasource) Validate() error {
	return dara.Validate(s)
}
