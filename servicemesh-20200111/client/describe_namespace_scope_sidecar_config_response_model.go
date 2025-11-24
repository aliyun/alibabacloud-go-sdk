// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNamespaceScopeSidecarConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNamespaceScopeSidecarConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNamespaceScopeSidecarConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNamespaceScopeSidecarConfigResponseBody) *DescribeNamespaceScopeSidecarConfigResponse
	GetBody() *DescribeNamespaceScopeSidecarConfigResponseBody
}

type DescribeNamespaceScopeSidecarConfigResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNamespaceScopeSidecarConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNamespaceScopeSidecarConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNamespaceScopeSidecarConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNamespaceScopeSidecarConfigResponse) GetBody() *DescribeNamespaceScopeSidecarConfigResponseBody {
	return s.Body
}

func (s *DescribeNamespaceScopeSidecarConfigResponse) SetHeaders(v map[string]*string) *DescribeNamespaceScopeSidecarConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponse) SetStatusCode(v int32) *DescribeNamespaceScopeSidecarConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponse) SetBody(v *DescribeNamespaceScopeSidecarConfigResponseBody) *DescribeNamespaceScopeSidecarConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
