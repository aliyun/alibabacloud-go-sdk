// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRecursionRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRecursionRecordResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRecursionRecordResponseBody) *UpdateRecursionRecordResponse
	GetBody() *UpdateRecursionRecordResponseBody
}

type UpdateRecursionRecordResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecursionRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRecursionRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionRecordResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecursionRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRecursionRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRecursionRecordResponse) GetBody() *UpdateRecursionRecordResponseBody {
	return s.Body
}

func (s *UpdateRecursionRecordResponse) SetHeaders(v map[string]*string) *UpdateRecursionRecordResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecursionRecordResponse) SetStatusCode(v int32) *UpdateRecursionRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecursionRecordResponse) SetBody(v *UpdateRecursionRecordResponseBody) *UpdateRecursionRecordResponse {
	s.Body = v
	return s
}

func (s *UpdateRecursionRecordResponse) Validate() error {
	return dara.Validate(s)
}
