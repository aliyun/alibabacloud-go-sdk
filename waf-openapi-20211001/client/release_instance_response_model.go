// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseInstanceResponseBody) *ReleaseInstanceResponse
	GetBody() *ReleaseInstanceResponseBody
}

type ReleaseInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseInstanceResponse) GetBody() *ReleaseInstanceResponseBody {
	return s.Body
}

func (s *ReleaseInstanceResponse) SetHeaders(v map[string]*string) *ReleaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseInstanceResponse) SetStatusCode(v int32) *ReleaseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseInstanceResponse) SetBody(v *ReleaseInstanceResponseBody) *ReleaseInstanceResponse {
	s.Body = v
	return s
}

func (s *ReleaseInstanceResponse) Validate() error {
	return dara.Validate(s)
}
