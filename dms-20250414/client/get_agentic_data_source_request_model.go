// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgenticDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceUuid(v string) *GetAgenticDataSourceRequest
	GetDataSourceUuid() *string
}

type GetAgenticDataSourceRequest struct {
	// The UUID of the data source to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou_a1b2c3d4e5f6
	DataSourceUuid *string `json:"DataSourceUuid,omitempty" xml:"DataSourceUuid,omitempty"`
}

func (s GetAgenticDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticDataSourceRequest) GoString() string {
	return s.String()
}

func (s *GetAgenticDataSourceRequest) GetDataSourceUuid() *string {
	return s.DataSourceUuid
}

func (s *GetAgenticDataSourceRequest) SetDataSourceUuid(v string) *GetAgenticDataSourceRequest {
	s.DataSourceUuid = &v
	return s
}

func (s *GetAgenticDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
