// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdaptorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAdaptorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAdaptorResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAdaptorResponseBody) *UpdateAdaptorResponse
	GetBody() *UpdateAdaptorResponseBody
}

type UpdateAdaptorResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAdaptorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAdaptorResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdaptorResponse) GoString() string {
	return s.String()
}

func (s *UpdateAdaptorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAdaptorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAdaptorResponse) GetBody() *UpdateAdaptorResponseBody {
	return s.Body
}

func (s *UpdateAdaptorResponse) SetHeaders(v map[string]*string) *UpdateAdaptorResponse {
	s.Headers = v
	return s
}

func (s *UpdateAdaptorResponse) SetStatusCode(v int32) *UpdateAdaptorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAdaptorResponse) SetBody(v *UpdateAdaptorResponseBody) *UpdateAdaptorResponse {
	s.Body = v
	return s
}

func (s *UpdateAdaptorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
