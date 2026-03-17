// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnterpriseCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEnterpriseCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEnterpriseCodeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEnterpriseCodeResponseBody) *UpdateEnterpriseCodeResponse
	GetBody() *UpdateEnterpriseCodeResponseBody
}

type UpdateEnterpriseCodeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEnterpriseCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEnterpriseCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnterpriseCodeResponse) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEnterpriseCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEnterpriseCodeResponse) GetBody() *UpdateEnterpriseCodeResponseBody {
	return s.Body
}

func (s *UpdateEnterpriseCodeResponse) SetHeaders(v map[string]*string) *UpdateEnterpriseCodeResponse {
	s.Headers = v
	return s
}

func (s *UpdateEnterpriseCodeResponse) SetStatusCode(v int32) *UpdateEnterpriseCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEnterpriseCodeResponse) SetBody(v *UpdateEnterpriseCodeResponseBody) *UpdateEnterpriseCodeResponse {
	s.Body = v
	return s
}

func (s *UpdateEnterpriseCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
