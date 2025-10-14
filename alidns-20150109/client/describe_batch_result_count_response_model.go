// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBatchResultCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBatchResultCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBatchResultCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBatchResultCountResponseBody) *DescribeBatchResultCountResponse
	GetBody() *DescribeBatchResultCountResponseBody
}

type DescribeBatchResultCountResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBatchResultCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBatchResultCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchResultCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeBatchResultCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBatchResultCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBatchResultCountResponse) GetBody() *DescribeBatchResultCountResponseBody {
	return s.Body
}

func (s *DescribeBatchResultCountResponse) SetHeaders(v map[string]*string) *DescribeBatchResultCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeBatchResultCountResponse) SetStatusCode(v int32) *DescribeBatchResultCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBatchResultCountResponse) SetBody(v *DescribeBatchResultCountResponseBody) *DescribeBatchResultCountResponse {
	s.Body = v
	return s
}

func (s *DescribeBatchResultCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
