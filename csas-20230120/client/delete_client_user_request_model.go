// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteClientUserRequest
	GetId() *string
}

type DeleteClientUserRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 27058
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteClientUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteClientUserRequest) GetId() *string {
	return s.Id
}

func (s *DeleteClientUserRequest) SetId(v string) *DeleteClientUserRequest {
	s.Id = &v
	return s
}

func (s *DeleteClientUserRequest) Validate() error {
	return dara.Validate(s)
}
