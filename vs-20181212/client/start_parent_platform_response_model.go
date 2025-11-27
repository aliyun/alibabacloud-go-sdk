// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartParentPlatformResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartParentPlatformResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartParentPlatformResponse
	GetStatusCode() *int32
	SetBody(v *StartParentPlatformResponseBody) *StartParentPlatformResponse
	GetBody() *StartParentPlatformResponseBody
}

type StartParentPlatformResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartParentPlatformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartParentPlatformResponse) String() string {
	return dara.Prettify(s)
}

func (s StartParentPlatformResponse) GoString() string {
	return s.String()
}

func (s *StartParentPlatformResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartParentPlatformResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartParentPlatformResponse) GetBody() *StartParentPlatformResponseBody {
	return s.Body
}

func (s *StartParentPlatformResponse) SetHeaders(v map[string]*string) *StartParentPlatformResponse {
	s.Headers = v
	return s
}

func (s *StartParentPlatformResponse) SetStatusCode(v int32) *StartParentPlatformResponse {
	s.StatusCode = &v
	return s
}

func (s *StartParentPlatformResponse) SetBody(v *StartParentPlatformResponseBody) *StartParentPlatformResponse {
	s.Body = v
	return s
}

func (s *StartParentPlatformResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
