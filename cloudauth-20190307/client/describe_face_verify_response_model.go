// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaceVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFaceVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFaceVerifyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFaceVerifyResponseBody) *DescribeFaceVerifyResponse
	GetBody() *DescribeFaceVerifyResponseBody
}

type DescribeFaceVerifyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFaceVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaceVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFaceVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFaceVerifyResponse) GetBody() *DescribeFaceVerifyResponseBody {
	return s.Body
}

func (s *DescribeFaceVerifyResponse) SetHeaders(v map[string]*string) *DescribeFaceVerifyResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaceVerifyResponse) SetStatusCode(v int32) *DescribeFaceVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFaceVerifyResponse) SetBody(v *DescribeFaceVerifyResponseBody) *DescribeFaceVerifyResponse {
	s.Body = v
	return s
}

func (s *DescribeFaceVerifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
