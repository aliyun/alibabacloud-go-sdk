// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveWhiteListStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveWhiteListStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveWhiteListStrategyResponse
	GetStatusCode() *int32
	SetBody(v *SaveWhiteListStrategyResponseBody) *SaveWhiteListStrategyResponse
	GetBody() *SaveWhiteListStrategyResponseBody
}

type SaveWhiteListStrategyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveWhiteListStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveWhiteListStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveWhiteListStrategyResponse) GoString() string {
	return s.String()
}

func (s *SaveWhiteListStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveWhiteListStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveWhiteListStrategyResponse) GetBody() *SaveWhiteListStrategyResponseBody {
	return s.Body
}

func (s *SaveWhiteListStrategyResponse) SetHeaders(v map[string]*string) *SaveWhiteListStrategyResponse {
	s.Headers = v
	return s
}

func (s *SaveWhiteListStrategyResponse) SetStatusCode(v int32) *SaveWhiteListStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveWhiteListStrategyResponse) SetBody(v *SaveWhiteListStrategyResponseBody) *SaveWhiteListStrategyResponse {
	s.Body = v
	return s
}

func (s *SaveWhiteListStrategyResponse) Validate() error {
	return dara.Validate(s)
}
