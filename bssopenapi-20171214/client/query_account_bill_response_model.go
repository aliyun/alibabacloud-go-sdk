// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountBillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAccountBillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAccountBillResponse
	GetStatusCode() *int32
	SetBody(v *QueryAccountBillResponseBody) *QueryAccountBillResponse
	GetBody() *QueryAccountBillResponseBody
}

type QueryAccountBillResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAccountBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAccountBillResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountBillResponse) GoString() string {
	return s.String()
}

func (s *QueryAccountBillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAccountBillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAccountBillResponse) GetBody() *QueryAccountBillResponseBody {
	return s.Body
}

func (s *QueryAccountBillResponse) SetHeaders(v map[string]*string) *QueryAccountBillResponse {
	s.Headers = v
	return s
}

func (s *QueryAccountBillResponse) SetStatusCode(v int32) *QueryAccountBillResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAccountBillResponse) SetBody(v *QueryAccountBillResponseBody) *QueryAccountBillResponse {
	s.Body = v
	return s
}

func (s *QueryAccountBillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
