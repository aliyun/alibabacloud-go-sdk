// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccessControlListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAccessControlListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAccessControlListResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAccessControlListResponseBody) *UpdateAccessControlListResponse
	GetBody() *UpdateAccessControlListResponseBody
}

type UpdateAccessControlListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAccessControlListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAccessControlListResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccessControlListResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccessControlListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAccessControlListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAccessControlListResponse) GetBody() *UpdateAccessControlListResponseBody {
	return s.Body
}

func (s *UpdateAccessControlListResponse) SetHeaders(v map[string]*string) *UpdateAccessControlListResponse {
	s.Headers = v
	return s
}

func (s *UpdateAccessControlListResponse) SetStatusCode(v int32) *UpdateAccessControlListResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAccessControlListResponse) SetBody(v *UpdateAccessControlListResponseBody) *UpdateAccessControlListResponse {
	s.Body = v
	return s
}

func (s *UpdateAccessControlListResponse) Validate() error {
	return dara.Validate(s)
}
