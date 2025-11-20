// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNavigationByFormTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNavigationByFormTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNavigationByFormTypeResponse
	GetStatusCode() *int32
	SetBody(v *ListNavigationByFormTypeResponseBody) *ListNavigationByFormTypeResponse
	GetBody() *ListNavigationByFormTypeResponseBody
}

type ListNavigationByFormTypeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNavigationByFormTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNavigationByFormTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNavigationByFormTypeResponse) GoString() string {
	return s.String()
}

func (s *ListNavigationByFormTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNavigationByFormTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNavigationByFormTypeResponse) GetBody() *ListNavigationByFormTypeResponseBody {
	return s.Body
}

func (s *ListNavigationByFormTypeResponse) SetHeaders(v map[string]*string) *ListNavigationByFormTypeResponse {
	s.Headers = v
	return s
}

func (s *ListNavigationByFormTypeResponse) SetStatusCode(v int32) *ListNavigationByFormTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNavigationByFormTypeResponse) SetBody(v *ListNavigationByFormTypeResponseBody) *ListNavigationByFormTypeResponse {
	s.Body = v
	return s
}

func (s *ListNavigationByFormTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
