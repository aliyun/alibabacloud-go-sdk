// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgDesensPlanDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgDesensPlanDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgDesensPlanDeleteResponse
	GetStatusCode() *int32
	SetBody(v *DsgDesensPlanDeleteResponseBody) *DsgDesensPlanDeleteResponse
	GetBody() *DsgDesensPlanDeleteResponseBody
}

type DsgDesensPlanDeleteResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgDesensPlanDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgDesensPlanDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanDeleteResponse) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgDesensPlanDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgDesensPlanDeleteResponse) GetBody() *DsgDesensPlanDeleteResponseBody {
	return s.Body
}

func (s *DsgDesensPlanDeleteResponse) SetHeaders(v map[string]*string) *DsgDesensPlanDeleteResponse {
	s.Headers = v
	return s
}

func (s *DsgDesensPlanDeleteResponse) SetStatusCode(v int32) *DsgDesensPlanDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgDesensPlanDeleteResponse) SetBody(v *DsgDesensPlanDeleteResponseBody) *DsgDesensPlanDeleteResponse {
	s.Body = v
	return s
}

func (s *DsgDesensPlanDeleteResponse) Validate() error {
	return dara.Validate(s)
}
