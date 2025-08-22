// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDcdnSubTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDcdnSubTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDcdnSubTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDcdnSubTaskResponseBody) *UpdateDcdnSubTaskResponse
	GetBody() *UpdateDcdnSubTaskResponseBody
}

type UpdateDcdnSubTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDcdnSubTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDcdnSubTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDcdnSubTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateDcdnSubTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDcdnSubTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDcdnSubTaskResponse) GetBody() *UpdateDcdnSubTaskResponseBody {
	return s.Body
}

func (s *UpdateDcdnSubTaskResponse) SetHeaders(v map[string]*string) *UpdateDcdnSubTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateDcdnSubTaskResponse) SetStatusCode(v int32) *UpdateDcdnSubTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDcdnSubTaskResponse) SetBody(v *UpdateDcdnSubTaskResponseBody) *UpdateDcdnSubTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateDcdnSubTaskResponse) Validate() error {
	return dara.Validate(s)
}
