// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewWuyingServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewWuyingServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewWuyingServerResponse
	GetStatusCode() *int32
	SetBody(v *RenewWuyingServerResponseBody) *RenewWuyingServerResponse
	GetBody() *RenewWuyingServerResponseBody
}

type RenewWuyingServerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewWuyingServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewWuyingServerResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewWuyingServerResponse) GoString() string {
	return s.String()
}

func (s *RenewWuyingServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewWuyingServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewWuyingServerResponse) GetBody() *RenewWuyingServerResponseBody {
	return s.Body
}

func (s *RenewWuyingServerResponse) SetHeaders(v map[string]*string) *RenewWuyingServerResponse {
	s.Headers = v
	return s
}

func (s *RenewWuyingServerResponse) SetStatusCode(v int32) *RenewWuyingServerResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewWuyingServerResponse) SetBody(v *RenewWuyingServerResponseBody) *RenewWuyingServerResponse {
	s.Body = v
	return s
}

func (s *RenewWuyingServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
