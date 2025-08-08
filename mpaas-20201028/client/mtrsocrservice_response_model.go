// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMTRSOCRServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MTRSOCRServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MTRSOCRServiceResponse
	GetStatusCode() *int32
	SetBody(v *MTRSOCRServiceResponseBody) *MTRSOCRServiceResponse
	GetBody() *MTRSOCRServiceResponseBody
}

type MTRSOCRServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MTRSOCRServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MTRSOCRServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s MTRSOCRServiceResponse) GoString() string {
	return s.String()
}

func (s *MTRSOCRServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MTRSOCRServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MTRSOCRServiceResponse) GetBody() *MTRSOCRServiceResponseBody {
	return s.Body
}

func (s *MTRSOCRServiceResponse) SetHeaders(v map[string]*string) *MTRSOCRServiceResponse {
	s.Headers = v
	return s
}

func (s *MTRSOCRServiceResponse) SetStatusCode(v int32) *MTRSOCRServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *MTRSOCRServiceResponse) SetBody(v *MTRSOCRServiceResponseBody) *MTRSOCRServiceResponse {
	s.Body = v
	return s
}

func (s *MTRSOCRServiceResponse) Validate() error {
	return dara.Validate(s)
}
