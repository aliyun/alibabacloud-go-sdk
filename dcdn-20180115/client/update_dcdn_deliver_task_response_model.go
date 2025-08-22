// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDcdnDeliverTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDcdnDeliverTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDcdnDeliverTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDcdnDeliverTaskResponseBody) *UpdateDcdnDeliverTaskResponse
	GetBody() *UpdateDcdnDeliverTaskResponseBody
}

type UpdateDcdnDeliverTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDcdnDeliverTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDcdnDeliverTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDcdnDeliverTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateDcdnDeliverTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDcdnDeliverTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDcdnDeliverTaskResponse) GetBody() *UpdateDcdnDeliverTaskResponseBody {
	return s.Body
}

func (s *UpdateDcdnDeliverTaskResponse) SetHeaders(v map[string]*string) *UpdateDcdnDeliverTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateDcdnDeliverTaskResponse) SetStatusCode(v int32) *UpdateDcdnDeliverTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDcdnDeliverTaskResponse) SetBody(v *UpdateDcdnDeliverTaskResponseBody) *UpdateDcdnDeliverTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateDcdnDeliverTaskResponse) Validate() error {
	return dara.Validate(s)
}
