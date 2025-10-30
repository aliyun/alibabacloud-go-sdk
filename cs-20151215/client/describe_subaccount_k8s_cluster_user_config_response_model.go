// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubaccountK8sClusterUserConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSubaccountK8sClusterUserConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSubaccountK8sClusterUserConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSubaccountK8sClusterUserConfigResponseBody) *DescribeSubaccountK8sClusterUserConfigResponse
	GetBody() *DescribeSubaccountK8sClusterUserConfigResponseBody
}

type DescribeSubaccountK8sClusterUserConfigResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSubaccountK8sClusterUserConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSubaccountK8sClusterUserConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubaccountK8sClusterUserConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubaccountK8sClusterUserConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSubaccountK8sClusterUserConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSubaccountK8sClusterUserConfigResponse) GetBody() *DescribeSubaccountK8sClusterUserConfigResponseBody {
	return s.Body
}

func (s *DescribeSubaccountK8sClusterUserConfigResponse) SetHeaders(v map[string]*string) *DescribeSubaccountK8sClusterUserConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubaccountK8sClusterUserConfigResponse) SetStatusCode(v int32) *DescribeSubaccountK8sClusterUserConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSubaccountK8sClusterUserConfigResponse) SetBody(v *DescribeSubaccountK8sClusterUserConfigResponseBody) *DescribeSubaccountK8sClusterUserConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeSubaccountK8sClusterUserConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
