// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIntegratedServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIntegratedServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIntegratedServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIntegratedServiceStatusResponseBody) *UpdateIntegratedServiceStatusResponse
	GetBody() *UpdateIntegratedServiceStatusResponseBody
}

type UpdateIntegratedServiceStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIntegratedServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIntegratedServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIntegratedServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateIntegratedServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIntegratedServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIntegratedServiceStatusResponse) GetBody() *UpdateIntegratedServiceStatusResponseBody {
	return s.Body
}

func (s *UpdateIntegratedServiceStatusResponse) SetHeaders(v map[string]*string) *UpdateIntegratedServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateIntegratedServiceStatusResponse) SetStatusCode(v int32) *UpdateIntegratedServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIntegratedServiceStatusResponse) SetBody(v *UpdateIntegratedServiceStatusResponseBody) *UpdateIntegratedServiceStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateIntegratedServiceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
