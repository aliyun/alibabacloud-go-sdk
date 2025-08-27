// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllBaseCityInfoQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllBaseCityInfoQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllBaseCityInfoQueryResponse
	GetStatusCode() *int32
	SetBody(v *AllBaseCityInfoQueryResponseBody) *AllBaseCityInfoQueryResponse
	GetBody() *AllBaseCityInfoQueryResponseBody
}

type AllBaseCityInfoQueryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllBaseCityInfoQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllBaseCityInfoQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s AllBaseCityInfoQueryResponse) GoString() string {
	return s.String()
}

func (s *AllBaseCityInfoQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllBaseCityInfoQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllBaseCityInfoQueryResponse) GetBody() *AllBaseCityInfoQueryResponseBody {
	return s.Body
}

func (s *AllBaseCityInfoQueryResponse) SetHeaders(v map[string]*string) *AllBaseCityInfoQueryResponse {
	s.Headers = v
	return s
}

func (s *AllBaseCityInfoQueryResponse) SetStatusCode(v int32) *AllBaseCityInfoQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *AllBaseCityInfoQueryResponse) SetBody(v *AllBaseCityInfoQueryResponseBody) *AllBaseCityInfoQueryResponse {
	s.Body = v
	return s
}

func (s *AllBaseCityInfoQueryResponse) Validate() error {
	return dara.Validate(s)
}
