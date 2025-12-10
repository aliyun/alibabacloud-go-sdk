// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentsUsersStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExperimentsUsersStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExperimentsUsersStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetExperimentsUsersStatisticsResponseBody) *GetExperimentsUsersStatisticsResponse
	GetBody() *GetExperimentsUsersStatisticsResponseBody
}

type GetExperimentsUsersStatisticsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExperimentsUsersStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExperimentsUsersStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentsUsersStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentsUsersStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExperimentsUsersStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExperimentsUsersStatisticsResponse) GetBody() *GetExperimentsUsersStatisticsResponseBody {
	return s.Body
}

func (s *GetExperimentsUsersStatisticsResponse) SetHeaders(v map[string]*string) *GetExperimentsUsersStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentsUsersStatisticsResponse) SetStatusCode(v int32) *GetExperimentsUsersStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentsUsersStatisticsResponse) SetBody(v *GetExperimentsUsersStatisticsResponseBody) *GetExperimentsUsersStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetExperimentsUsersStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
