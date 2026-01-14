// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDomainStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDomainStateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDomainStateResponseBody) *UpdateDomainStateResponse
	GetBody() *UpdateDomainStateResponseBody
}

type UpdateDomainStateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDomainStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDomainStateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainStateResponse) GoString() string {
	return s.String()
}

func (s *UpdateDomainStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDomainStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDomainStateResponse) GetBody() *UpdateDomainStateResponseBody {
	return s.Body
}

func (s *UpdateDomainStateResponse) SetHeaders(v map[string]*string) *UpdateDomainStateResponse {
	s.Headers = v
	return s
}

func (s *UpdateDomainStateResponse) SetStatusCode(v int32) *UpdateDomainStateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDomainStateResponse) SetBody(v *UpdateDomainStateResponseBody) *UpdateDomainStateResponse {
	s.Body = v
	return s
}

func (s *UpdateDomainStateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
