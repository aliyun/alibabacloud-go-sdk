// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIndexOnlineStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyIndexOnlineStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyIndexOnlineStrategyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyIndexOnlineStrategyResponseBody) *ModifyIndexOnlineStrategyResponse
	GetBody() *ModifyIndexOnlineStrategyResponseBody
}

type ModifyIndexOnlineStrategyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIndexOnlineStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIndexOnlineStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyIndexOnlineStrategyResponse) GoString() string {
	return s.String()
}

func (s *ModifyIndexOnlineStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyIndexOnlineStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyIndexOnlineStrategyResponse) GetBody() *ModifyIndexOnlineStrategyResponseBody {
	return s.Body
}

func (s *ModifyIndexOnlineStrategyResponse) SetHeaders(v map[string]*string) *ModifyIndexOnlineStrategyResponse {
	s.Headers = v
	return s
}

func (s *ModifyIndexOnlineStrategyResponse) SetStatusCode(v int32) *ModifyIndexOnlineStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIndexOnlineStrategyResponse) SetBody(v *ModifyIndexOnlineStrategyResponseBody) *ModifyIndexOnlineStrategyResponse {
	s.Body = v
	return s
}

func (s *ModifyIndexOnlineStrategyResponse) Validate() error {
	return dara.Validate(s)
}
