// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnDeliverTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDcdnDeliverTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDcdnDeliverTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDcdnDeliverTaskResponseBody) *DeleteDcdnDeliverTaskResponse
	GetBody() *DeleteDcdnDeliverTaskResponseBody
}

type DeleteDcdnDeliverTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDcdnDeliverTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDcdnDeliverTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnDeliverTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteDcdnDeliverTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDcdnDeliverTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDcdnDeliverTaskResponse) GetBody() *DeleteDcdnDeliverTaskResponseBody {
	return s.Body
}

func (s *DeleteDcdnDeliverTaskResponse) SetHeaders(v map[string]*string) *DeleteDcdnDeliverTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteDcdnDeliverTaskResponse) SetStatusCode(v int32) *DeleteDcdnDeliverTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDcdnDeliverTaskResponse) SetBody(v *DeleteDcdnDeliverTaskResponseBody) *DeleteDcdnDeliverTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteDcdnDeliverTaskResponse) Validate() error {
	return dara.Validate(s)
}
