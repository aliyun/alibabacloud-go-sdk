// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetContentRequest
	GetId() *int64
	SetType(v string) *GetContentRequest
	GetType() *string
}

type GetContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// song
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetContentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetContentRequest) GoString() string {
	return s.String()
}

func (s *GetContentRequest) GetId() *int64 {
	return s.Id
}

func (s *GetContentRequest) GetType() *string {
	return s.Type
}

func (s *GetContentRequest) SetId(v int64) *GetContentRequest {
	s.Id = &v
	return s
}

func (s *GetContentRequest) SetType(v string) *GetContentRequest {
	s.Type = &v
	return s
}

func (s *GetContentRequest) Validate() error {
	return dara.Validate(s)
}
