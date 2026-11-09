// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomAuthPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCustomAuthPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCustomAuthPermissionResponse
	GetStatusCode() *int32
	SetBody(v *AddCustomAuthPermissionResponseBody) *AddCustomAuthPermissionResponse
	GetBody() *AddCustomAuthPermissionResponseBody
}

type AddCustomAuthPermissionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCustomAuthPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCustomAuthPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCustomAuthPermissionResponse) GoString() string {
	return s.String()
}

func (s *AddCustomAuthPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCustomAuthPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCustomAuthPermissionResponse) GetBody() *AddCustomAuthPermissionResponseBody {
	return s.Body
}

func (s *AddCustomAuthPermissionResponse) SetHeaders(v map[string]*string) *AddCustomAuthPermissionResponse {
	s.Headers = v
	return s
}

func (s *AddCustomAuthPermissionResponse) SetStatusCode(v int32) *AddCustomAuthPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCustomAuthPermissionResponse) SetBody(v *AddCustomAuthPermissionResponseBody) *AddCustomAuthPermissionResponse {
	s.Body = v
	return s
}

func (s *AddCustomAuthPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
