// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileProtectEventStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFileProtectEventStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFileProtectEventStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFileProtectEventStatusResponseBody) *UpdateFileProtectEventStatusResponse
	GetBody() *UpdateFileProtectEventStatusResponseBody
}

type UpdateFileProtectEventStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFileProtectEventStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFileProtectEventStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileProtectEventStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateFileProtectEventStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFileProtectEventStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFileProtectEventStatusResponse) GetBody() *UpdateFileProtectEventStatusResponseBody {
	return s.Body
}

func (s *UpdateFileProtectEventStatusResponse) SetHeaders(v map[string]*string) *UpdateFileProtectEventStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateFileProtectEventStatusResponse) SetStatusCode(v int32) *UpdateFileProtectEventStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFileProtectEventStatusResponse) SetBody(v *UpdateFileProtectEventStatusResponseBody) *UpdateFileProtectEventStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateFileProtectEventStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
