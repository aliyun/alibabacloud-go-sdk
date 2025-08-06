// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTemplateGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTemplateGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListTemplateGroupResponseBody) *ListTemplateGroupResponse
	GetBody() *ListTemplateGroupResponseBody
}

type ListTemplateGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTemplateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTemplateGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateGroupResponse) GoString() string {
	return s.String()
}

func (s *ListTemplateGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTemplateGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTemplateGroupResponse) GetBody() *ListTemplateGroupResponseBody {
	return s.Body
}

func (s *ListTemplateGroupResponse) SetHeaders(v map[string]*string) *ListTemplateGroupResponse {
	s.Headers = v
	return s
}

func (s *ListTemplateGroupResponse) SetStatusCode(v int32) *ListTemplateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTemplateGroupResponse) SetBody(v *ListTemplateGroupResponseBody) *ListTemplateGroupResponse {
	s.Body = v
	return s
}

func (s *ListTemplateGroupResponse) Validate() error {
	return dara.Validate(s)
}
