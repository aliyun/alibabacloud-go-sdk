// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySettleBillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySettleBillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySettleBillResponse
	GetStatusCode() *int32
	SetBody(v *QuerySettleBillResponseBody) *QuerySettleBillResponse
	GetBody() *QuerySettleBillResponseBody
}

type QuerySettleBillResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySettleBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySettleBillResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySettleBillResponse) GoString() string {
	return s.String()
}

func (s *QuerySettleBillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySettleBillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySettleBillResponse) GetBody() *QuerySettleBillResponseBody {
	return s.Body
}

func (s *QuerySettleBillResponse) SetHeaders(v map[string]*string) *QuerySettleBillResponse {
	s.Headers = v
	return s
}

func (s *QuerySettleBillResponse) SetStatusCode(v int32) *QuerySettleBillResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySettleBillResponse) SetBody(v *QuerySettleBillResponseBody) *QuerySettleBillResponse {
	s.Body = v
	return s
}

func (s *QuerySettleBillResponse) Validate() error {
	return dara.Validate(s)
}
