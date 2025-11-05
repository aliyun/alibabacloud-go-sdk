// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLensServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLensServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLensServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLensServiceStatusResponseBody) *DescribeLensServiceStatusResponse
	GetBody() *DescribeLensServiceStatusResponseBody
}

type DescribeLensServiceStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLensServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLensServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLensServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeLensServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLensServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLensServiceStatusResponse) GetBody() *DescribeLensServiceStatusResponseBody {
	return s.Body
}

func (s *DescribeLensServiceStatusResponse) SetHeaders(v map[string]*string) *DescribeLensServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeLensServiceStatusResponse) SetStatusCode(v int32) *DescribeLensServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLensServiceStatusResponse) SetBody(v *DescribeLensServiceStatusResponseBody) *DescribeLensServiceStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeLensServiceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
