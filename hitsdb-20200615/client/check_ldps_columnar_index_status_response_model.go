// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckLdpsColumnarIndexStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckLdpsColumnarIndexStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckLdpsColumnarIndexStatusResponse
	GetStatusCode() *int32
	SetBody(v *CheckLdpsColumnarIndexStatusResponseBody) *CheckLdpsColumnarIndexStatusResponse
	GetBody() *CheckLdpsColumnarIndexStatusResponseBody
}

type CheckLdpsColumnarIndexStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckLdpsColumnarIndexStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckLdpsColumnarIndexStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckLdpsColumnarIndexStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckLdpsColumnarIndexStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckLdpsColumnarIndexStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckLdpsColumnarIndexStatusResponse) GetBody() *CheckLdpsColumnarIndexStatusResponseBody {
	return s.Body
}

func (s *CheckLdpsColumnarIndexStatusResponse) SetHeaders(v map[string]*string) *CheckLdpsColumnarIndexStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckLdpsColumnarIndexStatusResponse) SetStatusCode(v int32) *CheckLdpsColumnarIndexStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckLdpsColumnarIndexStatusResponse) SetBody(v *CheckLdpsColumnarIndexStatusResponseBody) *CheckLdpsColumnarIndexStatusResponse {
	s.Body = v
	return s
}

func (s *CheckLdpsColumnarIndexStatusResponse) Validate() error {
	return dara.Validate(s)
}
