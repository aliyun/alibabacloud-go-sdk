// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomizeFlowStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomizeFlowStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomizeFlowStrategyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomizeFlowStrategyResponseBody) *DeleteCustomizeFlowStrategyResponse
	GetBody() *DeleteCustomizeFlowStrategyResponseBody
}

type DeleteCustomizeFlowStrategyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomizeFlowStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomizeFlowStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomizeFlowStrategyResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomizeFlowStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomizeFlowStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomizeFlowStrategyResponse) GetBody() *DeleteCustomizeFlowStrategyResponseBody {
	return s.Body
}

func (s *DeleteCustomizeFlowStrategyResponse) SetHeaders(v map[string]*string) *DeleteCustomizeFlowStrategyResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomizeFlowStrategyResponse) SetStatusCode(v int32) *DeleteCustomizeFlowStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomizeFlowStrategyResponse) SetBody(v *DeleteCustomizeFlowStrategyResponseBody) *DeleteCustomizeFlowStrategyResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomizeFlowStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
