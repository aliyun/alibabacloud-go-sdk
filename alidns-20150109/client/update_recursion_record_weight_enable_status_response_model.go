// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionRecordWeightEnableStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRecursionRecordWeightEnableStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRecursionRecordWeightEnableStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRecursionRecordWeightEnableStatusResponseBody) *UpdateRecursionRecordWeightEnableStatusResponse
	GetBody() *UpdateRecursionRecordWeightEnableStatusResponseBody
}

type UpdateRecursionRecordWeightEnableStatusResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecursionRecordWeightEnableStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRecursionRecordWeightEnableStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionRecordWeightEnableStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecursionRecordWeightEnableStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRecursionRecordWeightEnableStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRecursionRecordWeightEnableStatusResponse) GetBody() *UpdateRecursionRecordWeightEnableStatusResponseBody {
	return s.Body
}

func (s *UpdateRecursionRecordWeightEnableStatusResponse) SetHeaders(v map[string]*string) *UpdateRecursionRecordWeightEnableStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecursionRecordWeightEnableStatusResponse) SetStatusCode(v int32) *UpdateRecursionRecordWeightEnableStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecursionRecordWeightEnableStatusResponse) SetBody(v *UpdateRecursionRecordWeightEnableStatusResponseBody) *UpdateRecursionRecordWeightEnableStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateRecursionRecordWeightEnableStatusResponse) Validate() error {
	return dara.Validate(s)
}
