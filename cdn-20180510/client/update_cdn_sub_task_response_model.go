// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCdnSubTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCdnSubTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCdnSubTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCdnSubTaskResponseBody) *UpdateCdnSubTaskResponse
	GetBody() *UpdateCdnSubTaskResponseBody
}

type UpdateCdnSubTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCdnSubTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCdnSubTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCdnSubTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateCdnSubTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCdnSubTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCdnSubTaskResponse) GetBody() *UpdateCdnSubTaskResponseBody {
	return s.Body
}

func (s *UpdateCdnSubTaskResponse) SetHeaders(v map[string]*string) *UpdateCdnSubTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateCdnSubTaskResponse) SetStatusCode(v int32) *UpdateCdnSubTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCdnSubTaskResponse) SetBody(v *UpdateCdnSubTaskResponseBody) *UpdateCdnSubTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateCdnSubTaskResponse) Validate() error {
	return dara.Validate(s)
}
