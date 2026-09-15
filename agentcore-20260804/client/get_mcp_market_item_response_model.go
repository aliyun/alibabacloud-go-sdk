// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpMarketItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMcpMarketItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMcpMarketItemResponse
	GetStatusCode() *int32
	SetBody(v *GetMcpMarketItemResponseBody) *GetMcpMarketItemResponse
	GetBody() *GetMcpMarketItemResponseBody
}

type GetMcpMarketItemResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMcpMarketItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMcpMarketItemResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMcpMarketItemResponse) GoString() string {
	return s.String()
}

func (s *GetMcpMarketItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMcpMarketItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMcpMarketItemResponse) GetBody() *GetMcpMarketItemResponseBody {
	return s.Body
}

func (s *GetMcpMarketItemResponse) SetHeaders(v map[string]*string) *GetMcpMarketItemResponse {
	s.Headers = v
	return s
}

func (s *GetMcpMarketItemResponse) SetStatusCode(v int32) *GetMcpMarketItemResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMcpMarketItemResponse) SetBody(v *GetMcpMarketItemResponseBody) *GetMcpMarketItemResponse {
	s.Body = v
	return s
}

func (s *GetMcpMarketItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
