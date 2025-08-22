// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDcdnDeliverTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDcdnDeliverTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDcdnDeliverTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateDcdnDeliverTaskResponseBody) *CreateDcdnDeliverTaskResponse
	GetBody() *CreateDcdnDeliverTaskResponseBody
}

type CreateDcdnDeliverTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDcdnDeliverTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDcdnDeliverTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDcdnDeliverTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDcdnDeliverTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDcdnDeliverTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDcdnDeliverTaskResponse) GetBody() *CreateDcdnDeliverTaskResponseBody {
	return s.Body
}

func (s *CreateDcdnDeliverTaskResponse) SetHeaders(v map[string]*string) *CreateDcdnDeliverTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDcdnDeliverTaskResponse) SetStatusCode(v int32) *CreateDcdnDeliverTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDcdnDeliverTaskResponse) SetBody(v *CreateDcdnDeliverTaskResponseBody) *CreateDcdnDeliverTaskResponse {
	s.Body = v
	return s
}

func (s *CreateDcdnDeliverTaskResponse) Validate() error {
	return dara.Validate(s)
}
