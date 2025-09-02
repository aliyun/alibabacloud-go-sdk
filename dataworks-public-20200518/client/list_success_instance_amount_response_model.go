// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSuccessInstanceAmountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSuccessInstanceAmountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSuccessInstanceAmountResponse
	GetStatusCode() *int32
	SetBody(v *ListSuccessInstanceAmountResponseBody) *ListSuccessInstanceAmountResponse
	GetBody() *ListSuccessInstanceAmountResponseBody
}

type ListSuccessInstanceAmountResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSuccessInstanceAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSuccessInstanceAmountResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSuccessInstanceAmountResponse) GoString() string {
	return s.String()
}

func (s *ListSuccessInstanceAmountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSuccessInstanceAmountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSuccessInstanceAmountResponse) GetBody() *ListSuccessInstanceAmountResponseBody {
	return s.Body
}

func (s *ListSuccessInstanceAmountResponse) SetHeaders(v map[string]*string) *ListSuccessInstanceAmountResponse {
	s.Headers = v
	return s
}

func (s *ListSuccessInstanceAmountResponse) SetStatusCode(v int32) *ListSuccessInstanceAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSuccessInstanceAmountResponse) SetBody(v *ListSuccessInstanceAmountResponseBody) *ListSuccessInstanceAmountResponse {
	s.Body = v
	return s
}

func (s *ListSuccessInstanceAmountResponse) Validate() error {
	return dara.Validate(s)
}
