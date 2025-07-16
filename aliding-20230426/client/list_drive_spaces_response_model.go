// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDriveSpacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDriveSpacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDriveSpacesResponse
	GetStatusCode() *int32
	SetBody(v *ListDriveSpacesResponseBody) *ListDriveSpacesResponse
	GetBody() *ListDriveSpacesResponseBody
}

type ListDriveSpacesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDriveSpacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDriveSpacesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDriveSpacesResponse) GoString() string {
	return s.String()
}

func (s *ListDriveSpacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDriveSpacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDriveSpacesResponse) GetBody() *ListDriveSpacesResponseBody {
	return s.Body
}

func (s *ListDriveSpacesResponse) SetHeaders(v map[string]*string) *ListDriveSpacesResponse {
	s.Headers = v
	return s
}

func (s *ListDriveSpacesResponse) SetStatusCode(v int32) *ListDriveSpacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDriveSpacesResponse) SetBody(v *ListDriveSpacesResponseBody) *ListDriveSpacesResponse {
	s.Body = v
	return s
}

func (s *ListDriveSpacesResponse) Validate() error {
	return dara.Validate(s)
}
