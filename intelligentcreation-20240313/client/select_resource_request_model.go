// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdempotentId(v string) *SelectResourceRequest
	GetIdempotentId() *string
}

type SelectResourceRequest struct {
	// example:
	//
	// 1111
	IdempotentId *string `json:"idempotentId,omitempty" xml:"idempotentId,omitempty"`
}

func (s SelectResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s SelectResourceRequest) GoString() string {
	return s.String()
}

func (s *SelectResourceRequest) GetIdempotentId() *string {
	return s.IdempotentId
}

func (s *SelectResourceRequest) SetIdempotentId(v string) *SelectResourceRequest {
	s.IdempotentId = &v
	return s
}

func (s *SelectResourceRequest) Validate() error {
	return dara.Validate(s)
}
