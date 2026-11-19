// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseExpiredPrepaidInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseExpiredPrepaidInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseExpiredPrepaidInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseExpiredPrepaidInstanceResponseBody) *ReleaseExpiredPrepaidInstanceResponse
	GetBody() *ReleaseExpiredPrepaidInstanceResponseBody
}

type ReleaseExpiredPrepaidInstanceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseExpiredPrepaidInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseExpiredPrepaidInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseExpiredPrepaidInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseExpiredPrepaidInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseExpiredPrepaidInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseExpiredPrepaidInstanceResponse) GetBody() *ReleaseExpiredPrepaidInstanceResponseBody {
	return s.Body
}

func (s *ReleaseExpiredPrepaidInstanceResponse) SetHeaders(v map[string]*string) *ReleaseExpiredPrepaidInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseExpiredPrepaidInstanceResponse) SetStatusCode(v int32) *ReleaseExpiredPrepaidInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseExpiredPrepaidInstanceResponse) SetBody(v *ReleaseExpiredPrepaidInstanceResponseBody) *ReleaseExpiredPrepaidInstanceResponse {
	s.Body = v
	return s
}

func (s *ReleaseExpiredPrepaidInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
