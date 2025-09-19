// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileProtectPluginStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFileProtectPluginStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFileProtectPluginStatusResponse
	GetStatusCode() *int32
	SetBody(v *ListFileProtectPluginStatusResponseBody) *ListFileProtectPluginStatusResponse
	GetBody() *ListFileProtectPluginStatusResponseBody
}

type ListFileProtectPluginStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFileProtectPluginStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFileProtectPluginStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectPluginStatusResponse) GoString() string {
	return s.String()
}

func (s *ListFileProtectPluginStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFileProtectPluginStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFileProtectPluginStatusResponse) GetBody() *ListFileProtectPluginStatusResponseBody {
	return s.Body
}

func (s *ListFileProtectPluginStatusResponse) SetHeaders(v map[string]*string) *ListFileProtectPluginStatusResponse {
	s.Headers = v
	return s
}

func (s *ListFileProtectPluginStatusResponse) SetStatusCode(v int32) *ListFileProtectPluginStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFileProtectPluginStatusResponse) SetBody(v *ListFileProtectPluginStatusResponseBody) *ListFileProtectPluginStatusResponse {
	s.Body = v
	return s
}

func (s *ListFileProtectPluginStatusResponse) Validate() error {
	return dara.Validate(s)
}
