// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotlineRecordDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHotlineRecordDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHotlineRecordDetailResponse
	GetStatusCode() *int32
	SetBody(v *ListHotlineRecordDetailResponseBody) *ListHotlineRecordDetailResponse
	GetBody() *ListHotlineRecordDetailResponseBody
}

type ListHotlineRecordDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotlineRecordDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotlineRecordDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHotlineRecordDetailResponse) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHotlineRecordDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHotlineRecordDetailResponse) GetBody() *ListHotlineRecordDetailResponseBody {
	return s.Body
}

func (s *ListHotlineRecordDetailResponse) SetHeaders(v map[string]*string) *ListHotlineRecordDetailResponse {
	s.Headers = v
	return s
}

func (s *ListHotlineRecordDetailResponse) SetStatusCode(v int32) *ListHotlineRecordDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotlineRecordDetailResponse) SetBody(v *ListHotlineRecordDetailResponseBody) *ListHotlineRecordDetailResponse {
	s.Body = v
	return s
}

func (s *ListHotlineRecordDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
