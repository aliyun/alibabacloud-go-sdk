// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLangfuseOrgResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLangfuseOrgResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLangfuseOrgResponse
	GetStatusCode() *int32
	SetBody(v *CreateLangfuseOrgResponseBody) *CreateLangfuseOrgResponse
	GetBody() *CreateLangfuseOrgResponseBody
}

type CreateLangfuseOrgResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLangfuseOrgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLangfuseOrgResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLangfuseOrgResponse) GoString() string {
	return s.String()
}

func (s *CreateLangfuseOrgResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLangfuseOrgResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLangfuseOrgResponse) GetBody() *CreateLangfuseOrgResponseBody {
	return s.Body
}

func (s *CreateLangfuseOrgResponse) SetHeaders(v map[string]*string) *CreateLangfuseOrgResponse {
	s.Headers = v
	return s
}

func (s *CreateLangfuseOrgResponse) SetStatusCode(v int32) *CreateLangfuseOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLangfuseOrgResponse) SetBody(v *CreateLangfuseOrgResponseBody) *CreateLangfuseOrgResponse {
	s.Body = v
	return s
}

func (s *CreateLangfuseOrgResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
