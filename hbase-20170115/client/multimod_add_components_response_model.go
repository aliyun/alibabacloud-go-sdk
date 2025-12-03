// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultimodAddComponentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MultimodAddComponentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MultimodAddComponentsResponse
	GetStatusCode() *int32
	SetBody(v *MultimodAddComponentsResponseBody) *MultimodAddComponentsResponse
	GetBody() *MultimodAddComponentsResponseBody
}

type MultimodAddComponentsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MultimodAddComponentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MultimodAddComponentsResponse) String() string {
	return dara.Prettify(s)
}

func (s MultimodAddComponentsResponse) GoString() string {
	return s.String()
}

func (s *MultimodAddComponentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MultimodAddComponentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MultimodAddComponentsResponse) GetBody() *MultimodAddComponentsResponseBody {
	return s.Body
}

func (s *MultimodAddComponentsResponse) SetHeaders(v map[string]*string) *MultimodAddComponentsResponse {
	s.Headers = v
	return s
}

func (s *MultimodAddComponentsResponse) SetStatusCode(v int32) *MultimodAddComponentsResponse {
	s.StatusCode = &v
	return s
}

func (s *MultimodAddComponentsResponse) SetBody(v *MultimodAddComponentsResponseBody) *MultimodAddComponentsResponse {
	s.Body = v
	return s
}

func (s *MultimodAddComponentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
