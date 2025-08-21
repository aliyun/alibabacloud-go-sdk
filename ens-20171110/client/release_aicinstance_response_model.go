// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseAICInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseAICInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseAICInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseAICInstanceResponseBody) *ReleaseAICInstanceResponse
	GetBody() *ReleaseAICInstanceResponseBody
}

type ReleaseAICInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseAICInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseAICInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseAICInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseAICInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseAICInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseAICInstanceResponse) GetBody() *ReleaseAICInstanceResponseBody {
	return s.Body
}

func (s *ReleaseAICInstanceResponse) SetHeaders(v map[string]*string) *ReleaseAICInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseAICInstanceResponse) SetStatusCode(v int32) *ReleaseAICInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseAICInstanceResponse) SetBody(v *ReleaseAICInstanceResponseBody) *ReleaseAICInstanceResponse {
	s.Body = v
	return s
}

func (s *ReleaseAICInstanceResponse) Validate() error {
	return dara.Validate(s)
}
