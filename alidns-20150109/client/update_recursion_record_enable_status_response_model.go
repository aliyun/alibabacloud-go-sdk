// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionRecordEnableStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRecursionRecordEnableStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRecursionRecordEnableStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRecursionRecordEnableStatusResponseBody) *UpdateRecursionRecordEnableStatusResponse
	GetBody() *UpdateRecursionRecordEnableStatusResponseBody
}

type UpdateRecursionRecordEnableStatusResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecursionRecordEnableStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRecursionRecordEnableStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionRecordEnableStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecursionRecordEnableStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRecursionRecordEnableStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRecursionRecordEnableStatusResponse) GetBody() *UpdateRecursionRecordEnableStatusResponseBody {
	return s.Body
}

func (s *UpdateRecursionRecordEnableStatusResponse) SetHeaders(v map[string]*string) *UpdateRecursionRecordEnableStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecursionRecordEnableStatusResponse) SetStatusCode(v int32) *UpdateRecursionRecordEnableStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecursionRecordEnableStatusResponse) SetBody(v *UpdateRecursionRecordEnableStatusResponseBody) *UpdateRecursionRecordEnableStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateRecursionRecordEnableStatusResponse) Validate() error {
	return dara.Validate(s)
}
