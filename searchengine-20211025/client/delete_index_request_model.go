// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSource(v string) *DeleteIndexRequest
	GetDataSource() *string
	SetDeleteDataSource(v bool) *DeleteIndexRequest
	GetDeleteDataSource() *bool
}

type DeleteIndexRequest struct {
	// The data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// ha-cn-pl32rf0js04_test
	DataSource *string `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
	// Specifies whether to delete the data source.
	//
	// example:
	//
	// true
	DeleteDataSource *bool `json:"deleteDataSource,omitempty" xml:"deleteDataSource,omitempty"`
}

func (s DeleteIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIndexRequest) GoString() string {
	return s.String()
}

func (s *DeleteIndexRequest) GetDataSource() *string {
	return s.DataSource
}

func (s *DeleteIndexRequest) GetDeleteDataSource() *bool {
	return s.DeleteDataSource
}

func (s *DeleteIndexRequest) SetDataSource(v string) *DeleteIndexRequest {
	s.DataSource = &v
	return s
}

func (s *DeleteIndexRequest) SetDeleteDataSource(v bool) *DeleteIndexRequest {
	s.DeleteDataSource = &v
	return s
}

func (s *DeleteIndexRequest) Validate() error {
	return dara.Validate(s)
}
