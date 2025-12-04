// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynDbsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSynDbsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSynDbsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSynDbsResponseBody) *DescribeSynDbsResponse
	GetBody() *DescribeSynDbsResponseBody
}

type DescribeSynDbsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSynDbsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSynDbsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynDbsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSynDbsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSynDbsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSynDbsResponse) GetBody() *DescribeSynDbsResponseBody {
	return s.Body
}

func (s *DescribeSynDbsResponse) SetHeaders(v map[string]*string) *DescribeSynDbsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSynDbsResponse) SetStatusCode(v int32) *DescribeSynDbsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSynDbsResponse) SetBody(v *DescribeSynDbsResponseBody) *DescribeSynDbsResponse {
	s.Body = v
	return s
}

func (s *DescribeSynDbsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
