// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionRecordWeightResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRecursionRecordWeightResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRecursionRecordWeightResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRecursionRecordWeightResponseBody) *UpdateRecursionRecordWeightResponse
	GetBody() *UpdateRecursionRecordWeightResponseBody
}

type UpdateRecursionRecordWeightResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecursionRecordWeightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRecursionRecordWeightResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionRecordWeightResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecursionRecordWeightResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRecursionRecordWeightResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRecursionRecordWeightResponse) GetBody() *UpdateRecursionRecordWeightResponseBody {
	return s.Body
}

func (s *UpdateRecursionRecordWeightResponse) SetHeaders(v map[string]*string) *UpdateRecursionRecordWeightResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecursionRecordWeightResponse) SetStatusCode(v int32) *UpdateRecursionRecordWeightResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecursionRecordWeightResponse) SetBody(v *UpdateRecursionRecordWeightResponseBody) *UpdateRecursionRecordWeightResponse {
	s.Body = v
	return s
}

func (s *UpdateRecursionRecordWeightResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
