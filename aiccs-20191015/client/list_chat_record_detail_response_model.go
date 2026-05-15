// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatRecordDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChatRecordDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChatRecordDetailResponse
	GetStatusCode() *int32
	SetBody(v *ListChatRecordDetailResponseBody) *ListChatRecordDetailResponse
	GetBody() *ListChatRecordDetailResponseBody
}

type ListChatRecordDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChatRecordDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChatRecordDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChatRecordDetailResponse) GoString() string {
	return s.String()
}

func (s *ListChatRecordDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChatRecordDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChatRecordDetailResponse) GetBody() *ListChatRecordDetailResponseBody {
	return s.Body
}

func (s *ListChatRecordDetailResponse) SetHeaders(v map[string]*string) *ListChatRecordDetailResponse {
	s.Headers = v
	return s
}

func (s *ListChatRecordDetailResponse) SetStatusCode(v int32) *ListChatRecordDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChatRecordDetailResponse) SetBody(v *ListChatRecordDetailResponseBody) *ListChatRecordDetailResponse {
	s.Body = v
	return s
}

func (s *ListChatRecordDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
