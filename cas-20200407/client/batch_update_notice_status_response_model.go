// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateNoticeStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchUpdateNoticeStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchUpdateNoticeStatusResponse
	GetStatusCode() *int32
	SetBody(v *BatchUpdateNoticeStatusResponseBody) *BatchUpdateNoticeStatusResponse
	GetBody() *BatchUpdateNoticeStatusResponseBody
}

type BatchUpdateNoticeStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUpdateNoticeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUpdateNoticeStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateNoticeStatusResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateNoticeStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchUpdateNoticeStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchUpdateNoticeStatusResponse) GetBody() *BatchUpdateNoticeStatusResponseBody {
	return s.Body
}

func (s *BatchUpdateNoticeStatusResponse) SetHeaders(v map[string]*string) *BatchUpdateNoticeStatusResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateNoticeStatusResponse) SetStatusCode(v int32) *BatchUpdateNoticeStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUpdateNoticeStatusResponse) SetBody(v *BatchUpdateNoticeStatusResponseBody) *BatchUpdateNoticeStatusResponse {
	s.Body = v
	return s
}

func (s *BatchUpdateNoticeStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
