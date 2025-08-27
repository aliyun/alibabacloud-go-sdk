// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarOrderQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CarOrderQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CarOrderQueryResponse
	GetStatusCode() *int32
	SetBody(v *CarOrderQueryResponseBody) *CarOrderQueryResponse
	GetBody() *CarOrderQueryResponseBody
}

type CarOrderQueryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CarOrderQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CarOrderQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s CarOrderQueryResponse) GoString() string {
	return s.String()
}

func (s *CarOrderQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CarOrderQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CarOrderQueryResponse) GetBody() *CarOrderQueryResponseBody {
	return s.Body
}

func (s *CarOrderQueryResponse) SetHeaders(v map[string]*string) *CarOrderQueryResponse {
	s.Headers = v
	return s
}

func (s *CarOrderQueryResponse) SetStatusCode(v int32) *CarOrderQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *CarOrderQueryResponse) SetBody(v *CarOrderQueryResponseBody) *CarOrderQueryResponse {
	s.Body = v
	return s
}

func (s *CarOrderQueryResponse) Validate() error {
	return dara.Validate(s)
}
