// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventCenterRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEventCenterRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEventCenterRecordResponse
	GetStatusCode() *int32
	SetBody(v *ListEventCenterRecordResponseBody) *ListEventCenterRecordResponse
	GetBody() *ListEventCenterRecordResponseBody
}

type ListEventCenterRecordResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEventCenterRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEventCenterRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEventCenterRecordResponse) GoString() string {
	return s.String()
}

func (s *ListEventCenterRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEventCenterRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEventCenterRecordResponse) GetBody() *ListEventCenterRecordResponseBody {
	return s.Body
}

func (s *ListEventCenterRecordResponse) SetHeaders(v map[string]*string) *ListEventCenterRecordResponse {
	s.Headers = v
	return s
}

func (s *ListEventCenterRecordResponse) SetStatusCode(v int32) *ListEventCenterRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEventCenterRecordResponse) SetBody(v *ListEventCenterRecordResponseBody) *ListEventCenterRecordResponse {
	s.Body = v
	return s
}

func (s *ListEventCenterRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
