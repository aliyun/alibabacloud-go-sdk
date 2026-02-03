// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartNisTrafficRankingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartNisTrafficRankingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartNisTrafficRankingResponse
	GetStatusCode() *int32
	SetBody(v *StartNisTrafficRankingResponseBody) *StartNisTrafficRankingResponse
	GetBody() *StartNisTrafficRankingResponseBody
}

type StartNisTrafficRankingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartNisTrafficRankingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartNisTrafficRankingResponse) String() string {
	return dara.Prettify(s)
}

func (s StartNisTrafficRankingResponse) GoString() string {
	return s.String()
}

func (s *StartNisTrafficRankingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartNisTrafficRankingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartNisTrafficRankingResponse) GetBody() *StartNisTrafficRankingResponseBody {
	return s.Body
}

func (s *StartNisTrafficRankingResponse) SetHeaders(v map[string]*string) *StartNisTrafficRankingResponse {
	s.Headers = v
	return s
}

func (s *StartNisTrafficRankingResponse) SetStatusCode(v int32) *StartNisTrafficRankingResponse {
	s.StatusCode = &v
	return s
}

func (s *StartNisTrafficRankingResponse) SetBody(v *StartNisTrafficRankingResponseBody) *StartNisTrafficRankingResponse {
	s.Body = v
	return s
}

func (s *StartNisTrafficRankingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
