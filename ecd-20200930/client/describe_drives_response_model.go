// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrivesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDrivesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDrivesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDrivesResponseBody) *DescribeDrivesResponse
	GetBody() *DescribeDrivesResponseBody
}

type DescribeDrivesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrivesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrivesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrivesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrivesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDrivesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDrivesResponse) GetBody() *DescribeDrivesResponseBody {
	return s.Body
}

func (s *DescribeDrivesResponse) SetHeaders(v map[string]*string) *DescribeDrivesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrivesResponse) SetStatusCode(v int32) *DescribeDrivesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrivesResponse) SetBody(v *DescribeDrivesResponseBody) *DescribeDrivesResponse {
	s.Body = v
	return s
}

func (s *DescribeDrivesResponse) Validate() error {
	return dara.Validate(s)
}
