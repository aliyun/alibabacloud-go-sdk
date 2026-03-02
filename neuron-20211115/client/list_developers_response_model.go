// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDevelopersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDevelopersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDevelopersResponse
	GetStatusCode() *int32
	SetBody(v *DeveloperPageResult) *ListDevelopersResponse
	GetBody() *DeveloperPageResult
}

type ListDevelopersResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeveloperPageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDevelopersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDevelopersResponse) GoString() string {
	return s.String()
}

func (s *ListDevelopersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDevelopersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDevelopersResponse) GetBody() *DeveloperPageResult {
	return s.Body
}

func (s *ListDevelopersResponse) SetHeaders(v map[string]*string) *ListDevelopersResponse {
	s.Headers = v
	return s
}

func (s *ListDevelopersResponse) SetStatusCode(v int32) *ListDevelopersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDevelopersResponse) SetBody(v *DeveloperPageResult) *ListDevelopersResponse {
	s.Body = v
	return s
}

func (s *ListDevelopersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
