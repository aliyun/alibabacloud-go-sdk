// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMmAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMmAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMmAppResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMmAppResponseBody) *DescribeMmAppResponse
	GetBody() *DescribeMmAppResponseBody
}

type DescribeMmAppResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMmAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMmAppResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMmAppResponse) GoString() string {
	return s.String()
}

func (s *DescribeMmAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMmAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMmAppResponse) GetBody() *DescribeMmAppResponseBody {
	return s.Body
}

func (s *DescribeMmAppResponse) SetHeaders(v map[string]*string) *DescribeMmAppResponse {
	s.Headers = v
	return s
}

func (s *DescribeMmAppResponse) SetStatusCode(v int32) *DescribeMmAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMmAppResponse) SetBody(v *DescribeMmAppResponseBody) *DescribeMmAppResponse {
	s.Body = v
	return s
}

func (s *DescribeMmAppResponse) Validate() error {
	return dara.Validate(s)
}
