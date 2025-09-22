// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckNFCBindPopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckNFCBindPopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckNFCBindPopResponse
	GetStatusCode() *int32
	SetBody(v *CheckNFCBindPopResponseBody) *CheckNFCBindPopResponse
	GetBody() *CheckNFCBindPopResponseBody
}

type CheckNFCBindPopResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckNFCBindPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckNFCBindPopResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckNFCBindPopResponse) GoString() string {
	return s.String()
}

func (s *CheckNFCBindPopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckNFCBindPopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckNFCBindPopResponse) GetBody() *CheckNFCBindPopResponseBody {
	return s.Body
}

func (s *CheckNFCBindPopResponse) SetHeaders(v map[string]*string) *CheckNFCBindPopResponse {
	s.Headers = v
	return s
}

func (s *CheckNFCBindPopResponse) SetStatusCode(v int32) *CheckNFCBindPopResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckNFCBindPopResponse) SetBody(v *CheckNFCBindPopResponseBody) *CheckNFCBindPopResponse {
	s.Body = v
	return s
}

func (s *CheckNFCBindPopResponse) Validate() error {
	return dara.Validate(s)
}
