// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveRequestsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSensitiveRequestsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSensitiveRequestsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSensitiveRequestsResponseBody) *DescribeSensitiveRequestsResponse
	GetBody() *DescribeSensitiveRequestsResponseBody
}

type DescribeSensitiveRequestsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSensitiveRequestsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSensitiveRequestsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveRequestsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveRequestsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSensitiveRequestsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSensitiveRequestsResponse) GetBody() *DescribeSensitiveRequestsResponseBody {
	return s.Body
}

func (s *DescribeSensitiveRequestsResponse) SetHeaders(v map[string]*string) *DescribeSensitiveRequestsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSensitiveRequestsResponse) SetStatusCode(v int32) *DescribeSensitiveRequestsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSensitiveRequestsResponse) SetBody(v *DescribeSensitiveRequestsResponseBody) *DescribeSensitiveRequestsResponse {
	s.Body = v
	return s
}

func (s *DescribeSensitiveRequestsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
