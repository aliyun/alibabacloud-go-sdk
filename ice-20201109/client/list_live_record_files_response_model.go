// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveRecordFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLiveRecordFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLiveRecordFilesResponse
	GetStatusCode() *int32
	SetBody(v *ListLiveRecordFilesResponseBody) *ListLiveRecordFilesResponse
	GetBody() *ListLiveRecordFilesResponseBody
}

type ListLiveRecordFilesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLiveRecordFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLiveRecordFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRecordFilesResponse) GoString() string {
	return s.String()
}

func (s *ListLiveRecordFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLiveRecordFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLiveRecordFilesResponse) GetBody() *ListLiveRecordFilesResponseBody {
	return s.Body
}

func (s *ListLiveRecordFilesResponse) SetHeaders(v map[string]*string) *ListLiveRecordFilesResponse {
	s.Headers = v
	return s
}

func (s *ListLiveRecordFilesResponse) SetStatusCode(v int32) *ListLiveRecordFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveRecordFilesResponse) SetBody(v *ListLiveRecordFilesResponseBody) *ListLiveRecordFilesResponse {
	s.Body = v
	return s
}

func (s *ListLiveRecordFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
