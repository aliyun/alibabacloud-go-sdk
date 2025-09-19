// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttackPathEventStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAttackPathEventStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAttackPathEventStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetAttackPathEventStatisticsResponseBody) *GetAttackPathEventStatisticsResponse
	GetBody() *GetAttackPathEventStatisticsResponseBody
}

type GetAttackPathEventStatisticsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAttackPathEventStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAttackPathEventStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAttackPathEventStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetAttackPathEventStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAttackPathEventStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAttackPathEventStatisticsResponse) GetBody() *GetAttackPathEventStatisticsResponseBody {
	return s.Body
}

func (s *GetAttackPathEventStatisticsResponse) SetHeaders(v map[string]*string) *GetAttackPathEventStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetAttackPathEventStatisticsResponse) SetStatusCode(v int32) *GetAttackPathEventStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAttackPathEventStatisticsResponse) SetBody(v *GetAttackPathEventStatisticsResponseBody) *GetAttackPathEventStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetAttackPathEventStatisticsResponse) Validate() error {
	return dara.Validate(s)
}
