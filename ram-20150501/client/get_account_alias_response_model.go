// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountAliasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccountAliasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccountAliasResponse
	GetStatusCode() *int32
	SetBody(v *GetAccountAliasResponseBody) *GetAccountAliasResponse
	GetBody() *GetAccountAliasResponseBody
}

type GetAccountAliasResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountAliasResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccountAliasResponse) GoString() string {
	return s.String()
}

func (s *GetAccountAliasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccountAliasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccountAliasResponse) GetBody() *GetAccountAliasResponseBody {
	return s.Body
}

func (s *GetAccountAliasResponse) SetHeaders(v map[string]*string) *GetAccountAliasResponse {
	s.Headers = v
	return s
}

func (s *GetAccountAliasResponse) SetStatusCode(v int32) *GetAccountAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountAliasResponse) SetBody(v *GetAccountAliasResponseBody) *GetAccountAliasResponse {
	s.Body = v
	return s
}

func (s *GetAccountAliasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
