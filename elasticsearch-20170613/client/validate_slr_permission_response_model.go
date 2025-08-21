// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateSlrPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateSlrPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateSlrPermissionResponse
	GetStatusCode() *int32
	SetBody(v *ValidateSlrPermissionResponseBody) *ValidateSlrPermissionResponse
	GetBody() *ValidateSlrPermissionResponseBody
}

type ValidateSlrPermissionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateSlrPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateSlrPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateSlrPermissionResponse) GoString() string {
	return s.String()
}

func (s *ValidateSlrPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateSlrPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateSlrPermissionResponse) GetBody() *ValidateSlrPermissionResponseBody {
	return s.Body
}

func (s *ValidateSlrPermissionResponse) SetHeaders(v map[string]*string) *ValidateSlrPermissionResponse {
	s.Headers = v
	return s
}

func (s *ValidateSlrPermissionResponse) SetStatusCode(v int32) *ValidateSlrPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateSlrPermissionResponse) SetBody(v *ValidateSlrPermissionResponseBody) *ValidateSlrPermissionResponse {
	s.Body = v
	return s
}

func (s *ValidateSlrPermissionResponse) Validate() error {
	return dara.Validate(s)
}
