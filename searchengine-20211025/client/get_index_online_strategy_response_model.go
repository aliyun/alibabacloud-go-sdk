// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIndexOnlineStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIndexOnlineStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIndexOnlineStrategyResponse
	GetStatusCode() *int32
	SetBody(v *GetIndexOnlineStrategyResponseBody) *GetIndexOnlineStrategyResponse
	GetBody() *GetIndexOnlineStrategyResponseBody
}

type GetIndexOnlineStrategyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIndexOnlineStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIndexOnlineStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIndexOnlineStrategyResponse) GoString() string {
	return s.String()
}

func (s *GetIndexOnlineStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIndexOnlineStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIndexOnlineStrategyResponse) GetBody() *GetIndexOnlineStrategyResponseBody {
	return s.Body
}

func (s *GetIndexOnlineStrategyResponse) SetHeaders(v map[string]*string) *GetIndexOnlineStrategyResponse {
	s.Headers = v
	return s
}

func (s *GetIndexOnlineStrategyResponse) SetStatusCode(v int32) *GetIndexOnlineStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIndexOnlineStrategyResponse) SetBody(v *GetIndexOnlineStrategyResponseBody) *GetIndexOnlineStrategyResponse {
	s.Body = v
	return s
}

func (s *GetIndexOnlineStrategyResponse) Validate() error {
	return dara.Validate(s)
}
