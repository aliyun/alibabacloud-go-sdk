// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPriceEntityListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPriceEntityListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPriceEntityListResponse
	GetStatusCode() *int32
	SetBody(v *QueryPriceEntityListResponseBody) *QueryPriceEntityListResponse
	GetBody() *QueryPriceEntityListResponseBody
}

type QueryPriceEntityListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPriceEntityListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPriceEntityListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPriceEntityListResponse) GoString() string {
	return s.String()
}

func (s *QueryPriceEntityListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPriceEntityListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPriceEntityListResponse) GetBody() *QueryPriceEntityListResponseBody {
	return s.Body
}

func (s *QueryPriceEntityListResponse) SetHeaders(v map[string]*string) *QueryPriceEntityListResponse {
	s.Headers = v
	return s
}

func (s *QueryPriceEntityListResponse) SetStatusCode(v int32) *QueryPriceEntityListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPriceEntityListResponse) SetBody(v *QueryPriceEntityListResponseBody) *QueryPriceEntityListResponse {
	s.Body = v
	return s
}

func (s *QueryPriceEntityListResponse) Validate() error {
	return dara.Validate(s)
}
