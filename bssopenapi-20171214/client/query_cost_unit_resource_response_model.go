// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCostUnitResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCostUnitResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCostUnitResourceResponse
	GetStatusCode() *int32
	SetBody(v *QueryCostUnitResourceResponseBody) *QueryCostUnitResourceResponse
	GetBody() *QueryCostUnitResourceResponseBody
}

type QueryCostUnitResourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCostUnitResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCostUnitResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCostUnitResourceResponse) GoString() string {
	return s.String()
}

func (s *QueryCostUnitResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCostUnitResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCostUnitResourceResponse) GetBody() *QueryCostUnitResourceResponseBody {
	return s.Body
}

func (s *QueryCostUnitResourceResponse) SetHeaders(v map[string]*string) *QueryCostUnitResourceResponse {
	s.Headers = v
	return s
}

func (s *QueryCostUnitResourceResponse) SetStatusCode(v int32) *QueryCostUnitResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCostUnitResourceResponse) SetBody(v *QueryCostUnitResourceResponseBody) *QueryCostUnitResourceResponse {
	s.Body = v
	return s
}

func (s *QueryCostUnitResourceResponse) Validate() error {
	return dara.Validate(s)
}
