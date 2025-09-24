// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCostUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCostUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCostUnitResponse
	GetStatusCode() *int32
	SetBody(v *QueryCostUnitResponseBody) *QueryCostUnitResponse
	GetBody() *QueryCostUnitResponseBody
}

type QueryCostUnitResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCostUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCostUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCostUnitResponse) GoString() string {
	return s.String()
}

func (s *QueryCostUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCostUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCostUnitResponse) GetBody() *QueryCostUnitResponseBody {
	return s.Body
}

func (s *QueryCostUnitResponse) SetHeaders(v map[string]*string) *QueryCostUnitResponse {
	s.Headers = v
	return s
}

func (s *QueryCostUnitResponse) SetStatusCode(v int32) *QueryCostUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCostUnitResponse) SetBody(v *QueryCostUnitResponseBody) *QueryCostUnitResponse {
	s.Body = v
	return s
}

func (s *QueryCostUnitResponse) Validate() error {
	return dara.Validate(s)
}
