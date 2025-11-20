// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddPermissionResponse
	GetStatusCode() *int32
	SetBody(v *AddPermissionResponseBody) *AddPermissionResponse
	GetBody() *AddPermissionResponseBody
}

type AddPermissionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s AddPermissionResponse) GoString() string {
	return s.String()
}

func (s *AddPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddPermissionResponse) GetBody() *AddPermissionResponseBody {
	return s.Body
}

func (s *AddPermissionResponse) SetHeaders(v map[string]*string) *AddPermissionResponse {
	s.Headers = v
	return s
}

func (s *AddPermissionResponse) SetStatusCode(v int32) *AddPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPermissionResponse) SetBody(v *AddPermissionResponseBody) *AddPermissionResponse {
	s.Body = v
	return s
}

func (s *AddPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
