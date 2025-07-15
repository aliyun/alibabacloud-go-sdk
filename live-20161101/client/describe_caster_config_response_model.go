// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCasterConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCasterConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCasterConfigResponseBody) *DescribeCasterConfigResponse
	GetBody() *DescribeCasterConfigResponseBody
}

type DescribeCasterConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCasterConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCasterConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeCasterConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCasterConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCasterConfigResponse) GetBody() *DescribeCasterConfigResponseBody {
	return s.Body
}

func (s *DescribeCasterConfigResponse) SetHeaders(v map[string]*string) *DescribeCasterConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeCasterConfigResponse) SetStatusCode(v int32) *DescribeCasterConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCasterConfigResponse) SetBody(v *DescribeCasterConfigResponseBody) *DescribeCasterConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeCasterConfigResponse) Validate() error {
	return dara.Validate(s)
}
