// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceUuid(v string) *DeleteAgenticDataSourceRequest
	GetDataSourceUuid() *string
}

type DeleteAgenticDataSourceRequest struct {
	// The UUID of the data source to delete. Ensure that no workloads are using the data source before deletion.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou_a1b2c3d4e5f6
	DataSourceUuid *string `json:"DataSourceUuid,omitempty" xml:"DataSourceUuid,omitempty"`
}

func (s DeleteAgenticDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgenticDataSourceRequest) GetDataSourceUuid() *string {
	return s.DataSourceUuid
}

func (s *DeleteAgenticDataSourceRequest) SetDataSourceUuid(v string) *DeleteAgenticDataSourceRequest {
	s.DataSourceUuid = &v
	return s
}

func (s *DeleteAgenticDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
