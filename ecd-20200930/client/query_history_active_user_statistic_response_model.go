// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHistoryActiveUserStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryHistoryActiveUserStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryHistoryActiveUserStatisticResponse
	GetStatusCode() *int32
	SetBody(v *QueryHistoryActiveUserStatisticResponseBody) *QueryHistoryActiveUserStatisticResponse
	GetBody() *QueryHistoryActiveUserStatisticResponseBody
}

type QueryHistoryActiveUserStatisticResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryHistoryActiveUserStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryHistoryActiveUserStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryActiveUserStatisticResponse) GoString() string {
	return s.String()
}

func (s *QueryHistoryActiveUserStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryHistoryActiveUserStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryHistoryActiveUserStatisticResponse) GetBody() *QueryHistoryActiveUserStatisticResponseBody {
	return s.Body
}

func (s *QueryHistoryActiveUserStatisticResponse) SetHeaders(v map[string]*string) *QueryHistoryActiveUserStatisticResponse {
	s.Headers = v
	return s
}

func (s *QueryHistoryActiveUserStatisticResponse) SetStatusCode(v int32) *QueryHistoryActiveUserStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHistoryActiveUserStatisticResponse) SetBody(v *QueryHistoryActiveUserStatisticResponseBody) *QueryHistoryActiveUserStatisticResponse {
	s.Body = v
	return s
}

func (s *QueryHistoryActiveUserStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
