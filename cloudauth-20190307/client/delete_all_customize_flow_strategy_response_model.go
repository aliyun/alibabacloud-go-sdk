// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAllCustomizeFlowStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAllCustomizeFlowStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAllCustomizeFlowStrategyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAllCustomizeFlowStrategyResponseBody) *DeleteAllCustomizeFlowStrategyResponse
	GetBody() *DeleteAllCustomizeFlowStrategyResponseBody
}

type DeleteAllCustomizeFlowStrategyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAllCustomizeFlowStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAllCustomizeFlowStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAllCustomizeFlowStrategyResponse) GoString() string {
	return s.String()
}

func (s *DeleteAllCustomizeFlowStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAllCustomizeFlowStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAllCustomizeFlowStrategyResponse) GetBody() *DeleteAllCustomizeFlowStrategyResponseBody {
	return s.Body
}

func (s *DeleteAllCustomizeFlowStrategyResponse) SetHeaders(v map[string]*string) *DeleteAllCustomizeFlowStrategyResponse {
	s.Headers = v
	return s
}

func (s *DeleteAllCustomizeFlowStrategyResponse) SetStatusCode(v int32) *DeleteAllCustomizeFlowStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAllCustomizeFlowStrategyResponse) SetBody(v *DeleteAllCustomizeFlowStrategyResponseBody) *DeleteAllCustomizeFlowStrategyResponse {
	s.Body = v
	return s
}

func (s *DeleteAllCustomizeFlowStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
