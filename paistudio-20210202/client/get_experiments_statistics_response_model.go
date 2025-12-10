// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentsStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExperimentsStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExperimentsStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetExperimentsStatisticsResponseBody) *GetExperimentsStatisticsResponse
	GetBody() *GetExperimentsStatisticsResponseBody
}

type GetExperimentsStatisticsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExperimentsStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExperimentsStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentsStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentsStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExperimentsStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExperimentsStatisticsResponse) GetBody() *GetExperimentsStatisticsResponseBody {
	return s.Body
}

func (s *GetExperimentsStatisticsResponse) SetHeaders(v map[string]*string) *GetExperimentsStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentsStatisticsResponse) SetStatusCode(v int32) *GetExperimentsStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentsStatisticsResponse) SetBody(v *GetExperimentsStatisticsResponseBody) *GetExperimentsStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetExperimentsStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
