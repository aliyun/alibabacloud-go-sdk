// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveWhiteListStrategyAssetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveWhiteListStrategyAssetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveWhiteListStrategyAssetsResponse
	GetStatusCode() *int32
	SetBody(v *SaveWhiteListStrategyAssetsResponseBody) *SaveWhiteListStrategyAssetsResponse
	GetBody() *SaveWhiteListStrategyAssetsResponseBody
}

type SaveWhiteListStrategyAssetsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveWhiteListStrategyAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveWhiteListStrategyAssetsResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveWhiteListStrategyAssetsResponse) GoString() string {
	return s.String()
}

func (s *SaveWhiteListStrategyAssetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveWhiteListStrategyAssetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveWhiteListStrategyAssetsResponse) GetBody() *SaveWhiteListStrategyAssetsResponseBody {
	return s.Body
}

func (s *SaveWhiteListStrategyAssetsResponse) SetHeaders(v map[string]*string) *SaveWhiteListStrategyAssetsResponse {
	s.Headers = v
	return s
}

func (s *SaveWhiteListStrategyAssetsResponse) SetStatusCode(v int32) *SaveWhiteListStrategyAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveWhiteListStrategyAssetsResponse) SetBody(v *SaveWhiteListStrategyAssetsResponseBody) *SaveWhiteListStrategyAssetsResponse {
	s.Body = v
	return s
}

func (s *SaveWhiteListStrategyAssetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
