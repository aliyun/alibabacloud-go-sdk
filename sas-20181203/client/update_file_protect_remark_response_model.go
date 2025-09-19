// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileProtectRemarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFileProtectRemarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFileProtectRemarkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFileProtectRemarkResponseBody) *UpdateFileProtectRemarkResponse
	GetBody() *UpdateFileProtectRemarkResponseBody
}

type UpdateFileProtectRemarkResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFileProtectRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFileProtectRemarkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileProtectRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateFileProtectRemarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFileProtectRemarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFileProtectRemarkResponse) GetBody() *UpdateFileProtectRemarkResponseBody {
	return s.Body
}

func (s *UpdateFileProtectRemarkResponse) SetHeaders(v map[string]*string) *UpdateFileProtectRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateFileProtectRemarkResponse) SetStatusCode(v int32) *UpdateFileProtectRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFileProtectRemarkResponse) SetBody(v *UpdateFileProtectRemarkResponseBody) *UpdateFileProtectRemarkResponse {
	s.Body = v
	return s
}

func (s *UpdateFileProtectRemarkResponse) Validate() error {
	return dara.Validate(s)
}
