// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpMarketItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMcpMarketItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMcpMarketItemsResponse
	GetStatusCode() *int32
	SetBody(v *ListMcpMarketItemsResponseBody) *ListMcpMarketItemsResponse
	GetBody() *ListMcpMarketItemsResponseBody
}

type ListMcpMarketItemsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMcpMarketItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMcpMarketItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMcpMarketItemsResponse) GoString() string {
	return s.String()
}

func (s *ListMcpMarketItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMcpMarketItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMcpMarketItemsResponse) GetBody() *ListMcpMarketItemsResponseBody {
	return s.Body
}

func (s *ListMcpMarketItemsResponse) SetHeaders(v map[string]*string) *ListMcpMarketItemsResponse {
	s.Headers = v
	return s
}

func (s *ListMcpMarketItemsResponse) SetStatusCode(v int32) *ListMcpMarketItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMcpMarketItemsResponse) SetBody(v *ListMcpMarketItemsResponseBody) *ListMcpMarketItemsResponse {
	s.Body = v
	return s
}

func (s *ListMcpMarketItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
