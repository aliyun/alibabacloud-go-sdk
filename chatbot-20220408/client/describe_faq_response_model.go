// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaqResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFaqResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFaqResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFaqResponseBody) *DescribeFaqResponse
	GetBody() *DescribeFaqResponseBody
}

type DescribeFaqResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFaqResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFaqResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaqResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaqResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFaqResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFaqResponse) GetBody() *DescribeFaqResponseBody {
	return s.Body
}

func (s *DescribeFaqResponse) SetHeaders(v map[string]*string) *DescribeFaqResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaqResponse) SetStatusCode(v int32) *DescribeFaqResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFaqResponse) SetBody(v *DescribeFaqResponseBody) *DescribeFaqResponse {
	s.Body = v
	return s
}

func (s *DescribeFaqResponse) Validate() error {
	return dara.Validate(s)
}
