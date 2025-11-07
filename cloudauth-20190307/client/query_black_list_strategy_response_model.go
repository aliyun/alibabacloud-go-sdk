// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBlackListStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryBlackListStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryBlackListStrategyResponse
	GetStatusCode() *int32
	SetBody(v *QueryBlackListStrategyResponseBody) *QueryBlackListStrategyResponse
	GetBody() *QueryBlackListStrategyResponseBody
}

type QueryBlackListStrategyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBlackListStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBlackListStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryBlackListStrategyResponse) GoString() string {
	return s.String()
}

func (s *QueryBlackListStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryBlackListStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryBlackListStrategyResponse) GetBody() *QueryBlackListStrategyResponseBody {
	return s.Body
}

func (s *QueryBlackListStrategyResponse) SetHeaders(v map[string]*string) *QueryBlackListStrategyResponse {
	s.Headers = v
	return s
}

func (s *QueryBlackListStrategyResponse) SetStatusCode(v int32) *QueryBlackListStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBlackListStrategyResponse) SetBody(v *QueryBlackListStrategyResponseBody) *QueryBlackListStrategyResponse {
	s.Body = v
	return s
}

func (s *QueryBlackListStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
