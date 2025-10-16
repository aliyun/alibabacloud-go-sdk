// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBundlesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBundlesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBundlesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBundlesResponseBody) *DescribeBundlesResponse
	GetBody() *DescribeBundlesResponseBody
}

type DescribeBundlesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBundlesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBundlesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBundlesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBundlesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBundlesResponse) GetBody() *DescribeBundlesResponseBody {
	return s.Body
}

func (s *DescribeBundlesResponse) SetHeaders(v map[string]*string) *DescribeBundlesResponse {
	s.Headers = v
	return s
}

func (s *DescribeBundlesResponse) SetStatusCode(v int32) *DescribeBundlesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBundlesResponse) SetBody(v *DescribeBundlesResponseBody) *DescribeBundlesResponse {
	s.Body = v
	return s
}

func (s *DescribeBundlesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
