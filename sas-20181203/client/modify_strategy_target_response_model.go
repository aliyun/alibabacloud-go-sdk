// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStrategyTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyStrategyTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyStrategyTargetResponse
	GetStatusCode() *int32
	SetBody(v *ModifyStrategyTargetResponseBody) *ModifyStrategyTargetResponse
	GetBody() *ModifyStrategyTargetResponseBody
}

type ModifyStrategyTargetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyStrategyTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyStrategyTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyStrategyTargetResponse) GoString() string {
	return s.String()
}

func (s *ModifyStrategyTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyStrategyTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyStrategyTargetResponse) GetBody() *ModifyStrategyTargetResponseBody {
	return s.Body
}

func (s *ModifyStrategyTargetResponse) SetHeaders(v map[string]*string) *ModifyStrategyTargetResponse {
	s.Headers = v
	return s
}

func (s *ModifyStrategyTargetResponse) SetStatusCode(v int32) *ModifyStrategyTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyStrategyTargetResponse) SetBody(v *ModifyStrategyTargetResponseBody) *ModifyStrategyTargetResponse {
	s.Body = v
	return s
}

func (s *ModifyStrategyTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
