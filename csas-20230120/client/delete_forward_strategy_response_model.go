// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteForwardStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteForwardStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteForwardStrategyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteForwardStrategyResponseBody) *DeleteForwardStrategyResponse
	GetBody() *DeleteForwardStrategyResponseBody
}

type DeleteForwardStrategyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteForwardStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteForwardStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteForwardStrategyResponse) GoString() string {
	return s.String()
}

func (s *DeleteForwardStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteForwardStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteForwardStrategyResponse) GetBody() *DeleteForwardStrategyResponseBody {
	return s.Body
}

func (s *DeleteForwardStrategyResponse) SetHeaders(v map[string]*string) *DeleteForwardStrategyResponse {
	s.Headers = v
	return s
}

func (s *DeleteForwardStrategyResponse) SetStatusCode(v int32) *DeleteForwardStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteForwardStrategyResponse) SetBody(v *DeleteForwardStrategyResponseBody) *DeleteForwardStrategyResponse {
	s.Body = v
	return s
}

func (s *DeleteForwardStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
