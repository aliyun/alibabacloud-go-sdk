// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRequestLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRequestLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRequestLogsResponse
	GetStatusCode() *int32
	SetBody(v *QueryRequestLogsResponseBody) *QueryRequestLogsResponse
	GetBody() *QueryRequestLogsResponseBody
}

type QueryRequestLogsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRequestLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRequestLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRequestLogsResponse) GoString() string {
	return s.String()
}

func (s *QueryRequestLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRequestLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRequestLogsResponse) GetBody() *QueryRequestLogsResponseBody {
	return s.Body
}

func (s *QueryRequestLogsResponse) SetHeaders(v map[string]*string) *QueryRequestLogsResponse {
	s.Headers = v
	return s
}

func (s *QueryRequestLogsResponse) SetStatusCode(v int32) *QueryRequestLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRequestLogsResponse) SetBody(v *QueryRequestLogsResponseBody) *QueryRequestLogsResponse {
	s.Body = v
	return s
}

func (s *QueryRequestLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
