// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteControlStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteControlStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteControlStrategyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteControlStrategyResponseBody) *DeleteControlStrategyResponse
	GetBody() *DeleteControlStrategyResponseBody
}

type DeleteControlStrategyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteControlStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteControlStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteControlStrategyResponse) GoString() string {
	return s.String()
}

func (s *DeleteControlStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteControlStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteControlStrategyResponse) GetBody() *DeleteControlStrategyResponseBody {
	return s.Body
}

func (s *DeleteControlStrategyResponse) SetHeaders(v map[string]*string) *DeleteControlStrategyResponse {
	s.Headers = v
	return s
}

func (s *DeleteControlStrategyResponse) SetStatusCode(v int32) *DeleteControlStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteControlStrategyResponse) SetBody(v *DeleteControlStrategyResponseBody) *DeleteControlStrategyResponse {
	s.Body = v
	return s
}

func (s *DeleteControlStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
