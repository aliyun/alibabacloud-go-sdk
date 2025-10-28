// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAliasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAliasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAliasResponse
	GetStatusCode() *int32
	SetBody(v *Alias) *UpdateAliasResponse
	GetBody() *Alias
}

type UpdateAliasResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Alias             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAliasResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAliasResponse) GoString() string {
	return s.String()
}

func (s *UpdateAliasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAliasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAliasResponse) GetBody() *Alias {
	return s.Body
}

func (s *UpdateAliasResponse) SetHeaders(v map[string]*string) *UpdateAliasResponse {
	s.Headers = v
	return s
}

func (s *UpdateAliasResponse) SetStatusCode(v int32) *UpdateAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAliasResponse) SetBody(v *Alias) *UpdateAliasResponse {
	s.Body = v
	return s
}

func (s *UpdateAliasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
