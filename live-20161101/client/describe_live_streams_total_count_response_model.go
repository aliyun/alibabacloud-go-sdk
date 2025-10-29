// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamsTotalCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamsTotalCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamsTotalCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamsTotalCountResponseBody) *DescribeLiveStreamsTotalCountResponse
	GetBody() *DescribeLiveStreamsTotalCountResponseBody
}

type DescribeLiveStreamsTotalCountResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamsTotalCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamsTotalCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsTotalCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsTotalCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamsTotalCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamsTotalCountResponse) GetBody() *DescribeLiveStreamsTotalCountResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamsTotalCountResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamsTotalCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamsTotalCountResponse) SetStatusCode(v int32) *DescribeLiveStreamsTotalCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamsTotalCountResponse) SetBody(v *DescribeLiveStreamsTotalCountResponseBody) *DescribeLiveStreamsTotalCountResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamsTotalCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
