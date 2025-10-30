// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataDomainsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataDomainsResponseBody) *ListDataDomainsResponse
	GetBody() *ListDataDomainsResponseBody
}

type ListDataDomainsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListDataDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataDomainsResponse) GetBody() *ListDataDomainsResponseBody {
	return s.Body
}

func (s *ListDataDomainsResponse) SetHeaders(v map[string]*string) *ListDataDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListDataDomainsResponse) SetStatusCode(v int32) *ListDataDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataDomainsResponse) SetBody(v *ListDataDomainsResponseBody) *ListDataDomainsResponse {
	s.Body = v
	return s
}

func (s *ListDataDomainsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
