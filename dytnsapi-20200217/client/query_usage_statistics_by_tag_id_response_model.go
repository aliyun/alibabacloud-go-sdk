// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUsageStatisticsByTagIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUsageStatisticsByTagIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUsageStatisticsByTagIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryUsageStatisticsByTagIdResponseBody) *QueryUsageStatisticsByTagIdResponse
	GetBody() *QueryUsageStatisticsByTagIdResponseBody
}

type QueryUsageStatisticsByTagIdResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUsageStatisticsByTagIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUsageStatisticsByTagIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUsageStatisticsByTagIdResponse) GoString() string {
	return s.String()
}

func (s *QueryUsageStatisticsByTagIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUsageStatisticsByTagIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUsageStatisticsByTagIdResponse) GetBody() *QueryUsageStatisticsByTagIdResponseBody {
	return s.Body
}

func (s *QueryUsageStatisticsByTagIdResponse) SetHeaders(v map[string]*string) *QueryUsageStatisticsByTagIdResponse {
	s.Headers = v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponse) SetStatusCode(v int32) *QueryUsageStatisticsByTagIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponse) SetBody(v *QueryUsageStatisticsByTagIdResponseBody) *QueryUsageStatisticsByTagIdResponse {
	s.Body = v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
