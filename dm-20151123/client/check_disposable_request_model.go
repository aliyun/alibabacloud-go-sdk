// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDisposableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *CheckDisposableRequest
	GetEmail() *string
}

type CheckDisposableRequest struct {
	// example:
	//
	// Txxxxm@xxxx.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
}

func (s CheckDisposableRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDisposableRequest) GoString() string {
	return s.String()
}

func (s *CheckDisposableRequest) GetEmail() *string {
	return s.Email
}

func (s *CheckDisposableRequest) SetEmail(v string) *CheckDisposableRequest {
	s.Email = &v
	return s
}

func (s *CheckDisposableRequest) Validate() error {
	return dara.Validate(s)
}
