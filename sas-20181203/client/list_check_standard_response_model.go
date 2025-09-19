// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckStandardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCheckStandardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCheckStandardResponse
	GetStatusCode() *int32
	SetBody(v *ListCheckStandardResponseBody) *ListCheckStandardResponse
	GetBody() *ListCheckStandardResponseBody
}

type ListCheckStandardResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCheckStandardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCheckStandardResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCheckStandardResponse) GoString() string {
	return s.String()
}

func (s *ListCheckStandardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCheckStandardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCheckStandardResponse) GetBody() *ListCheckStandardResponseBody {
	return s.Body
}

func (s *ListCheckStandardResponse) SetHeaders(v map[string]*string) *ListCheckStandardResponse {
	s.Headers = v
	return s
}

func (s *ListCheckStandardResponse) SetStatusCode(v int32) *ListCheckStandardResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCheckStandardResponse) SetBody(v *ListCheckStandardResponseBody) *ListCheckStandardResponse {
	s.Body = v
	return s
}

func (s *ListCheckStandardResponse) Validate() error {
	return dara.Validate(s)
}
