// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecordPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddRecordPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddRecordPermissionResponse
	GetStatusCode() *int32
	SetBody(v *AddRecordPermissionResponseBody) *AddRecordPermissionResponse
	GetBody() *AddRecordPermissionResponseBody
}

type AddRecordPermissionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRecordPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRecordPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s AddRecordPermissionResponse) GoString() string {
	return s.String()
}

func (s *AddRecordPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddRecordPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddRecordPermissionResponse) GetBody() *AddRecordPermissionResponseBody {
	return s.Body
}

func (s *AddRecordPermissionResponse) SetHeaders(v map[string]*string) *AddRecordPermissionResponse {
	s.Headers = v
	return s
}

func (s *AddRecordPermissionResponse) SetStatusCode(v int32) *AddRecordPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRecordPermissionResponse) SetBody(v *AddRecordPermissionResponseBody) *AddRecordPermissionResponse {
	s.Body = v
	return s
}

func (s *AddRecordPermissionResponse) Validate() error {
	return dara.Validate(s)
}
