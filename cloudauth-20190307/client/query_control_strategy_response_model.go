// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryControlStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryControlStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryControlStrategyResponse
	GetStatusCode() *int32
	SetBody(v *QueryControlStrategyResponseBody) *QueryControlStrategyResponse
	GetBody() *QueryControlStrategyResponseBody
}

type QueryControlStrategyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryControlStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryControlStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryControlStrategyResponse) GoString() string {
	return s.String()
}

func (s *QueryControlStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryControlStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryControlStrategyResponse) GetBody() *QueryControlStrategyResponseBody {
	return s.Body
}

func (s *QueryControlStrategyResponse) SetHeaders(v map[string]*string) *QueryControlStrategyResponse {
	s.Headers = v
	return s
}

func (s *QueryControlStrategyResponse) SetStatusCode(v int32) *QueryControlStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryControlStrategyResponse) SetBody(v *QueryControlStrategyResponseBody) *QueryControlStrategyResponse {
	s.Body = v
	return s
}

func (s *QueryControlStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
