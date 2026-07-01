// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOperationRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOperationRecordResponse
	GetStatusCode() *int32
	SetBody(v *ListOperationRecordResponseBody) *ListOperationRecordResponse
	GetBody() *ListOperationRecordResponseBody
}

type ListOperationRecordResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOperationRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOperationRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOperationRecordResponse) GoString() string {
	return s.String()
}

func (s *ListOperationRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOperationRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOperationRecordResponse) GetBody() *ListOperationRecordResponseBody {
	return s.Body
}

func (s *ListOperationRecordResponse) SetHeaders(v map[string]*string) *ListOperationRecordResponse {
	s.Headers = v
	return s
}

func (s *ListOperationRecordResponse) SetStatusCode(v int32) *ListOperationRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOperationRecordResponse) SetBody(v *ListOperationRecordResponseBody) *ListOperationRecordResponse {
	s.Body = v
	return s
}

func (s *ListOperationRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
