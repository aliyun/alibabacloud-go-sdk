// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAliasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindAliasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindAliasResponse
	GetStatusCode() *int32
	SetBody(v *BindAliasResponseBody) *BindAliasResponse
	GetBody() *BindAliasResponseBody
}

type BindAliasResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAliasResponse) String() string {
	return dara.Prettify(s)
}

func (s BindAliasResponse) GoString() string {
	return s.String()
}

func (s *BindAliasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindAliasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindAliasResponse) GetBody() *BindAliasResponseBody {
	return s.Body
}

func (s *BindAliasResponse) SetHeaders(v map[string]*string) *BindAliasResponse {
	s.Headers = v
	return s
}

func (s *BindAliasResponse) SetStatusCode(v int32) *BindAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAliasResponse) SetBody(v *BindAliasResponseBody) *BindAliasResponse {
	s.Body = v
	return s
}

func (s *BindAliasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
