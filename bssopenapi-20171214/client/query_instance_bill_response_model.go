// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceBillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryInstanceBillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryInstanceBillResponse
	GetStatusCode() *int32
	SetBody(v *QueryInstanceBillResponseBody) *QueryInstanceBillResponse
	GetBody() *QueryInstanceBillResponseBody
}

type QueryInstanceBillResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInstanceBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInstanceBillResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceBillResponse) GoString() string {
	return s.String()
}

func (s *QueryInstanceBillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryInstanceBillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryInstanceBillResponse) GetBody() *QueryInstanceBillResponseBody {
	return s.Body
}

func (s *QueryInstanceBillResponse) SetHeaders(v map[string]*string) *QueryInstanceBillResponse {
	s.Headers = v
	return s
}

func (s *QueryInstanceBillResponse) SetStatusCode(v int32) *QueryInstanceBillResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInstanceBillResponse) SetBody(v *QueryInstanceBillResponseBody) *QueryInstanceBillResponse {
	s.Body = v
	return s
}

func (s *QueryInstanceBillResponse) Validate() error {
	return dara.Validate(s)
}
