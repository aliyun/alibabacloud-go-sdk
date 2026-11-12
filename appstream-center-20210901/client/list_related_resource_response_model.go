// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRelatedResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRelatedResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRelatedResourceResponse
	GetStatusCode() *int32
	SetBody(v *ListRelatedResourceResponseBody) *ListRelatedResourceResponse
	GetBody() *ListRelatedResourceResponseBody
}

type ListRelatedResourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRelatedResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRelatedResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRelatedResourceResponse) GoString() string {
	return s.String()
}

func (s *ListRelatedResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRelatedResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRelatedResourceResponse) GetBody() *ListRelatedResourceResponseBody {
	return s.Body
}

func (s *ListRelatedResourceResponse) SetHeaders(v map[string]*string) *ListRelatedResourceResponse {
	s.Headers = v
	return s
}

func (s *ListRelatedResourceResponse) SetStatusCode(v int32) *ListRelatedResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRelatedResourceResponse) SetBody(v *ListRelatedResourceResponseBody) *ListRelatedResourceResponse {
	s.Body = v
	return s
}

func (s *ListRelatedResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
