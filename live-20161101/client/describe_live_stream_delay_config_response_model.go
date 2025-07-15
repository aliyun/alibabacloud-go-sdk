// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamDelayConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamDelayConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamDelayConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamDelayConfigResponseBody) *DescribeLiveStreamDelayConfigResponse
	GetBody() *DescribeLiveStreamDelayConfigResponseBody
}

type DescribeLiveStreamDelayConfigResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamDelayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamDelayConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamDelayConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamDelayConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamDelayConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamDelayConfigResponse) GetBody() *DescribeLiveStreamDelayConfigResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamDelayConfigResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamDelayConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamDelayConfigResponse) SetStatusCode(v int32) *DescribeLiveStreamDelayConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamDelayConfigResponse) SetBody(v *DescribeLiveStreamDelayConfigResponseBody) *DescribeLiveStreamDelayConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamDelayConfigResponse) Validate() error {
	return dara.Validate(s)
}
