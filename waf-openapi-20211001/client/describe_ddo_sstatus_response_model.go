// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDDoSStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDDoSStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDDoSStatusResponseBody) *DescribeDDoSStatusResponse
	GetBody() *DescribeDDoSStatusResponseBody
}

type DescribeDDoSStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDDoSStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDDoSStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDoSStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDDoSStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDDoSStatusResponse) GetBody() *DescribeDDoSStatusResponseBody {
	return s.Body
}

func (s *DescribeDDoSStatusResponse) SetHeaders(v map[string]*string) *DescribeDDoSStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDoSStatusResponse) SetStatusCode(v int32) *DescribeDDoSStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDDoSStatusResponse) SetBody(v *DescribeDDoSStatusResponseBody) *DescribeDDoSStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeDDoSStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
