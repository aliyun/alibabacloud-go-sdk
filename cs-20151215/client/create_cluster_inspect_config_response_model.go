// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterInspectConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateClusterInspectConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateClusterInspectConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateClusterInspectConfigResponseBody) *CreateClusterInspectConfigResponse
	GetBody() *CreateClusterInspectConfigResponseBody
}

type CreateClusterInspectConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClusterInspectConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClusterInspectConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterInspectConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterInspectConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateClusterInspectConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateClusterInspectConfigResponse) GetBody() *CreateClusterInspectConfigResponseBody {
	return s.Body
}

func (s *CreateClusterInspectConfigResponse) SetHeaders(v map[string]*string) *CreateClusterInspectConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterInspectConfigResponse) SetStatusCode(v int32) *CreateClusterInspectConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClusterInspectConfigResponse) SetBody(v *CreateClusterInspectConfigResponseBody) *CreateClusterInspectConfigResponse {
	s.Body = v
	return s
}

func (s *CreateClusterInspectConfigResponse) Validate() error {
	return dara.Validate(s)
}
