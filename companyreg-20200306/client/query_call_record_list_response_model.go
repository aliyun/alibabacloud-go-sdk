// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCallRecordListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCallRecordListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCallRecordListResponse
	GetStatusCode() *int32
	SetBody(v *QueryCallRecordListResponseBody) *QueryCallRecordListResponse
	GetBody() *QueryCallRecordListResponseBody
}

type QueryCallRecordListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCallRecordListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCallRecordListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCallRecordListResponse) GoString() string {
	return s.String()
}

func (s *QueryCallRecordListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCallRecordListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCallRecordListResponse) GetBody() *QueryCallRecordListResponseBody {
	return s.Body
}

func (s *QueryCallRecordListResponse) SetHeaders(v map[string]*string) *QueryCallRecordListResponse {
	s.Headers = v
	return s
}

func (s *QueryCallRecordListResponse) SetStatusCode(v int32) *QueryCallRecordListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCallRecordListResponse) SetBody(v *QueryCallRecordListResponseBody) *QueryCallRecordListResponse {
	s.Body = v
	return s
}

func (s *QueryCallRecordListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
