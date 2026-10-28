// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePxfsAccessTokensResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePxfsAccessTokensResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePxfsAccessTokensResponse
	GetStatusCode() *int32
	SetBody(v *DescribePxfsAccessTokensResponseBody) *DescribePxfsAccessTokensResponse
	GetBody() *DescribePxfsAccessTokensResponseBody
}

type DescribePxfsAccessTokensResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePxfsAccessTokensResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePxfsAccessTokensResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsAccessTokensResponse) GoString() string {
	return s.String()
}

func (s *DescribePxfsAccessTokensResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePxfsAccessTokensResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePxfsAccessTokensResponse) GetBody() *DescribePxfsAccessTokensResponseBody {
	return s.Body
}

func (s *DescribePxfsAccessTokensResponse) SetHeaders(v map[string]*string) *DescribePxfsAccessTokensResponse {
	s.Headers = v
	return s
}

func (s *DescribePxfsAccessTokensResponse) SetStatusCode(v int32) *DescribePxfsAccessTokensResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePxfsAccessTokensResponse) SetBody(v *DescribePxfsAccessTokensResponseBody) *DescribePxfsAccessTokensResponse {
	s.Body = v
	return s
}

func (s *DescribePxfsAccessTokensResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
