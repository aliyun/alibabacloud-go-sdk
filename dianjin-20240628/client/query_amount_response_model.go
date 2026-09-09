// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAmountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAmountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAmountResponse
	GetStatusCode() *int32
	SetBody(v *QueryAmountResponseBody) *QueryAmountResponse
	GetBody() *QueryAmountResponseBody
}

type QueryAmountResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAmountResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAmountResponse) GoString() string {
	return s.String()
}

func (s *QueryAmountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAmountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAmountResponse) GetBody() *QueryAmountResponseBody {
	return s.Body
}

func (s *QueryAmountResponse) SetHeaders(v map[string]*string) *QueryAmountResponse {
	s.Headers = v
	return s
}

func (s *QueryAmountResponse) SetStatusCode(v int32) *QueryAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAmountResponse) SetBody(v *QueryAmountResponseBody) *QueryAmountResponse {
	s.Body = v
	return s
}

func (s *QueryAmountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
