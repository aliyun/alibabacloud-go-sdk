// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartWuyingServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartWuyingServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartWuyingServerResponse
	GetStatusCode() *int32
	SetBody(v *StartWuyingServerResponseBody) *StartWuyingServerResponse
	GetBody() *StartWuyingServerResponseBody
}

type StartWuyingServerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartWuyingServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartWuyingServerResponse) String() string {
	return dara.Prettify(s)
}

func (s StartWuyingServerResponse) GoString() string {
	return s.String()
}

func (s *StartWuyingServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartWuyingServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartWuyingServerResponse) GetBody() *StartWuyingServerResponseBody {
	return s.Body
}

func (s *StartWuyingServerResponse) SetHeaders(v map[string]*string) *StartWuyingServerResponse {
	s.Headers = v
	return s
}

func (s *StartWuyingServerResponse) SetStatusCode(v int32) *StartWuyingServerResponse {
	s.StatusCode = &v
	return s
}

func (s *StartWuyingServerResponse) SetBody(v *StartWuyingServerResponseBody) *StartWuyingServerResponse {
	s.Body = v
	return s
}

func (s *StartWuyingServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
