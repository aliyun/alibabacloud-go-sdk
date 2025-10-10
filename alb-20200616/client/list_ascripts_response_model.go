// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAScriptsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAScriptsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAScriptsResponse
	GetStatusCode() *int32
	SetBody(v *ListAScriptsResponseBody) *ListAScriptsResponse
	GetBody() *ListAScriptsResponseBody
}

type ListAScriptsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAScriptsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAScriptsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAScriptsResponse) GoString() string {
	return s.String()
}

func (s *ListAScriptsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAScriptsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAScriptsResponse) GetBody() *ListAScriptsResponseBody {
	return s.Body
}

func (s *ListAScriptsResponse) SetHeaders(v map[string]*string) *ListAScriptsResponse {
	s.Headers = v
	return s
}

func (s *ListAScriptsResponse) SetStatusCode(v int32) *ListAScriptsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAScriptsResponse) SetBody(v *ListAScriptsResponseBody) *ListAScriptsResponse {
	s.Body = v
	return s
}

func (s *ListAScriptsResponse) Validate() error {
	return dara.Validate(s)
}
