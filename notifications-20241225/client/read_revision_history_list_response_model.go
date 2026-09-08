// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadRevisionHistoryListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadRevisionHistoryListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadRevisionHistoryListResponse
	GetStatusCode() *int32
	SetBody(v *ReadRevisionHistoryListResponseBody) *ReadRevisionHistoryListResponse
	GetBody() *ReadRevisionHistoryListResponseBody
}

type ReadRevisionHistoryListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadRevisionHistoryListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadRevisionHistoryListResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadRevisionHistoryListResponse) GoString() string {
	return s.String()
}

func (s *ReadRevisionHistoryListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadRevisionHistoryListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadRevisionHistoryListResponse) GetBody() *ReadRevisionHistoryListResponseBody {
	return s.Body
}

func (s *ReadRevisionHistoryListResponse) SetHeaders(v map[string]*string) *ReadRevisionHistoryListResponse {
	s.Headers = v
	return s
}

func (s *ReadRevisionHistoryListResponse) SetStatusCode(v int32) *ReadRevisionHistoryListResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadRevisionHistoryListResponse) SetBody(v *ReadRevisionHistoryListResponseBody) *ReadRevisionHistoryListResponse {
	s.Body = v
	return s
}

func (s *ReadRevisionHistoryListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
