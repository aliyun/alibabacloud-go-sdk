// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitialTranscodeConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitialTranscodeConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitialTranscodeConfigResponse
	GetStatusCode() *int32
	SetBody(v *InitialTranscodeConfigResponseBody) *InitialTranscodeConfigResponse
	GetBody() *InitialTranscodeConfigResponseBody
}

type InitialTranscodeConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitialTranscodeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitialTranscodeConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s InitialTranscodeConfigResponse) GoString() string {
	return s.String()
}

func (s *InitialTranscodeConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitialTranscodeConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitialTranscodeConfigResponse) GetBody() *InitialTranscodeConfigResponseBody {
	return s.Body
}

func (s *InitialTranscodeConfigResponse) SetHeaders(v map[string]*string) *InitialTranscodeConfigResponse {
	s.Headers = v
	return s
}

func (s *InitialTranscodeConfigResponse) SetStatusCode(v int32) *InitialTranscodeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *InitialTranscodeConfigResponse) SetBody(v *InitialTranscodeConfigResponseBody) *InitialTranscodeConfigResponse {
	s.Body = v
	return s
}

func (s *InitialTranscodeConfigResponse) Validate() error {
	return dara.Validate(s)
}
