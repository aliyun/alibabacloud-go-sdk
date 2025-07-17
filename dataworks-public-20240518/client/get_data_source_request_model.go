// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetDataSourceRequest
	GetId() *int64
}

type GetDataSourceRequest struct {
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 16035
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceRequest) GoString() string {
	return s.String()
}

func (s *GetDataSourceRequest) GetId() *int64 {
	return s.Id
}

func (s *GetDataSourceRequest) SetId(v int64) *GetDataSourceRequest {
	s.Id = &v
	return s
}

func (s *GetDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
