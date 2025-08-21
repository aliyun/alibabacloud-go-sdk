// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDdosCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDdosCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDdosCountResponseBody) *DescribeDdosCountResponse
	GetBody() *DescribeDdosCountResponseBody
}

type DescribeDdosCountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDdosCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDdosCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDdosCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDdosCountResponse) GetBody() *DescribeDdosCountResponseBody {
	return s.Body
}

func (s *DescribeDdosCountResponse) SetHeaders(v map[string]*string) *DescribeDdosCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosCountResponse) SetStatusCode(v int32) *DescribeDdosCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDdosCountResponse) SetBody(v *DescribeDdosCountResponseBody) *DescribeDdosCountResponse {
	s.Body = v
	return s
}

func (s *DescribeDdosCountResponse) Validate() error {
	return dara.Validate(s)
}
