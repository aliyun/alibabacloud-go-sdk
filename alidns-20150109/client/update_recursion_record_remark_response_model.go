// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionRecordRemarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRecursionRecordRemarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRecursionRecordRemarkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRecursionRecordRemarkResponseBody) *UpdateRecursionRecordRemarkResponse
	GetBody() *UpdateRecursionRecordRemarkResponseBody
}

type UpdateRecursionRecordRemarkResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecursionRecordRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRecursionRecordRemarkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionRecordRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecursionRecordRemarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRecursionRecordRemarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRecursionRecordRemarkResponse) GetBody() *UpdateRecursionRecordRemarkResponseBody {
	return s.Body
}

func (s *UpdateRecursionRecordRemarkResponse) SetHeaders(v map[string]*string) *UpdateRecursionRecordRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecursionRecordRemarkResponse) SetStatusCode(v int32) *UpdateRecursionRecordRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecursionRecordRemarkResponse) SetBody(v *UpdateRecursionRecordRemarkResponseBody) *UpdateRecursionRecordRemarkResponse {
	s.Body = v
	return s
}

func (s *UpdateRecursionRecordRemarkResponse) Validate() error {
	return dara.Validate(s)
}
