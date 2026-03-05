// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDayAmountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDayAmountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDayAmountResponse
	GetStatusCode() *int32
	SetBody(v *ListDayAmountResponseBody) *ListDayAmountResponse
	GetBody() *ListDayAmountResponseBody
}

type ListDayAmountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDayAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDayAmountResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDayAmountResponse) GoString() string {
	return s.String()
}

func (s *ListDayAmountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDayAmountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDayAmountResponse) GetBody() *ListDayAmountResponseBody {
	return s.Body
}

func (s *ListDayAmountResponse) SetHeaders(v map[string]*string) *ListDayAmountResponse {
	s.Headers = v
	return s
}

func (s *ListDayAmountResponse) SetStatusCode(v int32) *ListDayAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDayAmountResponse) SetBody(v *ListDayAmountResponseBody) *ListDayAmountResponse {
	s.Body = v
	return s
}

func (s *ListDayAmountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
