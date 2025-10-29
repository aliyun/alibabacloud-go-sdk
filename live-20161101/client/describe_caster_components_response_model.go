// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterComponentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCasterComponentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCasterComponentsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCasterComponentsResponseBody) *DescribeCasterComponentsResponse
	GetBody() *DescribeCasterComponentsResponseBody
}

type DescribeCasterComponentsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCasterComponentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCasterComponentsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterComponentsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCasterComponentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCasterComponentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCasterComponentsResponse) GetBody() *DescribeCasterComponentsResponseBody {
	return s.Body
}

func (s *DescribeCasterComponentsResponse) SetHeaders(v map[string]*string) *DescribeCasterComponentsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCasterComponentsResponse) SetStatusCode(v int32) *DescribeCasterComponentsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCasterComponentsResponse) SetBody(v *DescribeCasterComponentsResponseBody) *DescribeCasterComponentsResponse {
	s.Body = v
	return s
}

func (s *DescribeCasterComponentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
