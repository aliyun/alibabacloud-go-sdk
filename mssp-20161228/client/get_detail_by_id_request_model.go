// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDetailByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetDetailByIdRequest
	GetId() *int64
}

type GetDetailByIdRequest struct {
	// Primary key ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDetailByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDetailByIdRequest) GoString() string {
	return s.String()
}

func (s *GetDetailByIdRequest) GetId() *int64 {
	return s.Id
}

func (s *GetDetailByIdRequest) SetId(v int64) *GetDetailByIdRequest {
	s.Id = &v
	return s
}

func (s *GetDetailByIdRequest) Validate() error {
	return dara.Validate(s)
}
