// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMaintenanceActionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMaintenanceActionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMaintenanceActionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMaintenanceActionResponseBody) *DescribeMaintenanceActionResponse
	GetBody() *DescribeMaintenanceActionResponseBody
}

type DescribeMaintenanceActionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMaintenanceActionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMaintenanceActionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMaintenanceActionResponse) GoString() string {
	return s.String()
}

func (s *DescribeMaintenanceActionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMaintenanceActionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMaintenanceActionResponse) GetBody() *DescribeMaintenanceActionResponseBody {
	return s.Body
}

func (s *DescribeMaintenanceActionResponse) SetHeaders(v map[string]*string) *DescribeMaintenanceActionResponse {
	s.Headers = v
	return s
}

func (s *DescribeMaintenanceActionResponse) SetStatusCode(v int32) *DescribeMaintenanceActionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMaintenanceActionResponse) SetBody(v *DescribeMaintenanceActionResponseBody) *DescribeMaintenanceActionResponse {
	s.Body = v
	return s
}

func (s *DescribeMaintenanceActionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
