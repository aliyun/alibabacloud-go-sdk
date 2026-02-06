// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTTSDemoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTTSDemoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTTSDemoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTTSDemoResponseBody) *DescribeTTSDemoResponse
	GetBody() *DescribeTTSDemoResponseBody
}

type DescribeTTSDemoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTTSDemoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTTSDemoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTTSDemoResponse) GoString() string {
	return s.String()
}

func (s *DescribeTTSDemoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTTSDemoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTTSDemoResponse) GetBody() *DescribeTTSDemoResponseBody {
	return s.Body
}

func (s *DescribeTTSDemoResponse) SetHeaders(v map[string]*string) *DescribeTTSDemoResponse {
	s.Headers = v
	return s
}

func (s *DescribeTTSDemoResponse) SetStatusCode(v int32) *DescribeTTSDemoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTTSDemoResponse) SetBody(v *DescribeTTSDemoResponseBody) *DescribeTTSDemoResponse {
	s.Body = v
	return s
}

func (s *DescribeTTSDemoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
