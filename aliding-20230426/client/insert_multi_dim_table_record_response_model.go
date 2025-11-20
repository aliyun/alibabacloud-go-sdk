// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertMultiDimTableRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsertMultiDimTableRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsertMultiDimTableRecordResponse
	GetStatusCode() *int32
	SetBody(v *InsertMultiDimTableRecordResponseBody) *InsertMultiDimTableRecordResponse
	GetBody() *InsertMultiDimTableRecordResponseBody
}

type InsertMultiDimTableRecordResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertMultiDimTableRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertMultiDimTableRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s InsertMultiDimTableRecordResponse) GoString() string {
	return s.String()
}

func (s *InsertMultiDimTableRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsertMultiDimTableRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsertMultiDimTableRecordResponse) GetBody() *InsertMultiDimTableRecordResponseBody {
	return s.Body
}

func (s *InsertMultiDimTableRecordResponse) SetHeaders(v map[string]*string) *InsertMultiDimTableRecordResponse {
	s.Headers = v
	return s
}

func (s *InsertMultiDimTableRecordResponse) SetStatusCode(v int32) *InsertMultiDimTableRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertMultiDimTableRecordResponse) SetBody(v *InsertMultiDimTableRecordResponseBody) *InsertMultiDimTableRecordResponse {
	s.Body = v
	return s
}

func (s *InsertMultiDimTableRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
