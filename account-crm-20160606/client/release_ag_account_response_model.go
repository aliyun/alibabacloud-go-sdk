// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseAgAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseAgAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseAgAccountResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseAgAccountResponseBody) *ReleaseAgAccountResponse
	GetBody() *ReleaseAgAccountResponseBody
}

type ReleaseAgAccountResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseAgAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseAgAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseAgAccountResponse) GoString() string {
	return s.String()
}

func (s *ReleaseAgAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseAgAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseAgAccountResponse) GetBody() *ReleaseAgAccountResponseBody {
	return s.Body
}

func (s *ReleaseAgAccountResponse) SetHeaders(v map[string]*string) *ReleaseAgAccountResponse {
	s.Headers = v
	return s
}

func (s *ReleaseAgAccountResponse) SetStatusCode(v int32) *ReleaseAgAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseAgAccountResponse) SetBody(v *ReleaseAgAccountResponseBody) *ReleaseAgAccountResponse {
	s.Body = v
	return s
}

func (s *ReleaseAgAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
