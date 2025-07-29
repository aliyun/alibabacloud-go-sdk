// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterInspectConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteClusterInspectConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteClusterInspectConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteClusterInspectConfigResponseBody) *DeleteClusterInspectConfigResponse
	GetBody() *DeleteClusterInspectConfigResponseBody
}

type DeleteClusterInspectConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClusterInspectConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClusterInspectConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterInspectConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterInspectConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteClusterInspectConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteClusterInspectConfigResponse) GetBody() *DeleteClusterInspectConfigResponseBody {
	return s.Body
}

func (s *DeleteClusterInspectConfigResponse) SetHeaders(v map[string]*string) *DeleteClusterInspectConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterInspectConfigResponse) SetStatusCode(v int32) *DeleteClusterInspectConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClusterInspectConfigResponse) SetBody(v *DeleteClusterInspectConfigResponseBody) *DeleteClusterInspectConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteClusterInspectConfigResponse) Validate() error {
	return dara.Validate(s)
}
