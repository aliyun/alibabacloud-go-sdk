// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClearAuthInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClearAuthInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClearAuthInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClearAuthInfoResponseBody) *DescribeClearAuthInfoResponse
	GetBody() *DescribeClearAuthInfoResponseBody
}

type DescribeClearAuthInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClearAuthInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClearAuthInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClearAuthInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeClearAuthInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClearAuthInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClearAuthInfoResponse) GetBody() *DescribeClearAuthInfoResponseBody {
	return s.Body
}

func (s *DescribeClearAuthInfoResponse) SetHeaders(v map[string]*string) *DescribeClearAuthInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeClearAuthInfoResponse) SetStatusCode(v int32) *DescribeClearAuthInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClearAuthInfoResponse) SetBody(v *DescribeClearAuthInfoResponseBody) *DescribeClearAuthInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeClearAuthInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
