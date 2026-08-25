// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserProvisioningRdAccountStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserProvisioningRdAccountStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserProvisioningRdAccountStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetUserProvisioningRdAccountStatisticsResponseBody) *GetUserProvisioningRdAccountStatisticsResponse
	GetBody() *GetUserProvisioningRdAccountStatisticsResponseBody
}

type GetUserProvisioningRdAccountStatisticsResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserProvisioningRdAccountStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserProvisioningRdAccountStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserProvisioningRdAccountStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningRdAccountStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserProvisioningRdAccountStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserProvisioningRdAccountStatisticsResponse) GetBody() *GetUserProvisioningRdAccountStatisticsResponseBody {
	return s.Body
}

func (s *GetUserProvisioningRdAccountStatisticsResponse) SetHeaders(v map[string]*string) *GetUserProvisioningRdAccountStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetUserProvisioningRdAccountStatisticsResponse) SetStatusCode(v int32) *GetUserProvisioningRdAccountStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserProvisioningRdAccountStatisticsResponse) SetBody(v *GetUserProvisioningRdAccountStatisticsResponseBody) *GetUserProvisioningRdAccountStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetUserProvisioningRdAccountStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
