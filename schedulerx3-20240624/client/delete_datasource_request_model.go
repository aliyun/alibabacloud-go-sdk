// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteDatasourceRequest
	GetClusterId() *string
	SetDatasourceId(v int64) *DeleteDatasourceRequest
	GetDatasourceId() *int64
}

type DeleteDatasourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 195
	DatasourceId *int64 `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
}

func (s DeleteDatasourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatasourceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteDatasourceRequest) GetDatasourceId() *int64 {
	return s.DatasourceId
}

func (s *DeleteDatasourceRequest) SetClusterId(v string) *DeleteDatasourceRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteDatasourceRequest) SetDatasourceId(v int64) *DeleteDatasourceRequest {
	s.DatasourceId = &v
	return s
}

func (s *DeleteDatasourceRequest) Validate() error {
	return dara.Validate(s)
}
