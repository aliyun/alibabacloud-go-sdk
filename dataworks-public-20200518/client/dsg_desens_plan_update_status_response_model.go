// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgDesensPlanUpdateStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgDesensPlanUpdateStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgDesensPlanUpdateStatusResponse
	GetStatusCode() *int32
	SetBody(v *DsgDesensPlanUpdateStatusResponseBody) *DsgDesensPlanUpdateStatusResponse
	GetBody() *DsgDesensPlanUpdateStatusResponseBody
}

type DsgDesensPlanUpdateStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgDesensPlanUpdateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgDesensPlanUpdateStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanUpdateStatusResponse) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanUpdateStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgDesensPlanUpdateStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgDesensPlanUpdateStatusResponse) GetBody() *DsgDesensPlanUpdateStatusResponseBody {
	return s.Body
}

func (s *DsgDesensPlanUpdateStatusResponse) SetHeaders(v map[string]*string) *DsgDesensPlanUpdateStatusResponse {
	s.Headers = v
	return s
}

func (s *DsgDesensPlanUpdateStatusResponse) SetStatusCode(v int32) *DsgDesensPlanUpdateStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgDesensPlanUpdateStatusResponse) SetBody(v *DsgDesensPlanUpdateStatusResponseBody) *DsgDesensPlanUpdateStatusResponse {
	s.Body = v
	return s
}

func (s *DsgDesensPlanUpdateStatusResponse) Validate() error {
	return dara.Validate(s)
}
