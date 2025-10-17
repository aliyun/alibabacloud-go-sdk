// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDSEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDSEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDSEntityResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDSEntityResponseBody) *DescribeDSEntityResponse
	GetBody() *DescribeDSEntityResponseBody
}

type DescribeDSEntityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDSEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDSEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDSEntityResponse) GoString() string {
	return s.String()
}

func (s *DescribeDSEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDSEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDSEntityResponse) GetBody() *DescribeDSEntityResponseBody {
	return s.Body
}

func (s *DescribeDSEntityResponse) SetHeaders(v map[string]*string) *DescribeDSEntityResponse {
	s.Headers = v
	return s
}

func (s *DescribeDSEntityResponse) SetStatusCode(v int32) *DescribeDSEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDSEntityResponse) SetBody(v *DescribeDSEntityResponseBody) *DescribeDSEntityResponse {
	s.Body = v
	return s
}

func (s *DescribeDSEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
