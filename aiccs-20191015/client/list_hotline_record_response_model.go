// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotlineRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHotlineRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHotlineRecordResponse
	GetStatusCode() *int32
	SetBody(v *ListHotlineRecordResponseBody) *ListHotlineRecordResponse
	GetBody() *ListHotlineRecordResponseBody
}

type ListHotlineRecordResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotlineRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotlineRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHotlineRecordResponse) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHotlineRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHotlineRecordResponse) GetBody() *ListHotlineRecordResponseBody {
	return s.Body
}

func (s *ListHotlineRecordResponse) SetHeaders(v map[string]*string) *ListHotlineRecordResponse {
	s.Headers = v
	return s
}

func (s *ListHotlineRecordResponse) SetStatusCode(v int32) *ListHotlineRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotlineRecordResponse) SetBody(v *ListHotlineRecordResponseBody) *ListHotlineRecordResponse {
	s.Body = v
	return s
}

func (s *ListHotlineRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
