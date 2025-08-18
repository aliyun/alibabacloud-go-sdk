// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRecordResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRecordResponseBody) *UpdateRecordResponse
	GetBody() *UpdateRecordResponseBody
}

type UpdateRecordResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRecordResponse) GetBody() *UpdateRecordResponseBody {
	return s.Body
}

func (s *UpdateRecordResponse) SetHeaders(v map[string]*string) *UpdateRecordResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecordResponse) SetStatusCode(v int32) *UpdateRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecordResponse) SetBody(v *UpdateRecordResponseBody) *UpdateRecordResponse {
	s.Body = v
	return s
}

func (s *UpdateRecordResponse) Validate() error {
	return dara.Validate(s)
}
