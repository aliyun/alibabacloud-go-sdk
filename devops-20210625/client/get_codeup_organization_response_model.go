// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCodeupOrganizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCodeupOrganizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCodeupOrganizationResponse
	GetStatusCode() *int32
	SetBody(v *GetCodeupOrganizationResponseBody) *GetCodeupOrganizationResponse
	GetBody() *GetCodeupOrganizationResponseBody
}

type GetCodeupOrganizationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCodeupOrganizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCodeupOrganizationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCodeupOrganizationResponse) GoString() string {
	return s.String()
}

func (s *GetCodeupOrganizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCodeupOrganizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCodeupOrganizationResponse) GetBody() *GetCodeupOrganizationResponseBody {
	return s.Body
}

func (s *GetCodeupOrganizationResponse) SetHeaders(v map[string]*string) *GetCodeupOrganizationResponse {
	s.Headers = v
	return s
}

func (s *GetCodeupOrganizationResponse) SetStatusCode(v int32) *GetCodeupOrganizationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCodeupOrganizationResponse) SetBody(v *GetCodeupOrganizationResponseBody) *GetCodeupOrganizationResponse {
	s.Body = v
	return s
}

func (s *GetCodeupOrganizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
