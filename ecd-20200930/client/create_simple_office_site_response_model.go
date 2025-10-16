// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimpleOfficeSiteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSimpleOfficeSiteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSimpleOfficeSiteResponse
	GetStatusCode() *int32
	SetBody(v *CreateSimpleOfficeSiteResponseBody) *CreateSimpleOfficeSiteResponse
	GetBody() *CreateSimpleOfficeSiteResponseBody
}

type CreateSimpleOfficeSiteResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSimpleOfficeSiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSimpleOfficeSiteResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSimpleOfficeSiteResponse) GoString() string {
	return s.String()
}

func (s *CreateSimpleOfficeSiteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSimpleOfficeSiteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSimpleOfficeSiteResponse) GetBody() *CreateSimpleOfficeSiteResponseBody {
	return s.Body
}

func (s *CreateSimpleOfficeSiteResponse) SetHeaders(v map[string]*string) *CreateSimpleOfficeSiteResponse {
	s.Headers = v
	return s
}

func (s *CreateSimpleOfficeSiteResponse) SetStatusCode(v int32) *CreateSimpleOfficeSiteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSimpleOfficeSiteResponse) SetBody(v *CreateSimpleOfficeSiteResponseBody) *CreateSimpleOfficeSiteResponse {
	s.Body = v
	return s
}

func (s *CreateSimpleOfficeSiteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
