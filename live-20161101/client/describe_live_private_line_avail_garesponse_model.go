// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePrivateLineAvailGAResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLivePrivateLineAvailGAResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLivePrivateLineAvailGAResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLivePrivateLineAvailGAResponseBody) *DescribeLivePrivateLineAvailGAResponse
	GetBody() *DescribeLivePrivateLineAvailGAResponseBody
}

type DescribeLivePrivateLineAvailGAResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLivePrivateLineAvailGAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLivePrivateLineAvailGAResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePrivateLineAvailGAResponse) GoString() string {
	return s.String()
}

func (s *DescribeLivePrivateLineAvailGAResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLivePrivateLineAvailGAResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLivePrivateLineAvailGAResponse) GetBody() *DescribeLivePrivateLineAvailGAResponseBody {
	return s.Body
}

func (s *DescribeLivePrivateLineAvailGAResponse) SetHeaders(v map[string]*string) *DescribeLivePrivateLineAvailGAResponse {
	s.Headers = v
	return s
}

func (s *DescribeLivePrivateLineAvailGAResponse) SetStatusCode(v int32) *DescribeLivePrivateLineAvailGAResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLivePrivateLineAvailGAResponse) SetBody(v *DescribeLivePrivateLineAvailGAResponseBody) *DescribeLivePrivateLineAvailGAResponse {
	s.Body = v
	return s
}

func (s *DescribeLivePrivateLineAvailGAResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
