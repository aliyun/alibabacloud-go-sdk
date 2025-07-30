// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWarningConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWarningConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWarningConfigResponse
	GetStatusCode() *int32
	SetBody(v *ListWarningConfigResponseBody) *ListWarningConfigResponse
	GetBody() *ListWarningConfigResponseBody
}

type ListWarningConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWarningConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWarningConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWarningConfigResponse) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWarningConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWarningConfigResponse) GetBody() *ListWarningConfigResponseBody {
	return s.Body
}

func (s *ListWarningConfigResponse) SetHeaders(v map[string]*string) *ListWarningConfigResponse {
	s.Headers = v
	return s
}

func (s *ListWarningConfigResponse) SetStatusCode(v int32) *ListWarningConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWarningConfigResponse) SetBody(v *ListWarningConfigResponseBody) *ListWarningConfigResponse {
	s.Body = v
	return s
}

func (s *ListWarningConfigResponse) Validate() error {
	return dara.Validate(s)
}
