// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVulStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVulStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetVulStatisticsResponseBody) *GetVulStatisticsResponse
	GetBody() *GetVulStatisticsResponseBody
}

type GetVulStatisticsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVulStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVulStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVulStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetVulStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVulStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVulStatisticsResponse) GetBody() *GetVulStatisticsResponseBody {
	return s.Body
}

func (s *GetVulStatisticsResponse) SetHeaders(v map[string]*string) *GetVulStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetVulStatisticsResponse) SetStatusCode(v int32) *GetVulStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVulStatisticsResponse) SetBody(v *GetVulStatisticsResponseBody) *GetVulStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetVulStatisticsResponse) Validate() error {
	return dara.Validate(s)
}
