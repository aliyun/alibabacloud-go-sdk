// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResolveStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResolveStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResolveStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetResolveStatisticsResponseBody) *GetResolveStatisticsResponse
	GetBody() *GetResolveStatisticsResponseBody
}

type GetResolveStatisticsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResolveStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResolveStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResolveStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetResolveStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResolveStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResolveStatisticsResponse) GetBody() *GetResolveStatisticsResponseBody {
	return s.Body
}

func (s *GetResolveStatisticsResponse) SetHeaders(v map[string]*string) *GetResolveStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetResolveStatisticsResponse) SetStatusCode(v int32) *GetResolveStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResolveStatisticsResponse) SetBody(v *GetResolveStatisticsResponseBody) *GetResolveStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetResolveStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
