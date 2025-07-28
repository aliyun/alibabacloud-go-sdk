// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCertsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCertsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCertsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCertsResponseBody) *DescribeCertsResponse
	GetBody() *DescribeCertsResponseBody
}

type DescribeCertsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCertsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCertsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCertsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCertsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCertsResponse) GetBody() *DescribeCertsResponseBody {
	return s.Body
}

func (s *DescribeCertsResponse) SetHeaders(v map[string]*string) *DescribeCertsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCertsResponse) SetStatusCode(v int32) *DescribeCertsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCertsResponse) SetBody(v *DescribeCertsResponseBody) *DescribeCertsResponse {
	s.Body = v
	return s
}

func (s *DescribeCertsResponse) Validate() error {
	return dara.Validate(s)
}
