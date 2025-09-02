// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceAmountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceAmountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceAmountResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceAmountResponseBody) *ListInstanceAmountResponse
	GetBody() *ListInstanceAmountResponseBody
}

type ListInstanceAmountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceAmountResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAmountResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceAmountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceAmountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceAmountResponse) GetBody() *ListInstanceAmountResponseBody {
	return s.Body
}

func (s *ListInstanceAmountResponse) SetHeaders(v map[string]*string) *ListInstanceAmountResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceAmountResponse) SetStatusCode(v int32) *ListInstanceAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceAmountResponse) SetBody(v *ListInstanceAmountResponseBody) *ListInstanceAmountResponse {
	s.Body = v
	return s
}

func (s *ListInstanceAmountResponse) Validate() error {
	return dara.Validate(s)
}
