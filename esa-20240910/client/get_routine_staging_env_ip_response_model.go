// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineStagingEnvIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRoutineStagingEnvIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRoutineStagingEnvIpResponse
	GetStatusCode() *int32
	SetBody(v *GetRoutineStagingEnvIpResponseBody) *GetRoutineStagingEnvIpResponse
	GetBody() *GetRoutineStagingEnvIpResponseBody
}

type GetRoutineStagingEnvIpResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRoutineStagingEnvIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRoutineStagingEnvIpResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineStagingEnvIpResponse) GoString() string {
	return s.String()
}

func (s *GetRoutineStagingEnvIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRoutineStagingEnvIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRoutineStagingEnvIpResponse) GetBody() *GetRoutineStagingEnvIpResponseBody {
	return s.Body
}

func (s *GetRoutineStagingEnvIpResponse) SetHeaders(v map[string]*string) *GetRoutineStagingEnvIpResponse {
	s.Headers = v
	return s
}

func (s *GetRoutineStagingEnvIpResponse) SetStatusCode(v int32) *GetRoutineStagingEnvIpResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoutineStagingEnvIpResponse) SetBody(v *GetRoutineStagingEnvIpResponseBody) *GetRoutineStagingEnvIpResponse {
	s.Body = v
	return s
}

func (s *GetRoutineStagingEnvIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
