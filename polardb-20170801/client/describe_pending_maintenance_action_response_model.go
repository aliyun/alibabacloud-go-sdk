// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePendingMaintenanceActionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePendingMaintenanceActionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePendingMaintenanceActionResponse
	GetStatusCode() *int32
	SetBody(v *DescribePendingMaintenanceActionResponseBody) *DescribePendingMaintenanceActionResponse
	GetBody() *DescribePendingMaintenanceActionResponseBody
}

type DescribePendingMaintenanceActionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePendingMaintenanceActionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePendingMaintenanceActionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePendingMaintenanceActionResponse) GoString() string {
	return s.String()
}

func (s *DescribePendingMaintenanceActionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePendingMaintenanceActionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePendingMaintenanceActionResponse) GetBody() *DescribePendingMaintenanceActionResponseBody {
	return s.Body
}

func (s *DescribePendingMaintenanceActionResponse) SetHeaders(v map[string]*string) *DescribePendingMaintenanceActionResponse {
	s.Headers = v
	return s
}

func (s *DescribePendingMaintenanceActionResponse) SetStatusCode(v int32) *DescribePendingMaintenanceActionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponse) SetBody(v *DescribePendingMaintenanceActionResponseBody) *DescribePendingMaintenanceActionResponse {
	s.Body = v
	return s
}

func (s *DescribePendingMaintenanceActionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
