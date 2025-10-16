// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseExpiredInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseExpiredInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseExpiredInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseExpiredInstanceResponseBody) *ReleaseExpiredInstanceResponse
	GetBody() *ReleaseExpiredInstanceResponseBody
}

type ReleaseExpiredInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseExpiredInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseExpiredInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseExpiredInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseExpiredInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseExpiredInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseExpiredInstanceResponse) GetBody() *ReleaseExpiredInstanceResponseBody {
	return s.Body
}

func (s *ReleaseExpiredInstanceResponse) SetHeaders(v map[string]*string) *ReleaseExpiredInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseExpiredInstanceResponse) SetStatusCode(v int32) *ReleaseExpiredInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseExpiredInstanceResponse) SetBody(v *ReleaseExpiredInstanceResponseBody) *ReleaseExpiredInstanceResponse {
	s.Body = v
	return s
}

func (s *ReleaseExpiredInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
