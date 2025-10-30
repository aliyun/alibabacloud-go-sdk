// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineBizEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OfflineBizEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OfflineBizEntityResponse
	GetStatusCode() *int32
	SetBody(v *OfflineBizEntityResponseBody) *OfflineBizEntityResponse
	GetBody() *OfflineBizEntityResponseBody
}

type OfflineBizEntityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineBizEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineBizEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s OfflineBizEntityResponse) GoString() string {
	return s.String()
}

func (s *OfflineBizEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OfflineBizEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OfflineBizEntityResponse) GetBody() *OfflineBizEntityResponseBody {
	return s.Body
}

func (s *OfflineBizEntityResponse) SetHeaders(v map[string]*string) *OfflineBizEntityResponse {
	s.Headers = v
	return s
}

func (s *OfflineBizEntityResponse) SetStatusCode(v int32) *OfflineBizEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineBizEntityResponse) SetBody(v *OfflineBizEntityResponseBody) *OfflineBizEntityResponse {
	s.Body = v
	return s
}

func (s *OfflineBizEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
