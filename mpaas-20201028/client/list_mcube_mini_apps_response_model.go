// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeMiniAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMcubeMiniAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMcubeMiniAppsResponse
	GetStatusCode() *int32
	SetBody(v *ListMcubeMiniAppsResponseBody) *ListMcubeMiniAppsResponse
	GetBody() *ListMcubeMiniAppsResponseBody
}

type ListMcubeMiniAppsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMcubeMiniAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMcubeMiniAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeMiniAppsResponse) GoString() string {
	return s.String()
}

func (s *ListMcubeMiniAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMcubeMiniAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMcubeMiniAppsResponse) GetBody() *ListMcubeMiniAppsResponseBody {
	return s.Body
}

func (s *ListMcubeMiniAppsResponse) SetHeaders(v map[string]*string) *ListMcubeMiniAppsResponse {
	s.Headers = v
	return s
}

func (s *ListMcubeMiniAppsResponse) SetStatusCode(v int32) *ListMcubeMiniAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMcubeMiniAppsResponse) SetBody(v *ListMcubeMiniAppsResponseBody) *ListMcubeMiniAppsResponse {
	s.Body = v
	return s
}

func (s *ListMcubeMiniAppsResponse) Validate() error {
	return dara.Validate(s)
}
