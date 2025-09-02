// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgDesensPlanQueryListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgDesensPlanQueryListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgDesensPlanQueryListResponse
	GetStatusCode() *int32
	SetBody(v *DsgDesensPlanQueryListResponseBody) *DsgDesensPlanQueryListResponse
	GetBody() *DsgDesensPlanQueryListResponseBody
}

type DsgDesensPlanQueryListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgDesensPlanQueryListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgDesensPlanQueryListResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanQueryListResponse) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanQueryListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgDesensPlanQueryListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgDesensPlanQueryListResponse) GetBody() *DsgDesensPlanQueryListResponseBody {
	return s.Body
}

func (s *DsgDesensPlanQueryListResponse) SetHeaders(v map[string]*string) *DsgDesensPlanQueryListResponse {
	s.Headers = v
	return s
}

func (s *DsgDesensPlanQueryListResponse) SetStatusCode(v int32) *DsgDesensPlanQueryListResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgDesensPlanQueryListResponse) SetBody(v *DsgDesensPlanQueryListResponseBody) *DsgDesensPlanQueryListResponse {
	s.Body = v
	return s
}

func (s *DsgDesensPlanQueryListResponse) Validate() error {
	return dara.Validate(s)
}
