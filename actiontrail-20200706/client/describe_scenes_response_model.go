// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScenesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScenesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScenesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScenesResponseBody) *DescribeScenesResponse
	GetBody() *DescribeScenesResponseBody
}

type DescribeScenesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScenesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScenesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScenesResponse) GoString() string {
	return s.String()
}

func (s *DescribeScenesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScenesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScenesResponse) GetBody() *DescribeScenesResponseBody {
	return s.Body
}

func (s *DescribeScenesResponse) SetHeaders(v map[string]*string) *DescribeScenesResponse {
	s.Headers = v
	return s
}

func (s *DescribeScenesResponse) SetStatusCode(v int32) *DescribeScenesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScenesResponse) SetBody(v *DescribeScenesResponseBody) *DescribeScenesResponse {
	s.Body = v
	return s
}

func (s *DescribeScenesResponse) Validate() error {
	return dara.Validate(s)
}
