// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWuyingServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWuyingServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWuyingServerResponse
	GetStatusCode() *int32
	SetBody(v *CreateWuyingServerResponseBody) *CreateWuyingServerResponse
	GetBody() *CreateWuyingServerResponseBody
}

type CreateWuyingServerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWuyingServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWuyingServerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWuyingServerResponse) GoString() string {
	return s.String()
}

func (s *CreateWuyingServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWuyingServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWuyingServerResponse) GetBody() *CreateWuyingServerResponseBody {
	return s.Body
}

func (s *CreateWuyingServerResponse) SetHeaders(v map[string]*string) *CreateWuyingServerResponse {
	s.Headers = v
	return s
}

func (s *CreateWuyingServerResponse) SetStatusCode(v int32) *CreateWuyingServerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWuyingServerResponse) SetBody(v *CreateWuyingServerResponseBody) *CreateWuyingServerResponse {
	s.Body = v
	return s
}

func (s *CreateWuyingServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
