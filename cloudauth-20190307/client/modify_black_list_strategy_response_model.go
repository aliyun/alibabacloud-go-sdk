// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBlackListStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBlackListStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBlackListStrategyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBlackListStrategyResponseBody) *ModifyBlackListStrategyResponse
	GetBody() *ModifyBlackListStrategyResponseBody
}

type ModifyBlackListStrategyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBlackListStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBlackListStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBlackListStrategyResponse) GoString() string {
	return s.String()
}

func (s *ModifyBlackListStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBlackListStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBlackListStrategyResponse) GetBody() *ModifyBlackListStrategyResponseBody {
	return s.Body
}

func (s *ModifyBlackListStrategyResponse) SetHeaders(v map[string]*string) *ModifyBlackListStrategyResponse {
	s.Headers = v
	return s
}

func (s *ModifyBlackListStrategyResponse) SetStatusCode(v int32) *ModifyBlackListStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBlackListStrategyResponse) SetBody(v *ModifyBlackListStrategyResponseBody) *ModifyBlackListStrategyResponse {
	s.Body = v
	return s
}

func (s *ModifyBlackListStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
