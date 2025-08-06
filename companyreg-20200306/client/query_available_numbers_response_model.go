// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAvailableNumbersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAvailableNumbersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAvailableNumbersResponse
	GetStatusCode() *int32
	SetBody(v *QueryAvailableNumbersResponseBody) *QueryAvailableNumbersResponse
	GetBody() *QueryAvailableNumbersResponseBody
}

type QueryAvailableNumbersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAvailableNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAvailableNumbersResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAvailableNumbersResponse) GoString() string {
	return s.String()
}

func (s *QueryAvailableNumbersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAvailableNumbersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAvailableNumbersResponse) GetBody() *QueryAvailableNumbersResponseBody {
	return s.Body
}

func (s *QueryAvailableNumbersResponse) SetHeaders(v map[string]*string) *QueryAvailableNumbersResponse {
	s.Headers = v
	return s
}

func (s *QueryAvailableNumbersResponse) SetStatusCode(v int32) *QueryAvailableNumbersResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAvailableNumbersResponse) SetBody(v *QueryAvailableNumbersResponseBody) *QueryAvailableNumbersResponse {
	s.Body = v
	return s
}

func (s *QueryAvailableNumbersResponse) Validate() error {
	return dara.Validate(s)
}
