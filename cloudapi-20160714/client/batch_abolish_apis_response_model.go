// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAbolishApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchAbolishApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchAbolishApisResponse
	GetStatusCode() *int32
	SetBody(v *BatchAbolishApisResponseBody) *BatchAbolishApisResponse
	GetBody() *BatchAbolishApisResponseBody
}

type BatchAbolishApisResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchAbolishApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchAbolishApisResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchAbolishApisResponse) GoString() string {
	return s.String()
}

func (s *BatchAbolishApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchAbolishApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchAbolishApisResponse) GetBody() *BatchAbolishApisResponseBody {
	return s.Body
}

func (s *BatchAbolishApisResponse) SetHeaders(v map[string]*string) *BatchAbolishApisResponse {
	s.Headers = v
	return s
}

func (s *BatchAbolishApisResponse) SetStatusCode(v int32) *BatchAbolishApisResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchAbolishApisResponse) SetBody(v *BatchAbolishApisResponseBody) *BatchAbolishApisResponse {
	s.Body = v
	return s
}

func (s *BatchAbolishApisResponse) Validate() error {
	return dara.Validate(s)
}
