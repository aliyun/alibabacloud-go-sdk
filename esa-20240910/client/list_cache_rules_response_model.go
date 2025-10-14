// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCacheRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCacheRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCacheRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListCacheRulesResponseBody) *ListCacheRulesResponse
	GetBody() *ListCacheRulesResponseBody
}

type ListCacheRulesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCacheRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCacheRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCacheRulesResponse) GoString() string {
	return s.String()
}

func (s *ListCacheRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCacheRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCacheRulesResponse) GetBody() *ListCacheRulesResponseBody {
	return s.Body
}

func (s *ListCacheRulesResponse) SetHeaders(v map[string]*string) *ListCacheRulesResponse {
	s.Headers = v
	return s
}

func (s *ListCacheRulesResponse) SetStatusCode(v int32) *ListCacheRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCacheRulesResponse) SetBody(v *ListCacheRulesResponseBody) *ListCacheRulesResponse {
	s.Body = v
	return s
}

func (s *ListCacheRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
