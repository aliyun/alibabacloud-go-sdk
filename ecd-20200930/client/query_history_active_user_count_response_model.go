// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHistoryActiveUserCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryHistoryActiveUserCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryHistoryActiveUserCountResponse
	GetStatusCode() *int32
	SetBody(v *QueryHistoryActiveUserCountResponseBody) *QueryHistoryActiveUserCountResponse
	GetBody() *QueryHistoryActiveUserCountResponseBody
}

type QueryHistoryActiveUserCountResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryHistoryActiveUserCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryHistoryActiveUserCountResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryActiveUserCountResponse) GoString() string {
	return s.String()
}

func (s *QueryHistoryActiveUserCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryHistoryActiveUserCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryHistoryActiveUserCountResponse) GetBody() *QueryHistoryActiveUserCountResponseBody {
	return s.Body
}

func (s *QueryHistoryActiveUserCountResponse) SetHeaders(v map[string]*string) *QueryHistoryActiveUserCountResponse {
	s.Headers = v
	return s
}

func (s *QueryHistoryActiveUserCountResponse) SetStatusCode(v int32) *QueryHistoryActiveUserCountResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHistoryActiveUserCountResponse) SetBody(v *QueryHistoryActiveUserCountResponseBody) *QueryHistoryActiveUserCountResponse {
	s.Body = v
	return s
}

func (s *QueryHistoryActiveUserCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
