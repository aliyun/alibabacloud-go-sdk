// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectApproveCommandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RejectApproveCommandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RejectApproveCommandResponse
	GetStatusCode() *int32
	SetBody(v *RejectApproveCommandResponseBody) *RejectApproveCommandResponse
	GetBody() *RejectApproveCommandResponseBody
}

type RejectApproveCommandResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RejectApproveCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RejectApproveCommandResponse) String() string {
	return dara.Prettify(s)
}

func (s RejectApproveCommandResponse) GoString() string {
	return s.String()
}

func (s *RejectApproveCommandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RejectApproveCommandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RejectApproveCommandResponse) GetBody() *RejectApproveCommandResponseBody {
	return s.Body
}

func (s *RejectApproveCommandResponse) SetHeaders(v map[string]*string) *RejectApproveCommandResponse {
	s.Headers = v
	return s
}

func (s *RejectApproveCommandResponse) SetStatusCode(v int32) *RejectApproveCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *RejectApproveCommandResponse) SetBody(v *RejectApproveCommandResponseBody) *RejectApproveCommandResponse {
	s.Body = v
	return s
}

func (s *RejectApproveCommandResponse) Validate() error {
	return dara.Validate(s)
}
