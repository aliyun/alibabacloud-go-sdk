// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterFromGrafanaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterFromGrafanaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterFromGrafanaResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterFromGrafanaResponseBody) *ListClusterFromGrafanaResponse
	GetBody() *ListClusterFromGrafanaResponseBody
}

type ListClusterFromGrafanaResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterFromGrafanaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterFromGrafanaResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterFromGrafanaResponse) GoString() string {
	return s.String()
}

func (s *ListClusterFromGrafanaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterFromGrafanaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterFromGrafanaResponse) GetBody() *ListClusterFromGrafanaResponseBody {
	return s.Body
}

func (s *ListClusterFromGrafanaResponse) SetHeaders(v map[string]*string) *ListClusterFromGrafanaResponse {
	s.Headers = v
	return s
}

func (s *ListClusterFromGrafanaResponse) SetStatusCode(v int32) *ListClusterFromGrafanaResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterFromGrafanaResponse) SetBody(v *ListClusterFromGrafanaResponseBody) *ListClusterFromGrafanaResponse {
	s.Body = v
	return s
}

func (s *ListClusterFromGrafanaResponse) Validate() error {
	return dara.Validate(s)
}
