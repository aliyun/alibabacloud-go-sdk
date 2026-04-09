// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParameterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetParameterRequest
	GetId() *int64
}

type GetParameterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetParameterRequest) String() string {
	return dara.Prettify(s)
}

func (s GetParameterRequest) GoString() string {
	return s.String()
}

func (s *GetParameterRequest) GetId() *int64 {
	return s.Id
}

func (s *GetParameterRequest) SetId(v int64) *GetParameterRequest {
	s.Id = &v
	return s
}

func (s *GetParameterRequest) Validate() error {
	return dara.Validate(s)
}
