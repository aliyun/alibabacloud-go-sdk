// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamCountResponseBody) *DescribeLiveStreamCountResponse
	GetBody() *DescribeLiveStreamCountResponseBody
}

type DescribeLiveStreamCountResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamCountResponse) GetBody() *DescribeLiveStreamCountResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamCountResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamCountResponse) SetStatusCode(v int32) *DescribeLiveStreamCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamCountResponse) SetBody(v *DescribeLiveStreamCountResponseBody) *DescribeLiveStreamCountResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
