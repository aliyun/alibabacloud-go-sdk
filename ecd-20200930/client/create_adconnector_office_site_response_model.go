// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateADConnectorOfficeSiteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateADConnectorOfficeSiteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateADConnectorOfficeSiteResponse
	GetStatusCode() *int32
	SetBody(v *CreateADConnectorOfficeSiteResponseBody) *CreateADConnectorOfficeSiteResponse
	GetBody() *CreateADConnectorOfficeSiteResponseBody
}

type CreateADConnectorOfficeSiteResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateADConnectorOfficeSiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateADConnectorOfficeSiteResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateADConnectorOfficeSiteResponse) GoString() string {
	return s.String()
}

func (s *CreateADConnectorOfficeSiteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateADConnectorOfficeSiteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateADConnectorOfficeSiteResponse) GetBody() *CreateADConnectorOfficeSiteResponseBody {
	return s.Body
}

func (s *CreateADConnectorOfficeSiteResponse) SetHeaders(v map[string]*string) *CreateADConnectorOfficeSiteResponse {
	s.Headers = v
	return s
}

func (s *CreateADConnectorOfficeSiteResponse) SetStatusCode(v int32) *CreateADConnectorOfficeSiteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateADConnectorOfficeSiteResponse) SetBody(v *CreateADConnectorOfficeSiteResponseBody) *CreateADConnectorOfficeSiteResponse {
	s.Body = v
	return s
}

func (s *CreateADConnectorOfficeSiteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
