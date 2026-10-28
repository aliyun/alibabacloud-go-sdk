// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePxfsInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePxfsInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePxfsInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribePxfsInfoResponseBody) *DescribePxfsInfoResponse
	GetBody() *DescribePxfsInfoResponseBody
}

type DescribePxfsInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePxfsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePxfsInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribePxfsInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePxfsInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePxfsInfoResponse) GetBody() *DescribePxfsInfoResponseBody {
	return s.Body
}

func (s *DescribePxfsInfoResponse) SetHeaders(v map[string]*string) *DescribePxfsInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribePxfsInfoResponse) SetStatusCode(v int32) *DescribePxfsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePxfsInfoResponse) SetBody(v *DescribePxfsInfoResponseBody) *DescribePxfsInfoResponse {
	s.Body = v
	return s
}

func (s *DescribePxfsInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
