// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryApmGrafanaDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryApmGrafanaDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryApmGrafanaDataResponse
	GetStatusCode() *int32
	SetBody(v *QueryApmGrafanaDataResponseBody) *QueryApmGrafanaDataResponse
	GetBody() *QueryApmGrafanaDataResponseBody
}

type QueryApmGrafanaDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryApmGrafanaDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryApmGrafanaDataResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryApmGrafanaDataResponse) GoString() string {
	return s.String()
}

func (s *QueryApmGrafanaDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryApmGrafanaDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryApmGrafanaDataResponse) GetBody() *QueryApmGrafanaDataResponseBody {
	return s.Body
}

func (s *QueryApmGrafanaDataResponse) SetHeaders(v map[string]*string) *QueryApmGrafanaDataResponse {
	s.Headers = v
	return s
}

func (s *QueryApmGrafanaDataResponse) SetStatusCode(v int32) *QueryApmGrafanaDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryApmGrafanaDataResponse) SetBody(v *QueryApmGrafanaDataResponseBody) *QueryApmGrafanaDataResponse {
	s.Body = v
	return s
}

func (s *QueryApmGrafanaDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
