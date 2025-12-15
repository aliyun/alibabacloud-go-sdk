// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationMaintenanceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeActiveOperationMaintenanceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeActiveOperationMaintenanceConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeActiveOperationMaintenanceConfigResponseBody) *DescribeActiveOperationMaintenanceConfigResponse
	GetBody() *DescribeActiveOperationMaintenanceConfigResponseBody
}

type DescribeActiveOperationMaintenanceConfigResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeActiveOperationMaintenanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeActiveOperationMaintenanceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationMaintenanceConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationMaintenanceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeActiveOperationMaintenanceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeActiveOperationMaintenanceConfigResponse) GetBody() *DescribeActiveOperationMaintenanceConfigResponseBody {
	return s.Body
}

func (s *DescribeActiveOperationMaintenanceConfigResponse) SetHeaders(v map[string]*string) *DescribeActiveOperationMaintenanceConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeActiveOperationMaintenanceConfigResponse) SetStatusCode(v int32) *DescribeActiveOperationMaintenanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeActiveOperationMaintenanceConfigResponse) SetBody(v *DescribeActiveOperationMaintenanceConfigResponseBody) *DescribeActiveOperationMaintenanceConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeActiveOperationMaintenanceConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
