// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginForUuidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPluginForUuidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPluginForUuidResponse
	GetStatusCode() *int32
	SetBody(v *ListPluginForUuidResponseBody) *ListPluginForUuidResponse
	GetBody() *ListPluginForUuidResponseBody
}

type ListPluginForUuidResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPluginForUuidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPluginForUuidResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPluginForUuidResponse) GoString() string {
	return s.String()
}

func (s *ListPluginForUuidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPluginForUuidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPluginForUuidResponse) GetBody() *ListPluginForUuidResponseBody {
	return s.Body
}

func (s *ListPluginForUuidResponse) SetHeaders(v map[string]*string) *ListPluginForUuidResponse {
	s.Headers = v
	return s
}

func (s *ListPluginForUuidResponse) SetStatusCode(v int32) *ListPluginForUuidResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPluginForUuidResponse) SetBody(v *ListPluginForUuidResponseBody) *ListPluginForUuidResponse {
	s.Body = v
	return s
}

func (s *ListPluginForUuidResponse) Validate() error {
	return dara.Validate(s)
}
