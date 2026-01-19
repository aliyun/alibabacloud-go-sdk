// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationalUnitsForResourceServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOrganizationalUnitsForResourceServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOrganizationalUnitsForResourceServerResponse
	GetStatusCode() *int32
	SetBody(v *ListOrganizationalUnitsForResourceServerResponseBody) *ListOrganizationalUnitsForResourceServerResponse
	GetBody() *ListOrganizationalUnitsForResourceServerResponseBody
}

type ListOrganizationalUnitsForResourceServerResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOrganizationalUnitsForResourceServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOrganizationalUnitsForResourceServerResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitsForResourceServerResponse) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsForResourceServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOrganizationalUnitsForResourceServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOrganizationalUnitsForResourceServerResponse) GetBody() *ListOrganizationalUnitsForResourceServerResponseBody {
	return s.Body
}

func (s *ListOrganizationalUnitsForResourceServerResponse) SetHeaders(v map[string]*string) *ListOrganizationalUnitsForResourceServerResponse {
	s.Headers = v
	return s
}

func (s *ListOrganizationalUnitsForResourceServerResponse) SetStatusCode(v int32) *ListOrganizationalUnitsForResourceServerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrganizationalUnitsForResourceServerResponse) SetBody(v *ListOrganizationalUnitsForResourceServerResponseBody) *ListOrganizationalUnitsForResourceServerResponse {
	s.Body = v
	return s
}

func (s *ListOrganizationalUnitsForResourceServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
