// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImagePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyImagePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyImagePermissionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyImagePermissionResponseBody) *ModifyImagePermissionResponse
	GetBody() *ModifyImagePermissionResponseBody
}

type ModifyImagePermissionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyImagePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyImagePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyImagePermissionResponse) GoString() string {
	return s.String()
}

func (s *ModifyImagePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyImagePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyImagePermissionResponse) GetBody() *ModifyImagePermissionResponseBody {
	return s.Body
}

func (s *ModifyImagePermissionResponse) SetHeaders(v map[string]*string) *ModifyImagePermissionResponse {
	s.Headers = v
	return s
}

func (s *ModifyImagePermissionResponse) SetStatusCode(v int32) *ModifyImagePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyImagePermissionResponse) SetBody(v *ModifyImagePermissionResponseBody) *ModifyImagePermissionResponse {
	s.Body = v
	return s
}

func (s *ModifyImagePermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
