// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecordTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRecordTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRecordTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRecordTaskResponseBody) *UpdateRecordTaskResponse
	GetBody() *UpdateRecordTaskResponseBody
}

type UpdateRecordTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecordTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRecordTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecordTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRecordTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRecordTaskResponse) GetBody() *UpdateRecordTaskResponseBody {
	return s.Body
}

func (s *UpdateRecordTaskResponse) SetHeaders(v map[string]*string) *UpdateRecordTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecordTaskResponse) SetStatusCode(v int32) *UpdateRecordTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecordTaskResponse) SetBody(v *UpdateRecordTaskResponseBody) *UpdateRecordTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateRecordTaskResponse) Validate() error {
	return dara.Validate(s)
}
