// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveRecordVideoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLiveRecordVideoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLiveRecordVideoResponse
	GetStatusCode() *int32
	SetBody(v *ListLiveRecordVideoResponseBody) *ListLiveRecordVideoResponse
	GetBody() *ListLiveRecordVideoResponseBody
}

type ListLiveRecordVideoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLiveRecordVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLiveRecordVideoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRecordVideoResponse) GoString() string {
	return s.String()
}

func (s *ListLiveRecordVideoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLiveRecordVideoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLiveRecordVideoResponse) GetBody() *ListLiveRecordVideoResponseBody {
	return s.Body
}

func (s *ListLiveRecordVideoResponse) SetHeaders(v map[string]*string) *ListLiveRecordVideoResponse {
	s.Headers = v
	return s
}

func (s *ListLiveRecordVideoResponse) SetStatusCode(v int32) *ListLiveRecordVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveRecordVideoResponse) SetBody(v *ListLiveRecordVideoResponseBody) *ListLiveRecordVideoResponse {
	s.Body = v
	return s
}

func (s *ListLiveRecordVideoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
