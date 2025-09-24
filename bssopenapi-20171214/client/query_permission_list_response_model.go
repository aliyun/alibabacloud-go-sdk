// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPermissionListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPermissionListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPermissionListResponse
	GetStatusCode() *int32
	SetBody(v *QueryPermissionListResponseBody) *QueryPermissionListResponse
	GetBody() *QueryPermissionListResponseBody
}

type QueryPermissionListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPermissionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPermissionListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPermissionListResponse) GoString() string {
	return s.String()
}

func (s *QueryPermissionListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPermissionListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPermissionListResponse) GetBody() *QueryPermissionListResponseBody {
	return s.Body
}

func (s *QueryPermissionListResponse) SetHeaders(v map[string]*string) *QueryPermissionListResponse {
	s.Headers = v
	return s
}

func (s *QueryPermissionListResponse) SetStatusCode(v int32) *QueryPermissionListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPermissionListResponse) SetBody(v *QueryPermissionListResponseBody) *QueryPermissionListResponse {
	s.Body = v
	return s
}

func (s *QueryPermissionListResponse) Validate() error {
	return dara.Validate(s)
}
