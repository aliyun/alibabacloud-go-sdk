// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProdCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProdCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProdCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProdCountResponseBody) *DescribeProdCountResponse
	GetBody() *DescribeProdCountResponseBody
}

type DescribeProdCountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProdCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProdCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProdCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeProdCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProdCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProdCountResponse) GetBody() *DescribeProdCountResponseBody {
	return s.Body
}

func (s *DescribeProdCountResponse) SetHeaders(v map[string]*string) *DescribeProdCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeProdCountResponse) SetStatusCode(v int32) *DescribeProdCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProdCountResponse) SetBody(v *DescribeProdCountResponseBody) *DescribeProdCountResponse {
	s.Body = v
	return s
}

func (s *DescribeProdCountResponse) Validate() error {
	return dara.Validate(s)
}
