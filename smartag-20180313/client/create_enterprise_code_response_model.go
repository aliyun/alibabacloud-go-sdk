// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnterpriseCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEnterpriseCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEnterpriseCodeResponse
	GetStatusCode() *int32
	SetBody(v *CreateEnterpriseCodeResponseBody) *CreateEnterpriseCodeResponse
	GetBody() *CreateEnterpriseCodeResponseBody
}

type CreateEnterpriseCodeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEnterpriseCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEnterpriseCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseCodeResponse) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEnterpriseCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEnterpriseCodeResponse) GetBody() *CreateEnterpriseCodeResponseBody {
	return s.Body
}

func (s *CreateEnterpriseCodeResponse) SetHeaders(v map[string]*string) *CreateEnterpriseCodeResponse {
	s.Headers = v
	return s
}

func (s *CreateEnterpriseCodeResponse) SetStatusCode(v int32) *CreateEnterpriseCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEnterpriseCodeResponse) SetBody(v *CreateEnterpriseCodeResponseBody) *CreateEnterpriseCodeResponse {
	s.Body = v
	return s
}

func (s *CreateEnterpriseCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
