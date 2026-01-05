// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryStatisticResponse
	GetStatusCode() *int32
	SetBody(v *QueryStatisticResponseBody) *QueryStatisticResponse
	GetBody() *QueryStatisticResponseBody
}

type QueryStatisticResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryStatisticResponse) GoString() string {
	return s.String()
}

func (s *QueryStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryStatisticResponse) GetBody() *QueryStatisticResponseBody {
	return s.Body
}

func (s *QueryStatisticResponse) SetHeaders(v map[string]*string) *QueryStatisticResponse {
	s.Headers = v
	return s
}

func (s *QueryStatisticResponse) SetStatusCode(v int32) *QueryStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryStatisticResponse) SetBody(v *QueryStatisticResponseBody) *QueryStatisticResponse {
	s.Body = v
	return s
}

func (s *QueryStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
