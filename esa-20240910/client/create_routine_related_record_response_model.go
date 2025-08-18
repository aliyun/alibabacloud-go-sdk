// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineRelatedRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRoutineRelatedRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRoutineRelatedRecordResponse
	GetStatusCode() *int32
	SetBody(v *CreateRoutineRelatedRecordResponseBody) *CreateRoutineRelatedRecordResponse
	GetBody() *CreateRoutineRelatedRecordResponseBody
}

type CreateRoutineRelatedRecordResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRoutineRelatedRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRoutineRelatedRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineRelatedRecordResponse) GoString() string {
	return s.String()
}

func (s *CreateRoutineRelatedRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRoutineRelatedRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRoutineRelatedRecordResponse) GetBody() *CreateRoutineRelatedRecordResponseBody {
	return s.Body
}

func (s *CreateRoutineRelatedRecordResponse) SetHeaders(v map[string]*string) *CreateRoutineRelatedRecordResponse {
	s.Headers = v
	return s
}

func (s *CreateRoutineRelatedRecordResponse) SetStatusCode(v int32) *CreateRoutineRelatedRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRoutineRelatedRecordResponse) SetBody(v *CreateRoutineRelatedRecordResponseBody) *CreateRoutineRelatedRecordResponse {
	s.Body = v
	return s
}

func (s *CreateRoutineRelatedRecordResponse) Validate() error {
	return dara.Validate(s)
}
