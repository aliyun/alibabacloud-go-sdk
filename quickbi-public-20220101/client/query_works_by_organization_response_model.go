// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorksByOrganizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryWorksByOrganizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryWorksByOrganizationResponse
	GetStatusCode() *int32
	SetBody(v *QueryWorksByOrganizationResponseBody) *QueryWorksByOrganizationResponse
	GetBody() *QueryWorksByOrganizationResponseBody
}

type QueryWorksByOrganizationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryWorksByOrganizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryWorksByOrganizationResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryWorksByOrganizationResponse) GoString() string {
	return s.String()
}

func (s *QueryWorksByOrganizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryWorksByOrganizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryWorksByOrganizationResponse) GetBody() *QueryWorksByOrganizationResponseBody {
	return s.Body
}

func (s *QueryWorksByOrganizationResponse) SetHeaders(v map[string]*string) *QueryWorksByOrganizationResponse {
	s.Headers = v
	return s
}

func (s *QueryWorksByOrganizationResponse) SetStatusCode(v int32) *QueryWorksByOrganizationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryWorksByOrganizationResponse) SetBody(v *QueryWorksByOrganizationResponseBody) *QueryWorksByOrganizationResponse {
	s.Body = v
	return s
}

func (s *QueryWorksByOrganizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
