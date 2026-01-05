// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesForTagOptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourcesForTagOptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourcesForTagOptionResponse
	GetStatusCode() *int32
	SetBody(v *ListResourcesForTagOptionResponseBody) *ListResourcesForTagOptionResponse
	GetBody() *ListResourcesForTagOptionResponseBody
}

type ListResourcesForTagOptionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcesForTagOptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcesForTagOptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesForTagOptionResponse) GoString() string {
	return s.String()
}

func (s *ListResourcesForTagOptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourcesForTagOptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourcesForTagOptionResponse) GetBody() *ListResourcesForTagOptionResponseBody {
	return s.Body
}

func (s *ListResourcesForTagOptionResponse) SetHeaders(v map[string]*string) *ListResourcesForTagOptionResponse {
	s.Headers = v
	return s
}

func (s *ListResourcesForTagOptionResponse) SetStatusCode(v int32) *ListResourcesForTagOptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcesForTagOptionResponse) SetBody(v *ListResourcesForTagOptionResponseBody) *ListResourcesForTagOptionResponse {
	s.Body = v
	return s
}

func (s *ListResourcesForTagOptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
