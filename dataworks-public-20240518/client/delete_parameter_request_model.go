// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParameterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteParameterRequest
	GetId() *int64
}

type DeleteParameterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteParameterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteParameterRequest) GoString() string {
	return s.String()
}

func (s *DeleteParameterRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteParameterRequest) SetId(v int64) *DeleteParameterRequest {
	s.Id = &v
	return s
}

func (s *DeleteParameterRequest) Validate() error {
	return dara.Validate(s)
}
