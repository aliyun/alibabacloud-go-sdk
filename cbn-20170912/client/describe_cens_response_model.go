// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCensResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCensResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCensResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCensResponseBody) *DescribeCensResponse
	GetBody() *DescribeCensResponseBody
}

type DescribeCensResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCensResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCensResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCensResponse) GoString() string {
	return s.String()
}

func (s *DescribeCensResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCensResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCensResponse) GetBody() *DescribeCensResponseBody {
	return s.Body
}

func (s *DescribeCensResponse) SetHeaders(v map[string]*string) *DescribeCensResponse {
	s.Headers = v
	return s
}

func (s *DescribeCensResponse) SetStatusCode(v int32) *DescribeCensResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCensResponse) SetBody(v *DescribeCensResponseBody) *DescribeCensResponse {
	s.Body = v
	return s
}

func (s *DescribeCensResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
