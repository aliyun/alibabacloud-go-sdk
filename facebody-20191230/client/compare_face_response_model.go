// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareFaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompareFaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompareFaceResponse
	GetStatusCode() *int32
	SetBody(v *CompareFaceResponseBody) *CompareFaceResponse
	GetBody() *CompareFaceResponseBody
}

type CompareFaceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompareFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompareFaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CompareFaceResponse) GoString() string {
	return s.String()
}

func (s *CompareFaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompareFaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompareFaceResponse) GetBody() *CompareFaceResponseBody {
	return s.Body
}

func (s *CompareFaceResponse) SetHeaders(v map[string]*string) *CompareFaceResponse {
	s.Headers = v
	return s
}

func (s *CompareFaceResponse) SetStatusCode(v int32) *CompareFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CompareFaceResponse) SetBody(v *CompareFaceResponseBody) *CompareFaceResponse {
	s.Body = v
	return s
}

func (s *CompareFaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
