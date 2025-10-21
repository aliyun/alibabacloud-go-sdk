// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLongTtl(v int32) *UpdateMemoryRequest
	GetLongTtl() *int32
	SetPermanent(v bool) *UpdateMemoryRequest
	GetPermanent() *bool
	SetShortTtl(v int32) *UpdateMemoryRequest
	GetShortTtl() *int32
	SetStrategy(v []*string) *UpdateMemoryRequest
	GetStrategy() []*string
}

type UpdateMemoryRequest struct {
	// example:
	//
	// 365
	LongTtl *int32 `json:"longTtl,omitempty" xml:"longTtl,omitempty"`
	// example:
	//
	// true
	Permanent *bool `json:"permanent,omitempty" xml:"permanent,omitempty"`
	// example:
	//
	// 30
	ShortTtl *int32    `json:"shortTtl,omitempty" xml:"shortTtl,omitempty"`
	Strategy []*string `json:"strategy,omitempty" xml:"strategy,omitempty" type:"Repeated"`
}

func (s UpdateMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemoryRequest) GetLongTtl() *int32 {
	return s.LongTtl
}

func (s *UpdateMemoryRequest) GetPermanent() *bool {
	return s.Permanent
}

func (s *UpdateMemoryRequest) GetShortTtl() *int32 {
	return s.ShortTtl
}

func (s *UpdateMemoryRequest) GetStrategy() []*string {
	return s.Strategy
}

func (s *UpdateMemoryRequest) SetLongTtl(v int32) *UpdateMemoryRequest {
	s.LongTtl = &v
	return s
}

func (s *UpdateMemoryRequest) SetPermanent(v bool) *UpdateMemoryRequest {
	s.Permanent = &v
	return s
}

func (s *UpdateMemoryRequest) SetShortTtl(v int32) *UpdateMemoryRequest {
	s.ShortTtl = &v
	return s
}

func (s *UpdateMemoryRequest) SetStrategy(v []*string) *UpdateMemoryRequest {
	s.Strategy = v
	return s
}

func (s *UpdateMemoryRequest) Validate() error {
	return dara.Validate(s)
}
