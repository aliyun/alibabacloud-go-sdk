// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicesForPrivateAccessApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPolicesForPrivateAccessApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPolicesForPrivateAccessApplicationResponse
	GetStatusCode() *int32
	SetBody(v *ListPolicesForPrivateAccessApplicationResponseBody) *ListPolicesForPrivateAccessApplicationResponse
	GetBody() *ListPolicesForPrivateAccessApplicationResponseBody
}

type ListPolicesForPrivateAccessApplicationResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPolicesForPrivateAccessApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPolicesForPrivateAccessApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPolicesForPrivateAccessApplicationResponse) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPolicesForPrivateAccessApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPolicesForPrivateAccessApplicationResponse) GetBody() *ListPolicesForPrivateAccessApplicationResponseBody {
	return s.Body
}

func (s *ListPolicesForPrivateAccessApplicationResponse) SetHeaders(v map[string]*string) *ListPolicesForPrivateAccessApplicationResponse {
	s.Headers = v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponse) SetStatusCode(v int32) *ListPolicesForPrivateAccessApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponse) SetBody(v *ListPolicesForPrivateAccessApplicationResponseBody) *ListPolicesForPrivateAccessApplicationResponse {
	s.Body = v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
