// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLakebaseS3AccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLakebaseS3AccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLakebaseS3AccountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLakebaseS3AccountResponseBody) *DescribeLakebaseS3AccountResponse
	GetBody() *DescribeLakebaseS3AccountResponseBody
}

type DescribeLakebaseS3AccountResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLakebaseS3AccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLakebaseS3AccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakebaseS3AccountResponse) GoString() string {
	return s.String()
}

func (s *DescribeLakebaseS3AccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLakebaseS3AccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLakebaseS3AccountResponse) GetBody() *DescribeLakebaseS3AccountResponseBody {
	return s.Body
}

func (s *DescribeLakebaseS3AccountResponse) SetHeaders(v map[string]*string) *DescribeLakebaseS3AccountResponse {
	s.Headers = v
	return s
}

func (s *DescribeLakebaseS3AccountResponse) SetStatusCode(v int32) *DescribeLakebaseS3AccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLakebaseS3AccountResponse) SetBody(v *DescribeLakebaseS3AccountResponseBody) *DescribeLakebaseS3AccountResponse {
	s.Body = v
	return s
}

func (s *DescribeLakebaseS3AccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
