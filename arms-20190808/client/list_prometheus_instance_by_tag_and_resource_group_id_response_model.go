// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusInstanceByTagAndResourceGroupIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrometheusInstanceByTagAndResourceGroupIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrometheusInstanceByTagAndResourceGroupIdResponse
	GetStatusCode() *int32
	SetBody(v *ListPrometheusInstanceByTagAndResourceGroupIdResponseBody) *ListPrometheusInstanceByTagAndResourceGroupIdResponse
	GetBody() *ListPrometheusInstanceByTagAndResourceGroupIdResponseBody
}

type ListPrometheusInstanceByTagAndResourceGroupIdResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrometheusInstanceByTagAndResourceGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrometheusInstanceByTagAndResourceGroupIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusInstanceByTagAndResourceGroupIdResponse) GoString() string {
	return s.String()
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponse) GetBody() *ListPrometheusInstanceByTagAndResourceGroupIdResponseBody {
	return s.Body
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponse) SetHeaders(v map[string]*string) *ListPrometheusInstanceByTagAndResourceGroupIdResponse {
	s.Headers = v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponse) SetStatusCode(v int32) *ListPrometheusInstanceByTagAndResourceGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponse) SetBody(v *ListPrometheusInstanceByTagAndResourceGroupIdResponseBody) *ListPrometheusInstanceByTagAndResourceGroupIdResponse {
	s.Body = v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
