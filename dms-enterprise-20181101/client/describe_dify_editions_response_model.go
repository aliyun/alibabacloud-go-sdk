// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDifyEditionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDifyEditionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDifyEditionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDifyEditionsResponseBody) *DescribeDifyEditionsResponse
	GetBody() *DescribeDifyEditionsResponseBody
}

type DescribeDifyEditionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDifyEditionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDifyEditionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDifyEditionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDifyEditionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDifyEditionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDifyEditionsResponse) GetBody() *DescribeDifyEditionsResponseBody {
	return s.Body
}

func (s *DescribeDifyEditionsResponse) SetHeaders(v map[string]*string) *DescribeDifyEditionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDifyEditionsResponse) SetStatusCode(v int32) *DescribeDifyEditionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDifyEditionsResponse) SetBody(v *DescribeDifyEditionsResponseBody) *DescribeDifyEditionsResponse {
	s.Body = v
	return s
}

func (s *DescribeDifyEditionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
