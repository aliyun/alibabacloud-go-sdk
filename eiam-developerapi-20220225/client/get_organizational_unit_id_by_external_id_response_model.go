// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrganizationalUnitIdByExternalIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOrganizationalUnitIdByExternalIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOrganizationalUnitIdByExternalIdResponse
	GetStatusCode() *int32
	SetBody(v *GetOrganizationalUnitIdByExternalIdResponseBody) *GetOrganizationalUnitIdByExternalIdResponse
	GetBody() *GetOrganizationalUnitIdByExternalIdResponseBody
}

type GetOrganizationalUnitIdByExternalIdResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrganizationalUnitIdByExternalIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrganizationalUnitIdByExternalIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOrganizationalUnitIdByExternalIdResponse) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitIdByExternalIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOrganizationalUnitIdByExternalIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOrganizationalUnitIdByExternalIdResponse) GetBody() *GetOrganizationalUnitIdByExternalIdResponseBody {
	return s.Body
}

func (s *GetOrganizationalUnitIdByExternalIdResponse) SetHeaders(v map[string]*string) *GetOrganizationalUnitIdByExternalIdResponse {
	s.Headers = v
	return s
}

func (s *GetOrganizationalUnitIdByExternalIdResponse) SetStatusCode(v int32) *GetOrganizationalUnitIdByExternalIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrganizationalUnitIdByExternalIdResponse) SetBody(v *GetOrganizationalUnitIdByExternalIdResponseBody) *GetOrganizationalUnitIdByExternalIdResponse {
	s.Body = v
	return s
}

func (s *GetOrganizationalUnitIdByExternalIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
