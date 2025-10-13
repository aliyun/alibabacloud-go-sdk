// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebApplicationRevisionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWebApplicationRevisionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWebApplicationRevisionsResponse
	GetStatusCode() *int32
	SetBody(v *ListWebApplicationRevisionsBody) *ListWebApplicationRevisionsResponse
	GetBody() *ListWebApplicationRevisionsBody
}

type ListWebApplicationRevisionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWebApplicationRevisionsBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWebApplicationRevisionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWebApplicationRevisionsResponse) GoString() string {
	return s.String()
}

func (s *ListWebApplicationRevisionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWebApplicationRevisionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWebApplicationRevisionsResponse) GetBody() *ListWebApplicationRevisionsBody {
	return s.Body
}

func (s *ListWebApplicationRevisionsResponse) SetHeaders(v map[string]*string) *ListWebApplicationRevisionsResponse {
	s.Headers = v
	return s
}

func (s *ListWebApplicationRevisionsResponse) SetStatusCode(v int32) *ListWebApplicationRevisionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWebApplicationRevisionsResponse) SetBody(v *ListWebApplicationRevisionsBody) *ListWebApplicationRevisionsResponse {
	s.Body = v
	return s
}

func (s *ListWebApplicationRevisionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
