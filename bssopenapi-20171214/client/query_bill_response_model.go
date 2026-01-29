// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryBillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryBillResponse
	GetStatusCode() *int32
	SetBody(v *QueryBillResponseBody) *QueryBillResponse
	GetBody() *QueryBillResponseBody
}

type QueryBillResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBillResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryBillResponse) GoString() string {
	return s.String()
}

func (s *QueryBillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryBillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryBillResponse) GetBody() *QueryBillResponseBody {
	return s.Body
}

func (s *QueryBillResponse) SetHeaders(v map[string]*string) *QueryBillResponse {
	s.Headers = v
	return s
}

func (s *QueryBillResponse) SetStatusCode(v int32) *QueryBillResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBillResponse) SetBody(v *QueryBillResponseBody) *QueryBillResponse {
	s.Body = v
	return s
}

func (s *QueryBillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
