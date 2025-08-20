// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPullCashierResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PullCashierResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PullCashierResponse
	GetStatusCode() *int32
	SetBody(v *PullCashierResponseBody) *PullCashierResponse
	GetBody() *PullCashierResponseBody
}

type PullCashierResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PullCashierResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PullCashierResponse) String() string {
	return dara.Prettify(s)
}

func (s PullCashierResponse) GoString() string {
	return s.String()
}

func (s *PullCashierResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PullCashierResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PullCashierResponse) GetBody() *PullCashierResponseBody {
	return s.Body
}

func (s *PullCashierResponse) SetHeaders(v map[string]*string) *PullCashierResponse {
	s.Headers = v
	return s
}

func (s *PullCashierResponse) SetStatusCode(v int32) *PullCashierResponse {
	s.StatusCode = &v
	return s
}

func (s *PullCashierResponse) SetBody(v *PullCashierResponseBody) *PullCashierResponse {
	s.Body = v
	return s
}

func (s *PullCashierResponse) Validate() error {
	return dara.Validate(s)
}
