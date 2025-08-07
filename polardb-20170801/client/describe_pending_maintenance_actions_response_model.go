// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePendingMaintenanceActionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePendingMaintenanceActionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePendingMaintenanceActionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePendingMaintenanceActionsResponseBody) *DescribePendingMaintenanceActionsResponse
	GetBody() *DescribePendingMaintenanceActionsResponseBody
}

type DescribePendingMaintenanceActionsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePendingMaintenanceActionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePendingMaintenanceActionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePendingMaintenanceActionsResponse) GoString() string {
	return s.String()
}

func (s *DescribePendingMaintenanceActionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePendingMaintenanceActionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePendingMaintenanceActionsResponse) GetBody() *DescribePendingMaintenanceActionsResponseBody {
	return s.Body
}

func (s *DescribePendingMaintenanceActionsResponse) SetHeaders(v map[string]*string) *DescribePendingMaintenanceActionsResponse {
	s.Headers = v
	return s
}

func (s *DescribePendingMaintenanceActionsResponse) SetStatusCode(v int32) *DescribePendingMaintenanceActionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePendingMaintenanceActionsResponse) SetBody(v *DescribePendingMaintenanceActionsResponseBody) *DescribePendingMaintenanceActionsResponse {
	s.Body = v
	return s
}

func (s *DescribePendingMaintenanceActionsResponse) Validate() error {
	return dara.Validate(s)
}
