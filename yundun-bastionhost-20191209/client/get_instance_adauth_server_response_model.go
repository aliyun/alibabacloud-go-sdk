// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceADAuthServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceADAuthServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceADAuthServerResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceADAuthServerResponseBody) *GetInstanceADAuthServerResponse
	GetBody() *GetInstanceADAuthServerResponseBody
}

type GetInstanceADAuthServerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceADAuthServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceADAuthServerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceADAuthServerResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceADAuthServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceADAuthServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceADAuthServerResponse) GetBody() *GetInstanceADAuthServerResponseBody {
	return s.Body
}

func (s *GetInstanceADAuthServerResponse) SetHeaders(v map[string]*string) *GetInstanceADAuthServerResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceADAuthServerResponse) SetStatusCode(v int32) *GetInstanceADAuthServerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceADAuthServerResponse) SetBody(v *GetInstanceADAuthServerResponseBody) *GetInstanceADAuthServerResponse {
	s.Body = v
	return s
}

func (s *GetInstanceADAuthServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
