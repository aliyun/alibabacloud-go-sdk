// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataPermissionsResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataPermissionsResponseBody) *CreateDataPermissionsResponse
	GetBody() *CreateDataPermissionsResponseBody
}

type CreateDataPermissionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPermissionsResponse) GoString() string {
	return s.String()
}

func (s *CreateDataPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataPermissionsResponse) GetBody() *CreateDataPermissionsResponseBody {
	return s.Body
}

func (s *CreateDataPermissionsResponse) SetHeaders(v map[string]*string) *CreateDataPermissionsResponse {
	s.Headers = v
	return s
}

func (s *CreateDataPermissionsResponse) SetStatusCode(v int32) *CreateDataPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataPermissionsResponse) SetBody(v *CreateDataPermissionsResponseBody) *CreateDataPermissionsResponse {
	s.Body = v
	return s
}

func (s *CreateDataPermissionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
