// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserProvisioningStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserProvisioningStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserProvisioningStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetUserProvisioningStatisticsResponseBody) *GetUserProvisioningStatisticsResponse
	GetBody() *GetUserProvisioningStatisticsResponseBody
}

type GetUserProvisioningStatisticsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserProvisioningStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserProvisioningStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserProvisioningStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserProvisioningStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserProvisioningStatisticsResponse) GetBody() *GetUserProvisioningStatisticsResponseBody {
	return s.Body
}

func (s *GetUserProvisioningStatisticsResponse) SetHeaders(v map[string]*string) *GetUserProvisioningStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetUserProvisioningStatisticsResponse) SetStatusCode(v int32) *GetUserProvisioningStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserProvisioningStatisticsResponse) SetBody(v *GetUserProvisioningStatisticsResponseBody) *GetUserProvisioningStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetUserProvisioningStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
