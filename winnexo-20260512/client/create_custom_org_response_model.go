// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomOrgResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomOrgResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomOrgResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomOrgResponseBody) *CreateCustomOrgResponse
	GetBody() *CreateCustomOrgResponseBody
}

type CreateCustomOrgResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomOrgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomOrgResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomOrgResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomOrgResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomOrgResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomOrgResponse) GetBody() *CreateCustomOrgResponseBody {
	return s.Body
}

func (s *CreateCustomOrgResponse) SetHeaders(v map[string]*string) *CreateCustomOrgResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomOrgResponse) SetStatusCode(v int32) *CreateCustomOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomOrgResponse) SetBody(v *CreateCustomOrgResponseBody) *CreateCustomOrgResponse {
	s.Body = v
	return s
}

func (s *CreateCustomOrgResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
