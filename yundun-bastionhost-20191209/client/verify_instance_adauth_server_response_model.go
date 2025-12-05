// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyInstanceADAuthServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyInstanceADAuthServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyInstanceADAuthServerResponse
	GetStatusCode() *int32
	SetBody(v *VerifyInstanceADAuthServerResponseBody) *VerifyInstanceADAuthServerResponse
	GetBody() *VerifyInstanceADAuthServerResponseBody
}

type VerifyInstanceADAuthServerResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyInstanceADAuthServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyInstanceADAuthServerResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyInstanceADAuthServerResponse) GoString() string {
	return s.String()
}

func (s *VerifyInstanceADAuthServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyInstanceADAuthServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyInstanceADAuthServerResponse) GetBody() *VerifyInstanceADAuthServerResponseBody {
	return s.Body
}

func (s *VerifyInstanceADAuthServerResponse) SetHeaders(v map[string]*string) *VerifyInstanceADAuthServerResponse {
	s.Headers = v
	return s
}

func (s *VerifyInstanceADAuthServerResponse) SetStatusCode(v int32) *VerifyInstanceADAuthServerResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyInstanceADAuthServerResponse) SetBody(v *VerifyInstanceADAuthServerResponseBody) *VerifyInstanceADAuthServerResponse {
	s.Body = v
	return s
}

func (s *VerifyInstanceADAuthServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
