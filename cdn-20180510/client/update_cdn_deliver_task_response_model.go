// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCdnDeliverTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCdnDeliverTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCdnDeliverTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCdnDeliverTaskResponseBody) *UpdateCdnDeliverTaskResponse
	GetBody() *UpdateCdnDeliverTaskResponseBody
}

type UpdateCdnDeliverTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCdnDeliverTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCdnDeliverTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCdnDeliverTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateCdnDeliverTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCdnDeliverTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCdnDeliverTaskResponse) GetBody() *UpdateCdnDeliverTaskResponseBody {
	return s.Body
}

func (s *UpdateCdnDeliverTaskResponse) SetHeaders(v map[string]*string) *UpdateCdnDeliverTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateCdnDeliverTaskResponse) SetStatusCode(v int32) *UpdateCdnDeliverTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCdnDeliverTaskResponse) SetBody(v *UpdateCdnDeliverTaskResponseBody) *UpdateCdnDeliverTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateCdnDeliverTaskResponse) Validate() error {
	return dara.Validate(s)
}
