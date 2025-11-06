// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataSource interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v string) *DataSource
	GetDataSourceId() *string
	SetMountPath(v string) *DataSource
	GetMountPath() *string
	SetUri(v string) *DataSource
	GetUri() *string
}

type DataSource struct {
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	MountPath    *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	Uri          *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DataSource) String() string {
	return dara.Prettify(s)
}

func (s DataSource) GoString() string {
	return s.String()
}

func (s *DataSource) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *DataSource) GetMountPath() *string {
	return s.MountPath
}

func (s *DataSource) GetUri() *string {
	return s.Uri
}

func (s *DataSource) SetDataSourceId(v string) *DataSource {
	s.DataSourceId = &v
	return s
}

func (s *DataSource) SetMountPath(v string) *DataSource {
	s.MountPath = &v
	return s
}

func (s *DataSource) SetUri(v string) *DataSource {
	s.Uri = &v
	return s
}

func (s *DataSource) Validate() error {
	return dara.Validate(s)
}
