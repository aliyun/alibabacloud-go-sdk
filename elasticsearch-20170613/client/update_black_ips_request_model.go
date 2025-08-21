// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBlackIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateBlackIpsRequest
	GetClientToken() *string
}

type UpdateBlackIpsRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateBlackIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBlackIpsRequest) GoString() string {
	return s.String()
}

func (s *UpdateBlackIpsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateBlackIpsRequest) SetClientToken(v string) *UpdateBlackIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateBlackIpsRequest) Validate() error {
	return dara.Validate(s)
}
