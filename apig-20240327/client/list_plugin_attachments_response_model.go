// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginAttachmentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPluginAttachmentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPluginAttachmentsResponse
	GetStatusCode() *int32
	SetBody(v *ListPluginAttachmentsResponseBody) *ListPluginAttachmentsResponse
	GetBody() *ListPluginAttachmentsResponseBody
}

type ListPluginAttachmentsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPluginAttachmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPluginAttachmentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPluginAttachmentsResponse) GoString() string {
	return s.String()
}

func (s *ListPluginAttachmentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPluginAttachmentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPluginAttachmentsResponse) GetBody() *ListPluginAttachmentsResponseBody {
	return s.Body
}

func (s *ListPluginAttachmentsResponse) SetHeaders(v map[string]*string) *ListPluginAttachmentsResponse {
	s.Headers = v
	return s
}

func (s *ListPluginAttachmentsResponse) SetStatusCode(v int32) *ListPluginAttachmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPluginAttachmentsResponse) SetBody(v *ListPluginAttachmentsResponseBody) *ListPluginAttachmentsResponse {
	s.Body = v
	return s
}

func (s *ListPluginAttachmentsResponse) Validate() error {
	return dara.Validate(s)
}
