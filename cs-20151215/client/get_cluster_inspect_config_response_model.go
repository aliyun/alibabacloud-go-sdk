// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterInspectConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClusterInspectConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClusterInspectConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetClusterInspectConfigResponseBody) *GetClusterInspectConfigResponse
	GetBody() *GetClusterInspectConfigResponseBody
}

type GetClusterInspectConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterInspectConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterInspectConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClusterInspectConfigResponse) GoString() string {
	return s.String()
}

func (s *GetClusterInspectConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClusterInspectConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClusterInspectConfigResponse) GetBody() *GetClusterInspectConfigResponseBody {
	return s.Body
}

func (s *GetClusterInspectConfigResponse) SetHeaders(v map[string]*string) *GetClusterInspectConfigResponse {
	s.Headers = v
	return s
}

func (s *GetClusterInspectConfigResponse) SetStatusCode(v int32) *GetClusterInspectConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterInspectConfigResponse) SetBody(v *GetClusterInspectConfigResponseBody) *GetClusterInspectConfigResponse {
	s.Body = v
	return s
}

func (s *GetClusterInspectConfigResponse) Validate() error {
	return dara.Validate(s)
}
