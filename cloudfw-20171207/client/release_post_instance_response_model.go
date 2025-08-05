// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePostInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleasePostInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleasePostInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ReleasePostInstanceResponseBody) *ReleasePostInstanceResponse
	GetBody() *ReleasePostInstanceResponseBody
}

type ReleasePostInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleasePostInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleasePostInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleasePostInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleasePostInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleasePostInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleasePostInstanceResponse) GetBody() *ReleasePostInstanceResponseBody {
	return s.Body
}

func (s *ReleasePostInstanceResponse) SetHeaders(v map[string]*string) *ReleasePostInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleasePostInstanceResponse) SetStatusCode(v int32) *ReleasePostInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleasePostInstanceResponse) SetBody(v *ReleasePostInstanceResponseBody) *ReleasePostInstanceResponse {
	s.Body = v
	return s
}

func (s *ReleasePostInstanceResponse) Validate() error {
	return dara.Validate(s)
}
