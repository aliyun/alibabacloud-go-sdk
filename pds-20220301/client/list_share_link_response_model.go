// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListShareLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListShareLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListShareLinkResponse
	GetStatusCode() *int32
	SetBody(v *ListShareLinkResponseBody) *ListShareLinkResponse
	GetBody() *ListShareLinkResponseBody
}

type ListShareLinkResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListShareLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListShareLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s ListShareLinkResponse) GoString() string {
	return s.String()
}

func (s *ListShareLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListShareLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListShareLinkResponse) GetBody() *ListShareLinkResponseBody {
	return s.Body
}

func (s *ListShareLinkResponse) SetHeaders(v map[string]*string) *ListShareLinkResponse {
	s.Headers = v
	return s
}

func (s *ListShareLinkResponse) SetStatusCode(v int32) *ListShareLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *ListShareLinkResponse) SetBody(v *ListShareLinkResponseBody) *ListShareLinkResponse {
	s.Body = v
	return s
}

func (s *ListShareLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
