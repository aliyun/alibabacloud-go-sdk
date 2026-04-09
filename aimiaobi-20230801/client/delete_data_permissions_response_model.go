// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataPermissionsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataPermissionsResponseBody) *DeleteDataPermissionsResponse
	GetBody() *DeleteDataPermissionsResponseBody
}

type DeleteDataPermissionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataPermissionsResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataPermissionsResponse) GetBody() *DeleteDataPermissionsResponseBody {
	return s.Body
}

func (s *DeleteDataPermissionsResponse) SetHeaders(v map[string]*string) *DeleteDataPermissionsResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataPermissionsResponse) SetStatusCode(v int32) *DeleteDataPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataPermissionsResponse) SetBody(v *DeleteDataPermissionsResponseBody) *DeleteDataPermissionsResponse {
	s.Body = v
	return s
}

func (s *DeleteDataPermissionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
