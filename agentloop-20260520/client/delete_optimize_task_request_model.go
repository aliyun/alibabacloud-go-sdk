// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOptimizeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteOptimizeTaskRequest
	GetClientToken() *string
}

type DeleteOptimizeTaskRequest struct {
	// The idempotency token.
	//
	// example:
	//
	// a1b2c3d4-1234-5678-90ab-cdef12345678
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s DeleteOptimizeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOptimizeTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteOptimizeTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteOptimizeTaskRequest) SetClientToken(v string) *DeleteOptimizeTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteOptimizeTaskRequest) Validate() error {
	return dara.Validate(s)
}
