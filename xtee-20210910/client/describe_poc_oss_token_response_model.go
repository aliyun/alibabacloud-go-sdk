// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePocOssTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePocOssTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePocOssTokenResponse
	GetStatusCode() *int32
	SetBody(v *DescribePocOssTokenResponseBody) *DescribePocOssTokenResponse
	GetBody() *DescribePocOssTokenResponseBody
}

type DescribePocOssTokenResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePocOssTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePocOssTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePocOssTokenResponse) GoString() string {
	return s.String()
}

func (s *DescribePocOssTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePocOssTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePocOssTokenResponse) GetBody() *DescribePocOssTokenResponseBody {
	return s.Body
}

func (s *DescribePocOssTokenResponse) SetHeaders(v map[string]*string) *DescribePocOssTokenResponse {
	s.Headers = v
	return s
}

func (s *DescribePocOssTokenResponse) SetStatusCode(v int32) *DescribePocOssTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePocOssTokenResponse) SetBody(v *DescribePocOssTokenResponseBody) *DescribePocOssTokenResponse {
	s.Body = v
	return s
}

func (s *DescribePocOssTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
