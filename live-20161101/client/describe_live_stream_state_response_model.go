// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamStateResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamStateResponseBody) *DescribeLiveStreamStateResponse
	GetBody() *DescribeLiveStreamStateResponseBody
}

type DescribeLiveStreamStateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamStateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamStateResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamStateResponse) GetBody() *DescribeLiveStreamStateResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamStateResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamStateResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamStateResponse) SetStatusCode(v int32) *DescribeLiveStreamStateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamStateResponse) SetBody(v *DescribeLiveStreamStateResponseBody) *DescribeLiveStreamStateResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamStateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
