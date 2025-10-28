// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAliasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAliasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAliasResponse
	GetStatusCode() *int32
	SetBody(v *Alias) *CreateAliasResponse
	GetBody() *Alias
}

type CreateAliasResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Alias             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAliasResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAliasResponse) GoString() string {
	return s.String()
}

func (s *CreateAliasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAliasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAliasResponse) GetBody() *Alias {
	return s.Body
}

func (s *CreateAliasResponse) SetHeaders(v map[string]*string) *CreateAliasResponse {
	s.Headers = v
	return s
}

func (s *CreateAliasResponse) SetStatusCode(v int32) *CreateAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAliasResponse) SetBody(v *Alias) *CreateAliasResponse {
	s.Body = v
	return s
}

func (s *CreateAliasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
