// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUsageQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UsageQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UsageQueryResponse
	GetStatusCode() *int32
	SetBody(v *UsageQueryResponseBody) *UsageQueryResponse
	GetBody() *UsageQueryResponseBody
}

type UsageQueryResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UsageQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UsageQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s UsageQueryResponse) GoString() string {
	return s.String()
}

func (s *UsageQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UsageQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UsageQueryResponse) GetBody() *UsageQueryResponseBody {
	return s.Body
}

func (s *UsageQueryResponse) SetHeaders(v map[string]*string) *UsageQueryResponse {
	s.Headers = v
	return s
}

func (s *UsageQueryResponse) SetStatusCode(v int32) *UsageQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *UsageQueryResponse) SetBody(v *UsageQueryResponseBody) *UsageQueryResponse {
	s.Body = v
	return s
}

func (s *UsageQueryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
