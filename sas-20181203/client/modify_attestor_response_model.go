// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAttestorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAttestorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAttestorResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAttestorResponseBody) *ModifyAttestorResponse
	GetBody() *ModifyAttestorResponseBody
}

type ModifyAttestorResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAttestorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAttestorResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAttestorResponse) GoString() string {
	return s.String()
}

func (s *ModifyAttestorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAttestorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAttestorResponse) GetBody() *ModifyAttestorResponseBody {
	return s.Body
}

func (s *ModifyAttestorResponse) SetHeaders(v map[string]*string) *ModifyAttestorResponse {
	s.Headers = v
	return s
}

func (s *ModifyAttestorResponse) SetStatusCode(v int32) *ModifyAttestorResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAttestorResponse) SetBody(v *ModifyAttestorResponseBody) *ModifyAttestorResponse {
	s.Body = v
	return s
}

func (s *ModifyAttestorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
