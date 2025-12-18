// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRemediationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRemediationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRemediationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRemediationResponseBody) *UpdateRemediationResponse
	GetBody() *UpdateRemediationResponseBody
}

type UpdateRemediationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRemediationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRemediationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRemediationResponse) GoString() string {
	return s.String()
}

func (s *UpdateRemediationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRemediationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRemediationResponse) GetBody() *UpdateRemediationResponseBody {
	return s.Body
}

func (s *UpdateRemediationResponse) SetHeaders(v map[string]*string) *UpdateRemediationResponse {
	s.Headers = v
	return s
}

func (s *UpdateRemediationResponse) SetStatusCode(v int32) *UpdateRemediationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRemediationResponse) SetBody(v *UpdateRemediationResponseBody) *UpdateRemediationResponse {
	s.Body = v
	return s
}

func (s *UpdateRemediationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
