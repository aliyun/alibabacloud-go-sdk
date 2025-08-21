// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCateInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *ListCateInfoRequest
	GetType() *string
}

type ListCateInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// song
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCateInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCateInfoRequest) GoString() string {
	return s.String()
}

func (s *ListCateInfoRequest) GetType() *string {
	return s.Type
}

func (s *ListCateInfoRequest) SetType(v string) *ListCateInfoRequest {
	s.Type = &v
	return s
}

func (s *ListCateInfoRequest) Validate() error {
	return dara.Validate(s)
}
