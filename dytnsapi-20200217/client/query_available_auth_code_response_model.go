// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAvailableAuthCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAvailableAuthCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAvailableAuthCodeResponse
	GetStatusCode() *int32
	SetBody(v *QueryAvailableAuthCodeResponseBody) *QueryAvailableAuthCodeResponse
	GetBody() *QueryAvailableAuthCodeResponseBody
}

type QueryAvailableAuthCodeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAvailableAuthCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAvailableAuthCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAvailableAuthCodeResponse) GoString() string {
	return s.String()
}

func (s *QueryAvailableAuthCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAvailableAuthCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAvailableAuthCodeResponse) GetBody() *QueryAvailableAuthCodeResponseBody {
	return s.Body
}

func (s *QueryAvailableAuthCodeResponse) SetHeaders(v map[string]*string) *QueryAvailableAuthCodeResponse {
	s.Headers = v
	return s
}

func (s *QueryAvailableAuthCodeResponse) SetStatusCode(v int32) *QueryAvailableAuthCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAvailableAuthCodeResponse) SetBody(v *QueryAvailableAuthCodeResponseBody) *QueryAvailableAuthCodeResponse {
	s.Body = v
	return s
}

func (s *QueryAvailableAuthCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
