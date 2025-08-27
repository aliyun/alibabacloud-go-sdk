// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarOrderListQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CarOrderListQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CarOrderListQueryResponse
	GetStatusCode() *int32
	SetBody(v *CarOrderListQueryResponseBody) *CarOrderListQueryResponse
	GetBody() *CarOrderListQueryResponseBody
}

type CarOrderListQueryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CarOrderListQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CarOrderListQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s CarOrderListQueryResponse) GoString() string {
	return s.String()
}

func (s *CarOrderListQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CarOrderListQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CarOrderListQueryResponse) GetBody() *CarOrderListQueryResponseBody {
	return s.Body
}

func (s *CarOrderListQueryResponse) SetHeaders(v map[string]*string) *CarOrderListQueryResponse {
	s.Headers = v
	return s
}

func (s *CarOrderListQueryResponse) SetStatusCode(v int32) *CarOrderListQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *CarOrderListQueryResponse) SetBody(v *CarOrderListQueryResponseBody) *CarOrderListQueryResponse {
	s.Body = v
	return s
}

func (s *CarOrderListQueryResponse) Validate() error {
	return dara.Validate(s)
}
