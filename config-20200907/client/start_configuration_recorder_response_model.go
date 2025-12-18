// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartConfigurationRecorderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartConfigurationRecorderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartConfigurationRecorderResponse
	GetStatusCode() *int32
	SetBody(v *StartConfigurationRecorderResponseBody) *StartConfigurationRecorderResponse
	GetBody() *StartConfigurationRecorderResponseBody
}

type StartConfigurationRecorderResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartConfigurationRecorderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartConfigurationRecorderResponse) String() string {
	return dara.Prettify(s)
}

func (s StartConfigurationRecorderResponse) GoString() string {
	return s.String()
}

func (s *StartConfigurationRecorderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartConfigurationRecorderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartConfigurationRecorderResponse) GetBody() *StartConfigurationRecorderResponseBody {
	return s.Body
}

func (s *StartConfigurationRecorderResponse) SetHeaders(v map[string]*string) *StartConfigurationRecorderResponse {
	s.Headers = v
	return s
}

func (s *StartConfigurationRecorderResponse) SetStatusCode(v int32) *StartConfigurationRecorderResponse {
	s.StatusCode = &v
	return s
}

func (s *StartConfigurationRecorderResponse) SetBody(v *StartConfigurationRecorderResponseBody) *StartConfigurationRecorderResponse {
	s.Body = v
	return s
}

func (s *StartConfigurationRecorderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
