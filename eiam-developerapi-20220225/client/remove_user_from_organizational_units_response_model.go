// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserFromOrganizationalUnitsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveUserFromOrganizationalUnitsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveUserFromOrganizationalUnitsResponse
	GetStatusCode() *int32
}

type RemoveUserFromOrganizationalUnitsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s RemoveUserFromOrganizationalUnitsResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserFromOrganizationalUnitsResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserFromOrganizationalUnitsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveUserFromOrganizationalUnitsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveUserFromOrganizationalUnitsResponse) SetHeaders(v map[string]*string) *RemoveUserFromOrganizationalUnitsResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserFromOrganizationalUnitsResponse) SetStatusCode(v int32) *RemoveUserFromOrganizationalUnitsResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUserFromOrganizationalUnitsResponse) Validate() error {
	return dara.Validate(s)
}
