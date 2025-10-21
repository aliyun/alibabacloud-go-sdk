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
	SetPermanent(v bool) *CreateMemoryRequest
	GetPermanent() *bool
	SetShortTtl(v int32) *CreateMemoryRequest
	GetShortTtl() *int32
	SetStrategy(v []*string) *CreateMemoryRequest
	GetStrategy() []*string
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
	// example:
	//
	// false
	Permanent *bool `json:"permanent,omitempty" xml:"permanent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30
	ShortTtl *int32 `json:"shortTtl,omitempty" xml:"shortTtl,omitempty"`
	// example:
	//
	// [\"Preference\", \"Facts\"]
	Strategy []*string `json:"strategy,omitempty" xml:"strategy,omitempty" type:"Repeated"`
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

func (s *CreateMemoryRequest) GetPermanent() *bool {
	return s.Permanent
}

func (s *CreateMemoryRequest) GetShortTtl() *int32 {
	return s.ShortTtl
}

func (s *CreateMemoryRequest) GetStrategy() []*string {
	return s.Strategy
}

func (s *CreateMemoryRequest) SetLongTtl(v int32) *CreateMemoryRequest {
	s.LongTtl = &v
	return s
}

func (s *CreateMemoryRequest) SetName(v string) *CreateMemoryRequest {
	s.Name = &v
	return s
}

func (s *CreateMemoryRequest) SetPermanent(v bool) *CreateMemoryRequest {
	s.Permanent = &v
	return s
}

func (s *CreateMemoryRequest) SetShortTtl(v int32) *CreateMemoryRequest {
	s.ShortTtl = &v
	return s
}

func (s *CreateMemoryRequest) SetStrategy(v []*string) *CreateMemoryRequest {
	s.Strategy = v
	return s
}

func (s *CreateMemoryRequest) Validate() error {
	return dara.Validate(s)
}
