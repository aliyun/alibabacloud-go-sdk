// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePageResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePageResponseBody) *UpdatePageResponse
	GetBody() *UpdatePageResponseBody
}

type UpdatePageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePageResponse) GoString() string {
	return s.String()
}

func (s *UpdatePageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePageResponse) GetBody() *UpdatePageResponseBody {
	return s.Body
}

func (s *UpdatePageResponse) SetHeaders(v map[string]*string) *UpdatePageResponse {
	s.Headers = v
	return s
}

func (s *UpdatePageResponse) SetStatusCode(v int32) *UpdatePageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePageResponse) SetBody(v *UpdatePageResponseBody) *UpdatePageResponse {
	s.Body = v
	return s
}

func (s *UpdatePageResponse) Validate() error {
	return dara.Validate(s)
}
