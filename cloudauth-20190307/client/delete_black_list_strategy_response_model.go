// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBlackListStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBlackListStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBlackListStrategyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBlackListStrategyResponseBody) *DeleteBlackListStrategyResponse
	GetBody() *DeleteBlackListStrategyResponseBody
}

type DeleteBlackListStrategyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBlackListStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBlackListStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBlackListStrategyResponse) GoString() string {
	return s.String()
}

func (s *DeleteBlackListStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBlackListStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBlackListStrategyResponse) GetBody() *DeleteBlackListStrategyResponseBody {
	return s.Body
}

func (s *DeleteBlackListStrategyResponse) SetHeaders(v map[string]*string) *DeleteBlackListStrategyResponse {
	s.Headers = v
	return s
}

func (s *DeleteBlackListStrategyResponse) SetStatusCode(v int32) *DeleteBlackListStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBlackListStrategyResponse) SetBody(v *DeleteBlackListStrategyResponseBody) *DeleteBlackListStrategyResponse {
	s.Body = v
	return s
}

func (s *DeleteBlackListStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
