// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterCheckItemWarningStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClusterCheckItemWarningStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClusterCheckItemWarningStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetClusterCheckItemWarningStatisticsResponseBody) *GetClusterCheckItemWarningStatisticsResponse
	GetBody() *GetClusterCheckItemWarningStatisticsResponseBody
}

type GetClusterCheckItemWarningStatisticsResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterCheckItemWarningStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterCheckItemWarningStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClusterCheckItemWarningStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetClusterCheckItemWarningStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClusterCheckItemWarningStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClusterCheckItemWarningStatisticsResponse) GetBody() *GetClusterCheckItemWarningStatisticsResponseBody {
	return s.Body
}

func (s *GetClusterCheckItemWarningStatisticsResponse) SetHeaders(v map[string]*string) *GetClusterCheckItemWarningStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetClusterCheckItemWarningStatisticsResponse) SetStatusCode(v int32) *GetClusterCheckItemWarningStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterCheckItemWarningStatisticsResponse) SetBody(v *GetClusterCheckItemWarningStatisticsResponseBody) *GetClusterCheckItemWarningStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetClusterCheckItemWarningStatisticsResponse) Validate() error {
	return dara.Validate(s)
}
