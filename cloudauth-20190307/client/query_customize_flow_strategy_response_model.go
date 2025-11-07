// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCustomizeFlowStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCustomizeFlowStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCustomizeFlowStrategyResponse
	GetStatusCode() *int32
	SetBody(v *QueryCustomizeFlowStrategyResponseBody) *QueryCustomizeFlowStrategyResponse
	GetBody() *QueryCustomizeFlowStrategyResponseBody
}

type QueryCustomizeFlowStrategyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCustomizeFlowStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCustomizeFlowStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomizeFlowStrategyResponse) GoString() string {
	return s.String()
}

func (s *QueryCustomizeFlowStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCustomizeFlowStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCustomizeFlowStrategyResponse) GetBody() *QueryCustomizeFlowStrategyResponseBody {
	return s.Body
}

func (s *QueryCustomizeFlowStrategyResponse) SetHeaders(v map[string]*string) *QueryCustomizeFlowStrategyResponse {
	s.Headers = v
	return s
}

func (s *QueryCustomizeFlowStrategyResponse) SetStatusCode(v int32) *QueryCustomizeFlowStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCustomizeFlowStrategyResponse) SetBody(v *QueryCustomizeFlowStrategyResponseBody) *QueryCustomizeFlowStrategyResponse {
	s.Body = v
	return s
}

func (s *QueryCustomizeFlowStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
