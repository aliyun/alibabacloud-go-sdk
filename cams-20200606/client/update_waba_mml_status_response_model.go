// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWabaMmlStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWabaMmlStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWabaMmlStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWabaMmlStatusResponseBody) *UpdateWabaMmlStatusResponse
	GetBody() *UpdateWabaMmlStatusResponseBody
}

type UpdateWabaMmlStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWabaMmlStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWabaMmlStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWabaMmlStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateWabaMmlStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWabaMmlStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWabaMmlStatusResponse) GetBody() *UpdateWabaMmlStatusResponseBody {
	return s.Body
}

func (s *UpdateWabaMmlStatusResponse) SetHeaders(v map[string]*string) *UpdateWabaMmlStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateWabaMmlStatusResponse) SetStatusCode(v int32) *UpdateWabaMmlStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWabaMmlStatusResponse) SetBody(v *UpdateWabaMmlStatusResponseBody) *UpdateWabaMmlStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateWabaMmlStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
