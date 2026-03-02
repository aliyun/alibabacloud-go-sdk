// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComponentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListComponentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListComponentsResponse
	GetStatusCode() *int32
	SetBody(v *ComponentPageResult) *ListComponentsResponse
	GetBody() *ComponentPageResult
}

type ListComponentsResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ComponentPageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListComponentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListComponentsResponse) GoString() string {
	return s.String()
}

func (s *ListComponentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListComponentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListComponentsResponse) GetBody() *ComponentPageResult {
	return s.Body
}

func (s *ListComponentsResponse) SetHeaders(v map[string]*string) *ListComponentsResponse {
	s.Headers = v
	return s
}

func (s *ListComponentsResponse) SetStatusCode(v int32) *ListComponentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListComponentsResponse) SetBody(v *ComponentPageResult) *ListComponentsResponse {
	s.Body = v
	return s
}

func (s *ListComponentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
