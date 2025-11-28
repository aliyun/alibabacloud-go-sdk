// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTotalStatInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTotalStatInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTotalStatInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTotalStatInfoResponseBody) *DescribeTotalStatInfoResponse
	GetBody() *DescribeTotalStatInfoResponseBody
}

type DescribeTotalStatInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTotalStatInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTotalStatInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTotalStatInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeTotalStatInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTotalStatInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTotalStatInfoResponse) GetBody() *DescribeTotalStatInfoResponseBody {
	return s.Body
}

func (s *DescribeTotalStatInfoResponse) SetHeaders(v map[string]*string) *DescribeTotalStatInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeTotalStatInfoResponse) SetStatusCode(v int32) *DescribeTotalStatInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTotalStatInfoResponse) SetBody(v *DescribeTotalStatInfoResponseBody) *DescribeTotalStatInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeTotalStatInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
