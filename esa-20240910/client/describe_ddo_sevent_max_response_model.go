// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSEventMaxResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDDoSEventMaxResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDDoSEventMaxResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDDoSEventMaxResponseBody) *DescribeDDoSEventMaxResponse
	GetBody() *DescribeDDoSEventMaxResponseBody
}

type DescribeDDoSEventMaxResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDDoSEventMaxResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDDoSEventMaxResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSEventMaxResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDoSEventMaxResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDDoSEventMaxResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDDoSEventMaxResponse) GetBody() *DescribeDDoSEventMaxResponseBody {
	return s.Body
}

func (s *DescribeDDoSEventMaxResponse) SetHeaders(v map[string]*string) *DescribeDDoSEventMaxResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDoSEventMaxResponse) SetStatusCode(v int32) *DescribeDDoSEventMaxResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDDoSEventMaxResponse) SetBody(v *DescribeDDoSEventMaxResponseBody) *DescribeDDoSEventMaxResponse {
	s.Body = v
	return s
}

func (s *DescribeDDoSEventMaxResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
