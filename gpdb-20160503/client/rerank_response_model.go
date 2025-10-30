// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerankResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RerankResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RerankResponse
	GetStatusCode() *int32
	SetBody(v *RerankResponseBody) *RerankResponse
	GetBody() *RerankResponseBody
}

type RerankResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RerankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RerankResponse) String() string {
	return dara.Prettify(s)
}

func (s RerankResponse) GoString() string {
	return s.String()
}

func (s *RerankResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RerankResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RerankResponse) GetBody() *RerankResponseBody {
	return s.Body
}

func (s *RerankResponse) SetHeaders(v map[string]*string) *RerankResponse {
	s.Headers = v
	return s
}

func (s *RerankResponse) SetStatusCode(v int32) *RerankResponse {
	s.StatusCode = &v
	return s
}

func (s *RerankResponse) SetBody(v *RerankResponseBody) *RerankResponse {
	s.Body = v
	return s
}

func (s *RerankResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
