// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceADAuthServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceADAuthServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceADAuthServerResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceADAuthServerResponseBody) *ModifyInstanceADAuthServerResponse
	GetBody() *ModifyInstanceADAuthServerResponseBody
}

type ModifyInstanceADAuthServerResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceADAuthServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceADAuthServerResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceADAuthServerResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceADAuthServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceADAuthServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceADAuthServerResponse) GetBody() *ModifyInstanceADAuthServerResponseBody {
	return s.Body
}

func (s *ModifyInstanceADAuthServerResponse) SetHeaders(v map[string]*string) *ModifyInstanceADAuthServerResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceADAuthServerResponse) SetStatusCode(v int32) *ModifyInstanceADAuthServerResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceADAuthServerResponse) SetBody(v *ModifyInstanceADAuthServerResponseBody) *ModifyInstanceADAuthServerResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceADAuthServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
