// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateListResponse
	GetStatusCode() *int32
	SetBody(v *UpdateListResponseBody) *UpdateListResponse
	GetBody() *UpdateListResponseBody
}

type UpdateListResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateListResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateListResponse) GoString() string {
	return s.String()
}

func (s *UpdateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateListResponse) GetBody() *UpdateListResponseBody {
	return s.Body
}

func (s *UpdateListResponse) SetHeaders(v map[string]*string) *UpdateListResponse {
	s.Headers = v
	return s
}

func (s *UpdateListResponse) SetStatusCode(v int32) *UpdateListResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateListResponse) SetBody(v *UpdateListResponseBody) *UpdateListResponse {
	s.Body = v
	return s
}

func (s *UpdateListResponse) Validate() error {
	return dara.Validate(s)
}
