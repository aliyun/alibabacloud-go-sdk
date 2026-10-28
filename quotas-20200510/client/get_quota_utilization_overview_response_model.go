// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaUtilizationOverviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQuotaUtilizationOverviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQuotaUtilizationOverviewResponse
	GetStatusCode() *int32
	SetBody(v *GetQuotaUtilizationOverviewResponseBody) *GetQuotaUtilizationOverviewResponse
	GetBody() *GetQuotaUtilizationOverviewResponseBody
}

type GetQuotaUtilizationOverviewResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQuotaUtilizationOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQuotaUtilizationOverviewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaUtilizationOverviewResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaUtilizationOverviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQuotaUtilizationOverviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQuotaUtilizationOverviewResponse) GetBody() *GetQuotaUtilizationOverviewResponseBody {
	return s.Body
}

func (s *GetQuotaUtilizationOverviewResponse) SetHeaders(v map[string]*string) *GetQuotaUtilizationOverviewResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaUtilizationOverviewResponse) SetStatusCode(v int32) *GetQuotaUtilizationOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaUtilizationOverviewResponse) SetBody(v *GetQuotaUtilizationOverviewResponseBody) *GetQuotaUtilizationOverviewResponse {
	s.Body = v
	return s
}

func (s *GetQuotaUtilizationOverviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
