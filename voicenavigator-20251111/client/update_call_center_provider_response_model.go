// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCallCenterProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCallCenterProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCallCenterProviderResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCallCenterProviderResponseBody) *UpdateCallCenterProviderResponse
	GetBody() *UpdateCallCenterProviderResponseBody
}

type UpdateCallCenterProviderResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCallCenterProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCallCenterProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCallCenterProviderResponse) GoString() string {
	return s.String()
}

func (s *UpdateCallCenterProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCallCenterProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCallCenterProviderResponse) GetBody() *UpdateCallCenterProviderResponseBody {
	return s.Body
}

func (s *UpdateCallCenterProviderResponse) SetHeaders(v map[string]*string) *UpdateCallCenterProviderResponse {
	s.Headers = v
	return s
}

func (s *UpdateCallCenterProviderResponse) SetStatusCode(v int32) *UpdateCallCenterProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCallCenterProviderResponse) SetBody(v *UpdateCallCenterProviderResponseBody) *UpdateCallCenterProviderResponse {
	s.Body = v
	return s
}

func (s *UpdateCallCenterProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
