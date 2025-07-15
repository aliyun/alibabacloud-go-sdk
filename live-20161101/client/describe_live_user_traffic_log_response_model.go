// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveUserTrafficLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveUserTrafficLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveUserTrafficLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveUserTrafficLogResponseBody) *DescribeLiveUserTrafficLogResponse
	GetBody() *DescribeLiveUserTrafficLogResponseBody
}

type DescribeLiveUserTrafficLogResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveUserTrafficLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveUserTrafficLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserTrafficLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserTrafficLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveUserTrafficLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveUserTrafficLogResponse) GetBody() *DescribeLiveUserTrafficLogResponseBody {
	return s.Body
}

func (s *DescribeLiveUserTrafficLogResponse) SetHeaders(v map[string]*string) *DescribeLiveUserTrafficLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveUserTrafficLogResponse) SetStatusCode(v int32) *DescribeLiveUserTrafficLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveUserTrafficLogResponse) SetBody(v *DescribeLiveUserTrafficLogResponseBody) *DescribeLiveUserTrafficLogResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveUserTrafficLogResponse) Validate() error {
	return dara.Validate(s)
}
