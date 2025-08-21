// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFileInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFileInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFileInfoResponseBody) *UpdateFileInfoResponse
	GetBody() *UpdateFileInfoResponseBody
}

type UpdateFileInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFileInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFileInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateFileInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFileInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFileInfoResponse) GetBody() *UpdateFileInfoResponseBody {
	return s.Body
}

func (s *UpdateFileInfoResponse) SetHeaders(v map[string]*string) *UpdateFileInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateFileInfoResponse) SetStatusCode(v int32) *UpdateFileInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFileInfoResponse) SetBody(v *UpdateFileInfoResponseBody) *UpdateFileInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateFileInfoResponse) Validate() error {
	return dara.Validate(s)
}
