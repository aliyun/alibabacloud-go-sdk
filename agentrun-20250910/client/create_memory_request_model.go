// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLongTtl(v int32) *CreateMemoryRequest
	GetLongTtl() *int32
	SetName(v string) *CreateMemoryRequest
	GetName() *string
	SetShortTtl(v int32) *CreateMemoryRequest
	GetShortTtl() *int32
}

type CreateMemoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 365
	LongTtl *int32 `json:"longTtl,omitempty" xml:"longTtl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-memory
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30
	ShortTtl *int32 `json:"shortTtl,omitempty" xml:"shortTtl,omitempty"`
}

func (s CreateMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryRequest) GoString() string {
	return s.String()
}

func (s *CreateMemoryRequest) GetLongTtl() *int32 {
	return s.LongTtl
}

func (s *CreateMemoryRequest) GetName() *string {
	return s.Name
}

func (s *CreateMemoryRequest) GetShortTtl() *int32 {
	return s.ShortTtl
}

func (s *CreateMemoryRequest) SetLongTtl(v int32) *CreateMemoryRequest {
	s.LongTtl = &v
	return s
}

func (s *CreateMemoryRequest) SetName(v string) *CreateMemoryRequest {
	s.Name = &v
	return s
}

func (s *CreateMemoryRequest) SetShortTtl(v int32) *CreateMemoryRequest {
	s.ShortTtl = &v
	return s
}

func (s *CreateMemoryRequest) Validate() error {
	return dara.Validate(s)
}
