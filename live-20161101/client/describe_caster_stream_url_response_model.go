// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterStreamUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCasterStreamUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCasterStreamUrlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCasterStreamUrlResponseBody) *DescribeCasterStreamUrlResponse
	GetBody() *DescribeCasterStreamUrlResponseBody
}

type DescribeCasterStreamUrlResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCasterStreamUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCasterStreamUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterStreamUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeCasterStreamUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCasterStreamUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCasterStreamUrlResponse) GetBody() *DescribeCasterStreamUrlResponseBody {
	return s.Body
}

func (s *DescribeCasterStreamUrlResponse) SetHeaders(v map[string]*string) *DescribeCasterStreamUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeCasterStreamUrlResponse) SetStatusCode(v int32) *DescribeCasterStreamUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCasterStreamUrlResponse) SetBody(v *DescribeCasterStreamUrlResponseBody) *DescribeCasterStreamUrlResponse {
	s.Body = v
	return s
}

func (s *DescribeCasterStreamUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
