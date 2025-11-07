// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomizeFlowStrategyListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCustomizeFlowStrategyListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCustomizeFlowStrategyListResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCustomizeFlowStrategyListResponseBody) *ModifyCustomizeFlowStrategyListResponse
	GetBody() *ModifyCustomizeFlowStrategyListResponseBody
}

type ModifyCustomizeFlowStrategyListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCustomizeFlowStrategyListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCustomizeFlowStrategyListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomizeFlowStrategyListResponse) GoString() string {
	return s.String()
}

func (s *ModifyCustomizeFlowStrategyListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCustomizeFlowStrategyListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCustomizeFlowStrategyListResponse) GetBody() *ModifyCustomizeFlowStrategyListResponseBody {
	return s.Body
}

func (s *ModifyCustomizeFlowStrategyListResponse) SetHeaders(v map[string]*string) *ModifyCustomizeFlowStrategyListResponse {
	s.Headers = v
	return s
}

func (s *ModifyCustomizeFlowStrategyListResponse) SetStatusCode(v int32) *ModifyCustomizeFlowStrategyListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCustomizeFlowStrategyListResponse) SetBody(v *ModifyCustomizeFlowStrategyListResponseBody) *ModifyCustomizeFlowStrategyListResponse {
	s.Body = v
	return s
}

func (s *ModifyCustomizeFlowStrategyListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
