// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDelayConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDelayConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDelayConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDelayConfigResponseBody) *DescribeLiveDelayConfigResponse
	GetBody() *DescribeLiveDelayConfigResponseBody
}

type DescribeLiveDelayConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDelayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDelayConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDelayConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDelayConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDelayConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDelayConfigResponse) GetBody() *DescribeLiveDelayConfigResponseBody {
	return s.Body
}

func (s *DescribeLiveDelayConfigResponse) SetHeaders(v map[string]*string) *DescribeLiveDelayConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDelayConfigResponse) SetStatusCode(v int32) *DescribeLiveDelayConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDelayConfigResponse) SetBody(v *DescribeLiveDelayConfigResponseBody) *DescribeLiveDelayConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDelayConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
