// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAreaMapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAreaMapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAreaMapResponse
	GetStatusCode() *int32
	SetBody(v *ListAreaMapResponseBody) *ListAreaMapResponse
	GetBody() *ListAreaMapResponseBody
}

type ListAreaMapResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAreaMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAreaMapResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAreaMapResponse) GoString() string {
	return s.String()
}

func (s *ListAreaMapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAreaMapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAreaMapResponse) GetBody() *ListAreaMapResponseBody {
	return s.Body
}

func (s *ListAreaMapResponse) SetHeaders(v map[string]*string) *ListAreaMapResponse {
	s.Headers = v
	return s
}

func (s *ListAreaMapResponse) SetStatusCode(v int32) *ListAreaMapResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAreaMapResponse) SetBody(v *ListAreaMapResponseBody) *ListAreaMapResponse {
	s.Body = v
	return s
}

func (s *ListAreaMapResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
