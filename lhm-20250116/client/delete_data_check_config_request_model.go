// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataCheckConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteDataCheckConfigRequest
	GetId() *int64
}

type DeleteDataCheckConfigRequest struct {
	// The configuration ID. You can obtain this ID by calling the GetDataCheckConfig operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s DeleteDataCheckConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataCheckConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataCheckConfigRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteDataCheckConfigRequest) SetId(v int64) *DeleteDataCheckConfigRequest {
	s.Id = &v
	return s
}

func (s *DeleteDataCheckConfigRequest) Validate() error {
	return dara.Validate(s)
}
