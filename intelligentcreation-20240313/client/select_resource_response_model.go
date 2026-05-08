// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SelectResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SelectResourceResponse
	GetStatusCode() *int32
	SetBody(v *SelectResourceResponseBody) *SelectResourceResponse
	GetBody() *SelectResourceResponseBody
}

type SelectResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SelectResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SelectResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s SelectResourceResponse) GoString() string {
	return s.String()
}

func (s *SelectResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SelectResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SelectResourceResponse) GetBody() *SelectResourceResponseBody {
	return s.Body
}

func (s *SelectResourceResponse) SetHeaders(v map[string]*string) *SelectResourceResponse {
	s.Headers = v
	return s
}

func (s *SelectResourceResponse) SetStatusCode(v int32) *SelectResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *SelectResourceResponse) SetBody(v *SelectResourceResponseBody) *SelectResourceResponse {
	s.Body = v
	return s
}

func (s *SelectResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
