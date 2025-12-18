// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigurationRecorderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConfigurationRecorderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConfigurationRecorderResponse
	GetStatusCode() *int32
	SetBody(v *GetConfigurationRecorderResponseBody) *GetConfigurationRecorderResponse
	GetBody() *GetConfigurationRecorderResponseBody
}

type GetConfigurationRecorderResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConfigurationRecorderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConfigurationRecorderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConfigurationRecorderResponse) GoString() string {
	return s.String()
}

func (s *GetConfigurationRecorderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConfigurationRecorderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConfigurationRecorderResponse) GetBody() *GetConfigurationRecorderResponseBody {
	return s.Body
}

func (s *GetConfigurationRecorderResponse) SetHeaders(v map[string]*string) *GetConfigurationRecorderResponse {
	s.Headers = v
	return s
}

func (s *GetConfigurationRecorderResponse) SetStatusCode(v int32) *GetConfigurationRecorderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConfigurationRecorderResponse) SetBody(v *GetConfigurationRecorderResponseBody) *GetConfigurationRecorderResponse {
	s.Body = v
	return s
}

func (s *GetConfigurationRecorderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
