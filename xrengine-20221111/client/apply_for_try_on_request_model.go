// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyForTryOnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *ApplyForTryOnRequest
	GetContent() *string
	SetJwtToken(v string) *ApplyForTryOnRequest
	GetJwtToken() *string
}

type ApplyForTryOnRequest struct {
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s ApplyForTryOnRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyForTryOnRequest) GoString() string {
	return s.String()
}

func (s *ApplyForTryOnRequest) GetContent() *string {
	return s.Content
}

func (s *ApplyForTryOnRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *ApplyForTryOnRequest) SetContent(v string) *ApplyForTryOnRequest {
	s.Content = &v
	return s
}

func (s *ApplyForTryOnRequest) SetJwtToken(v string) *ApplyForTryOnRequest {
	s.JwtToken = &v
	return s
}

func (s *ApplyForTryOnRequest) Validate() error {
	return dara.Validate(s)
}
