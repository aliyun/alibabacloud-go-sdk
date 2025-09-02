// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v int64) *DeleteDataSourceRequest
	GetDataSourceId() *int64
}

type DeleteDataSourceRequest struct {
	// The data source ID. You can call the [ListDataSources](https://help.aliyun.com/document_detail/211431.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
}

func (s DeleteDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceRequest) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *DeleteDataSourceRequest) SetDataSourceId(v int64) *DeleteDataSourceRequest {
	s.DataSourceId = &v
	return s
}

func (s *DeleteDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
