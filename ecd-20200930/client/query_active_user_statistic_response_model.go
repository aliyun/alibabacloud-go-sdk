// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryActiveUserStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryActiveUserStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryActiveUserStatisticResponse
	GetStatusCode() *int32
	SetBody(v *QueryActiveUserStatisticResponseBody) *QueryActiveUserStatisticResponse
	GetBody() *QueryActiveUserStatisticResponseBody
}

type QueryActiveUserStatisticResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryActiveUserStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryActiveUserStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryActiveUserStatisticResponse) GoString() string {
	return s.String()
}

func (s *QueryActiveUserStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryActiveUserStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryActiveUserStatisticResponse) GetBody() *QueryActiveUserStatisticResponseBody {
	return s.Body
}

func (s *QueryActiveUserStatisticResponse) SetHeaders(v map[string]*string) *QueryActiveUserStatisticResponse {
	s.Headers = v
	return s
}

func (s *QueryActiveUserStatisticResponse) SetStatusCode(v int32) *QueryActiveUserStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryActiveUserStatisticResponse) SetBody(v *QueryActiveUserStatisticResponseBody) *QueryActiveUserStatisticResponse {
	s.Body = v
	return s
}

func (s *QueryActiveUserStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
