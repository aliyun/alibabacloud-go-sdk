// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGateVerifyRecordListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryGateVerifyRecordListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryGateVerifyRecordListResponse
	GetStatusCode() *int32
	SetBody(v *QueryGateVerifyRecordListResponseBody) *QueryGateVerifyRecordListResponse
	GetBody() *QueryGateVerifyRecordListResponseBody
}

type QueryGateVerifyRecordListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGateVerifyRecordListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGateVerifyRecordListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryGateVerifyRecordListResponse) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyRecordListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryGateVerifyRecordListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryGateVerifyRecordListResponse) GetBody() *QueryGateVerifyRecordListResponseBody {
	return s.Body
}

func (s *QueryGateVerifyRecordListResponse) SetHeaders(v map[string]*string) *QueryGateVerifyRecordListResponse {
	s.Headers = v
	return s
}

func (s *QueryGateVerifyRecordListResponse) SetStatusCode(v int32) *QueryGateVerifyRecordListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGateVerifyRecordListResponse) SetBody(v *QueryGateVerifyRecordListResponseBody) *QueryGateVerifyRecordListResponse {
	s.Body = v
	return s
}

func (s *QueryGateVerifyRecordListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
