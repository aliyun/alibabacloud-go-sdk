// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSaasServiceVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSaasServiceVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSaasServiceVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSaasServiceVersionResponseBody) *UpdateSaasServiceVersionResponse
	GetBody() *UpdateSaasServiceVersionResponseBody
}

type UpdateSaasServiceVersionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSaasServiceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSaasServiceVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSaasServiceVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateSaasServiceVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSaasServiceVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSaasServiceVersionResponse) GetBody() *UpdateSaasServiceVersionResponseBody {
	return s.Body
}

func (s *UpdateSaasServiceVersionResponse) SetHeaders(v map[string]*string) *UpdateSaasServiceVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateSaasServiceVersionResponse) SetStatusCode(v int32) *UpdateSaasServiceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSaasServiceVersionResponse) SetBody(v *UpdateSaasServiceVersionResponseBody) *UpdateSaasServiceVersionResponse {
	s.Body = v
	return s
}

func (s *UpdateSaasServiceVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
