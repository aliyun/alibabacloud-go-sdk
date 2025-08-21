// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRenderingProjectInstanceStateMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRenderingProjectInstanceStateMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRenderingProjectInstanceStateMetricsResponse
	GetStatusCode() *int32
	SetBody(v *GetRenderingProjectInstanceStateMetricsResponseBody) *GetRenderingProjectInstanceStateMetricsResponse
	GetBody() *GetRenderingProjectInstanceStateMetricsResponseBody
}

type GetRenderingProjectInstanceStateMetricsResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRenderingProjectInstanceStateMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRenderingProjectInstanceStateMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRenderingProjectInstanceStateMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetRenderingProjectInstanceStateMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRenderingProjectInstanceStateMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRenderingProjectInstanceStateMetricsResponse) GetBody() *GetRenderingProjectInstanceStateMetricsResponseBody {
	return s.Body
}

func (s *GetRenderingProjectInstanceStateMetricsResponse) SetHeaders(v map[string]*string) *GetRenderingProjectInstanceStateMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetRenderingProjectInstanceStateMetricsResponse) SetStatusCode(v int32) *GetRenderingProjectInstanceStateMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRenderingProjectInstanceStateMetricsResponse) SetBody(v *GetRenderingProjectInstanceStateMetricsResponseBody) *GetRenderingProjectInstanceStateMetricsResponse {
	s.Body = v
	return s
}

func (s *GetRenderingProjectInstanceStateMetricsResponse) Validate() error {
	return dara.Validate(s)
}
