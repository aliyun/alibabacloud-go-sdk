// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTokenResult interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTokenResult
	GetRequestId() *string
	SetToken(v string) *CreateTokenResult
	GetToken() *string
}

type CreateTokenResult struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Token     *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s CreateTokenResult) String() string {
	return dara.Prettify(s)
}

func (s CreateTokenResult) GoString() string {
	return s.String()
}

func (s *CreateTokenResult) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTokenResult) GetToken() *string {
	return s.Token
}

func (s *CreateTokenResult) SetRequestId(v string) *CreateTokenResult {
	s.RequestId = &v
	return s
}

func (s *CreateTokenResult) SetToken(v string) *CreateTokenResult {
	s.Token = &v
	return s
}

func (s *CreateTokenResult) Validate() error {
	return dara.Validate(s)
}
