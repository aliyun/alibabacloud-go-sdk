// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyCoordinationWithCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyCoordinationWithCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyCoordinationWithCodeResponse
	GetStatusCode() *int32
	SetBody(v *ApplyCoordinationWithCodeResponseBody) *ApplyCoordinationWithCodeResponse
	GetBody() *ApplyCoordinationWithCodeResponseBody
}

type ApplyCoordinationWithCodeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyCoordinationWithCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyCoordinationWithCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyCoordinationWithCodeResponse) GoString() string {
	return s.String()
}

func (s *ApplyCoordinationWithCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyCoordinationWithCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyCoordinationWithCodeResponse) GetBody() *ApplyCoordinationWithCodeResponseBody {
	return s.Body
}

func (s *ApplyCoordinationWithCodeResponse) SetHeaders(v map[string]*string) *ApplyCoordinationWithCodeResponse {
	s.Headers = v
	return s
}

func (s *ApplyCoordinationWithCodeResponse) SetStatusCode(v int32) *ApplyCoordinationWithCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponse) SetBody(v *ApplyCoordinationWithCodeResponseBody) *ApplyCoordinationWithCodeResponse {
	s.Body = v
	return s
}

func (s *ApplyCoordinationWithCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
