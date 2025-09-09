// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptApproveCommandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AcceptApproveCommandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AcceptApproveCommandResponse
	GetStatusCode() *int32
	SetBody(v *AcceptApproveCommandResponseBody) *AcceptApproveCommandResponse
	GetBody() *AcceptApproveCommandResponseBody
}

type AcceptApproveCommandResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AcceptApproveCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AcceptApproveCommandResponse) String() string {
	return dara.Prettify(s)
}

func (s AcceptApproveCommandResponse) GoString() string {
	return s.String()
}

func (s *AcceptApproveCommandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AcceptApproveCommandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AcceptApproveCommandResponse) GetBody() *AcceptApproveCommandResponseBody {
	return s.Body
}

func (s *AcceptApproveCommandResponse) SetHeaders(v map[string]*string) *AcceptApproveCommandResponse {
	s.Headers = v
	return s
}

func (s *AcceptApproveCommandResponse) SetStatusCode(v int32) *AcceptApproveCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *AcceptApproveCommandResponse) SetBody(v *AcceptApproveCommandResponseBody) *AcceptApproveCommandResponse {
	s.Body = v
	return s
}

func (s *AcceptApproveCommandResponse) Validate() error {
	return dara.Validate(s)
}
