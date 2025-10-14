// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineRelatedRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRoutineRelatedRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRoutineRelatedRecordResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRoutineRelatedRecordResponseBody) *DeleteRoutineRelatedRecordResponse
	GetBody() *DeleteRoutineRelatedRecordResponseBody
}

type DeleteRoutineRelatedRecordResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRoutineRelatedRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRoutineRelatedRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineRelatedRecordResponse) GoString() string {
	return s.String()
}

func (s *DeleteRoutineRelatedRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRoutineRelatedRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRoutineRelatedRecordResponse) GetBody() *DeleteRoutineRelatedRecordResponseBody {
	return s.Body
}

func (s *DeleteRoutineRelatedRecordResponse) SetHeaders(v map[string]*string) *DeleteRoutineRelatedRecordResponse {
	s.Headers = v
	return s
}

func (s *DeleteRoutineRelatedRecordResponse) SetStatusCode(v int32) *DeleteRoutineRelatedRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRoutineRelatedRecordResponse) SetBody(v *DeleteRoutineRelatedRecordResponseBody) *DeleteRoutineRelatedRecordResponse {
	s.Body = v
	return s
}

func (s *DeleteRoutineRelatedRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
