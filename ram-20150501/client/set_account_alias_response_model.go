// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAccountAliasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetAccountAliasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetAccountAliasResponse
	GetStatusCode() *int32
	SetBody(v *SetAccountAliasResponseBody) *SetAccountAliasResponse
	GetBody() *SetAccountAliasResponseBody
}

type SetAccountAliasResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAccountAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAccountAliasResponse) String() string {
	return dara.Prettify(s)
}

func (s SetAccountAliasResponse) GoString() string {
	return s.String()
}

func (s *SetAccountAliasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetAccountAliasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetAccountAliasResponse) GetBody() *SetAccountAliasResponseBody {
	return s.Body
}

func (s *SetAccountAliasResponse) SetHeaders(v map[string]*string) *SetAccountAliasResponse {
	s.Headers = v
	return s
}

func (s *SetAccountAliasResponse) SetStatusCode(v int32) *SetAccountAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAccountAliasResponse) SetBody(v *SetAccountAliasResponseBody) *SetAccountAliasResponse {
	s.Body = v
	return s
}

func (s *SetAccountAliasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
