// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAvatarResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdempotentId(v string) *QueryAvatarResourceRequest
	GetIdempotentId() *string
}

type QueryAvatarResourceRequest struct {
	// example:
	//
	// 11111
	IdempotentId *string `json:"idempotentId,omitempty" xml:"idempotentId,omitempty"`
}

func (s QueryAvatarResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAvatarResourceRequest) GoString() string {
	return s.String()
}

func (s *QueryAvatarResourceRequest) GetIdempotentId() *string {
	return s.IdempotentId
}

func (s *QueryAvatarResourceRequest) SetIdempotentId(v string) *QueryAvatarResourceRequest {
	s.IdempotentId = &v
	return s
}

func (s *QueryAvatarResourceRequest) Validate() error {
	return dara.Validate(s)
}
