// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataSourceBloodRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v string) *DataSourceBloodRequest
	GetDataSourceId() *string
}

type DataSourceBloodRequest struct {
	// Data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 44051300991327000048
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
}

func (s DataSourceBloodRequest) String() string {
	return dara.Prettify(s)
}

func (s DataSourceBloodRequest) GoString() string {
	return s.String()
}

func (s *DataSourceBloodRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *DataSourceBloodRequest) SetDataSourceId(v string) *DataSourceBloodRequest {
	s.DataSourceId = &v
	return s
}

func (s *DataSourceBloodRequest) Validate() error {
	return dara.Validate(s)
}
