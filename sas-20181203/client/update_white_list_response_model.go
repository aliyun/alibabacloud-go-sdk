// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWhiteListResponseBody) *UpdateWhiteListResponse
	GetBody() *UpdateWhiteListResponseBody
}

type UpdateWhiteListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhiteListResponse) GoString() string {
	return s.String()
}

func (s *UpdateWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWhiteListResponse) GetBody() *UpdateWhiteListResponseBody {
	return s.Body
}

func (s *UpdateWhiteListResponse) SetHeaders(v map[string]*string) *UpdateWhiteListResponse {
	s.Headers = v
	return s
}

func (s *UpdateWhiteListResponse) SetStatusCode(v int32) *UpdateWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWhiteListResponse) SetBody(v *UpdateWhiteListResponseBody) *UpdateWhiteListResponse {
	s.Body = v
	return s
}

func (s *UpdateWhiteListResponse) Validate() error {
	return dara.Validate(s)
}
