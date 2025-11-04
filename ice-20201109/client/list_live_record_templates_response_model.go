// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveRecordTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLiveRecordTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLiveRecordTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListLiveRecordTemplatesResponseBody) *ListLiveRecordTemplatesResponse
	GetBody() *ListLiveRecordTemplatesResponseBody
}

type ListLiveRecordTemplatesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLiveRecordTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLiveRecordTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRecordTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListLiveRecordTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLiveRecordTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLiveRecordTemplatesResponse) GetBody() *ListLiveRecordTemplatesResponseBody {
	return s.Body
}

func (s *ListLiveRecordTemplatesResponse) SetHeaders(v map[string]*string) *ListLiveRecordTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListLiveRecordTemplatesResponse) SetStatusCode(v int32) *ListLiveRecordTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveRecordTemplatesResponse) SetBody(v *ListLiveRecordTemplatesResponseBody) *ListLiveRecordTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListLiveRecordTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
