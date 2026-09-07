// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEndUserHistoryUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryEndUserHistoryUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryEndUserHistoryUsageResponse
	GetStatusCode() *int32
	SetBody(v *QueryEndUserHistoryUsageResponseBody) *QueryEndUserHistoryUsageResponse
	GetBody() *QueryEndUserHistoryUsageResponseBody
}

type QueryEndUserHistoryUsageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEndUserHistoryUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEndUserHistoryUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryEndUserHistoryUsageResponse) GoString() string {
	return s.String()
}

func (s *QueryEndUserHistoryUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryEndUserHistoryUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryEndUserHistoryUsageResponse) GetBody() *QueryEndUserHistoryUsageResponseBody {
	return s.Body
}

func (s *QueryEndUserHistoryUsageResponse) SetHeaders(v map[string]*string) *QueryEndUserHistoryUsageResponse {
	s.Headers = v
	return s
}

func (s *QueryEndUserHistoryUsageResponse) SetStatusCode(v int32) *QueryEndUserHistoryUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEndUserHistoryUsageResponse) SetBody(v *QueryEndUserHistoryUsageResponseBody) *QueryEndUserHistoryUsageResponse {
	s.Body = v
	return s
}

func (s *QueryEndUserHistoryUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
