// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityIdentifyResultStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSecurityIdentifyResultStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSecurityIdentifyResultStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSecurityIdentifyResultStatusResponseBody) *UpdateSecurityIdentifyResultStatusResponse
	GetBody() *UpdateSecurityIdentifyResultStatusResponseBody
}

type UpdateSecurityIdentifyResultStatusResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSecurityIdentifyResultStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSecurityIdentifyResultStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityIdentifyResultStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateSecurityIdentifyResultStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSecurityIdentifyResultStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSecurityIdentifyResultStatusResponse) GetBody() *UpdateSecurityIdentifyResultStatusResponseBody {
	return s.Body
}

func (s *UpdateSecurityIdentifyResultStatusResponse) SetHeaders(v map[string]*string) *UpdateSecurityIdentifyResultStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateSecurityIdentifyResultStatusResponse) SetStatusCode(v int32) *UpdateSecurityIdentifyResultStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSecurityIdentifyResultStatusResponse) SetBody(v *UpdateSecurityIdentifyResultStatusResponseBody) *UpdateSecurityIdentifyResultStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateSecurityIdentifyResultStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
